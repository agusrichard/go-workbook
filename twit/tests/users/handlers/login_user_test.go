package userHandlerTests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"twit/configs"
	"twit/models/responses"
	"twit/servers"
	"twit/utils"

	"github.com/stretchr/testify/suite"
)

type HandlerLoginUserSuite struct {
	suite.Suite
	testingServer   *httptest.Server
	cleanupExecutor utils.TruncateTableExecutor
}

func (suite *HandlerLoginUserSuite) SetupTest() {
	router := servers.SetupServer()
	testingServer := httptest.NewServer(router)

	suite.testingServer = testingServer

	cleanupExecutor := utils.InitTruncateTableExecutor(configs.DB)
	suite.cleanupExecutor = cleanupExecutor
}

func (suite *HandlerLoginUserSuite) TearDownTest() {
	defer suite.testingServer.Close()
	defer suite.cleanupExecutor.TruncateTable([]string{"users"})
}

func (suite *HandlerLoginUserSuite) TestLoginUserPositive() {
	// Register user first
	requestBody, err := json.Marshal(map[string]string{
		"username": "username",
		"email":    "email",
		"password": "password",
	})
	suite.NoError(err, "There should be no errors when create requestBody")

	response, err := http.Post(fmt.Sprintf("%s/auth/register", suite.testingServer.URL), "application/json", bytes.NewBuffer(requestBody))
	suite.Equal(http.StatusOK, response.StatusCode)

	defer response.Body.Close()
	body := responses.Response{}
	json.NewDecoder(response.Body).Decode(&body)
	suite.Equal("Success to register user", body.Message)
	suite.Equal(true, body.Success)

	// Login User
	requestBody, err = json.Marshal(map[string]string{
		"email":    "email",
		"password": "password",
	})
	suite.NoError(err, "There should be no errors when create requestBody")

	response, err = http.Post(fmt.Sprintf("%s/auth/login", suite.testingServer.URL), "application/json", bytes.NewBuffer(requestBody))
	suite.Equal(http.StatusOK, response.StatusCode)

	defer response.Body.Close()
	body = responses.Response{}
	json.NewDecoder(response.Body).Decode(&body)
	suite.Equal("Success to login", body.Message)
	suite.Equal(true, body.Success)
}

func TestHandlerLoginUserSuite(t *testing.T) {
	suite.Run(t, new(HandlerLoginUserSuite))
}
