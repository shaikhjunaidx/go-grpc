package main

import (
	"context"
	"io"
	"log"
	"time"

	"github.com/shaikhjunaidx/go-grpc/proto"
)

func callHelloBidirectionalStream(client proto.GreetServiceClient, names *proto.NamesList) {
	log.Printf("Bidirectional Streaming has started...")
	stream, err := client.SayHelloBidirectionalStreaming(context.Background())

	if err != nil {
		log.Fatalf("Could not send names: %v", err)
	}

	waitch := make(chan struct{})

	go func() {
		for {
			message, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("Error while streaming to the server: %v", err)
			}
			log.Println(message)
		}
		close(waitch)
	}()

	for _, name := range names.Names {
		request := &proto.HelloRequest{Name: name}

		if err := stream.Send(request); err != nil {
			log.Fatalf("Error while sending the request %v", err)
		}

		time.Sleep(2 * time.Second)
	}

	stream.CloseSend()
	<-waitch
	log.Printf("Bidirectional streaming finished")
}
