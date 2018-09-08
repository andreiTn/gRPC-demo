package main

import (
	"net"
	"log"
	"google.golang.org/grpc"
	"github.com/andreiTn/gRPC-unary/pb"
	"context"
	"fmt"
)

type Server struct{}

func (*Server) PingPong(ctx context.Context, req *pingpong.PPRequest) (*pingpong.PPResponse, error)  {
	fmt.Println("Request: ", req.Ping)
	return &pingpong.PPResponse{
		Pong: "pong",
	}, nil
}

func main() {
	listner, err := net.Listen("tcp", "0.0.0.0:3100")

	if err != nil {
		log.Fatalf("There was an error: %v", err)
	}

	server := grpc.NewServer()
	pingpong.RegisterPPServiceServer(server, &Server{})

	err = server.Serve(listner)

	if err != nil {
		log.Fatalf("Could not serve: %v", err)
	}

}