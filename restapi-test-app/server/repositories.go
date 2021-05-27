package server

import (
	"github.com/jmoiron/sqlx"
	"restapi-tested-app/repositories"
)

type Repositories struct {
	TweetRepository repositories.TweetRepository
}

func SetupRepositories(db *sqlx.DB) *Repositories {
	tweetRepository := repositories.InitializeTweetRepository(db)

	return &Repositories{
		TweetRepository: tweetRepository,
	}
}