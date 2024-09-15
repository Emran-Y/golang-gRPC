package main

import (
	"log"
	"time"

	pb "github.com/emran-y/golang-grpc/proto"
)

func (s *helloServer) SayHelloServerStreaming(req *pb.NameList, stream pb.GreetingService_SayHelloServerStreamingServer) error {
	names := req.Name

	log.Println(names)

	for _, name := range names {
		res := &pb.HelloResponse{
			Message: "Hello " + name,
		}

		if err := stream.Send(res); err != nil {
			return err
		}

		time.Sleep(2 * time.Second)

	}

	return nil
}
