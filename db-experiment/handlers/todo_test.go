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

var resultv1 *interface{}

type testCase struct{
	name string
	request []byte
	expected struct {
		statusCode int
		response model.Response
	}
}

func setupTestTodoV1(b testing.TB) (func(b testing.TB), *httptest.Server) {
	b.Log("Setup benchmarking")
	router := gin.Default()

	configs := config.GetConfig()
	db := config.ConnectDB(configs)

	r := repository.InitializeTodoRepository(db)
	u := usecase.InitializeTodoUsecase(r)
	h := InitializeTodoHandler(u)

	router.POST("/v1/todos", h.CreateTodo())
	router.GET("/v1/todos", h.GetAllTodos())
	router.GET("/v1/todos/:id", h.GetTodoByID())
	router.GET("/v1/todos/filter", h.FilterTodos())
	router.PUT("/v1/todos", h.UpdateTodo())
	router.DELETE("/v1/todos/:id", h.DeleteTodo())

	testingServer := httptest.NewServer(router)
	cleanupExecutor := util.InitTruncateTableExecutor(db)

	return func(b testing.TB) {
		b.Log("Teardown benchmarking")
		defer testingServer.Close()
		defer cleanupExecutor.TruncateTable([]string{"todos"})
	}, testingServer
}

func TestTodoHandler_CreateTodo(t *testing.T) {
	teardown, testingServer := setupTestTodoV1(t)
	defer teardown(t)

	var cases []testCase

	var validInputPositiveCases []testCase
	for i := 0; i < 100; i++ {
		rb, err := json.Marshal(map[string]interface{}{
			"username": fmt.Sprintf("username%d", i),
			"title":    fmt.Sprintf("title%d", i),
			"description": fmt.Sprintf("description%d", i),
		})
		if err != nil {
			t.Fatal("there shouldn't be any error when marshalling request body")
		}

		validInputPositiveCases = append(validInputPositiveCases, testCase{
			name: fmt.Sprintf("validInputPositiveCases_%d", i),
			request: rb,
			expected: struct {
				statusCode int
				response model.Response
			}{
				statusCode: http.StatusOK,
				response: model.Response{
					Success: true,
					Message: "Success to create todo",
				},
			},
		})
	}

	var emptyDescriptionPositiveCases []testCase
	for i := 0; i < 100; i++ {
		rb, err := json.Marshal(map[string]interface{}{
			"username": fmt.Sprintf("username%d", i),
			"title":    fmt.Sprintf("title%d", i),
		})
		if err != nil {
			t.Fatal("there shouldn't be any error when marshalling request body")
		}

		emptyDescriptionPositiveCases = append(emptyDescriptionPositiveCases, testCase{
			name: fmt.Sprintf("emptyDescriptionPositiveCases_%d", i),
			request: rb,
			expected: struct {
				statusCode int
				response model.Response
			}{
				statusCode: http.StatusOK,
				response: model.Response{
					Success: true,
					Message: "Success to create todo",
				},
			},
		})
	}

	var emptyAllFieldsNegativeCases []testCase
	for i := 0; i < 100; i++ {
		rb, err := json.Marshal(map[string]interface{}{})
		if err != nil {
			t.Fatal("there shouldn't be any error when marshalling request body")
		}

		emptyAllFieldsNegativeCases = append(emptyAllFieldsNegativeCases, testCase{
			name: fmt.Sprintf("emptyAllFieldsNegativeCases_%d", i),
			request: rb,
			expected: struct {
				statusCode int
				response model.Response
			}{
				statusCode: http.StatusInternalServerError,
				response: model.Response{
					Success: false,
				},
			},
		})
	}

	cases = append(cases, validInputPositiveCases...)
	cases = append(cases, emptyDescriptionPositiveCases...)
	cases = append(cases, emptyAllFieldsNegativeCases...)

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			resp, statusCode := runCreateTodoV1(t, testingServer, tc.request)
			if statusCode != tc.expected.statusCode {
				t.Fatalf("expect status code %v got %v", tc.expected.statusCode, statusCode)
			}

			if resp.Success != tc.expected.response.Success {
				t.Fatalf("expect status code %v got %v", tc.expected.statusCode, statusCode)
			}
		})
	}
}

func BenchmarkTodoHandler_CreateTodo(b *testing.B) {
	b.StopTimer () // call the function of the test time count stop pressure

	var r interface{}
	teardown, testingServer := setupTestTodoV1(b)
	defer teardown(b)

	b.StartTimer () // re-start time
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		requestBody, err := json.Marshal(map[string]interface{}{
			"username": fmt.Sprintf("username%d", i),
			"title":    fmt.Sprintf("title%d", i),
			"description": fmt.Sprintf("description%d", i),
		})
		if err != nil {
			b.Fatal(err)
		}
		b.StartTimer()

		r, _ = runCreateTodoV1(b, testingServer, requestBody)
	}

	resultv1 = &r
}

func runCreateTodoV1(b testing.TB, testingServer *httptest.Server, requestBody []byte) (model.Response, int) {
	response, err := http.Post(fmt.Sprintf("%s/v1/todos", testingServer.URL), "application/json", bytes.NewBuffer(requestBody))

	if err != nil {
		b.Fatal("error run create todo v1")
	}
	defer response.Body.Close()
	body := model.Response{}
	json.NewDecoder(response.Body).Decode(&body)

	return body, response.StatusCode
}