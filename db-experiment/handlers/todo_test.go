package handler

import (
	"db-experiment/config"
	model "db-experiment/models"
	"db-experiment/server"
	"db-experiment/util"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"net/http/httptest"
	"testing"
)

func useless(n int) int {
	return n
}

var router *gin.Engine
var testingServer *httptest.Server
var configs *model.Config
var db *sqlx.DB
var cleanupExecutor util.TruncateTableExecutor

func setup(b *testing.B) func(b *testing.B) {
	var router = server.SetupServer()
	var testingServer = httptest.NewServer(router)
	var configs = config.GetConfig()
	var db = config.ConnectDB(configs)
	var cleanupExecutor = util.InitTruncateTableExecutor(db)

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