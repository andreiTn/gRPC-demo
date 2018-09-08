package gRPC_stuff

import (
	"log"
)

type GRPC_Stuff_Logger interface {
	DialError(error)
	ListenErr(err error)
	ServeError(err error)
}

type Logger struct {}

func (*Logger) DialError(err error) {
	log.Fatalf("Could not dial: %v", err)
}
func (*Logger) ListenErr(err error)  {
	log.Fatalf("Listen error: %v", err)
}
func (*Logger) ServeError(err error)  {
	log.Fatalf("gRPC serve error: %v", err)
}
