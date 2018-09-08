package main

import (
	"github.com/andreiTn/gRPC-stuff"
	"net"
	"google.golang.org/grpc"
	"github.com/andreiTn/gRPC-stuff/client_streaming/pb"
	"io"
	"log"
)

var logger = gRPC_stuff.Logger{}

//Stream(StreamService_StreamServer) error
type Server struct {}

func (*Server) Stream(stream clsrv.StreamService_StreamServer) error {
	for {
		res, err := stream.Recv()

		if err != nil {
			if err == io.EOF {
				stream.SendAndClose(&clsrv.Res{
					Msg: "Done... closing stream",
				})
				goto END
			}
			log.Fatalf("Chunk error: %v", err)
		}
		log.Printf("Chunk: %s", res)
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

	clsrv.RegisterStreamServiceServer(gRPCSrv, &Server{})

	gRPCSrv.Serve(lis)

}