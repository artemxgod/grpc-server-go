package hello

import (
	"context"
	"log"
)

type Server struct {
	UnimplementedHelloServiceServer
}

func (s *Server) Talk(ctx context.Context, in *Message) (*Message, error) {
	log.Printf("Message received from %s", in.Content)
	return &Message{Content: "Greetings from server"}, nil
}