package main

import (
	"context"
	"github.com/andreiTn/gRPC-stuff"
	"github.com/andreiTn/gRPC-stuff/server_streaming/pb"
	"google.golang.org/grpc"
	"io"
	"log"
)

var logger = gRPC_stuff.Logger{}

func main() {
	conn, err := grpc.Dial("0.0.0.0:3101", grpc.WithInsecure())

	if err != nil {
		logger.DialError(err)
	}
	defer conn.Close()

	client := srvstr.NewMsgServiceClient(conn)

	//GetMsg(ctx context.Context, in *MsgReq, opts ...grpc.CallOption) (MsgService_GetMsgClient, error)
	msgStream, err := client.GetMsg(context.Background(), &srvstr.MsgReq{
		Rows: 12,
	})

	if err != nil {
		log.Fatalf("Stream failed: %v", err)
	}

	for {
		res, err := msgStream.Recv()

		if err != nil {
			if err == io.EOF {
				log.Printf("End of stream")
				goto END
			}
			log.Fatalf("Chunk error: %v", err)
		}
		log.Printf("Chunk: %s", res)
	}

END:
	log.Println("Exit...")
}
