package repositories

import (
	"github.com/jmoiron/sqlx"
)

type tweetRepository struct {
	db *sqlx.DB
}

type TweetRepository interface {
	CreateTweet()
}

func InitializeRepository(db *sqlx.DB) TweetRepository {
	return &tweetRepository{db}
}

func (repository *tweetRepository) CreateTweet() {

}