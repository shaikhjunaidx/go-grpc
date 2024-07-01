package main

import (
	"context"

	"github.com/shaikhjunaidx/go-grpc/proto"
)

func (s *helloServer) SayHello(ctx context.Context, request *proto.NoParam) (*proto.HelloResponse, error) {
	return &proto.HelloResponse{
		Message: "Hello",
	}, nil
}
