package main

import (
	"fmt"
	"io"

	pb "github.com/emran-y/golang-grpc/proto"
)

func (s *helloServer) SayHelloClientStreaming(stream pb.GreetingService_SayHelloClientStreamingServer) error {
	var messages []string

	fmt.Println("Called server")

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		messages = append(messages, "Hello"+req.GetName())
	}

	stream.SendAndClose(&pb.MessageList{Message: messages})

	return nil

}
