package main

import (
	"google.golang.org/grpc"
	"github.com/andreiTn/gRPC-stuff"
	"context"
	"log"
	"strconv"
	"github.com/andreiTn/gRPC-stuff/bi-streaming/pb"
	"io"
	"fmt"
)

var logger = gRPC_stuff.Logger{}

func main()  {
	conn, err := grpc.Dial("0.0.0.0:3102", grpc.WithInsecure())

	if err != nil {
		logger.DialError(err)
	}
	defer conn.Close()

	client := bid.NewChatServiceClient(conn)

	chat, err := client.Message(context.Background())

	if err != nil {
		log.Fatalf("Stream error: %v", err)
	}

	waitc := make(chan struct{})

	// send some messages
	go func() {
		for i := 0; i <= 10; i++ {
			chat.Send(&bid.Req{
				Msg: "req: " + strconv.Itoa(i),
			})
		}
		chat.CloseSend()
	}()

	// receive some messages
	go func() {
		for {
			res, err := chat.Recv()

			if err != nil {
				if err == io.EOF {
					goto END
				}
				log.Fatalf("Response err: %v", err)
				close(waitc)
			}
			log.Printf("Response: %s", res.GetMsg())
		}
		END:
			log.Println("Exit")
			close(waitc)

	}()

	// block until everything is done
	<-waitc
	fmt.Println("All done...")
}