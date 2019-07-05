package api

import (
	"net"

	"github.com/cad/ovpm/api/pb"
	"github.com/sirupsen/logrus"
	"github.com/vishvananda/netlink"
	"golang.org/x/net/context"
)

// RouteService ...
type RouteService struct{}

// List ...
func (s *RouteService) List(ctx context.Context, req *pb.RouteListRequest) (*pb.RouteListResponse, error) {
	logrus.Debug("rpc call: route list")

	// get loopback interface
	link, err := netlink.LinkByName("tun0")
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	// bring the interface up
	if err := netlink.LinkSetUp(link); err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	rts, err := netlink.RouteList(link, netlink.FAMILY_V4)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	var routes []*pb.Route
	for _, rt := range rts {
		route := &pb.Route{
			Cidr: rt.Dst.String(),
			Via:  rt.Gw.String(),
			Dev:  "tun0",
		}
		routes = append(routes, route)
	}

	return &pb.RouteListResponse{
		Routes: routes,
	}, nil
}

// Create ...
func (s *RouteService) Create(ctx context.Context, req *pb.RouteAddRequest) (*pb.RouteAddResponse, error) {
	logrus.Debugf("rpc call: route create: %s", req.Cidr)

	if req.Dev == "" {
		req.Dev = "tun0"
	}

	// get loopback interface
	link, err := netlink.LinkByName(req.Dev)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	// bring the interface up
	if err := netlink.LinkSetUp(link); err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	// Add route
	_, dst, err := net.ParseCIDR(req.Cidr)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}
	gw := net.ParseIP(req.Via)
	route := netlink.Route{LinkIndex: link.Attrs().Index, Dst: dst, Gw: gw}
	if err := netlink.RouteAdd(&route); err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	return &pb.RouteAddResponse{
		Route: &pb.Route{
			Cidr: req.Cidr,
			Via:  req.Via,
			Dev:  req.Dev,
		},
	}, nil
}

// Delete ...
func (s *RouteService) Delete(ctx context.Context, req *pb.RouteDeleteRequest) (*pb.RouteDeleteResponse, error) {
	logrus.Debugf("rpc call: route delete: %s", req.Cidr)

	if req.Dev == "" {
		req.Dev = "tun0"
	}

	// get loopback interface
	link, err := netlink.LinkByName(req.Dev)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	// bring the interface up
	if err := netlink.LinkSetUp(link); err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	// Delete route
	_, dst, err := net.ParseCIDR(req.Cidr)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}
	route := netlink.Route{LinkIndex: link.Attrs().Index, Dst: dst}
	if err := netlink.RouteDel(&route); err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	return &pb.RouteDeleteResponse{}, nil
}
