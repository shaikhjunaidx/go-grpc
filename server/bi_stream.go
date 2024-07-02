package main

import (
	"io"
	"log"

	"github.com/shaikhjunaidx/go-grpc/proto"
)

func (s *helloServer) SayHelloBidirectionalStreaming(stream proto.GreetService_SayHelloBidirectionalStreamingServer) error {
	for {
		request, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		log.Printf("Got request with name: %v", request.Name)

		response := &proto.HelloResponse{
			Message: "Hello " + request.Name,
		}

		if err := stream.Send(response); err != nil {
			return nil
		}
	}
}
