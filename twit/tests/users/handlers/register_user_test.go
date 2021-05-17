package userHandlerTests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"twit/configs"
	"twit/models"
	"twit/servers"
	"twit/utils"

	"github.com/stretchr/testify/suite"
)

type HandlerRegisterUserSuite struct {
	suite.Suite
	testingServer   *httptest.Server
	cleanupExecutor utils.TruncateTableExecutor
}

func (suite *HandlerRegisterUserSuite) SetupTest() {
	router := servers.SetupServer()
	testingServer := httptest.NewServer(router)

	suite.testingServer = testingServer

	cleanupExecutor := utils.InitTruncateTableExecutor(configs.DB)
	suite.cleanupExecutor = cleanupExecutor
}

func (suite *HandlerRegisterUserSuite) TearDownTest() {
	defer suite.testingServer.Close()
	defer suite.cleanupExecutor.TruncateTable([]string{"users"})
}

func (suite *HandlerRegisterUserSuite) TestRegisterSingleUserPositive() {
	requestBody, err := json.Marshal(map[string]string{
		"username": "username",
		"email":    "email",
		"password": "password",
	})
	suite.NoError(err, "There should be no errors when create requestBody")

	response, err := http.Post(fmt.Sprintf("%s/auth/register", suite.testingServer.URL), "application/json", bytes.NewBuffer(requestBody))
	suite.Equal(http.StatusOK, response.StatusCode)

	defer response.Body.Close()
	body := models.Response{}
	json.NewDecoder(response.Body).Decode(&body)
	suite.Equal("Success to register user", body.Message)
	suite.Equal(true, body.Success)
}

func (suite *HandlerRegisterUserSuite) TestRegisterSameUserTwiceNegative() {
	// First attempt to register user (Positive -- No Error)
	requestBody, err := json.Marshal(map[string]string{
		"username": "username",
		"email":    "email",
		"password": "password",
	})
	suite.NoError(err, "There should be no errors when create requestBody")

	response, err := http.Post(fmt.Sprintf("%s/auth/register", suite.testingServer.URL), "application/json", bytes.NewBuffer(requestBody))
	suite.Equal(http.StatusOK, response.StatusCode)

	defer response.Body.Close()
	body := models.Response{}
	json.NewDecoder(response.Body).Decode(&body)
	suite.Equal("Success to register user", body.Message)
	suite.Equal(true, body.Success)

	// Second attempt to register the same user (Negative test)
	requestBody, err = json.Marshal(map[string]string{
		"username": "username",
		"email":    "email",
		"password": "password",
	})
	suite.NoError(err, "There should be no errors when create requestBody")

	response, err = http.Post(fmt.Sprintf("%s/auth/register", suite.testingServer.URL), "application/json", bytes.NewBuffer(requestBody))
	suite.Equal(http.StatusBadRequest, response.StatusCode)

	defer response.Body.Close()
	body = models.Response{}
	json.NewDecoder(response.Body).Decode(&body)
	suite.Equal("This email has been registered. Choose another one", body.Message)
	suite.Equal(false, body.Success)
}

func (suite *HandlerRegisterUserSuite) TestRegisterUserNoEmailNegative() {
	requestBody, err := json.Marshal(map[string]string{
		"username": "username",
		"email":    "",
		"password": "password",
	})
	suite.NoError(err, "There should be no errors when create requestBody")

	response, err := http.Post(fmt.Sprintf("%s/auth/register", suite.testingServer.URL), "application/json", bytes.NewBuffer(requestBody))
	suite.Equal(http.StatusBadRequest, response.StatusCode)

	defer response.Body.Close()
	body := models.Response{}
	json.NewDecoder(response.Body).Decode(&body)
	suite.Equal("Please provide email or password", body.Message)
	suite.Equal(false, body.Success)
}

func (suite *HandlerRegisterUserSuite) TestRegisterUserNoPasswordNegative() {
	requestBody, err := json.Marshal(map[string]string{
		"username": "username",
		"email":    "email",
		"password": "",
	})
	suite.NoError(err, "There should be no errors when create requestBody")

	response, err := http.Post(fmt.Sprintf("%s/auth/register", suite.testingServer.URL), "application/json", bytes.NewBuffer(requestBody))
	suite.Equal(http.StatusBadRequest, response.StatusCode)

	defer response.Body.Close()
	body := models.Response{}
	json.NewDecoder(response.Body).Decode(&body)
	suite.Equal("Please provide email or password", body.Message)
	suite.Equal(false, body.Success)
}

func (suite *HandlerRegisterUserSuite) TestRegisterUserNoUsernamePositive() {
	requestBody, err := json.Marshal(map[string]string{
		"username": "",
		"email":    "email",
		"password": "password",
	})
	suite.NoError(err, "There should be no errors when create requestBody")

	response, err := http.Post(fmt.Sprintf("%s/auth/register", suite.testingServer.URL), "application/json", bytes.NewBuffer(requestBody))
	suite.Equal(http.StatusOK, response.StatusCode)

	defer response.Body.Close()
	body := models.Response{}
	json.NewDecoder(response.Body).Decode(&body)
	suite.Equal("Success to register user", body.Message)
	suite.Equal(true, body.Success)
}

func TestHandlerRegisterUserSuite(t *testing.T) {
	suite.Run(t, new(HandlerRegisterUserSuite))
}
