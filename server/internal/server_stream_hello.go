package internal

import (
	"fmt"
	"time"

	"github.com/nandanurseptama/golang-grpc-tutorial/server"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (*grpcServer) ServerStreamHello(
	req *server.HelloRequest,
	handler server.HelloService_ServerStreamHelloServer,
) error {

	if req.Message == "" {
		return status.Error(codes.InvalidArgument, "`message` cannot empty")
	}

	handler.Send(&server.HelloResponse{
		Message: fmt.Sprintf("Reply : %s", req.Message),
	})

	time.Sleep(time.Second * 3)

	for i := 0; i < 5; i++ {
		handler.Send(&server.HelloResponse{
			Message: fmt.Sprintf(
				"Reply : %s, current time is %s",
				req.Message,
				time.Now().UTC().Format(time.DateTime),
			),
		})
		time.Sleep(time.Second)
	}
	return nil
}
