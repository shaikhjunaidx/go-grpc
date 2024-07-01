package main

import (
	"context"
	"log"
	"time"

	"github.com/shaikhjunaidx/go-grpc/proto"
)

func callSayHello(client proto.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	response, err := client.SayHello(ctx, &proto.NoParam{})

	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}

	log.Printf("%s", response.Message)
}