// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type UserUpdateRequest_GWPref int32

const (
	UserUpdateRequest_NOPREF UserUpdateRequest_GWPref = 0
	UserUpdateRequest_NOGW   UserUpdateRequest_GWPref = 1
	UserUpdateRequest_GW     UserUpdateRequest_GWPref = 2
)

var UserUpdateRequest_GWPref_name = map[int32]string{
	0: "NOPREF",
	1: "NOGW",
	2: "GW",
}

var UserUpdateRequest_GWPref_value = map[string]int32{
	"NOPREF": 0,
	"NOGW":   1,
	"GW":     2,
}

func (x UserUpdateRequest_GWPref) String() string {
	return proto.EnumName(UserUpdateRequest_GWPref_name, int32(x))
}

func (UserUpdateRequest_GWPref) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2, 0}
}

type UserUpdateRequest_StaticPref int32

const (
	UserUpdateRequest_NOPREFSTATIC UserUpdateRequest_StaticPref = 0
	UserUpdateRequest_NOSTATIC     UserUpdateRequest_StaticPref = 1
	UserUpdateRequest_STATIC       UserUpdateRequest_StaticPref = 2
)

var UserUpdateRequest_StaticPref_name = map[int32]string{
	0: "NOPREFSTATIC",
	1: "NOSTATIC",
	2: "STATIC",
}

var UserUpdateRequest_StaticPref_value = map[string]int32{
	"NOPREFSTATIC": 0,
	"NOSTATIC":     1,
	"STATIC":       2,
}

func (x UserUpdateRequest_StaticPref) String() string {
	return proto.EnumName(UserUpdateRequest_StaticPref_name, int32(x))
}

func (UserUpdateRequest_StaticPref) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2, 1}
}

type UserUpdateRequest_AdminPref int32

const (
	UserUpdateRequest_NOPREFADMIN UserUpdateRequest_AdminPref = 0
	UserUpdateRequest_NOADMIN     UserUpdateRequest_AdminPref = 1
	UserUpdateRequest_ADMIN       UserUpdateRequest_AdminPref = 2
)

var UserUpdateRequest_AdminPref_name = map[int32]string{
	0: "NOPREFADMIN",
	1: "NOADMIN",
	2: "ADMIN",
}

var UserUpdateRequest_AdminPref_value = map[string]int32{
	"NOPREFADMIN": 0,
	"NOADMIN":     1,
	"ADMIN":       2,
}

func (x UserUpdateRequest_AdminPref) String() string {
	return proto.EnumName(UserUpdateRequest_AdminPref_name, int32(x))
}

func (UserUpdateRequest_AdminPref) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2, 2}
}

type UserListRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserListRequest) Reset()         { *m = UserListRequest{} }
func (m *UserListRequest) String() string { return proto.CompactTextString(m) }
func (*UserListRequest) ProtoMessage()    {}
func (*UserListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *UserListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserListRequest.Unmarshal(m, b)
}
func (m *UserListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserListRequest.Marshal(b, m, deterministic)
}
func (m *UserListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserListRequest.Merge(m, src)
}
func (m *UserListRequest) XXX_Size() int {
	return xxx_messageInfo_UserListRequest.Size(m)
}
func (m *UserListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserListRequest proto.InternalMessageInfo

type UserCreateRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	NoGw                 bool     `protobuf:"varint,3,opt,name=no_gw,json=noGw,proto3" json:"no_gw,omitempty"`
	HostId               uint32   `protobuf:"varint,4,opt,name=host_id,json=hostId,proto3" json:"host_id,omitempty"`
	IsAdmin              bool     `protobuf:"varint,5,opt,name=is_admin,json=isAdmin,proto3" json:"is_admin,omitempty"`
	IsDevice             bool     `protobuf:"varint,6,opt,name=is_device,json=isDevice,proto3" json:"is_device,omitempty"`
	DeviceSubnet         string   `protobuf:"bytes,7,opt,name=device_subnet,json=deviceSubnet,proto3" json:"device_subnet,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserCreateRequest) Reset()         { *m = UserCreateRequest{} }
func (m *UserCreateRequest) String() string { return proto.CompactTextString(m) }
func (*UserCreateRequest) ProtoMessage()    {}
func (*UserCreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *UserCreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserCreateRequest.Unmarshal(m, b)
}
func (m *UserCreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserCreateRequest.Marshal(b, m, deterministic)
}
func (m *UserCreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserCreateRequest.Merge(m, src)
}
func (m *UserCreateRequest) XXX_Size() int {
	return xxx_messageInfo_UserCreateRequest.Size(m)
}
func (m *UserCreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserCreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserCreateRequest proto.InternalMessageInfo

func (m *UserCreateRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UserCreateRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *UserCreateRequest) GetNoGw() bool {
	if m != nil {
		return m.NoGw
	}
	return false
}

func (m *UserCreateRequest) GetHostId() uint32 {
	if m != nil {
		return m.HostId
	}
	return 0
}

func (m *UserCreateRequest) GetIsAdmin() bool {
	if m != nil {
		return m.IsAdmin
	}
	return false
}

func (m *UserCreateRequest) GetIsDevice() bool {
	if m != nil {
		return m.IsDevice
	}
	return false
}

func (m *UserCreateRequest) GetDeviceSubnet() string {
	if m != nil {
		return m.DeviceSubnet
	}
	return ""
}

type UserUpdateRequest struct {
	Username             string                       `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string                       `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Gwpref               UserUpdateRequest_GWPref     `protobuf:"varint,3,opt,name=gwpref,proto3,enum=pb.UserUpdateRequest_GWPref" json:"gwpref,omitempty"`
	HostId               uint32                       `protobuf:"varint,4,opt,name=host_id,json=hostId,proto3" json:"host_id,omitempty"`
	StaticPref           UserUpdateRequest_StaticPref `protobuf:"varint,5,opt,name=static_pref,json=staticPref,proto3,enum=pb.UserUpdateRequest_StaticPref" json:"static_pref,omitempty"`
	AdminPref            UserUpdateRequest_AdminPref  `protobuf:"varint,6,opt,name=admin_pref,json=adminPref,proto3,enum=pb.UserUpdateRequest_AdminPref" json:"admin_pref,omitempty"`
	DeviceSubnet         string                       `protobuf:"bytes,7,opt,name=device_subnet,json=deviceSubnet,proto3" json:"device_subnet,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *UserUpdateRequest) Reset()         { *m = UserUpdateRequest{} }
func (m *UserUpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UserUpdateRequest) ProtoMessage()    {}
func (*UserUpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *UserUpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserUpdateRequest.Unmarshal(m, b)
}
func (m *UserUpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserUpdateRequest.Marshal(b, m, deterministic)
}
func (m *UserUpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserUpdateRequest.Merge(m, src)
}
func (m *UserUpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UserUpdateRequest.Size(m)
}
func (m *UserUpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserUpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserUpdateRequest proto.InternalMessageInfo

func (m *UserUpdateRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UserUpdateRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *UserUpdateRequest) GetGwpref() UserUpdateRequest_GWPref {
	if m != nil {
		return m.Gwpref
	}
	return UserUpdateRequest_NOPREF
}

func (m *UserUpdateRequest) GetHostId() uint32 {
	if m != nil {
		return m.HostId
	}
	return 0
}

func (m *UserUpdateRequest) GetStaticPref() UserUpdateRequest_StaticPref {
	if m != nil {
		return m.StaticPref
	}
	return UserUpdateRequest_NOPREFSTATIC
}

func (m *UserUpdateRequest) GetAdminPref() UserUpdateRequest_AdminPref {
	if m != nil {
		return m.AdminPref
	}
	return UserUpdateRequest_NOPREFADMIN
}

func (m *UserUpdateRequest) GetDeviceSubnet() string {
	if m != nil {
		return m.DeviceSubnet
	}
	return ""
}

type UserDeleteRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserDeleteRequest) Reset()         { *m = UserDeleteRequest{} }
func (m *UserDeleteRequest) String() string { return proto.CompactTextString(m) }
func (*UserDeleteRequest) ProtoMessage()    {}
func (*UserDeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{3}
}

func (m *UserDeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserDeleteRequest.Unmarshal(m, b)
}
func (m *UserDeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserDeleteRequest.Marshal(b, m, deterministic)
}
func (m *UserDeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserDeleteRequest.Merge(m, src)
}
func (m *UserDeleteRequest) XXX_Size() int {
	return xxx_messageInfo_UserDeleteRequest.Size(m)
}
func (m *UserDeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserDeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserDeleteRequest proto.InternalMessageInfo

func (m *UserDeleteRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type UserRenewRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserRenewRequest) Reset()         { *m = UserRenewRequest{} }
func (m *UserRenewRequest) String() string { return proto.CompactTextString(m) }
func (*UserRenewRequest) ProtoMessage()    {}
func (*UserRenewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{4}
}

func (m *UserRenewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRenewRequest.Unmarshal(m, b)
}
func (m *UserRenewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRenewRequest.Marshal(b, m, deterministic)
}
func (m *UserRenewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRenewRequest.Merge(m, src)
}
func (m *UserRenewRequest) XXX_Size() int {
	return xxx_messageInfo_UserRenewRequest.Size(m)
}
func (m *UserRenewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRenewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserRenewRequest proto.InternalMessageInfo

func (m *UserRenewRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type UserGetRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserGetRequest) Reset()         { *m = UserGetRequest{} }
func (m *UserGetRequest) String() string { return proto.CompactTextString(m) }
func (*UserGetRequest) ProtoMessage()    {}
func (*UserGetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{5}
}

func (m *UserGetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserGetRequest.Unmarshal(m, b)
}
func (m *UserGetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserGetRequest.Marshal(b, m, deterministic)
}
func (m *UserGetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserGetRequest.Merge(m, src)
}
func (m *UserGetRequest) XXX_Size() int {
	return xxx_messageInfo_UserGetRequest.Size(m)
}
func (m *UserGetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserGetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserGetRequest proto.InternalMessageInfo

func (m *UserGetRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type UserGenConfigRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserGenConfigRequest) Reset()         { *m = UserGenConfigRequest{} }
func (m *UserGenConfigRequest) String() string { return proto.CompactTextString(m) }
func (*UserGenConfigRequest) ProtoMessage()    {}
func (*UserGenConfigRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{6}
}

func (m *UserGenConfigRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserGenConfigRequest.Unmarshal(m, b)
}
func (m *UserGenConfigRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserGenConfigRequest.Marshal(b, m, deterministic)
}
func (m *UserGenConfigRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserGenConfigRequest.Merge(m, src)
}
func (m *UserGenConfigRequest) XXX_Size() int {
	return xxx_messageInfo_UserGenConfigRequest.Size(m)
}
func (m *UserGenConfigRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserGenConfigRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserGenConfigRequest proto.InternalMessageInfo

func (m *UserGenConfigRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type UserResponse struct {
	Users                []*UserResponse_User `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *UserResponse) Reset()         { *m = UserResponse{} }
func (m *UserResponse) String() string { return proto.CompactTextString(m) }
func (*UserResponse) ProtoMessage()    {}
func (*UserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{7}
}

func (m *UserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserResponse.Unmarshal(m, b)
}
func (m *UserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserResponse.Marshal(b, m, deterministic)
}
func (m *UserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserResponse.Merge(m, src)
}
func (m *UserResponse) XXX_Size() int {
	return xxx_messageInfo_UserResponse.Size(m)
}
func (m *UserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserResponse proto.InternalMessageInfo

func (m *UserResponse) GetUsers() []*UserResponse_User {
	if m != nil {
		return m.Users
	}
	return nil
}

type UserResponse_User struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	ServerSerialNumber   string   `protobuf:"bytes,2,opt,name=server_serial_number,json=serverSerialNumber,proto3" json:"server_serial_number,omitempty"`
	Cert                 string   `protobuf:"bytes,3,opt,name=cert,proto3" json:"cert,omitempty"`
	CreatedAt            string   `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	IpNet                string   `protobuf:"bytes,5,opt,name=ip_net,json=ipNet,proto3" json:"ip_net,omitempty"`
	NoGw                 bool     `protobuf:"varint,6,opt,name=no_gw,json=noGw,proto3" json:"no_gw,omitempty"`
	HostId               uint32   `protobuf:"varint,7,opt,name=host_id,json=hostId,proto3" json:"host_id,omitempty"`
	IsAdmin              bool     `protobuf:"varint,8,opt,name=is_admin,json=isAdmin,proto3" json:"is_admin,omitempty"`
	IsConnected          bool     `protobuf:"varint,9,opt,name=is_connected,json=isConnected,proto3" json:"is_connected,omitempty"`
	ConnectedSince       string   `protobuf:"bytes,10,opt,name=connected_since,json=connectedSince,proto3" json:"connected_since,omitempty"`
	BytesSent            uint64   `protobuf:"varint,11,opt,name=bytes_sent,json=bytesSent,proto3" json:"bytes_sent,omitempty"`
	BytesReceived        uint64   `protobuf:"varint,12,opt,name=bytes_received,json=bytesReceived,proto3" json:"bytes_received,omitempty"`
	ExpiresAt            string   `protobuf:"bytes,13,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	IsDevice             bool     `protobuf:"varint,14,opt,name=is_device,json=isDevice,proto3" json:"is_device,omitempty"`
	DeviceSubnet         string   `protobuf:"bytes,15,opt,name=device_subnet,json=deviceSubnet,proto3" json:"device_subnet,omitempty"`
	DeviceVSubnet        string   `protobuf:"bytes,16,opt,name=device_v_subnet,json=deviceVSubnet,proto3" json:"device_v_subnet,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserResponse_User) Reset()         { *m = UserResponse_User{} }
func (m *UserResponse_User) String() string { return proto.CompactTextString(m) }
func (*UserResponse_User) ProtoMessage()    {}
func (*UserResponse_User) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{7, 0}
}

func (m *UserResponse_User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserResponse_User.Unmarshal(m, b)
}
func (m *UserResponse_User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserResponse_User.Marshal(b, m, deterministic)
}
func (m *UserResponse_User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserResponse_User.Merge(m, src)
}
func (m *UserResponse_User) XXX_Size() int {
	return xxx_messageInfo_UserResponse_User.Size(m)
}
func (m *UserResponse_User) XXX_DiscardUnknown() {
	xxx_messageInfo_UserResponse_User.DiscardUnknown(m)
}

var xxx_messageInfo_UserResponse_User proto.InternalMessageInfo

func (m *UserResponse_User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UserResponse_User) GetServerSerialNumber() string {
	if m != nil {
		return m.ServerSerialNumber
	}
	return ""
}

func (m *UserResponse_User) GetCert() string {
	if m != nil {
		return m.Cert
	}
	return ""
}

func (m *UserResponse_User) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *UserResponse_User) GetIpNet() string {
	if m != nil {
		return m.IpNet
	}
	return ""
}

func (m *UserResponse_User) GetNoGw() bool {
	if m != nil {
		return m.NoGw
	}
	return false
}

func (m *UserResponse_User) GetHostId() uint32 {
	if m != nil {
		return m.HostId
	}
	return 0
}

func (m *UserResponse_User) GetIsAdmin() bool {
	if m != nil {
		return m.IsAdmin
	}
	return false
}

func (m *UserResponse_User) GetIsConnected() bool {
	if m != nil {
		return m.IsConnected
	}
	return false
}

func (m *UserResponse_User) GetConnectedSince() string {
	if m != nil {
		return m.ConnectedSince
	}
	return ""
}

func (m *UserResponse_User) GetBytesSent() uint64 {
	if m != nil {
		return m.BytesSent
	}
	return 0
}

func (m *UserResponse_User) GetBytesReceived() uint64 {
	if m != nil {
		return m.BytesReceived
	}
	return 0
}

func (m *UserResponse_User) GetExpiresAt() string {
	if m != nil {
		return m.ExpiresAt
	}
	return ""
}

func (m *UserResponse_User) GetIsDevice() bool {
	if m != nil {
		return m.IsDevice
	}
	return false
}

func (m *UserResponse_User) GetDeviceSubnet() string {
	if m != nil {
		return m.DeviceSubnet
	}
	return ""
}

func (m *UserResponse_User) GetDeviceVSubnet() string {
	if m != nil {
		return m.DeviceVSubnet
	}
	return ""
}

type UserGenConfigResponse struct {
	ClientConfig         string   `protobuf:"bytes,1,opt,name=client_config,json=clientConfig,proto3" json:"client_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserGenConfigResponse) Reset()         { *m = UserGenConfigResponse{} }
func (m *UserGenConfigResponse) String() string { return proto.CompactTextString(m) }
func (*UserGenConfigResponse) ProtoMessage()    {}
func (*UserGenConfigResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{8}
}

func (m *UserGenConfigResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserGenConfigResponse.Unmarshal(m, b)
}
func (m *UserGenConfigResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserGenConfigResponse.Marshal(b, m, deterministic)
}
func (m *UserGenConfigResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserGenConfigResponse.Merge(m, src)
}
func (m *UserGenConfigResponse) XXX_Size() int {
	return xxx_messageInfo_UserGenConfigResponse.Size(m)
}
func (m *UserGenConfigResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserGenConfigResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserGenConfigResponse proto.InternalMessageInfo

func (m *UserGenConfigResponse) GetClientConfig() string {
	if m != nil {
		return m.ClientConfig
	}
	return ""
}

func init() {
	proto.RegisterEnum("pb.UserUpdateRequest_GWPref", UserUpdateRequest_GWPref_name, UserUpdateRequest_GWPref_value)
	proto.RegisterEnum("pb.UserUpdateRequest_StaticPref", UserUpdateRequest_StaticPref_name, UserUpdateRequest_StaticPref_value)
	proto.RegisterEnum("pb.UserUpdateRequest_AdminPref", UserUpdateRequest_AdminPref_name, UserUpdateRequest_AdminPref_value)
	proto.RegisterType((*UserListRequest)(nil), "pb.UserListRequest")
	proto.RegisterType((*UserCreateRequest)(nil), "pb.UserCreateRequest")
	proto.RegisterType((*UserUpdateRequest)(nil), "pb.UserUpdateRequest")
	proto.RegisterType((*UserDeleteRequest)(nil), "pb.UserDeleteRequest")
	proto.RegisterType((*UserRenewRequest)(nil), "pb.UserRenewRequest")
	proto.RegisterType((*UserGetRequest)(nil), "pb.UserGetRequest")
	proto.RegisterType((*UserGenConfigRequest)(nil), "pb.UserGenConfigRequest")
	proto.RegisterType((*UserResponse)(nil), "pb.UserResponse")
	proto.RegisterType((*UserResponse_User)(nil), "pb.UserResponse.User")
	proto.RegisterType((*UserGenConfigResponse)(nil), "pb.UserGenConfigResponse")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 913 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x96, 0x4d, 0x8f, 0xdb, 0x44,
	0x18, 0xc7, 0xd7, 0x59, 0xc7, 0x89, 0x9f, 0xbc, 0x79, 0x67, 0x77, 0xa9, 0x37, 0xdb, 0x42, 0xea,
	0x8a, 0x12, 0x15, 0x94, 0xc0, 0xd2, 0x03, 0xaa, 0x10, 0x52, 0xb4, 0x0b, 0xe9, 0x4a, 0x90, 0x2d,
	0x4e, 0xcb, 0x1e, 0x2d, 0xc7, 0x7e, 0x36, 0x8c, 0x94, 0x1d, 0x1b, 0xcf, 0x24, 0x81, 0x6b, 0xbf,
	0x02, 0x5f, 0x88, 0x6f, 0xc0, 0x01, 0x89, 0x03, 0x67, 0x3e, 0x08, 0x9a, 0x19, 0x27, 0xc1, 0xdb,
	0x14, 0x72, 0xe8, 0x6d, 0x9e, 0xb7, 0x9f, 0x66, 0xc6, 0xff, 0xfc, 0x27, 0x00, 0x73, 0x8e, 0x59,
	0x2f, 0xcd, 0x12, 0x91, 0x90, 0x52, 0x3a, 0x69, 0xdf, 0x9f, 0x26, 0xc9, 0x74, 0x86, 0xfd, 0x30,
	0xa5, 0xfd, 0x90, 0xb1, 0x44, 0x84, 0x82, 0x26, 0x8c, 0xeb, 0x0e, 0xef, 0x00, 0x5a, 0xaf, 0x38,
	0x66, 0xdf, 0x52, 0x2e, 0x7c, 0xfc, 0x69, 0x8e, 0x5c, 0x78, 0x7f, 0x1a, 0x70, 0x20, 0x73, 0xe7,
	0x19, 0x86, 0x02, 0xf3, 0x2c, 0x69, 0x43, 0x55, 0x82, 0x59, 0x78, 0x8b, 0xae, 0xd1, 0x31, 0xba,
	0xb6, 0xbf, 0x8e, 0x65, 0x2d, 0x0d, 0x39, 0x5f, 0x26, 0x59, 0xec, 0x96, 0x74, 0x6d, 0x15, 0x93,
	0x43, 0x28, 0xb3, 0x24, 0x98, 0x2e, 0xdd, 0xfd, 0x8e, 0xd1, 0xad, 0xfa, 0x26, 0x4b, 0x86, 0x4b,
	0x72, 0x0f, 0x2a, 0x3f, 0x26, 0x5c, 0x04, 0x34, 0x76, 0xcd, 0x8e, 0xd1, 0x6d, 0xf8, 0x96, 0x0c,
	0x2f, 0x63, 0x72, 0x02, 0x55, 0xca, 0x83, 0x30, 0xbe, 0xa5, 0xcc, 0x2d, 0xab, 0x81, 0x0a, 0xe5,
	0x03, 0x19, 0x92, 0x53, 0xb0, 0x29, 0x0f, 0x62, 0x5c, 0xd0, 0x08, 0x5d, 0x4b, 0xd5, 0xaa, 0x94,
	0x5f, 0xa8, 0x98, 0x3c, 0x82, 0x86, 0xae, 0x04, 0x7c, 0x3e, 0x61, 0x28, 0xdc, 0x8a, 0xda, 0x46,
	0x5d, 0x27, 0xc7, 0x2a, 0xe7, 0xfd, 0xb5, 0xaf, 0x0f, 0xf6, 0x2a, 0x8d, 0xdf, 0xc1, 0xc1, 0x9e,
	0x82, 0x35, 0x5d, 0xa6, 0x19, 0xde, 0xa8, 0x93, 0x35, 0xcf, 0xee, 0xf7, 0xd2, 0x49, 0xef, 0x0d,
	0x7c, 0x6f, 0x78, 0xfd, 0x22, 0xc3, 0x1b, 0x3f, 0xef, 0x7d, 0xfb, 0xc9, 0x07, 0x50, 0xe3, 0xf2,
	0xd3, 0x44, 0x81, 0x62, 0x96, 0x15, 0xb3, 0xb3, 0x9d, 0x39, 0x56, 0x8d, 0x8a, 0x0b, 0x7c, 0xbd,
	0x26, 0x5f, 0x01, 0xa8, 0x9b, 0xd3, 0x04, 0x4b, 0x11, 0x3e, 0xd8, 0x4e, 0x50, 0x57, 0xaa, 0x00,
	0x76, 0xb8, 0x5a, 0xee, 0x76, 0x89, 0x8f, 0xc1, 0xd2, 0x47, 0x22, 0x00, 0xd6, 0xe8, 0xea, 0x85,
	0xff, 0xf5, 0x37, 0xce, 0x1e, 0xa9, 0x82, 0x39, 0xba, 0x1a, 0x5e, 0x3b, 0x06, 0xb1, 0xa0, 0x34,
	0xbc, 0x76, 0x4a, 0xde, 0x17, 0x00, 0x9b, 0x6d, 0x12, 0x07, 0xea, 0xba, 0x77, 0xfc, 0x72, 0xf0,
	0xf2, 0xf2, 0xdc, 0xd9, 0x23, 0x75, 0xa8, 0x8e, 0xae, 0xf2, 0xc8, 0x90, 0xac, 0x7c, 0x5d, 0xf2,
	0x9e, 0x82, 0xbd, 0xde, 0x1e, 0x69, 0x41, 0x4d, 0x0f, 0x0e, 0x2e, 0xbe, 0xbb, 0x1c, 0x39, 0x7b,
	0xa4, 0x06, 0x95, 0xd1, 0x95, 0x0e, 0x0c, 0x62, 0x43, 0x59, 0x2f, 0x4b, 0x5e, 0x5f, 0x7f, 0xdb,
	0x0b, 0x9c, 0xe1, 0x4e, 0xdf, 0xd6, 0xeb, 0x81, 0x23, 0x07, 0x7c, 0x64, 0xb8, 0xdc, 0xa5, 0xff,
	0x13, 0x68, 0xca, 0xfe, 0x21, 0x8a, 0x5d, 0xba, 0xcf, 0xe0, 0x48, 0x77, 0xb3, 0xf3, 0x84, 0xdd,
	0xd0, 0xe9, 0x2e, 0x33, 0xbf, 0x99, 0x50, 0xd7, 0x5b, 0xe2, 0x69, 0xc2, 0x38, 0x92, 0x8f, 0xa1,
	0x2c, 0x8b, 0xdc, 0x35, 0x3a, 0xfb, 0xdd, 0xda, 0xd9, 0xf1, 0xea, 0x5b, 0xae, 0x1a, 0x74, 0xa0,
	0x7b, 0xda, 0xaf, 0x4d, 0x30, 0x65, 0xfc, 0x9f, 0x82, 0xfe, 0x14, 0x8e, 0x38, 0x66, 0x0b, 0xcc,
	0x02, 0x8e, 0x19, 0x0d, 0x67, 0x01, 0x9b, 0xdf, 0x4e, 0x30, 0xcb, 0xc5, 0x4d, 0x74, 0x6d, 0xac,
	0x4a, 0x23, 0x55, 0x21, 0x04, 0xcc, 0x08, 0x33, 0xa1, 0x44, 0x6e, 0xfb, 0x6a, 0x4d, 0x1e, 0x00,
	0x44, 0xca, 0x1c, 0xe2, 0x20, 0x14, 0x4a, 0xc7, 0xb6, 0x6f, 0xe7, 0x99, 0x81, 0x20, 0xc7, 0x60,
	0xd1, 0x34, 0x90, 0x02, 0x2a, 0xab, 0x52, 0x99, 0xa6, 0x23, 0x14, 0x1b, 0x27, 0xb0, 0xb6, 0x3b,
	0x41, 0xe5, 0xad, 0x4e, 0x50, 0x2d, 0x3a, 0xc1, 0x43, 0xa8, 0x53, 0x1e, 0x44, 0x09, 0x63, 0x18,
	0x09, 0x8c, 0x5d, 0x5b, 0x95, 0x6b, 0x94, 0x9f, 0xaf, 0x52, 0xe4, 0x23, 0x68, 0xad, 0xeb, 0x01,
	0xa7, 0x2c, 0x42, 0x17, 0xd4, 0x5e, 0x9a, 0xeb, 0xf4, 0x58, 0x66, 0xe5, 0x51, 0x26, 0xbf, 0x08,
	0xe4, 0x01, 0x47, 0x26, 0xdc, 0x5a, 0xc7, 0xe8, 0x9a, 0xbe, 0xad, 0x32, 0x63, 0x64, 0x82, 0x7c,
	0x08, 0x4d, 0x5d, 0xce, 0x30, 0x42, 0xba, 0xc0, 0xd8, 0xad, 0xab, 0x96, 0x86, 0xca, 0xfa, 0x79,
	0x52, 0x52, 0xf0, 0xe7, 0x94, 0x66, 0xc8, 0xe5, 0x85, 0x34, 0xf4, 0x85, 0xe4, 0x99, 0x81, 0x28,
	0x5a, 0x57, 0xf3, 0xff, 0xac, 0xab, 0xf5, 0xe6, 0xaf, 0x8e, 0x3c, 0x86, 0x56, 0xde, 0xb4, 0x58,
	0xb5, 0x39, 0xaa, 0x2d, 0x9f, 0xfd, 0x21, 0xff, 0x75, 0x7e, 0x09, 0xc7, 0x77, 0x64, 0x97, 0x4b,
	0xe9, 0x11, 0x34, 0xa2, 0x19, 0x45, 0x26, 0xe4, 0xbd, 0xdd, 0xd0, 0x69, 0xae, 0x8c, 0xba, 0x4e,
	0xea, 0xe6, 0xb3, 0xdf, 0x4d, 0xa8, 0xc9, 0xf1, 0x31, 0x66, 0x6a, 0x6b, 0xcf, 0xc1, 0x94, 0x0f,
	0x03, 0x39, 0x5c, 0x09, 0xef, 0x5f, 0xcf, 0x44, 0xdb, 0xb9, 0xab, 0x46, 0xef, 0xe4, 0xf5, 0x1f,
	0x7f, 0xff, 0x5a, 0x3a, 0x24, 0x07, 0xea, 0xad, 0x59, 0x7c, 0xd6, 0x97, 0xb2, 0xeb, 0xcf, 0x24,
	0xe1, 0x7b, 0xb0, 0xf4, 0x73, 0x42, 0xd6, 0x22, 0x2e, 0x3c, 0x2f, 0x5b, 0x68, 0xef, 0x2b, 0x9a,
	0xeb, 0x1d, 0x16, 0x68, 0x5a, 0x65, 0xcf, 0x8c, 0x27, 0x12, 0xa9, 0x3d, 0x6d, 0x83, 0x2c, 0x78,
	0xdc, 0xce, 0xc8, 0xb9, 0x9a, 0xca, 0x91, 0xda, 0x3f, 0x36, 0xc8, 0x82, 0x9f, 0xec, 0x8c, 0x8c,
	0xd5, 0x94, 0x44, 0x8e, 0xa0, 0xac, 0x1c, 0x86, 0x1c, 0x6d, 0x46, 0x37, 0x86, 0xb3, 0x05, 0xf8,
	0x40, 0x01, 0xef, 0x79, 0xa4, 0x00, 0xcc, 0xe4, 0x90, 0xe4, 0x3d, 0x87, 0xfd, 0x21, 0x0a, 0x42,
	0x56, 0x73, 0x1b, 0x3b, 0xda, 0xc2, 0x3a, 0x55, 0xac, 0x63, 0xcf, 0x29, 0xb0, 0xa6, 0x28, 0x24,
	0x29, 0x02, 0x7b, 0x2d, 0x13, 0xe2, 0x6e, 0x78, 0x45, 0xc3, 0x6a, 0x9f, 0x6c, 0xa9, 0xe4, 0xf8,
	0x87, 0x0a, 0x7f, 0xea, 0xbd, 0x77, 0x07, 0xcf, 0xb4, 0xc4, 0x9e, 0x19, 0x4f, 0x26, 0x96, 0xfa,
	0x97, 0xf1, 0xf9, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xe2, 0x9c, 0xe1, 0x82, 0x95, 0x08, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	List(ctx context.Context, in *UserListRequest, opts ...grpc.CallOption) (*UserResponse, error)
	Create(ctx context.Context, in *UserCreateRequest, opts ...grpc.CallOption) (*UserResponse, error)
	Update(ctx context.Context, in *UserUpdateRequest, opts ...grpc.CallOption) (*UserResponse, error)
	Delete(ctx context.Context, in *UserDeleteRequest, opts ...grpc.CallOption) (*UserResponse, error)
	Renew(ctx context.Context, in *UserRenewRequest, opts ...grpc.CallOption) (*UserResponse, error)
	Get(ctx context.Context, in *UserGetRequest, opts ...grpc.CallOption) (*UserResponse, error)
	GenConfig(ctx context.Context, in *UserGenConfigRequest, opts ...grpc.CallOption) (*UserGenConfigResponse, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) List(ctx context.Context, in *UserListRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/pb.UserService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Create(ctx context.Context, in *UserCreateRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/pb.UserService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Update(ctx context.Context, in *UserUpdateRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/pb.UserService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Delete(ctx context.Context, in *UserDeleteRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/pb.UserService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Renew(ctx context.Context, in *UserRenewRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/pb.UserService/Renew", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Get(ctx context.Context, in *UserGetRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/pb.UserService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GenConfig(ctx context.Context, in *UserGenConfigRequest, opts ...grpc.CallOption) (*UserGenConfigResponse, error) {
	out := new(UserGenConfigResponse)
	err := c.cc.Invoke(ctx, "/pb.UserService/GenConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	List(context.Context, *UserListRequest) (*UserResponse, error)
	Create(context.Context, *UserCreateRequest) (*UserResponse, error)
	Update(context.Context, *UserUpdateRequest) (*UserResponse, error)
	Delete(context.Context, *UserDeleteRequest) (*UserResponse, error)
	Renew(context.Context, *UserRenewRequest) (*UserResponse, error)
	Get(context.Context, *UserGetRequest) (*UserResponse, error)
	GenConfig(context.Context, *UserGenConfigRequest) (*UserGenConfigResponse, error)
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).List(ctx, req.(*UserListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Create(ctx, req.(*UserCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Update(ctx, req.(*UserUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Delete(ctx, req.(*UserDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Renew_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRenewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Renew(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/Renew",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Renew(ctx, req.(*UserRenewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Get(ctx, req.(*UserGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GenConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserGenConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GenConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/GenConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GenConfig(ctx, req.(*UserGenConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _UserService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _UserService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _UserService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _UserService_Delete_Handler,
		},
		{
			MethodName: "Renew",
			Handler:    _UserService_Renew_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _UserService_Get_Handler,
		},
		{
			MethodName: "GenConfig",
			Handler:    _UserService_GenConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
