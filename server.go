package main

import (
	"github.com/artemxgod/grpc-server-go/hello"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	// listening to server
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Server has failed: %+v ", err)
	}

	// creating grpc server
	grpcServer := grpc.NewServer()
	// handler that we implemented
	s := hello.Server{}

	// register service
	hello.RegisterHelloServiceServer(grpcServer, &s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}