package chat

import (
	"context"
	"fmt"
	"log"
)

type Server struct{}

func (s Server) SayHello(ctx context.Context, in *Message) (*Message, error) {
	log.Printf("Receive message body from client: %s", in.Body)
	return &Message{Body: fmt.Sprintf("Pesan: %s sudah diterima", in.Body)}, nil
}

func (s Server) mustEmbedUnimplementedChatServiceServer() {}
