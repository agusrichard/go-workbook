package server

import "restapi-tested-app/handlers"

type Handlers struct {
	TweetHandler handlers.TweetHandler
}

func setupHandlers(uscs *Usecases) *Handlers {
	tweetHandlers := handlers.InitializeTweetHandler(uscs.TweetUsecase)

	return &Handlers{
		TweetHandler: tweetHandlers,
	}
}