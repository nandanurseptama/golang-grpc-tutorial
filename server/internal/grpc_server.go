package internal

import "github.com/nandanurseptama/golang-grpc-tutorial/server"

type grpcServer struct {
	server.UnimplementedHelloServiceServer
}

func NewGrpcServer() *grpcServer {
	return &grpcServer{}
}
