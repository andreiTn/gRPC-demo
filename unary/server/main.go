package main

import (
	"net"
	"log"
	"context"
	"google.golang.org/grpc"
	"github.com/andreiTn/gRPC-stuff/unary/pb"
)

type Server struct{}

// Implement the PPServiceServer interface
// Protobuf: rpc PingPong(PPRequest) returns (PPResponse) {};
func (*Server) PingPong(ctx context.Context, req *pingpong.PPRequest) (*pingpong.PPResponse, error)  {
	log.Printf("Request: ", req.Ping)
	return &pingpong.PPResponse{
		Pong: "pong",
	}, nil
}

func main() {
	// Create a tcp listener
	listener, err := net.Listen("tcp", "0.0.0.0:3100")

	if err != nil {
		log.Fatalf("There was an error: %v", err)
	}
	// Create and register the gRPC server
	server := grpc.NewServer()
	pingpong.RegisterPPServiceServer(server, &Server{})

	err = server.Serve(listener)

	if err != nil {
		log.Fatalf("Could not serve: %v", err)
	}
}