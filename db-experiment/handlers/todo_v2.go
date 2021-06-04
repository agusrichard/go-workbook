package handler

import (
	model "db-experiment/models"
	usecase "db-experiment/usecases"
	"db-experiment/util"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type todoHandlerV2 struct {
	todoUsecase usecase.TodoUsecaseV2
}

type TodoHandlerV2 interface {
	CreateTodo() gin.HandlerFunc
	GetAllTodos() gin.HandlerFunc
	GetTodoByID() gin.HandlerFunc
	FilterTodos() gin.HandlerFunc
	UpdateTodo() gin.HandlerFunc
	DeleteTodo() gin.HandlerFunc
}

func InitializeTodoHandlerV2(u usecase.TodoUsecaseV2) TodoHandlerV2 {
	return &todoHandlerV2{
		todoUsecase: u,
	}
}

func (h *todoHandlerV2) CreateTodo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var err error
		var todo model.TodoShape

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

func (h *todoHandlerV2) GetAllTodos() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		todos, err := h.todoUsecase.GetAllTodos()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, model.Response{
				Success: false,
				Message: err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, model.Response{
			Success: true,
			Message: "Success to get all todos",
			Data: todos,
		})
	}
}

func (h *todoHandlerV2) GetTodoByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		todo, err := h.todoUsecase.GetTodoByID(id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, model.Response{
				Success: false,
				Message: err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, model.Response{
			Success: true,
			Message: "Success to get todo by id",
			Data: todo,
		})
	}
}

func (h *todoHandlerV2) FilterTodos() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var query model.Query
		var filters []model.Filter

		err := ctx.ShouldBind(&query)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, model.Response{
				Success: false,
				Message: "failed to parse filter",
			})
			return
		}

		err = json.Unmarshal([]byte(query.FilterString), &filters)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, model.Response{
				Success: false,
				Message: "failed to parse filter",
			})
			return
		}

		queryFilter, err := util.CreateQueryFilter(&filters, nil)
		fmt.Println("query filter", queryFilter)

		todos, err := h.todoUsecase.FilterTodos(queryFilter, query.Skip, query.Take)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, model.Response{
				Success: false,
				Message: err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, model.Response{
			Success: true,
			Message: "Success to get all todos filtered",
			Data: todos,
		})
	}
}

func (h *todoHandlerV2) UpdateTodo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var err error
		var todo model.TodoShape

		if err = ctx.ShouldBindJSON(&todo); err != nil {
			ctx.JSON(http.StatusBadRequest, model.Response{
				Success: false,
				Message: "failed to parse todo",
			})
			return
		}

		err = h.todoUsecase.UpdateTodo(&todo)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, model.Response{
				Success: false,
				Message: err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, model.Response{
			Success: true,
			Message: "Success to update todo",
		})
	}
}

func (h *todoHandlerV2) DeleteTodo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		err = h.todoUsecase.DeleteTodo(id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, model.Response{
				Success: false,
				Message: err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, model.Response{
			Success: true,
			Message: "Success to delete todo",
		})
	}
}