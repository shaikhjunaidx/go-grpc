package main

import (
	"context"
	"log"
	"time"

	"github.com/shaikhjunaidx/go-grpc/proto"
)

func callSayHelloClientStream(client proto.GreetServiceClient, names *proto.NamesList) {
	log.Printf("Client streaming started...")
	stream, err := client.SayHelloClientStreaming(context.Background())

	if err != nil {
		log.Fatalf("Could not send names: %v", names)
	}

	for _, name := range names.Names {
		request := &proto.HelloRequest{Name: name}

		if err := stream.Send(request); err != nil {
			log.Fatalf("Error while sending: %v", err)
		}

		log.Printf("Sent the request with name: %s", name)
		time.Sleep(2 * time.Second)
	}

	response, err := stream.CloseAndRecv()
	log.Printf("Client streaming finished")

	if err != nil {
		log.Fatalf("Error while `Close and Receive`: %v", err)
	}

	log.Printf("Received messages: %v", response.Messages)
}