package usecases

import (
	"github.com/stretchr/testify/suite"
	"restapi-tested-app/entities"
	"restapi-tested-app/mocks"
	"restapi-tested-app/utils"
	"testing"
)

type tweetUsecaseSuite struct {
	suite.Suite
	repository *mocks.TweetRepository
	usecase TweetUsecase
	cleanupExecutor utils.TruncateTableExecutor
}

func (suite *tweetUsecaseSuite) SetupTest() {
	repository := new(mocks.TweetRepository)
	usecase := InitializeTweetUsecase(repository)


	suite.repository = repository
	suite.usecase = usecase
}

func (suite *tweetUsecaseSuite) TestCreateTweetPositive() {
	tweet := entities.Tweet{
		Username: "username",
		Text: "text",
	}

	suite.repository.On("CreateTweet", &tweet).Return(nil)

	suite.usecase.CreateTweet(&tweet)
	suite.repository.AssertExpectations(suite.T())
}

func (suite *tweetUsecaseSuite) TestCreateTweetNilPointerNegative() {
	err := suite.usecase.CreateTweet(nil)
	suite.Error(err.Err, "error when create tweet with nil pointer")
	suite.repository.AssertExpectations(suite.T())
}

func TestTweetUsecase(t *testing.T) {
	suite.Run(t, new(tweetUsecaseSuite))
}