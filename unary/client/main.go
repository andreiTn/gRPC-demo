package main

import (
	"google.golang.org/grpc"
	"github.com/andreiTn/gRPC-unary/pb"
	"log"
	"context"
	"fmt"
)

func main() {
	conn, err := grpc.Dial("0.0.0.0:3100", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not dial: %v", err)
	}
	defer conn.Close()

	client := pingpong.NewPPServiceClient(conn)


	response, err := client.PingPong(context.Background(), &pingpong.PPRequest{
		Ping: "ping",
	})

	if err != nil {
		log.Fatalf("Request error: %v", err)
	}
	fmt.Println(response.Pong)
}
