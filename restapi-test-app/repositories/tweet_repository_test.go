package repositories

import (
	"github.com/stretchr/testify/suite"
	"restapi-tested-app/config"
	"restapi-tested-app/entities"
	"restapi-tested-app/utils"
	"testing"
)

type tweetRepositorySuite struct {
	suite.Suite
	repository TweetRepository
	cleanupExecutor utils.TruncateTableExecutor
}

func (suite *tweetRepositorySuite) SetupTest() {
	configs := config.GetConfig()
	db := config.ConnectDB(configs)
	repository := InitializeTweetRepository(db)

	suite.repository = repository

	suite.cleanupExecutor = utils.InitTruncateTableExecutor(db)
}

func (suite *tweetRepositorySuite) TearDownTest() {
	defer suite.cleanupExecutor.TruncateTable([]string{"tweets"})
}

func (suite *tweetRepositorySuite) TestCreateTweetPositive() {
	tweet := entities.Tweet{
		Username: "username",
		Text: "text",
	}

	err := suite.repository.CreateTweet(&tweet)
	suite.NoError(err, "no error when create tweet with valid input")
}

func (suite *tweetRepositorySuite) TestCreateTweetNilPointerNegative() {
	err := suite.repository.CreateTweet(nil)
	suite.Error(err, "create error with nil input returns error")
}

func (suite *tweetRepositorySuite) TestCreateTweetEmptyFieldsPositive() {
	var tweet entities.Tweet
	err := suite.repository.CreateTweet(&tweet)
	suite.NoError(err, "no error when create tweet with empty fields")
}

func TestTweetRepository(t *testing.T) {
	suite.Run(t, new(tweetRepositorySuite))
}