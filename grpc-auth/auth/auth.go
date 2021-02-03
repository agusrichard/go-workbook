package auth

import (
	"log"

	"golang.org/x/net/context"
)

type Server struct{}

func (s *Server) Register(ctx context.Context, request *RegisterRequest) (*RegisterResponse, error) {
	log.Printf("Receive request from client: %v", request)
	return &RegisterResponse{Success: true, Message: "Hello from the server register"}, nil
}

func (s *Server) Login(ctx context.Context, request *LoginRequest) (*LoginResponse, error) {
	log.Printf("Receive request from client: %v", request)
	return &LoginResponse{Success: true, Message: "Hello from the server login"}, nil
}

func (s *Server) ValidateToken(ctx context.Context, request *ValidateTokenRequest) (*ValidateTokenResponse, error) {
	log.Printf("Receive request from client: %v", request)
	return &ValidateTokenResponse{Success: true, Message: "Hello from the server validate token"}, nil
}
