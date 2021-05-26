package server

import (
	"github.com/jmoiron/sqlx"
	"restapi-tested-app/repositories"
)

type Repositories struct {
	TweetRepository repositories.TweetRepository
}

func setupRepositories(db *sqlx.DB) *Repositories {
	tweetRepository := repositories.InitializeRepository(db)

	return &Repositories{
		TweetRepository: tweetRepository,
	}
}