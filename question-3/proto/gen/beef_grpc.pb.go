// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: proto/beef.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Beef_BeefSummary_FullMethodName = "/beef.Beef/BeefSummary"
)

// BeefClient is the client API for Beef service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BeefClient interface {
	BeefSummary(ctx context.Context, in *BeefSummaryRequest, opts ...grpc.CallOption) (*BeefSummaryResponse, error)
}

type beefClient struct {
	cc grpc.ClientConnInterface
}

func NewBeefClient(cc grpc.ClientConnInterface) BeefClient {
	return &beefClient{cc}
}

func (c *beefClient) BeefSummary(ctx context.Context, in *BeefSummaryRequest, opts ...grpc.CallOption) (*BeefSummaryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BeefSummaryResponse)
	err := c.cc.Invoke(ctx, Beef_BeefSummary_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BeefServer is the server API for Beef service.
// All implementations must embed UnimplementedBeefServer
// for forward compatibility.
type BeefServer interface {
	BeefSummary(context.Context, *BeefSummaryRequest) (*BeefSummaryResponse, error)
	mustEmbedUnimplementedBeefServer()
}

// UnimplementedBeefServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBeefServer struct{}

func (UnimplementedBeefServer) BeefSummary(context.Context, *BeefSummaryRequest) (*BeefSummaryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BeefSummary not implemented")
}
func (UnimplementedBeefServer) mustEmbedUnimplementedBeefServer() {}
func (UnimplementedBeefServer) testEmbeddedByValue()              {}

// UnsafeBeefServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BeefServer will
// result in compilation errors.
type UnsafeBeefServer interface {
	mustEmbedUnimplementedBeefServer()
}

func RegisterBeefServer(s grpc.ServiceRegistrar, srv BeefServer) {
	// If the following call pancis, it indicates UnimplementedBeefServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Beef_ServiceDesc, srv)
}

func _Beef_BeefSummary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BeefSummaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BeefServer).BeefSummary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Beef_BeefSummary_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BeefServer).BeefSummary(ctx, req.(*BeefSummaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Beef_ServiceDesc is the grpc.ServiceDesc for Beef service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Beef_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "beef.Beef",
	HandlerType: (*BeefServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BeefSummary",
			Handler:    _Beef_BeefSummary_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/beef.proto",
}
