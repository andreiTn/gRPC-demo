package main

import (
	"google.golang.org/grpc"
	"github.com/andreiTn/gRPC-stuff"
	"context"
	"github.com/andreiTn/gRPC-stuff/client_streaming/pb"
	"log"
	"strconv"
)

var logger = gRPC_stuff.Logger{}

func main()  {
	conn, err := grpc.Dial("0.0.0.0:3102", grpc.WithInsecure())

	if err != nil {
		logger.DialError(err)
	}
	defer conn.Close()

	client := clsrv.NewStreamServiceClient(conn)

	stream, err := client.Stream(context.Background())

	if err != nil {
		log.Fatalf("Stream error: %v", err)
	}

	for i := 0; i <= 10; i++ {
		stream.Send(&clsrv.Req{
			Msg: "req: " + strconv.Itoa(i),
		})
	}
	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Response err: %v", err)
	}

	log.Printf("Response: %v", res.Msg)
}