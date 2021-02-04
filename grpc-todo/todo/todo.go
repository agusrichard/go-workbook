package todo

import (
	context "context"
	"fmt"
	"grpc-todo/models"
	"grpc-todo/usecase"
	"log"
)

type Server struct {
	todoUsecase usecase.TodoUsecase
}

func InitServer(todoUsecase usecase.TodoUsecase) Server {
	return Server{
		todoUsecase,
	}
}

func (s *Server) CreateTodo(ctx context.Context, request *CreateTodoRequest) (*CreateTodoResponse, error) {
	todo := models.Todo{
		Title:       request.Title,
		Description: request.Description,
		UserID:      request.UserID,
	}
	_, err := s.todoUsecase.CreateTodo(todo)
	if err != nil {
		log.Println("Error create todo", err)
		return &CreateTodoResponse{Success: false, Message: fmt.Sprintf("%v", err)}, err
	}
	return &CreateTodoResponse{Success: true, Message: "Success to create todo"}, nil
}

func (s *Server) GetTodos(ctx context.Context, request *GetTodosRequest) (*GetTodosResponse, error) {
	return &GetTodosResponse{Success: false, Message: "GetTodos", Data: ""}, nil
}
