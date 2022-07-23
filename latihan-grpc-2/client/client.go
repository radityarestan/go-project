package main

import (
	"context"
	"fmt"
	"latihan-grpc-2/chat"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var conn *grpc.ClientConn

	conn, err := grpc.Dial(":9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	for i := 0; i < 10; i++ {
		response, err := c.SayHello(context.Background(), &chat.Message{Body: fmt.Sprintf("Pesan: Hello from client-2 iterasi ke-%d", i)})

		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}

		log.Printf("Response from server: %s", response.Body)
		time.Sleep(time.Second)
	}
}
