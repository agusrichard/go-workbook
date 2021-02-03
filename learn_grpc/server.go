package main

import (
	"fmt"
	"learn-grpc/chat"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Go gRPC Beginners Tutorial")

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %s", err)
	}

	s := chat.Server{}

	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to server gRPC server over port 9000: %v", err)
	}

}
