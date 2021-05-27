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
	usecase := new(mocks.TweetUsecase)
	handler := InitializeTweetHandler(usecase)

	router := gin.Default()
	router.POST("/tweet", utils.ServeHTTP(handler.CreateTweet))
	router.GET("/tweet", utils.ServeHTTP(handler.GetAllTweets))
	testingServer := httptest.NewServer(router)

	suite.testingServer = testingServer
	suite.usecase = usecase
	suite.handler = handler
}

func (suite *tweetHandlerSuite) TearDownTest() {
	defer suite.testingServer.Close()
}

func (suite *tweetHandlerSuite) TestCreateTweet_Positive() {
	tweet := entities.Tweet{
		Username: "username",
		Text: "text",
	}

	suite.usecase.On("CreateTweet", &tweet).Return(nil)

	requestBody, err := json.Marshal(&tweet)
	suite.NoError(err, "can not marshal struct to json")

	response, err := http.Post(fmt.Sprintf("%s/tweet", suite.testingServer.URL), "application/json", bytes.NewBuffer(requestBody))
	suite.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	responseBody := entities.Response{}
	json.NewDecoder(response.Body).Decode(&responseBody)

	suite.Equal(http.StatusCreated, response.StatusCode)
	suite.Equal(responseBody.Message, "Success to create tweet")
	suite.usecase.AssertExpectations(suite.T())
}

func (suite *tweetHandlerSuite) TestGetAllTweets_Positive() {
	tweets := []entities.Tweet{
		{
			Username: "username",
			Text: "text",
		},
		{
			Username: "username",
			Text: "text",
		},
		{
			Username: "username",
			Text: "text",
		},
	}
	fmt.Println("tweets", tweets)

	suite.usecase.On("GetAllTweets").Return(&tweets, nil)

	response, err := http.Get(fmt.Sprintf("%s/tweet", suite.testingServer.URL))
	suite.NoError(err, "no error when calling this endpoint")
	defer response.Body.Close()

	responseBody := entities.Response{}
	json.NewDecoder(response.Body).Decode(&responseBody)

	suite.Equal(http.StatusOK, response.StatusCode)
	suite.Equal(responseBody.Message, "Success to get all tweets")
	suite.usecase.AssertExpectations(suite.T())
}

func TestTweetHandler(t *testing.T) {
	suite.Run(t, new(tweetHandlerSuite))
}