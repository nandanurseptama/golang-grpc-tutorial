package main

import (
	"fmt"
	"log"
	"net"

	"github.com/nandanurseptama/golang-grpc-tutorial/server"
	"github.com/nandanurseptama/golang-grpc-tutorial/server/internal"
	"google.golang.org/grpc"
)

func main() {

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8080))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	grpcServer := internal.NewGrpcServer()

	server.RegisterHelloServiceServer(s, grpcServer)

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
