package main

import (
	"context"
	"github.com/andreiTn/gRPC-stuff/unary/pb"
	"google.golang.org/grpc"
	"log"
	"net"
	"github.com/andreiTn/gRPC-stuff"
)

var logger = gRPC_stuff.Logger{}

type Server struct{}

// Implement the PPServiceServer interface
// Protobuf: rpc PingPong(PPRequest) returns (PPResponse) {};
func (*Server) PingPong(ctx context.Context, req *pingpong.PPRequest) (*pingpong.PPResponse, error) {
	log.Printf("Request: %s", req.Ping)
	return &pingpong.PPResponse{
		Pong: "pong",
	}, nil
}

func main() {
	// Create a tcp listener
	listener, err := net.Listen("tcp", "0.0.0.0:3100")

	if err != nil {
		logger.ListenErr(err)
	}
	// Create and register the gRPC server
	server := grpc.NewServer()
	pingpong.RegisterPPServiceServer(server, &Server{})

	err = server.Serve(listener)

	if err != nil {
		logger.ServeError(err)
	}
}
