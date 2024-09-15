package main

import (
	"context"
	"log"
	"time"

	pb "github.com/emran-y/golang-grpc/proto"
)

func callSayHelloClientStreaming(client pb.GreetingServiceClient, names *pb.NameList) {
	log.Println("SayHello Client streaming had started")

	stream, err := client.SayHelloClientStreaming(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	for _, name := range names.Name {

		if err := stream.Send(&pb.HelloRequest{
			Name: name,
		}); err != nil {
			log.Fatal(err)
		}

		log.Printf("Request send with the name %v", name)
		time.Sleep(2 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	log.Printf("Streaming has finished")
	if err != nil {
		log.Printf("Error while receving %v", err)
	}

	log.Printf("Message Received %v", res)

}
