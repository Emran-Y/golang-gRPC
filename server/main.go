package main

import (
	"log"
	"net"

	pb "github.com/emran-y/golang-grpc/proto"
	"google.golang.org/grpc"
)

const port = ":8080"

// Define your helloServer struct that implements the server interface
type helloServer struct {
	pb.GreetingServiceServer
}

func main() {
	// Start listening on the defined port
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create a new gRPC server
	grpcServer := grpc.NewServer()

	// Register the GreetingServiceServer with the server instance
	pb.RegisterGreetingServiceServer(grpcServer, &helloServer{})

	log.Printf("server started at %v", lis.Addr())

	// Start serving on the listener
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
