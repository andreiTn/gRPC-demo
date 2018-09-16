package main

import (
	"github.com/andreiTn/gRPC-stuff"
	"net"
	"google.golang.org/grpc"
	"github.com/andreiTn/gRPC-stuff/bi-streaming/pb"
	"io"
	"log"
)

var logger = gRPC_stuff.Logger{}

//Stream(StreamService_StreamServer) error
type Server struct {}

func (*Server) Message(chat bid.ChatService_MessageServer) error  {
	for {
		message, err := chat.Recv()

		if err != nil {
			if err == io.EOF {
				goto END
			}
			log.Fatalf("Chunk error: ^v", err)
		}
		log.Printf("Got message: %s", message.GetMsg())

		err = chat.Send(&bid.Res{
			Msg: "aloha",
		})
		if err != nil {
			log.Fatalf("Send error: ^v", err)
		}
	}

	END:
		log.Println("Exit ...")
		return nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:3102")

	if err != nil {
		logger.ListenErr(err)
	}

	gRPCSrv := grpc.NewServer()

	bid.RegisterChatServiceServer(gRPCSrv, &Server{})
	gRPCSrv.Serve(lis)

}