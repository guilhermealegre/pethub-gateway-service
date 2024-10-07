// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.3
// source: logging_service_logging.proto

package logging_service_logging

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Logging_Log_FullMethodName = "/logging_service_logging.Logging/Log"
)

// LoggingClient is the client API for Logging service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LoggingClient interface {
	Log(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type loggingClient struct {
	cc grpc.ClientConnInterface
}

func NewLoggingClient(cc grpc.ClientConnInterface) LoggingClient {
	return &loggingClient{cc}
}

func (c *loggingClient) Log(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Logging_Log_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoggingServer is the server API for Logging service.
// All implementations must embed UnimplementedLoggingServer
// for forward compatibility
type LoggingServer interface {
	Log(context.Context, *LogRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedLoggingServer()
}

// UnimplementedLoggingServer must be embedded to have forward compatible implementations.
type UnimplementedLoggingServer struct {
}

func (UnimplementedLoggingServer) Log(context.Context, *LogRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Log not implemented")
}
func (UnimplementedLoggingServer) mustEmbedUnimplementedLoggingServer() {}

// UnsafeLoggingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LoggingServer will
// result in compilation errors.
type UnsafeLoggingServer interface {
	mustEmbedUnimplementedLoggingServer()
}

func RegisterLoggingServer(s grpc.ServiceRegistrar, srv LoggingServer) {
	s.RegisterService(&Logging_ServiceDesc, srv)
}

func _Logging_Log_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoggingServer).Log(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Logging_Log_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoggingServer).Log(ctx, req.(*LogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Logging_ServiceDesc is the grpc.ServiceDesc for Logging service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Logging_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "logging_service_logging.Logging",
	HandlerType: (*LoggingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Log",
			Handler:    _Logging_Log_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "logging_service_logging.proto",
}