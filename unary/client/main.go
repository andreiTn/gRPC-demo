package main

import (
	"context"
	"github.com/andreiTn/gRPC-stuff/unary/pb"
	"google.golang.org/grpc"
	"log"
	"github.com/andreiTn/gRPC-stuff"
)

var logger = gRPC_stuff.Logger{}

func main() {
	// Connect to a gRPC server
	conn, err := grpc.Dial("0.0.0.0:3100", grpc.WithInsecure())

	if err != nil {
		logger.DialError(err)
	}
	// Close connection at the end
	defer conn.Close()

	// Create a new client
	client := pingpong.NewPPServiceClient(conn)

	sendPing(client)
}

func sendPing(client pingpong.PPServiceClient) {
	// Call the PingPong method defined in server.go
	// PingPong(context.Context, *pingpong.PPRequest) (*pingpong.PPResponse, error)
	response, err := client.PingPong(context.Background(), &pingpong.PPRequest{
		Ping: "ping",
	})

	if err != nil {
		log.Fatalf("Request error: %v", err)
	}
	log.Printf("Response: %v", response.Pong)
}
