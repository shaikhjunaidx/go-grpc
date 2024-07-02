package main

import (
	"log"
	"time"

	"github.com/shaikhjunaidx/go-grpc/proto"
)

func (s *helloServer) SayHelloServerStreaming(request *proto.NamesList, stream proto.GreetService_SayHelloServerStreamingServer) error {
	log.Printf("Got a request with these names: %v", request.Names)

	for _, name := range request.Names {
		response := &proto.HelloResponse{
			Message: "Hello " + name,
		}
		if err := stream.Send(response); err != nil {
			return err
		}
		time.Sleep(2 * time.Second)
	}
	return nil
}
