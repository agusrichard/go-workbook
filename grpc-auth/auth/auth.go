package auth

import (
	"grpc-auth/usecase"
	"log"

	"golang.org/x/net/context"
)

type Server struct {
	userUsecase usecase.UserUsecase
}

func InitServer(userUsecase usecase.UserUsecase) Server {
	return Server{
		userUsecase,
	}
}

func (s *Server) Register(ctx context.Context, request *RegisterRequest) (*RegisterResponse, error) {
	_, err := s.userUsecase.Register(request.Username, request.Password)
	if err != nil {
		log.Println("Error to register user", err)
		return &RegisterResponse{Success: false, Message: "Failed to register"}, err
	}
	return &RegisterResponse{Success: true, Message: "Succeed to register"}, nil
}

func (s *Server) Login(ctx context.Context, request *LoginRequest) (*LoginResponse, error) {
	token, err := s.userUsecase.Login(request.Username, request.Password)
	if err != nil {
		log.Println("Error to register user", err)
		return &LoginResponse{Success: false, Message: "Failed to login", Token: ""}, err
	}
	return &LoginResponse{Success: true, Message: "Login success", Token: token}, nil
}

func (s *Server) ValidateToken(ctx context.Context, request *ValidateTokenRequest) (*ValidateTokenResponse, error) {
	log.Printf("Receive request from client: %v", request)
	return &ValidateTokenResponse{Success: true, Message: "Hello from the server validate token"}, nil
}
