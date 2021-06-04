package handler

import (
	"db-experiment/config"
	repository "db-experiment/repositories"
	usecase "db-experiment/usecases"
	"db-experiment/util"
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"testing"
)

func useless(n int) int {
	return n
}


func setup(b *testing.B) func(b *testing.B) {
	router := gin.Default()

	configs := config.GetConfig()
	db := config.ConnectDB(configs)

	r := repository.InitializeTodoRepository(db)
	u := usecase.InitializeTodoUsecase(r)
	h := InitializeTodoHandler(u)

	router.POST("", h.CreateTodo())
	router.GET("", h.GetAllTodos())
	router.GET("/:id", h.GetTodoByID())
	router.GET("/filter", h.FilterTodos())
	router.PUT("", h.UpdateTodo())
	router.DELETE("/:id", h.DeleteTodo())

	testingServer := httptest.NewServer(router)
	cleanupExecutor := util.InitTruncateTableExecutor(db)

	b.Log("Setup benchmarking")
	return func(b *testing.B) {
		b.Log("Teardown benchmarking")
		defer testingServer.Close()
		defer cleanupExecutor.TruncateTable([]string{"todos"})
	}
}

func BenchmarkTodoHandler_CreateTodo(b *testing.B) {
	teardown := setup(b)
	defer teardown(b)
	for n := 0; n < b.N; n++ {
		useless(1)
	}
}