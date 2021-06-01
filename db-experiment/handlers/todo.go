package handler

import usecase "db-experiment/usecases"

type todoHandler struct {
	usecase usecase.TodoUsecase
}

type TodoHandler interface {

}

func InitializeTodoHandler(u usecase.TodoUsecase) TodoHandler {
	return &todoHandler{
		usecase: u,
	}
}