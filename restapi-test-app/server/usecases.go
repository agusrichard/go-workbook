package server

import "restapi-tested-app/usecases"

type Usecases struct {
	TweetUsecase usecases.TweetUsecase
}

func setupUsecases(repos *Repositories) *Usecases {
	tweetUsecase := usecases.InitializeTweetUsecase(repos.TweetRepository)

	return &Usecases{
		TweetUsecase: tweetUsecase,
	}
}