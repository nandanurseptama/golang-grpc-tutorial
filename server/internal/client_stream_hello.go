package internal

import (
	"fmt"
	"io"
	"log/slog"

	"github.com/nandanurseptama/golang-grpc-tutorial/server"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (*grpcServer) ClientStreamHello(stream server.HelloService_ClientStreamHelloServer) error {
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&server.HelloResponse{
				Message: "client force end streaming",
			})
		}

		if err != nil {
			return status.Error(codes.Internal, fmt.Sprintf("failed receive message : %s", err.Error()))
		}

		slog.Info("received message", slog.String("message", req.Message))

		if req.Message == "client_stop_streaming" {
			return stream.SendAndClose(&server.HelloResponse{
				Message: "client request to stop streaming",
			})
		}

	}

}
