package main

import (
	"fmt"
	"log"
	"net"

	"grpc-auth/auth"
	"grpc-auth/config"
	"grpc-auth/repository"
	"grpc-auth/usecase"

	"google.golang.org/grpc"
)

func main() {
	db := config.ConnectDB()
	userRepository := repository.InitUserRepository(db)
	userUsecase := usecase.InitUserUsecase(userRepository)

	s := auth.InitServer(userUsecase)

	grpcServer := grpc.NewServer()

	auth.RegisterAuthServiceServer(grpcServer, &s)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Println("failed to listen: %v", err)
	}
	fmt.Println("Listen to port 9000")

	if err := grpcServer.Serve(lis); err != nil {
		log.Println("failed to serve: %s", err)
	}
}
