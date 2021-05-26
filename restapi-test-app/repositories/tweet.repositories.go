package repositories

import (
	"github.com/jmoiron/sqlx"
	"restapi-tested-app/entities"
)

type tweetRepository struct {
	db *sqlx.DB
}

type TweetRepository interface {
	GetAllTweets() (*[]entities.Tweet, error)
	CreateTweet(tweet *entities.Tweet) error
}

func InitializeRepository(db *sqlx.DB) TweetRepository {
	return &tweetRepository{db}
}

func (repository *tweetRepository) GetAllTweets() (*[]entities.Tweet, error) {
	var result []entities.Tweet
	rows, err := repository.db.Queryx(`SELECT id, username, text, created_at, modified_at FROM tweets`)

	for rows.Next() {
		var tweet entities.Tweet
		err = rows.StructScan(&tweet)
		result = append(result, tweet)
	}

	return &result, err
}

func (repository *tweetRepository) CreateTweet(tweet *entities.Tweet) error {
	var err error

	tx, errTx := repository.db.Beginx()
	if errTx != nil {
	} else {
		err = insertTweet(tx, tweet)
		if err != nil {
		}
	}

	if err == nil {
		tx.Commit()
	} else {
		tx.Rollback()
	}

	return err
}

func insertTweet(tx *sqlx.Tx, tweet *entities.Tweet) error {
	_, err := tx.NamedExec(`
		INSERT INTO tweets(username, text)
		VALUES (:username, :text)
	`, tweet)

	return err
}