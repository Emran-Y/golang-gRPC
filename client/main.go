package main

import (
	"log"

	pb "github.com/emran-y/golang-grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewGreetingServiceClient(conn)

	// callSayHello(client)

	names := &pb.NameList{
		Name: []string{
			"Alice", "Bob", "Chris",
		},
	}

	// callSayHelloServerStream(client, names)

	callSayHelloClientStreaming(client, names)

	defer conn.Close()
}
