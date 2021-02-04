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
	log.Println("Create todo")
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
	log.Println("Get todo")
	todos, err := s.todoUsecase.GetTodos(request.UserID)
	if err != nil {
		log.Println("Error create todo", err)
		return &GetTodosResponse{Success: false, Message: fmt.Sprintf("%v", err)}, err
	}
	return &GetTodosResponse{Success: true, Message: "Success to get todos", Data: todos}, nil
}

func (s *Server) UpdateTodo(ctx context.Context, request *UpdateTodoRequest) (*UpdateTodoResponse, error) {
	log.Println("Update todo")
	todo := models.Todo{
		Title:       request.Title,
		Description: request.Description,
		ID:          request.Id,
	}
	_, err := s.todoUsecase.UpdateTodo(todo)
	if err != nil {
		log.Println("Error update todo", err)
		return &UpdateTodoResponse{Success: false, Message: "Failed to update"}, err
	}

	return &UpdateTodoResponse{Success: true, Message: "Success to update todo"}, err
}

func (s *Server) DeleteTodo(ctx context.Context, request *DeleteTodoRequest) (*DeleteTodoResponse, error) {
	log.Println("Delete todo")
	_, err := s.todoUsecase.DeleteTodo(request.Id)
	if err != nil {
		log.Println("Error update todo", err)
		return &DeleteTodoResponse{Success: false, Message: "Failed to delete"}, err
	}

	return &DeleteTodoResponse{Success: true, Message: "Success to delete todo"}, err
}
