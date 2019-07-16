package ovpm

import (
	"fmt"
	"net"
	"time"

	passlib "gopkg.in/hlandau/passlib.v1"

	"github.com/asaskevich/govalidator"
	"github.com/cad/ovpm/pki"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

// dbRevokedModel is a database model for revoked VPN users.
type dbRevokedModel struct {
	gorm.Model
	SerialNumber string
}

// dbUserModel is database model for VPN users.
type dbUserModel struct {
	gorm.Model
	ServerID uint
	Server   dbServerModel

	Username           string `gorm:"unique_index"`
	Cert               string // not user writable
	ServerSerialNumber string // not user writable
	Hash               string
	Key                string // not user writable
	NoGW               bool
	HostID             uint32 // not user writable
	Admin              bool
	Device             bool
	DeviceSubNet       string
	AuthToken          string // auth token
}

// User represents a vpn user.
type User struct {
	dbUserModel // persisted fields

	isConnected    bool
	connectedSince time.Time
	bytesReceived  uint64
	bytesSent      uint64
}

func (u *dbUserModel) setPassword(password string) error {
	hashedPassword, err := passlib.Hash(password)
	if err != nil {
		return fmt.Errorf("can not set password: %v", err)
	}

	u.Hash = hashedPassword
	return nil
}

// RenewToken generates a new AuthToken and sets it to the db.
func (u *User) RenewToken() (string, error) {
	token := uuid.New().String()
	u.AuthToken = token
	db.Save(u.dbUserModel)
	if db.Error != nil {
		return "", db.Error
	}
	return token, nil
}

// ValidateToken returns whether the given token is valid or not.
func (u *User) ValidateToken(token string) bool {
	if u.AuthToken == "" {
		return false
	}
	return u.AuthToken == token
}

// CheckPassword returns whether the given password is correct for the user.
func (u *User) CheckPassword(password string) bool {
	_, err := passlib.Verify(password, u.Hash)
	if err != nil {
		logrus.Error(err)
		return false
	}
	return true
}

// GetUser finds and returns the user with the given username from database.
func GetUser(username string) (*User, error) {
	user := dbUserModel{}
	db.Where(&dbUserModel{Username: username}).First(&user)
	if db.NewRecord(&user) {
		// user is not found
		return nil, fmt.Errorf("user not found: %s", username)
	}
	return &User{dbUserModel: user}, nil
}

// GetUserByToken finds and returns the user with the given token from database.
func GetUserByToken(token string) (*User, error) {
	if token == "" {
		return nil, fmt.Errorf("token can not be empty")
	}

	user := dbUserModel{}
	db.Where(&dbUserModel{AuthToken: token}).First(&user)
	if db.NewRecord(&user) {
		// user is not found
		return nil, fmt.Errorf("user not found by token: <token>")
	}
	return &User{dbUserModel: user}, nil
}

// GetAllUsers returns all recorded users in the database.
func GetAllUsers() ([]*User, error) {
	var users []*User
	var dbUsers []*dbUserModel
	db.Find(&dbUsers)
	for _, u := range dbUsers {
		users = append(users, &User{dbUserModel: *u})
	}
	return users, nil
}

// CreateNewUser creates a new user with the given username and password in the database.
// If nogw is true, then ovpm doesn't push vpn server as the default gw for the user.
//
// It also generates the necessary client keys and signs certificates with the current
// server's CA.
func CreateNewUser(username, password string, nogw bool, hostid uint32, admin bool, device bool, deviceSubNet string) (*User, error) {
	svr := TheServer()
	if !svr.IsInitialized() {
		return nil, fmt.Errorf("you first need to create server")
	}
	// Validate user input.
	if govalidator.IsNull(username) {
		return nil, fmt.Errorf("validation error: %s can not be null", username)
	}
	if !govalidator.Matches(username, "^([\\w\\.]+)$") { // allow alphanumeric, underscore and dot
		return nil, fmt.Errorf("validation error: `%s` can only contain letters, numbers, underscores and dots", username)
	}
	if username == "root" {
		return nil, fmt.Errorf("forbidden: username root is reserved and can not be used")
	}

	ca, err := svr.GetSystemCA()
	if err != nil {
		return nil, err
	}

	clientCert, err := pki.NewClientCertHolder(ca, username)
	if err != nil {
		return nil, fmt.Errorf("can not create client cert %s: %v", username, err)
	}

	if hostid != 0 {
		ip := HostID2IP(hostid)
		if ip == nil {
			return nil, fmt.Errorf("host id doesn't represent an ip %d", hostid)
		}

		network := net.IPNet{IP: net.ParseIP(svr.Net).To4(), Mask: net.IPMask(net.ParseIP(svr.Mask).To4())}
		if !network.Contains(ip) {
			return nil, fmt.Errorf("ip %s, is out of vpn network %s", ip, network.String())
		}

		if hostIDsContains(getStaticHostIDs(), hostid) {
			return nil, fmt.Errorf("ip %s is already allocated", ip)
		}

		// Check if requested ip is allocated to the VPN server itself.
		serverNet := net.IPNet{
			IP:   net.ParseIP(svr.Net).To4(),
			Mask: net.IPMask(net.ParseIP(svr.Mask).To4()),
		}

		ip, ipnet, err := net.ParseCIDR(serverNet.String())
		if err != nil {
			return nil, fmt.Errorf("can not parse: %v", err)
		}
		if hostid == IP2HostID(ipnet.IP)+1 { // If it's VPN server's IP addr, then don't allow it.
			return nil, fmt.Errorf("can't assign server's ip address to a user")
		}
	}
	if deviceSubNet == "" {
		deviceSubNet = DefaultDeviceSubNetNetwork
	}
	user := dbUserModel{
		Username:           username,
		Cert:               clientCert.Cert,
		Key:                clientCert.Key,
		ServerSerialNumber: svr.SerialNumber,
		NoGW:               nogw,
		HostID:             hostid,
		Admin:              admin,
		Device:             device,
		DeviceSubNet:       deviceSubNet,
	}
	user.setPassword(password)

	db.Create(&user)
	if db.NewRecord(&user) {
		// user is still not created
		return nil, fmt.Errorf("can not create user in database: %s", user.Username)
	}
	logrus.Infof("user created: %s", username)

	// EmitWithRestart server config
	if err = svr.EmitWithRestart(); err != nil {
		return nil, err
	}
	return &User{dbUserModel: user}, nil
}

// Update updates the user's attributes and writes them to the database.
//
// How this method works is similiar to PUT semantics of REST. It sets the user record fields to the provided function arguments.
func (u *User) Update(password string, nogw bool, hostid uint32, admin bool, deviceSubNet string) error {
	svr := TheServer()
	if !svr.IsInitialized() {
		return fmt.Errorf("you first need to create server")
	}

	// If password is provided; set it. If not; leave it as it is.
	if password != "" {
		u.setPassword(password)
	}

	u.NoGW = nogw
	u.HostID = hostid
	u.Admin = admin
	u.DeviceSubNet = deviceSubNet

	if hostid != 0 {
		ip := HostID2IP(hostid)
		if ip == nil {
			return fmt.Errorf("host id doesn't represent an ip %d", hostid)
		}

		network := net.IPNet{IP: net.ParseIP(svr.Net).To4(), Mask: net.IPMask(net.ParseIP(svr.Mask).To4())}
		if !network.Contains(ip) {
			return fmt.Errorf("ip %s, is out of vpn network %s", ip, network.String())
		}

		if u.HostID != hostid && hostIDsContains(getStaticHostIDs(), hostid) {
			return fmt.Errorf("ip %s is already allocated", ip)
		}
	}
	db.Save(u.dbUserModel)

	return svr.EmitWithRestart()
}

// Delete deletes a user by the given username from the database.
func (u *User) Delete() error {
	if db.NewRecord(u.dbUserModel) {
		// user is not found
		return fmt.Errorf("user is not initialized: %s", u.Username)
	}
	crt, err := pki.ReadCertFromPEM(u.Cert)
	if err != nil {
		return fmt.Errorf("can not get user's certificate: %v", err)
	}
	db.Create(&dbRevokedModel{
		SerialNumber: crt.SerialNumber.Text(16),
	})
	db.Unscoped().Delete(u.dbUserModel)
	logrus.Infof("user deleted: %s", u.GetUsername())

	if err = TheServer().EmitWithRestart(); err != nil {
		return err
	}
	u = nil // delete the existing user struct
	return nil
}

// ResetPassword resets the users password into the provided password.
func (u *User) ResetPassword(password string) error {
	err := u.dbUserModel.setPassword(password)
	if err != nil {
		// user password can not be updated
		return fmt.Errorf("user password can not be updated %s: %v", u.Username, err)
	}
	db.Save(u.dbUserModel)
	if err = TheServer().EmitWithRestart(); err != nil {
		return err
	}

	logrus.Infof("user password reset: %s", u.GetUsername())
	return nil
}

// Renew creates a key and a ceritificate signed by the current server's CA.
//
// This is often used to sign users when the current CA is changed while there are
// still  existing users in the database.
//
// Also it can be used when a user cert is expired or user's private key stolen, missing etc.
func (u *User) Renew() error {
	svr := TheServer()
	if !svr.IsInitialized() {
		return fmt.Errorf("you first need to create server")
	}
	ca, err := svr.GetSystemCA()
	if err != nil {
		return err
	}

	clientCert, err := pki.NewClientCertHolder(ca, u.Username)
	if err != nil {
		return fmt.Errorf("can not create client cert %s: %v", u.Username, err)
	}

	u.Cert = clientCert.Cert
	u.Key = clientCert.Key
	u.ServerSerialNumber = svr.SerialNumber

	db.Save(u.dbUserModel)
	if err = svr.EmitWithRestart(); err != nil {
		return err
	}

	logrus.Infof("user renewed cert: %s", u.GetUsername())
	return nil
}

// GetUsername returns user's username.
func (u *User) GetUsername() string {
	return u.Username
}

// GetCert returns user's public certificate.
func (u *User) GetCert() string {
	return u.Cert
}

// ExpiresAt returns user's certificate expiration date time.
func (u *User) ExpiresAt() time.Time {
	crt, err := pki.ReadCertFromPEM(u.Cert)
	if err != nil {
		logrus.Fatalf("can't parse cert: %v", err)
	}
	return crt.NotAfter
}

// GetServerSerialNumber returns user's server serial number.
func (u *User) GetServerSerialNumber() string {
	return u.ServerSerialNumber
}

// GetCreatedAt returns user's creation time.
func (u *User) GetCreatedAt() string {
	return u.CreatedAt.Format(time.RFC3339)
}

// getIP returns user's vpn ip addr.
func (u *User) getIP() net.IP {
	users := getNonStaticHostUsers()
	staticHostIDs := getStaticHostIDs()
	svr := TheServer()
	mask := net.IPMask(net.ParseIP(svr.Mask).To4())
	network := net.ParseIP(svr.Net).To4().Mask(mask)

	// If the user has static ip address, return it immediately.
	if u.HostID != 0 {
		return HostID2IP(u.HostID)
	}

	// Calculate dynamic ip addresses from a deterministic address pool.
	freeHostID := 0
	for _, user := range users {
		// Skip, if user is supposed to have static ip.
		if user.HostID != 0 {
			continue
		}

		// Try the next available host id.
		hostID := IP2HostID(network) + uint32(freeHostID)
		for hostIDsContains(staticHostIDs, hostID+2) {
			freeHostID++ // Increase the host id and try again until it is available.
			hostID = IP2HostID(network) + uint32(freeHostID)
		}
		if user.ID == u.ID {
			return HostID2IP(hostID + 2)
		}
		freeHostID++
	}
	return nil
}

// GetIPNet returns user's vpn ip network. (e.g. 192.168.0.1/24)
func (u *User) GetIPNet() string {
	svr := TheServer()

	mask := net.IPMask(net.ParseIP(svr.Mask).To4())

	ipn := net.IPNet{
		IP:   u.getIP(),
		Mask: mask,
	}
	return ipn.String()
}

// IsNoGW returns whether user is set to get the vpn server as their default gateway.
func (u *User) IsNoGW() bool {
	return u.NoGW
}

// GetHostID returns user's Host ID.
func (u *User) GetHostID() uint32 {
	return u.HostID
}

// IsAdmin returns whether user is admin or not.
func (u *User) IsAdmin() bool {
	return u.Admin
}

// IsDevice returns whether user is device or not.
func (u *User) IsDevice() bool {
	return u.Device
}

// DeviceSubNet returns user's subnet ip network. (e.g. 192.168.1.1/24)
func (u *User) GetDeviceSubNet() string {
	return u.DeviceSubNet
}

// GetDeviceVSubNet returns user's virtual subnet for device.
func (u *User) GetDeviceVSubNet() string {
	ip := u.getIP()
	return net.IPv4(ip[0], ip[2]+1, ip[3], 0).String()
}

func (u *User) getKey() string {
	return u.Key
}

// ConnectionStatus returns information about user's connection to the VPN server.
func (u *User) ConnectionStatus() (isConnected bool, connectedSince time.Time, bytesSent uint64, bytesReceived uint64) {
	var found *clEntry

	svr := TheServer()

	// Open the status log file.
	f, err := svr.openFunc(_DefaultStatusLogPath)
	if err != nil {
		panic(err)
	}

	cl, _ := svr.parseStatusLogFunc(f) // client list from OpenVPN status log
	for _, c := range cl {
		if c.CommonName == u.Username {
			found = &c
		}
	}
	if found == nil {
		return false, time.Time{}, 0, 0
	}
	return true, found.ConnectedSince, found.BytesSent, found.BytesReceived
}

func getStaticHostUsers() []*User {
	var users []*User
	var dbUsers []*dbUserModel
	db.Unscoped().Not(dbUserModel{HostID: 0}).Find(&dbUsers)
	for _, u := range dbUsers {
		users = append(users, &User{dbUserModel: *u})
	}
	return users
}

func getNonStaticHostUsers() []*User {
	var users []*User
	var dbUsers []*dbUserModel
	db.Unscoped().Where(dbUserModel{HostID: 0}).Find(&dbUsers)
	for _, u := range dbUsers {
		users = append(users, &User{dbUserModel: *u})
	}

	return users
}

func getStaticHostIDs() []uint32 {
	var ids []uint32
	users := getStaticHostUsers()
	for _, user := range users {
		ids = append(ids, user.HostID)
	}

	return ids
}

func hostIDsContains(s []uint32, e uint32) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
