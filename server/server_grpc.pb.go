// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.0
// source: server/server.proto

package server

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// HelloServiceClient is the client API for HelloService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HelloServiceClient interface {
	// Unary RPC Request
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
	// Server streaming RPC
	ServerStreamHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (HelloService_ServerStreamHelloClient, error)
	// Client streaming RPC
	ClientStreamHello(ctx context.Context, opts ...grpc.CallOption) (HelloService_ClientStreamHelloClient, error)
	// Bidirectional Streaming RPC
	StreamHello(ctx context.Context, opts ...grpc.CallOption) (HelloService_StreamHelloClient, error)
}

type helloServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHelloServiceClient(cc grpc.ClientConnInterface) HelloServiceClient {
	return &helloServiceClient{cc}
}

func (c *helloServiceClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/server.HelloService/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloServiceClient) ServerStreamHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (HelloService_ServerStreamHelloClient, error) {
	stream, err := c.cc.NewStream(ctx, &HelloService_ServiceDesc.Streams[0], "/server.HelloService/ServerStreamHello", opts...)
	if err != nil {
		return nil, err
	}
	x := &helloServiceServerStreamHelloClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type HelloService_ServerStreamHelloClient interface {
	Recv() (*HelloResponse, error)
	grpc.ClientStream
}

type helloServiceServerStreamHelloClient struct {
	grpc.ClientStream
}

func (x *helloServiceServerStreamHelloClient) Recv() (*HelloResponse, error) {
	m := new(HelloResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *helloServiceClient) ClientStreamHello(ctx context.Context, opts ...grpc.CallOption) (HelloService_ClientStreamHelloClient, error) {
	stream, err := c.cc.NewStream(ctx, &HelloService_ServiceDesc.Streams[1], "/server.HelloService/ClientStreamHello", opts...)
	if err != nil {
		return nil, err
	}
	x := &helloServiceClientStreamHelloClient{stream}
	return x, nil
}

type HelloService_ClientStreamHelloClient interface {
	Send(*HelloRequest) error
	CloseAndRecv() (*HelloResponse, error)
	grpc.ClientStream
}

type helloServiceClientStreamHelloClient struct {
	grpc.ClientStream
}

func (x *helloServiceClientStreamHelloClient) Send(m *HelloRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *helloServiceClientStreamHelloClient) CloseAndRecv() (*HelloResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(HelloResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *helloServiceClient) StreamHello(ctx context.Context, opts ...grpc.CallOption) (HelloService_StreamHelloClient, error) {
	stream, err := c.cc.NewStream(ctx, &HelloService_ServiceDesc.Streams[2], "/server.HelloService/StreamHello", opts...)
	if err != nil {
		return nil, err
	}
	x := &helloServiceStreamHelloClient{stream}
	return x, nil
}

type HelloService_StreamHelloClient interface {
	Send(*HelloRequest) error
	Recv() (*HelloResponse, error)
	grpc.ClientStream
}

type helloServiceStreamHelloClient struct {
	grpc.ClientStream
}

func (x *helloServiceStreamHelloClient) Send(m *HelloRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *helloServiceStreamHelloClient) Recv() (*HelloResponse, error) {
	m := new(HelloResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// HelloServiceServer is the server API for HelloService service.
// All implementations must embed UnimplementedHelloServiceServer
// for forward compatibility
type HelloServiceServer interface {
	// Unary RPC Request
	SayHello(context.Context, *HelloRequest) (*HelloResponse, error)
	// Server streaming RPC
	ServerStreamHello(*HelloRequest, HelloService_ServerStreamHelloServer) error
	// Client streaming RPC
	ClientStreamHello(HelloService_ClientStreamHelloServer) error
	// Bidirectional Streaming RPC
	StreamHello(HelloService_StreamHelloServer) error
	mustEmbedUnimplementedHelloServiceServer()
}

// UnimplementedHelloServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHelloServiceServer struct {
}

func (UnimplementedHelloServiceServer) SayHello(context.Context, *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedHelloServiceServer) ServerStreamHello(*HelloRequest, HelloService_ServerStreamHelloServer) error {
	return status.Errorf(codes.Unimplemented, "method ServerStreamHello not implemented")
}
func (UnimplementedHelloServiceServer) ClientStreamHello(HelloService_ClientStreamHelloServer) error {
	return status.Errorf(codes.Unimplemented, "method ClientStreamHello not implemented")
}
func (UnimplementedHelloServiceServer) StreamHello(HelloService_StreamHelloServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamHello not implemented")
}
func (UnimplementedHelloServiceServer) mustEmbedUnimplementedHelloServiceServer() {}

// UnsafeHelloServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HelloServiceServer will
// result in compilation errors.
type UnsafeHelloServiceServer interface {
	mustEmbedUnimplementedHelloServiceServer()
}

func RegisterHelloServiceServer(s grpc.ServiceRegistrar, srv HelloServiceServer) {
	s.RegisterService(&HelloService_ServiceDesc, srv)
}

func _HelloService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.HelloService/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServiceServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelloService_ServerStreamHello_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(HelloRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HelloServiceServer).ServerStreamHello(m, &helloServiceServerStreamHelloServer{stream})
}

type HelloService_ServerStreamHelloServer interface {
	Send(*HelloResponse) error
	grpc.ServerStream
}

type helloServiceServerStreamHelloServer struct {
	grpc.ServerStream
}

func (x *helloServiceServerStreamHelloServer) Send(m *HelloResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _HelloService_ClientStreamHello_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HelloServiceServer).ClientStreamHello(&helloServiceClientStreamHelloServer{stream})
}

type HelloService_ClientStreamHelloServer interface {
	SendAndClose(*HelloResponse) error
	Recv() (*HelloRequest, error)
	grpc.ServerStream
}

type helloServiceClientStreamHelloServer struct {
	grpc.ServerStream
}

func (x *helloServiceClientStreamHelloServer) SendAndClose(m *HelloResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *helloServiceClientStreamHelloServer) Recv() (*HelloRequest, error) {
	m := new(HelloRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _HelloService_StreamHello_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HelloServiceServer).StreamHello(&helloServiceStreamHelloServer{stream})
}

type HelloService_StreamHelloServer interface {
	Send(*HelloResponse) error
	Recv() (*HelloRequest, error)
	grpc.ServerStream
}

type helloServiceStreamHelloServer struct {
	grpc.ServerStream
}

func (x *helloServiceStreamHelloServer) Send(m *HelloResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *helloServiceStreamHelloServer) Recv() (*HelloRequest, error) {
	m := new(HelloRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// HelloService_ServiceDesc is the grpc.ServiceDesc for HelloService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HelloService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "server.HelloService",
	HandlerType: (*HelloServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _HelloService_SayHello_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ServerStreamHello",
			Handler:       _HelloService_ServerStreamHello_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ClientStreamHello",
			Handler:       _HelloService_ClientStreamHello_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "StreamHello",
			Handler:       _HelloService_StreamHello_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "server/server.proto",
}
