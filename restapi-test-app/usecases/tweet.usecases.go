package usecases

import (
	"restapi-tested-app/entities"
	"restapi-tested-app/repositories"
	"time"
)

type tweetUsecase struct {
	tweetRepository repositories.TweetRepository
}

type TweetUsecase interface {
	GetAllTweets() (*[]entities.Tweet, error)
	GetTweetByID(id int) (*entities.Tweet, error)
	SearchTextByText(text string) (*[]entities.Tweet, error)
	CreateTweet(tweet *entities.Tweet) error
	UpdateTweet(tweet *entities.Tweet) error
	DeleteTweet(id int) error
}

func InitializeTweetUsecase(repository repositories.TweetRepository) TweetUsecase {
	return &tweetUsecase{repository}
}

func (usecase *tweetUsecase) GetAllTweets() (*[]entities.Tweet, error) {
	return usecase.tweetRepository.GetAllTweets()
}

func (usecase *tweetUsecase) GetTweetByID(id int) (*entities.Tweet, error) {
	return usecase.tweetRepository.GetTweetByID(id)
}

func (usecase *tweetUsecase) SearchTextByText(text string) (*[]entities.Tweet, error) {
	return usecase.tweetRepository.SearchTweetByText("%"+text+"%")
}

func (usecase *tweetUsecase) CreateTweet(tweet *entities.Tweet) error {
	err := usecase.tweetRepository.CreateTweet(tweet)
	return err
}

func (usecase *tweetUsecase) UpdateTweet(tweet *entities.Tweet) error {
	tweet.ModifiedAt = time.Now()
	err := usecase.tweetRepository.UpdateTweet(tweet)
	return err
}

func (usecase *tweetUsecase) DeleteTweet(id int) error {
	return usecase.tweetRepository.DeleteTweet(id)
}