package main

import (
	"io"
	"log"

	"github.com/shaikhjunaidx/go-grpc/proto"
)

func (s *helloServer) SayHelloClientStreaming(stream proto.GreetService_SayHelloClientStreamingServer) error {
	var messages []string

	for {
		request, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&proto.MessageList{Messages: messages})
		}

		if err != nil {
			return err
		}

		log.Printf("Got request with name: %v", request.Name)
		messages = append(messages, "Hello ", request.Name, "\n")
	}
}
