package handler

import (
	"bytes"
	"db-experiment/config"
	model "db-experiment/models"
	repository "db-experiment/repositories"
	usecase "db-experiment/usecases"
	"db-experiment/util"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

var result *interface{}

func setupTestTodoV2(b testing.TB) (func(b testing.TB), *httptest.Server) {
	b.Log("Setup benchmarking")
	router := gin.Default()

	configs := config.GetConfig()
	db := config.ConnectDB(configs)

	r := repository.InitializeTodoRepositoryV2(db)
	u := usecase.InitializeTodoUsecaseV2(r)
	h := InitializeTodoHandlerV2(u)

	router.POST("/v2/todos", h.CreateTodo())
	router.GET("/v2/todos", h.GetAllTodos())
	router.GET("/v2/todos/:id", h.GetTodoByID())
	router.GET("/v2/todos/filter", h.FilterTodos())
	router.PUT("/v2/todos", h.UpdateTodo())
	router.DELETE("/v2/todos/:id", h.DeleteTodo())

	testingServer := httptest.NewServer(router)
	cleanupExecutor := util.InitTruncateTableExecutor(db)

	return func(b testing.TB) {
		b.Log("Teardown benchmarking")
		defer testingServer.Close()
		defer cleanupExecutor.TruncateTable([]string{"todos"})
	}, testingServer
}

func runCreateTodoV2(b testing.TB, testingServer *httptest.Server, requestBody []byte) model.Response {
	response, err := http.Post(fmt.Sprintf("%s/v2/todos", testingServer.URL), "application/json", bytes.NewBuffer(requestBody))

	if err != nil {
		b.Fatal()
	}
	defer response.Body.Close()
	body := model.Response{}
	json.NewDecoder(response.Body).Decode(&body)

	return body
}

func BenchmarkTodoHandler_CreateTodoV2(b *testing.B) {
	var r interface{}
	teardown, testingServer := setupTestTodoV2(b)
	defer teardown(b)

	requestBody, err := json.Marshal(map[string]string{
		"username": "username",
		"title":    "title",
		"description": "description",
	})
	if err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		r = runCreateTodoV2(b, testingServer, requestBody)
	}

	result = &r
}