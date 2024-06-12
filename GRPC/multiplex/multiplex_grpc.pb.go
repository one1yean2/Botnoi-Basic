// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: multiplex.proto

package multiplex

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	MultiplexService_Process_FullMethodName = "/multiplex.MultiplexService/Process"
)

// MultiplexServiceClient is the client API for MultiplexService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MultiplexServiceClient interface {
	Process(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type multiplexServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMultiplexServiceClient(cc grpc.ClientConnInterface) MultiplexServiceClient {
	return &multiplexServiceClient{cc}
}

func (c *multiplexServiceClient) Process(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, MultiplexService_Process_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MultiplexServiceServer is the server API for MultiplexService service.
// All implementations must embed UnimplementedMultiplexServiceServer
// for forward compatibility
type MultiplexServiceServer interface {
	Process(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedMultiplexServiceServer()
}

// UnimplementedMultiplexServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMultiplexServiceServer struct {
}

func (UnimplementedMultiplexServiceServer) Process(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Process not implemented")
}
func (UnimplementedMultiplexServiceServer) mustEmbedUnimplementedMultiplexServiceServer() {}

// UnsafeMultiplexServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MultiplexServiceServer will
// result in compilation errors.
type UnsafeMultiplexServiceServer interface {
	mustEmbedUnimplementedMultiplexServiceServer()
}

func RegisterMultiplexServiceServer(s grpc.ServiceRegistrar, srv MultiplexServiceServer) {
	s.RegisterService(&MultiplexService_ServiceDesc, srv)
}

func _MultiplexService_Process_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MultiplexServiceServer).Process(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MultiplexService_Process_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MultiplexServiceServer).Process(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// MultiplexService_ServiceDesc is the grpc.ServiceDesc for MultiplexService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MultiplexService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "multiplex.MultiplexService",
	HandlerType: (*MultiplexServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Process",
			Handler:    _MultiplexService_Process_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "multiplex.proto",
}