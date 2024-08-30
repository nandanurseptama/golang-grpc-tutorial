package internal

import (
	"context"
	"fmt"

	"github.com/nandanurseptama/golang-grpc-tutorial/server"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Say Hello handler
//
// SayHello is Unary RPC
// it's like Request and Response like http/1 request
func (*grpcServer) SayHello(
	ctx context.Context,
	req *server.HelloRequest,
) (*server.HelloResponse, error) {

	// handling when username empty
	if req.Username == "" {
		return nil, status.Error(codes.InvalidArgument, "`username` required")
	}

	return &server.HelloResponse{
		Message: fmt.Sprintf("Hello %s", req.Username),
	}, nil
}
