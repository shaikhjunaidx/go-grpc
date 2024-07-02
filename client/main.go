package main

import (
	"log"

	"github.com/shaikhjunaidx/go-grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.NewClient("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Connection not successful: %v", err)
	}

	defer conn.Close()

	client := proto.NewGreetServiceClient(conn)

	names := &proto.NamesList{
		Names: []string{"Alice", "Bob", "Cam"},
	}

	// callSayHello(client)

	// callSayHelloServerStream(client, names)

	callSayHelloClientStream(client, names)
}
