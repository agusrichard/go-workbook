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

func (suite *tweetUsecaseSuite) TestCreateTweet_Positive() {
	tweet := entities.Tweet{
		Username: "username",
		Text: "text",
	}

	suite.repository.On("CreateTweet", &tweet).Return(nil)

	err := suite.usecase.CreateTweet(&tweet)
	suite.Nil(err, "err is a nil pointer so no error in this process")
	suite.repository.AssertExpectations(suite.T())
}

func (suite *tweetUsecaseSuite) TestCreateTweet_NilPointer_Negative() {
	err := suite.usecase.CreateTweet(nil)
	suite.Error(err.(*entities.AppError).Err, "error when create tweet with nil pointer")
	suite.repository.AssertExpectations(suite.T())
}

func (suite *tweetUsecaseSuite) TestGetAllTweets_EmptySlice_Positive() {
	emptyTweets := []entities.Tweet(nil)
	suite.repository.On("GetAllTweets").Return(&emptyTweets, nil)
	tweets, err := suite.usecase.GetAllTweets()
	suite.NoError(err, "no error when get all tweets")
	suite.Equal(len(*tweets), 0, "tweets is a empty slice object")
}

func (suite *tweetUsecaseSuite) TestGetAllTweets_FilledSlice_Positive() {
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
	suite.repository.On("GetAllTweets").Return(&tweets, nil)
	result, err := suite.usecase.GetAllTweets()
	suite.NoError(err, "no error when get all tweets")
	suite.Equal(len(*result), len(tweets), "tweets and result should have the same length")
	suite.Equal(*result, tweets, "result and tweets are the same")
}

func TestTweetUsecase(t *testing.T) {
	suite.Run(t, new(tweetUsecaseSuite))
}