package usecases

import (
	"restapi-tested-app/entities"
	"restapi-tested-app/repositories"
)

type tweetUsecase struct {
	tweetRepository repositories.TweetRepository
}

type TweetUsecase interface {
	GetAllTweets() (*[]entities.Tweet, error)
	CreateTweet(tweet *entities.Tweet) error
}

func InitializeTweetUsecase(repository repositories.TweetRepository) TweetUsecase {
	return &tweetUsecase{repository}
}

func (usecase *tweetUsecase) GetAllTweets() (*[]entities.Tweet, error) {
	return usecase.tweetRepository.GetAllTweets()
}

func (usecase *tweetUsecase) CreateTweet(tweet *entities.Tweet) error {
	err := usecase.tweetRepository.CreateTweet(tweet)
	return err
}