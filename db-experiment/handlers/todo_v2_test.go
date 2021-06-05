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
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

var resultv2 *interface{}

func setupTestTodoV2(b testing.TB) (func(b testing.TB), *httptest.Server) {
	b.Log("Setup benchmarking")
	router := gin.New()

	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard

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

func TestTodoHandler_CreateTodoV2(t *testing.T) {
	teardown, testingServer := setupTestTodoV2(t)
	defer teardown(t)

	var cases []testCase

	var validInputPositiveCases []testCase
	for i := 0; i < 10; i++ {
		rb, err := json.Marshal(map[string]interface{}{
			"username":    fmt.Sprintf("username%d", i),
			"title":       fmt.Sprintf("title%d", i),
			"description": fmt.Sprintf("description%d", i),
		})
		if err != nil {
			t.Fatal("there shouldn't be any error when marshalling request body")
		}

		validInputPositiveCases = append(validInputPositiveCases, testCase{
			name:    fmt.Sprintf("validInputPositiveCases_%d", i),
			request: rb,
			expected: struct {
				statusCode int
				response   model.Response
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
	for i := 0; i < 10; i++ {
		rb, err := json.Marshal(map[string]interface{}{
			"username": fmt.Sprintf("username%d", i),
			"title":    fmt.Sprintf("title%d", i),
		})
		if err != nil {
			t.Fatal("there shouldn't be any error when marshalling request body")
		}

		emptyDescriptionPositiveCases = append(emptyDescriptionPositiveCases, testCase{
			name:    fmt.Sprintf("emptyDescriptionPositiveCases_%d", i),
			request: rb,
			expected: struct {
				statusCode int
				response   model.Response
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
	for i := 0; i < 10; i++ {
		rb, err := json.Marshal(map[string]interface{}{})
		if err != nil {
			t.Fatal("there shouldn't be any error when marshalling request body")
		}

		emptyAllFieldsNegativeCases = append(emptyAllFieldsNegativeCases, testCase{
			name:    fmt.Sprintf("emptyAllFieldsNegativeCases_%d", i),
			request: rb,
			expected: struct {
				statusCode int
				response   model.Response
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
			resp, statusCode := runCreateTodoV2(t, testingServer, tc.request)
			if statusCode != tc.expected.statusCode {
				t.Fatalf("expect status code %v got %v", tc.expected.statusCode, statusCode)
			}

			if resp.Success != tc.expected.response.Success {
				t.Fatalf("expect status code %v got %v", tc.expected.statusCode, statusCode)
			}
		})
	}
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
		r, _ = runCreateTodoV2(b, testingServer, requestBody)
	}

	resultv2 = &r
}

func runCreateTodoV2(tb testing.TB, testingServer *httptest.Server, requestBody []byte) (model.Response, int) {
	response, err := http.Post(fmt.Sprintf("%s/v2/todos", testingServer.URL), "application/json", bytes.NewBuffer(requestBody))

	if err != nil {
		tb.Fatal("error run create todo v2")
	}
	defer response.Body.Close()
	body := model.Response{}
	json.NewDecoder(response.Body).Decode(&body)

	return body, response.StatusCode
}

func TestTodoHandler_GetAllTodosV2(t *testing.T) {
	teardown, testingServer := setupTestTodoV2(t)
	defer teardown(t)

	requestBody, err := json.Marshal(map[string]interface{}{
		"username":    "username",
		"title":       "title",
		"description": "description",
	})
	if err != nil {
		t.Fatal(err)
	}

	// Insert 10 data points in table todos
	numOfRecords := 10
	for i := 0; i < numOfRecords; i++ {
		runCreateTodoV2(t, testingServer, requestBody)
	}

	response, statusCode := runGetAllTodosV2(t, testingServer)
	if statusCode != http.StatusOK {
		t.Fatalf("expect http status code 200 but got %d", statusCode)
	}

	if !response.Success {
		t.Fatal("expect to success get all data, got failed")
	}

	if response.Message != "Success to get all todos" {
		t.Fatalf("expect %v as message got %v", "Success to get all todos", response.Message)
	}
}

func BenchmarkTodoHandler_GetAllTodosV2(b *testing.B) {
	b.StopTimer() // call the function of the test time count stop pressure

	var r interface{}
	teardown, testingServer := setupTestTodoV2(b)
	defer teardown(b)

	requestBody, err := json.Marshal(map[string]interface{}{
		"username":    "username",
		"title":       "title",
		"description": "description",
	})
	if err != nil {
		b.Fatal(err)
	}

	// Insert 10 data points in table todos
	numOfRecords := 10
	for i := 0; i < numOfRecords; i++ {
		runCreateTodoV2(b, testingServer, requestBody)
	}

	b.StartTimer() // re-start time
	for i := 0; i < b.N; i++ {
		r, _ = runGetAllTodosV2(b, testingServer)
	}

	resultv2 = &r
}

func runGetAllTodosV2(tb testing.TB, testingServer *httptest.Server) (model.Response, int) {
	response, err := http.Get(fmt.Sprintf("%s/v2/todos", testingServer.URL))

	if err != nil {
		tb.Fatal("error run get todos v2")
	}
	defer response.Body.Close()
	body := model.Response{}
	json.NewDecoder(response.Body).Decode(&body)

	return body, response.StatusCode
}