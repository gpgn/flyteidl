// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/service/signal.proto

package service

import (
	context "context"
	fmt "fmt"
	admin "github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/admin"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

func init() { proto.RegisterFile("flyteidl/service/signal.proto", fileDescriptor_ca211d25a1023377) }

var fileDescriptor_ca211d25a1023377 = []byte{
	// 534 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0xc1, 0x6a, 0xd4, 0x50,
	0x14, 0x65, 0x2a, 0x0a, 0x46, 0x44, 0x8d, 0xe2, 0x22, 0x2a, 0xdc, 0xce, 0x4a, 0xa4, 0x4d, 0xaa,
	0x16, 0x65, 0x74, 0xd5, 0x0a, 0x16, 0x8b, 0x20, 0x74, 0x76, 0x22, 0x96, 0x3b, 0x79, 0x37, 0x2f,
	0xd7, 0x26, 0xef, 0xc5, 0xf7, 0x6e, 0xa6, 0x1d, 0xa4, 0x1b, 0x3f, 0x40, 0xa4, 0x7e, 0x89, 0x5f,
	0xa0, 0xff, 0xe0, 0xde, 0x95, 0xe0, 0x6f, 0x48, 0x92, 0xce, 0xe0, 0xd4, 0x19, 0x37, 0x79, 0x90,
	0x73, 0xee, 0x39, 0x37, 0xe7, 0xe4, 0x05, 0x77, 0xb2, 0x62, 0x22, 0xc4, 0xaa, 0x48, 0x3c, 0xb9,
	0x31, 0xa7, 0x94, 0x78, 0xd6, 0x06, 0x8b, 0xb8, 0x72, 0x56, 0x6c, 0x78, 0x75, 0x0a, 0xc7, 0xa7,
	0x70, 0x74, 0x5b, 0x5b, 0xab, 0x0b, 0x4a, 0xb0, 0xe2, 0x04, 0x8d, 0xb1, 0x82, 0xc2, 0xd6, 0xf8,
	0x8e, 0x1f, 0xdd, 0x9a, 0xc9, 0xa1, 0x2a, 0xd9, 0xcc, 0x89, 0x45, 0x6b, 0xed, 0x91, 0xae, 0x6b,
	0x32, 0xeb, 0xfe, 0x10, 0xb5, 0x26, 0x97, 0xd8, 0xaa, 0x1d, 0xff, 0x57, 0xea, 0xc1, 0xef, 0xf3,
	0xc1, 0xe5, 0x61, 0x3b, 0x3e, 0xec, 0xac, 0xc3, 0xcf, 0xbd, 0xe0, 0xda, 0x0e, 0xc9, 0x2b, 0xf7,
	0xcc, 0x11, 0x0a, 0x75, 0x60, 0x78, 0x37, 0x9e, 0xed, 0xd8, 0x7a, 0xc6, 0xdd, 0xfb, 0xbf, 0x88,
	0x7b, 0xf4, 0xbe, 0x26, 0x2f, 0xd1, 0xcd, 0xc5, 0xcc, 0xfe, 0xe0, 0x64, 0xeb, 0x51, 0xb4, 0xb9,
	0x47, 0xe2, 0x98, 0xc6, 0x04, 0x08, 0xdd, 0xd6, 0x6b, 0x90, 0x36, 0xe3, 0x6c, 0x34, 0xb0, 0x00,
	0x67, 0xcd, 0x53, 0x59, 0xf2, 0x60, 0xac, 0x00, 0x1d, 0xb1, 0x97, 0x38, 0xfc, 0xb4, 0x12, 0x5c,
	0x7a, 0xc9, 0x5e, 0x3a, 0x25, 0x1f, 0xae, 0x2e, 0xb6, 0x68, 0x28, 0xd3, 0x2d, 0xa2, 0xe5, 0x94,
	0xfe, 0xb7, 0xde, 0xc9, 0xd6, 0x8b, 0x68, 0xe7, 0x39, 0x49, 0x9a, 0x77, 0x16, 0x8d, 0x7b, 0xb7,
	0x0e, 0x28, 0xca, 0xd8, 0x70, 0x1b, 0x11, 0x94, 0x28, 0x69, 0xde, 0x60, 0x92, 0x13, 0xb0, 0xa9,
	0x6a, 0x99, 0xb2, 0x58, 0x41, 0xc6, 0x85, 0x90, 0xf3, 0xf1, 0xc7, 0x1f, 0xbf, 0xbe, 0xac, 0x94,
	0xe1, 0x41, 0xdb, 0xd5, 0xf8, 0xfe, 0x69, 0x19, 0x3e, 0xf9, 0x70, 0x68, 0xdd, 0x41, 0x56, 0xd8,
	0xc3, 0x7d, 0x3a, 0xa2, 0xb4, 0x6e, 0x14, 0xf7, 0x59, 0x35, 0xb1, 0xbf, 0xa3, 0x54, 0x8e, 0x97,
	0xe1, 0xca, 0x96, 0xc8, 0x66, 0x29, 0x6c, 0xb0, 0xa4, 0xe3, 0xf0, 0xeb, 0x4a, 0x70, 0x71, 0x48,
	0xa7, 0x79, 0x84, 0xb0, 0xf8, 0x5b, 0x87, 0x34, 0x4b, 0x63, 0xf5, 0x3f, 0x0c, 0x5f, 0x59, 0xe3,
	0xa9, 0xff, 0xb3, 0x77, 0xb2, 0xf5, 0xbd, 0x17, 0x5d, 0x1f, 0x92, 0xcc, 0xba, 0x81, 0x31, 0x16,
	0x35, 0xc5, 0xbb, 0xdb, 0xc1, 0xb9, 0xcd, 0x8d, 0x8d, 0xf0, 0x69, 0x30, 0xd8, 0x23, 0xa9, 0x9d,
	0x21, 0x05, 0x99, 0x75, 0x30, 0x42, 0x05, 0xae, 0x73, 0x00, 0xc9, 0x51, 0xa0, 0xc4, 0x09, 0xe4,
	0x38, 0x26, 0xc8, 0x90, 0x0b, 0x52, 0xcd, 0x38, 0xab, 0xf6, 0x77, 0x8b, 0x77, 0xa9, 0xd1, 0x18,
	0x84, 0x6f, 0x83, 0x37, 0x73, 0x1a, 0x38, 0xaf, 0xe0, 0x28, 0x23, 0x47, 0x26, 0x25, 0x0f, 0x68,
	0x80, 0x15, 0x19, 0xe1, 0x14, 0x0b, 0x68, 0x4e, 0x99, 0x74, 0x9c, 0x1c, 0x3d, 0x60, 0xe1, 0x08,
	0xd5, 0x04, 0x46, 0x44, 0x06, 0x1c, 0x69, 0xf6, 0x42, 0x8e, 0x54, 0x57, 0xc6, 0x8d, 0xfe, 0x95,
	0x33, 0x65, 0x3c, 0xe9, 0xdd, 0xdb, 0x1e, 0xbc, 0x7e, 0xac, 0x59, 0xf2, 0x7a, 0x14, 0xa7, 0xb6,
	0x4c, 0xda, 0x3c, 0xac, 0xd3, 0xc9, 0xec, 0x2a, 0x69, 0x32, 0x49, 0x35, 0x5a, 0xd7, 0x36, 0x39,
	0x7b, 0x59, 0x47, 0x17, 0xda, 0xbb, 0xf2, 0xf0, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa7, 0xc1,
	0xaa, 0x6b, 0xc7, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SignalServiceClient is the client API for SignalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SignalServiceClient interface {
	// Fetches or creates a :ref:`ref_flyteidl.admin.Signal`.
	GetOrCreateSignal(ctx context.Context, in *admin.SignalGetOrCreateRequest, opts ...grpc.CallOption) (*admin.Signal, error)
	// Fetch a list of :ref:`ref_flyteidl.admin.Signal` definitions.
	ListSignals(ctx context.Context, in *admin.SignalListRequest, opts ...grpc.CallOption) (*admin.SignalList, error)
	// Sets the value on a :ref:`ref_flyteidl.admin.Signal` definition
	SetSignal(ctx context.Context, in *admin.SignalSetRequest, opts ...grpc.CallOption) (*admin.SignalSetResponse, error)
}

type signalServiceClient struct {
	cc *grpc.ClientConn
}

func NewSignalServiceClient(cc *grpc.ClientConn) SignalServiceClient {
	return &signalServiceClient{cc}
}

func (c *signalServiceClient) GetOrCreateSignal(ctx context.Context, in *admin.SignalGetOrCreateRequest, opts ...grpc.CallOption) (*admin.Signal, error) {
	out := new(admin.Signal)
	err := c.cc.Invoke(ctx, "/flyteidl.service.SignalService/GetOrCreateSignal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signalServiceClient) ListSignals(ctx context.Context, in *admin.SignalListRequest, opts ...grpc.CallOption) (*admin.SignalList, error) {
	out := new(admin.SignalList)
	err := c.cc.Invoke(ctx, "/flyteidl.service.SignalService/ListSignals", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signalServiceClient) SetSignal(ctx context.Context, in *admin.SignalSetRequest, opts ...grpc.CallOption) (*admin.SignalSetResponse, error) {
	out := new(admin.SignalSetResponse)
	err := c.cc.Invoke(ctx, "/flyteidl.service.SignalService/SetSignal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SignalServiceServer is the server API for SignalService service.
type SignalServiceServer interface {
	// Fetches or creates a :ref:`ref_flyteidl.admin.Signal`.
	GetOrCreateSignal(context.Context, *admin.SignalGetOrCreateRequest) (*admin.Signal, error)
	// Fetch a list of :ref:`ref_flyteidl.admin.Signal` definitions.
	ListSignals(context.Context, *admin.SignalListRequest) (*admin.SignalList, error)
	// Sets the value on a :ref:`ref_flyteidl.admin.Signal` definition
	SetSignal(context.Context, *admin.SignalSetRequest) (*admin.SignalSetResponse, error)
}

// UnimplementedSignalServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSignalServiceServer struct {
}

func (*UnimplementedSignalServiceServer) GetOrCreateSignal(ctx context.Context, req *admin.SignalGetOrCreateRequest) (*admin.Signal, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrCreateSignal not implemented")
}
func (*UnimplementedSignalServiceServer) ListSignals(ctx context.Context, req *admin.SignalListRequest) (*admin.SignalList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSignals not implemented")
}
func (*UnimplementedSignalServiceServer) SetSignal(ctx context.Context, req *admin.SignalSetRequest) (*admin.SignalSetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetSignal not implemented")
}

func RegisterSignalServiceServer(s *grpc.Server, srv SignalServiceServer) {
	s.RegisterService(&_SignalService_serviceDesc, srv)
}

func _SignalService_GetOrCreateSignal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(admin.SignalGetOrCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignalServiceServer).GetOrCreateSignal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flyteidl.service.SignalService/GetOrCreateSignal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignalServiceServer).GetOrCreateSignal(ctx, req.(*admin.SignalGetOrCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SignalService_ListSignals_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(admin.SignalListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignalServiceServer).ListSignals(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flyteidl.service.SignalService/ListSignals",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignalServiceServer).ListSignals(ctx, req.(*admin.SignalListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SignalService_SetSignal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(admin.SignalSetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignalServiceServer).SetSignal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flyteidl.service.SignalService/SetSignal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignalServiceServer).SetSignal(ctx, req.(*admin.SignalSetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SignalService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "flyteidl.service.SignalService",
	HandlerType: (*SignalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOrCreateSignal",
			Handler:    _SignalService_GetOrCreateSignal_Handler,
		},
		{
			MethodName: "ListSignals",
			Handler:    _SignalService_ListSignals_Handler,
		},
		{
			MethodName: "SetSignal",
			Handler:    _SignalService_SetSignal_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "flyteidl/service/signal.proto",
}