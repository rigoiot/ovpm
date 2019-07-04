package api

import (
	"github.com/cad/ovpm/api/pb"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

// RouteService ...
type RouteService struct{}

// List ...
func (s *RouteService) List(ctx context.Context, req *pb.RouteListRequest) (*pb.RouteListResponse, error) {
	logrus.Debug("rpc call: route list")

	return &pb.RouteListResponse{}, nil
}

// Create ...
func (s *RouteService) Create(ctx context.Context, req *pb.RouteAddRequest) (*pb.RouteAddResponse, error) {
	logrus.Debugf("rpc call: route create: %s", req.Cidr)

	return &pb.RouteAddResponse{}, nil
}

// Delete ...
func (s *RouteService) Delete(ctx context.Context, req *pb.RouteDeleteRequest) (*pb.RouteDeleteResponse, error) {
	logrus.Debugf("rpc call: route delete: %s", req.Cidr)

	return &pb.RouteDeleteResponse{}, nil
}
