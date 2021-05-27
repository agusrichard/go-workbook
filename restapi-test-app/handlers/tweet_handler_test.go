package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"restapi-tested-app/entities"
	"restapi-tested-app/mocks"
	"restapi-tested-app/utils"
	"testing"
)

type tweetHandlerSuite struct {
	suite.Suite
	usecase *mocks.TweetUsecase
	handler TweetHandler
	testingServer   *httptest.Server
}

func (suite *tweetHandlerSuite) SetupTest() {
	fmt.Println("setup test")
	usecase := new(mocks.TweetUsecase)
	handler := InitializeTweetHandler(usecase)

	router := gin.Default()
	router.POST("/tweet", utils.ServeHTTP(handler.CreateTweet))
	testingServer := httptest.NewServer(router)

	suite.testingServer = testingServer
	suite.usecase = usecase
	suite.handler = handler
}

func (suite *tweetHandlerSuite) TearDownTest() {
	fmt.Println("teardown test")
	defer suite.testingServer.Close()
}

func (suite *tweetHandlerSuite) TestTweetHandler() {
	tweet := entities.Tweet{
		Username: "username",
		Text: "text",
	}

	suite.usecase.On("CreateTweet", &tweet).Return(nil)

	requestBody, err := json.Marshal(&tweet)
	if err != nil {
		suite.T().Fatalf("can not marshal struct to json")
	}

	response, err := http.Post(fmt.Sprintf("%s/tweet", suite.testingServer.URL), "application/json", bytes.NewBuffer(requestBody))
	defer response.Body.Close()
	responseBody := entities.Response{}
	json.NewDecoder(response.Body).Decode(&responseBody)

	suite.Equal(http.StatusCreated, response.StatusCode)
	suite.Equal(responseBody.Message, "Success to create tweet")
	suite.usecase.AssertExpectations(suite.T())
}

func TestTweetHandler(t *testing.T) {
	suite.Run(t, new(tweetHandlerSuite))
}