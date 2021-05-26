package usecases

import "restapi-tested-app/repositories"

type tweetUsecase struct {
	tweetRepository repositories.TweetRepository
}

type TweetUsecase interface {
	CreateTweet()
}

func InitializeTweetUsecase(repository repositories.TweetRepository) TweetUsecase {
	return &tweetUsecase{repository}
}

func (usecase *tweetUsecase) CreateTweet() {

}