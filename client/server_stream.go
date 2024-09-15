package main

import (
	"context"
	"io"
	"log"

	pb "github.com/emran-y/golang-grpc/proto"
)

func callSayHelloServerStream(client pb.GreetingServiceClient, names *pb.NameList) {

	stream, err := client.SayHelloServerStreaming(context.Background(), names)

	if err != nil {
		log.Fatal(err)
	}

	for {
		res, err := stream.Recv()

		if err != nil {
			log.Fatal(err)
		}

		if err == io.EOF {
			break
		}

		log.Println(res.GetMessage())
	}

	log.Printf("Streaming finished")

}
