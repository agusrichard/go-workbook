package handler

import (
	model "db-experiment/models"
	usecase "db-experiment/usecases"
	"github.com/gin-gonic/gin"
	"net/http"
)

type todoHandler struct {
	todoUsecase usecase.TodoUsecase
}

type TodoHandler interface {
	CreateTodo() gin.HandlerFunc
}

func InitializeTodoHandler(u usecase.TodoUsecase) TodoHandler {
	return &todoHandler{
		todoUsecase: u,
	}
}

func (h *todoHandler) CreateTodo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var err error
		var todo model.Todo

		if err = ctx.ShouldBindJSON(&todo); err != nil {
			ctx.JSON(http.StatusBadRequest, model.Response{
				Success: false,
				Message: "failed to parse todo",
			})
			return
		}

		err = h.todoUsecase.CreateTodo(&todo)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, model.Response{
				Success: false,
				Message: err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, model.Response{
			Success: true,
			Message: "Success to create todo",
		})
	}
}