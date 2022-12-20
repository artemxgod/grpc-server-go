package main

import (
	"context"
	"log"

	"github.com/artemxgod/grpc-server-go/hello"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	// connecting to server
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	
	// creating a client
	c := hello.NewHelloServiceClient(conn)
	message, err := c.Talk(context.Background(), &hello.Message{Content: "Artem Kain"})
	if err != nil {
		log.Printf("Failed to talk %+v", err)
	}
	log.Println(message.Content)
}