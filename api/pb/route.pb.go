// Code generated by protoc-gen-go. DO NOT EDIT.
// source: route.proto

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

type RouteAddRequest struct {
	Cidr                 string   `protobuf:"bytes,1,opt,name=cidr,proto3" json:"cidr,omitempty"`
	Via                  string   `protobuf:"bytes,2,opt,name=via,proto3" json:"via,omitempty"`
	Dev                  string   `protobuf:"bytes,3,opt,name=dev,proto3" json:"dev,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RouteAddRequest) Reset()         { *m = RouteAddRequest{} }
func (m *RouteAddRequest) String() string { return proto.CompactTextString(m) }
func (*RouteAddRequest) ProtoMessage()    {}
func (*RouteAddRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0984d49a362b6b9f, []int{0}
}

func (m *RouteAddRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteAddRequest.Unmarshal(m, b)
}
func (m *RouteAddRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteAddRequest.Marshal(b, m, deterministic)
}
func (m *RouteAddRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteAddRequest.Merge(m, src)
}
func (m *RouteAddRequest) XXX_Size() int {
	return xxx_messageInfo_RouteAddRequest.Size(m)
}
func (m *RouteAddRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteAddRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RouteAddRequest proto.InternalMessageInfo

func (m *RouteAddRequest) GetCidr() string {
	if m != nil {
		return m.Cidr
	}
	return ""
}

func (m *RouteAddRequest) GetVia() string {
	if m != nil {
		return m.Via
	}
	return ""
}

func (m *RouteAddRequest) GetDev() string {
	if m != nil {
		return m.Dev
	}
	return ""
}

type RouteListRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RouteListRequest) Reset()         { *m = RouteListRequest{} }
func (m *RouteListRequest) String() string { return proto.CompactTextString(m) }
func (*RouteListRequest) ProtoMessage()    {}
func (*RouteListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0984d49a362b6b9f, []int{1}
}

func (m *RouteListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteListRequest.Unmarshal(m, b)
}
func (m *RouteListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteListRequest.Marshal(b, m, deterministic)
}
func (m *RouteListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteListRequest.Merge(m, src)
}
func (m *RouteListRequest) XXX_Size() int {
	return xxx_messageInfo_RouteListRequest.Size(m)
}
func (m *RouteListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RouteListRequest proto.InternalMessageInfo

type RouteDeleteRequest struct {
	Cidr                 string   `protobuf:"bytes,1,opt,name=cidr,proto3" json:"cidr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RouteDeleteRequest) Reset()         { *m = RouteDeleteRequest{} }
func (m *RouteDeleteRequest) String() string { return proto.CompactTextString(m) }
func (*RouteDeleteRequest) ProtoMessage()    {}
func (*RouteDeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0984d49a362b6b9f, []int{2}
}

func (m *RouteDeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteDeleteRequest.Unmarshal(m, b)
}
func (m *RouteDeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteDeleteRequest.Marshal(b, m, deterministic)
}
func (m *RouteDeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteDeleteRequest.Merge(m, src)
}
func (m *RouteDeleteRequest) XXX_Size() int {
	return xxx_messageInfo_RouteDeleteRequest.Size(m)
}
func (m *RouteDeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteDeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RouteDeleteRequest proto.InternalMessageInfo

func (m *RouteDeleteRequest) GetCidr() string {
	if m != nil {
		return m.Cidr
	}
	return ""
}

type Route struct {
	Cidr                 string   `protobuf:"bytes,1,opt,name=cidr,proto3" json:"cidr,omitempty"`
	Via                  string   `protobuf:"bytes,2,opt,name=via,proto3" json:"via,omitempty"`
	Dev                  string   `protobuf:"bytes,3,opt,name=dev,proto3" json:"dev,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Route) Reset()         { *m = Route{} }
func (m *Route) String() string { return proto.CompactTextString(m) }
func (*Route) ProtoMessage()    {}
func (*Route) Descriptor() ([]byte, []int) {
	return fileDescriptor_0984d49a362b6b9f, []int{3}
}

func (m *Route) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Route.Unmarshal(m, b)
}
func (m *Route) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Route.Marshal(b, m, deterministic)
}
func (m *Route) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Route.Merge(m, src)
}
func (m *Route) XXX_Size() int {
	return xxx_messageInfo_Route.Size(m)
}
func (m *Route) XXX_DiscardUnknown() {
	xxx_messageInfo_Route.DiscardUnknown(m)
}

var xxx_messageInfo_Route proto.InternalMessageInfo

func (m *Route) GetCidr() string {
	if m != nil {
		return m.Cidr
	}
	return ""
}

func (m *Route) GetVia() string {
	if m != nil {
		return m.Via
	}
	return ""
}

func (m *Route) GetDev() string {
	if m != nil {
		return m.Dev
	}
	return ""
}

type RouteAddResponse struct {
	Route                *Route   `protobuf:"bytes,1,opt,name=route,proto3" json:"route,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RouteAddResponse) Reset()         { *m = RouteAddResponse{} }
func (m *RouteAddResponse) String() string { return proto.CompactTextString(m) }
func (*RouteAddResponse) ProtoMessage()    {}
func (*RouteAddResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0984d49a362b6b9f, []int{4}
}

func (m *RouteAddResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteAddResponse.Unmarshal(m, b)
}
func (m *RouteAddResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteAddResponse.Marshal(b, m, deterministic)
}
func (m *RouteAddResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteAddResponse.Merge(m, src)
}
func (m *RouteAddResponse) XXX_Size() int {
	return xxx_messageInfo_RouteAddResponse.Size(m)
}
func (m *RouteAddResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteAddResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RouteAddResponse proto.InternalMessageInfo

func (m *RouteAddResponse) GetRoute() *Route {
	if m != nil {
		return m.Route
	}
	return nil
}

type RouteListResponse struct {
	Routes               []*Route `protobuf:"bytes,1,rep,name=routes,proto3" json:"routes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RouteListResponse) Reset()         { *m = RouteListResponse{} }
func (m *RouteListResponse) String() string { return proto.CompactTextString(m) }
func (*RouteListResponse) ProtoMessage()    {}
func (*RouteListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0984d49a362b6b9f, []int{5}
}

func (m *RouteListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteListResponse.Unmarshal(m, b)
}
func (m *RouteListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteListResponse.Marshal(b, m, deterministic)
}
func (m *RouteListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteListResponse.Merge(m, src)
}
func (m *RouteListResponse) XXX_Size() int {
	return xxx_messageInfo_RouteListResponse.Size(m)
}
func (m *RouteListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RouteListResponse proto.InternalMessageInfo

func (m *RouteListResponse) GetRoutes() []*Route {
	if m != nil {
		return m.Routes
	}
	return nil
}

type RouteDeleteResponse struct {
	Route                *Route   `protobuf:"bytes,1,opt,name=route,proto3" json:"route,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RouteDeleteResponse) Reset()         { *m = RouteDeleteResponse{} }
func (m *RouteDeleteResponse) String() string { return proto.CompactTextString(m) }
func (*RouteDeleteResponse) ProtoMessage()    {}
func (*RouteDeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0984d49a362b6b9f, []int{6}
}

func (m *RouteDeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteDeleteResponse.Unmarshal(m, b)
}
func (m *RouteDeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteDeleteResponse.Marshal(b, m, deterministic)
}
func (m *RouteDeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteDeleteResponse.Merge(m, src)
}
func (m *RouteDeleteResponse) XXX_Size() int {
	return xxx_messageInfo_RouteDeleteResponse.Size(m)
}
func (m *RouteDeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteDeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RouteDeleteResponse proto.InternalMessageInfo

func (m *RouteDeleteResponse) GetRoute() *Route {
	if m != nil {
		return m.Route
	}
	return nil
}

func init() {
	proto.RegisterType((*RouteAddRequest)(nil), "pb.RouteAddRequest")
	proto.RegisterType((*RouteListRequest)(nil), "pb.RouteListRequest")
	proto.RegisterType((*RouteDeleteRequest)(nil), "pb.RouteDeleteRequest")
	proto.RegisterType((*Route)(nil), "pb.Route")
	proto.RegisterType((*RouteAddResponse)(nil), "pb.RouteAddResponse")
	proto.RegisterType((*RouteListResponse)(nil), "pb.RouteListResponse")
	proto.RegisterType((*RouteDeleteResponse)(nil), "pb.RouteDeleteResponse")
}

func init() { proto.RegisterFile("route.proto", fileDescriptor_0984d49a362b6b9f) }

var fileDescriptor_0984d49a362b6b9f = []byte{
	// 341 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xc1, 0x6a, 0xf2, 0x40,
	0x14, 0x85, 0x49, 0xd4, 0x80, 0xd7, 0x1f, 0x7e, 0xbd, 0x5a, 0x9b, 0x06, 0x41, 0x3b, 0x2b, 0x71,
	0x61, 0xa8, 0x82, 0x0b, 0x37, 0xa5, 0xb4, 0x9b, 0x42, 0xa1, 0x34, 0xdd, 0x75, 0x17, 0x9d, 0x8b,
	0x04, 0x24, 0x93, 0x26, 0x63, 0x1e, 0xa0, 0x0f, 0xd0, 0x4d, 0x1f, 0xad, 0xaf, 0xd0, 0x07, 0x29,
	0xb9, 0x09, 0xea, 0xb4, 0x50, 0x4a, 0x77, 0xc3, 0x99, 0x73, 0xbe, 0x93, 0x9b, 0x3b, 0xd0, 0x4a,
	0xd5, 0x4e, 0xd3, 0x34, 0x49, 0x95, 0x56, 0x68, 0x27, 0x2b, 0x6f, 0xb0, 0x51, 0x6a, 0xb3, 0x25,
	0x3f, 0x4c, 0x22, 0x3f, 0x8c, 0x63, 0xa5, 0x43, 0x1d, 0xa9, 0x38, 0x2b, 0x1d, 0xe2, 0x16, 0xfe,
	0x07, 0x45, 0xe0, 0x4a, 0xca, 0x80, 0x9e, 0x77, 0x94, 0x69, 0x44, 0xa8, 0xaf, 0x23, 0x99, 0xba,
	0xd6, 0xc8, 0x1a, 0x37, 0x03, 0x3e, 0x63, 0x1b, 0x6a, 0x79, 0x14, 0xba, 0x36, 0x4b, 0xc5, 0xb1,
	0x50, 0x24, 0xe5, 0x6e, 0xad, 0x54, 0x24, 0xe5, 0x02, 0xa1, 0xcd, 0xa8, 0xbb, 0x28, 0xd3, 0x15,
	0x4b, 0x8c, 0x01, 0x59, 0xbb, 0xa1, 0x2d, 0x69, 0xfa, 0xa1, 0x41, 0x5c, 0x42, 0x83, 0x9d, 0x7f,
	0xae, 0x9f, 0x57, 0xf5, 0x3c, 0x49, 0x96, 0xa8, 0x38, 0x23, 0x1c, 0x42, 0x83, 0x7f, 0x07, 0xc3,
	0x5a, 0xb3, 0xe6, 0x34, 0x59, 0x4d, 0xd9, 0x14, 0x94, 0xba, 0x58, 0x40, 0xe7, 0xe8, 0x9b, 0xab,
	0xd4, 0x39, 0x38, 0x7c, 0x9b, 0xb9, 0xd6, 0xa8, 0x66, 0xc6, 0xaa, 0x0b, 0xb1, 0x80, 0xae, 0x31,
	0xd7, 0x2f, 0xfb, 0x66, 0xaf, 0x36, 0xfc, 0x63, 0xe1, 0x91, 0xd2, 0x3c, 0x5a, 0x13, 0x3e, 0x80,
	0x73, 0x9d, 0x52, 0xa8, 0x09, 0xbb, 0x7b, 0xf3, 0x61, 0x17, 0x5e, 0xcf, 0x14, 0xcb, 0x1a, 0x31,
	0x78, 0x79, 0xff, 0x78, 0xb3, 0xfb, 0xa2, 0xc3, 0x4b, 0xcd, 0x2f, 0x7c, 0x86, 0xfb, 0xa1, 0x94,
	0x4b, 0x6b, 0x82, 0xf7, 0x50, 0x2f, 0xc6, 0xc1, 0x43, 0xf6, 0x68, 0x23, 0xde, 0xc9, 0x17, 0xb5,
	0x42, 0x7a, 0x8c, 0xec, 0x21, 0x9a, 0xc8, 0x6d, 0x01, 0x7a, 0x02, 0xa7, 0x9c, 0x13, 0xfb, 0xfb,
	0xb0, 0xb1, 0x50, 0xef, 0xf4, 0x9b, 0x5e, 0x61, 0x87, 0x8c, 0x3d, 0x13, 0x3d, 0x13, 0x2b, 0xd9,
	0xb5, 0xb4, 0x26, 0x2b, 0x87, 0x9f, 0xe1, 0xfc, 0x33, 0x00, 0x00, 0xff, 0xff, 0xf0, 0x3e, 0xcb,
	0xe2, 0xb7, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RouteServiceClient is the client API for RouteService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RouteServiceClient interface {
	Create(ctx context.Context, in *RouteAddRequest, opts ...grpc.CallOption) (*RouteAddResponse, error)
	List(ctx context.Context, in *RouteListRequest, opts ...grpc.CallOption) (*RouteListResponse, error)
	Delete(ctx context.Context, in *RouteDeleteRequest, opts ...grpc.CallOption) (*RouteDeleteResponse, error)
}

type routeServiceClient struct {
	cc *grpc.ClientConn
}

func NewRouteServiceClient(cc *grpc.ClientConn) RouteServiceClient {
	return &routeServiceClient{cc}
}

func (c *routeServiceClient) Create(ctx context.Context, in *RouteAddRequest, opts ...grpc.CallOption) (*RouteAddResponse, error) {
	out := new(RouteAddResponse)
	err := c.cc.Invoke(ctx, "/pb.RouteService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routeServiceClient) List(ctx context.Context, in *RouteListRequest, opts ...grpc.CallOption) (*RouteListResponse, error) {
	out := new(RouteListResponse)
	err := c.cc.Invoke(ctx, "/pb.RouteService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routeServiceClient) Delete(ctx context.Context, in *RouteDeleteRequest, opts ...grpc.CallOption) (*RouteDeleteResponse, error) {
	out := new(RouteDeleteResponse)
	err := c.cc.Invoke(ctx, "/pb.RouteService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RouteServiceServer is the server API for RouteService service.
type RouteServiceServer interface {
	Create(context.Context, *RouteAddRequest) (*RouteAddResponse, error)
	List(context.Context, *RouteListRequest) (*RouteListResponse, error)
	Delete(context.Context, *RouteDeleteRequest) (*RouteDeleteResponse, error)
}

func RegisterRouteServiceServer(s *grpc.Server, srv RouteServiceServer) {
	s.RegisterService(&_RouteService_serviceDesc, srv)
}

func _RouteService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RouteAddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.RouteService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteServiceServer).Create(ctx, req.(*RouteAddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RouteService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RouteListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.RouteService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteServiceServer).List(ctx, req.(*RouteListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RouteService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RouteDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.RouteService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteServiceServer).Delete(ctx, req.(*RouteDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RouteService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.RouteService",
	HandlerType: (*RouteServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _RouteService_Create_Handler,
		},
		{
			MethodName: "List",
			Handler:    _RouteService_List_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _RouteService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "route.proto",
}
