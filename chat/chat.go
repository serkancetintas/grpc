package chat

import (
	"github.com/serkancetintas/grpc/chat/github.com/serkancetintas/grpc/gen"
	"golang.org/x/net/context"
	"log"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, in *gen.Message) (*gen.Message, error) {
	log.Printf("Receive message body from client: %s", in.Body)
	return &gen.Message{Body: "Hello From the Server!"}, nil
}
