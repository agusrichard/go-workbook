package repositories

import (
	"fmt"
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

func (suite *tweetRepositorySuite) SetupSuite() {
	fmt.Println("SetupSuite")
	configs := config.GetConfig()
	db := config.ConnectDB(configs)
	repository := InitializeTweetRepository(db)

	suite.repository = repository

	suite.cleanupExecutor = utils.InitTruncateTableExecutor(db)
}

func (suite *tweetRepositorySuite) TearDownTest() {
	fmt.Println("TearDownTest")
	defer suite.cleanupExecutor.TruncateTable([]string{"tweets"})
}

func (suite *tweetRepositorySuite) TestCreateTweet_Positive() {
	tweet := entities.Tweet{
		Username: "username",
		Text: "text",
	}

	err := suite.repository.CreateTweet(&tweet)
	suite.NoError(err, "no error when create tweet with valid input")
}

func (suite *tweetRepositorySuite) TestCreateTweet_NilPointer_Negative() {
	err := suite.repository.CreateTweet(nil)
	suite.Error(err, "create error with nil input returns error")
}

func (suite *tweetRepositorySuite) TestCreateTweet_EmptyFields_Positive() {
	var tweet entities.Tweet
	err := suite.repository.CreateTweet(&tweet)
	suite.NoError(err, "no error when create tweet with empty fields")
}

func (suite *tweetRepositorySuite) TestGetAllTweets_EmptySlice_Positive() {
	tweets, err := suite.repository.GetAllTweets()
	suite.NoError(err, "no error when get all tweets when the table is empty")
	suite.Equal(len(*tweets), 0, "length of tweets should be 0, since it is empty slice")
	suite.Equal(*tweets, []entities.Tweet(nil), "tweets is an empty slice")
}

func (suite *tweetRepositorySuite) TestGetAllTweets_FilledRecords_Positive() {
	tweet := entities.Tweet{
		Username: "username",
		Text: "text",
	}

	// inserting 3 tweets to be queried later
	err := suite.repository.CreateTweet(&tweet)
	suite.NoError(err, "no error when create tweet with valid input")
	err = suite.repository.CreateTweet(&tweet)
	suite.NoError(err, "no error when create tweet with valid input")
	err = suite.repository.CreateTweet(&tweet)
	suite.NoError(err, "no error when create tweet with valid input")

	tweets, err := suite.repository.GetAllTweets()
	suite.NoError(err, "no error when get all tweets when the table is empty")
	suite.Equal(len(*tweets), 3, "insert 3 records before the all data, so it should contain three tweets")
}

func (suite *tweetRepositorySuite) TestGetTweetByID_NotFound_Negative() {
	id := 1

	_, err := suite.repository.GetTweetByID(id)
	suite.Error(err, "error sql not found")
	suite.Equal(err.Error(), "sql: no rows in result set")
}

func (suite *tweetRepositorySuite) TestGetTweetByID_Exists_Positive() {
	id := 1
	tweet := entities.Tweet{
		Username: "username",
		Text: "text",
	}

	// will create a record with id 1
	err := suite.repository.CreateTweet(&tweet)
	suite.NoError(err, "no error when create tweet with valid input")

	result, err := suite.repository.GetTweetByID(id)
	suite.NoError(err, "no error because tweet is found")
	suite.Equal(tweet.Username, (*result).Username, "should be equal between result and tweet")
	suite.Equal(tweet.Text, (*result).Text, "should be equal between result and tweet")
}

func TestTweetRepository(t *testing.T) {
	suite.Run(t, new(tweetRepositorySuite))
}