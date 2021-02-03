package main

import (
	"fmt"
	"log"
	"net"

	"grpc-auth/auth"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Listen to port 9000")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := auth.Server{}

	grpcServer := grpc.NewServer()

	auth.RegisterAuthServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
