package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/emran-y/golang-grpc/proto"
)

func callSayHello(client pb.GreetingServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	res, err := client.SayHello(ctx, &pb.NoParam{})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res.Message)
}
