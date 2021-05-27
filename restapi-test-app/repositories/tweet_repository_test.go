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

func (suite *tweetRepositorySuite) TestCreateTweet() {
	tweet := entities.Tweet{
		Username: "username",
		Text: "text",
	}

	err := suite.repository.CreateTweet(&tweet)
	suite.NoError(err)
}

func TestRepositoryRegisterUserSuite(t *testing.T) {
	suite.Run(t, new(tweetRepositorySuite))
}