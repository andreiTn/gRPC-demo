package main

import (
	"net"
	"log"
	"google.golang.org/grpc"
	"github.com/andreiTn/gRPC-stuff/server_streaming/pb"
	"strconv"
	"time"
	"github.com/andreiTn/gRPC-stuff"
)

var logger = gRPC_stuff.Logger{}

type Server struct {}


//GetMsg(*MsgReq, MsgService_GetMsgServer) error
func (*Server) GetMsg(req *srvstr.MsgReq, stream srvstr.MsgService_GetMsgServer) (error)  {
	log.Printf("Get rows: %d", req.Rows)

	for i := 0; i <= int(req.Rows); i ++  {
		message := "Row: " + strconv.Itoa(i)

		stream.Send(&srvstr.MsgRes{
			Row: message,
		})
		time.Sleep(500 * time.Millisecond)
	}
	// Return nil to end the stream respnse
	return nil
}
func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:3101")

	if err != nil {
		logger.ListenErr(err)
	}

	grpcSrv := grpc.NewServer()
	srvstr.RegisterMsgServiceServer(grpcSrv, &Server{})

	err = grpcSrv.Serve(listener)

	if err != nil {
		logger.ServeError(err)
	}
}