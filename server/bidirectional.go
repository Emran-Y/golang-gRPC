package main

import (
	"fmt"
	"io"

	pb "github.com/emran-y/golang-grpc/proto"
)

func (s *helloServer) SayHelloBiDirectionalStreaming(stream pb.GreetingService_SayHelloBiDirectionalStreamingServer) error {
	for {
		res, err := stream.Recv()

		if err == io.EOF {
			return err
		}

		if err != nil {
			return err
		}

		fmt.Println("res: ", res.GetName())

		if err := stream.Send(&pb.HelloResponse{
			Message: res.GetName(),
		}); err != nil {
			return nil
		}

	}
}
