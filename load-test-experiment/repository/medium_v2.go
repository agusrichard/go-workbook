package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"load-test-experiment/model"
)

type mediumV2Repository struct {
	db *sqlx.DB
}

type MediumV2Repository interface {
	Create(m *model.MediumV2Model) error
	Get(filterQuery string, skip, take int) (*[]model.MediumV2Model, error)
}

func NewMediumV2Repository(db *sqlx.DB) MediumV2Repository {
	return &mediumV2Repository{db}
}

func (r *mediumV2Repository) Create(m *model.MediumV2Model) error {
	if m == nil {
		return errors.New("MediumV2Repository: Create: m is nil;")
	}

	tx, err := r.db.Beginx()
	if err != nil {
		return errors.Wrap(err, "MediumV2Repository: Create: failed to initiate transaction;")
	}

	err = insertM(tx, m)
	if err != nil {
		tx.Rollback()
		return errors.Wrap(err, "MediumV2Repository: Create: failed to insert m in repository;")
	}

	tx.Commit()

	return nil
}

func insertM(tx *sqlx.Tx, m *model.MediumV2Model) error {
	_, err := tx.NamedExec(`
		INSERT INTO medium_table(field_one, field_two, field_three, field_four)
		VALUES (:field_one, :field_two, :field_three, :field_four);
;	`, m)

	return err
}


func (r *mediumV2Repository) Get(filterQuery string, skip, take int) (*[]model.MediumV2Model, error) {
	var result []model.MediumV2Model

	query := fmt.Sprintf(`
		SELECT id, field_one, field_two, field_three, field_four
		FROM medium_table
		%s
		OFFSET %d
		LIMIT %d;
	`, filterQuery, skip, take)

	err := r.db.Select(&result, query)
	if err != nil {
		return nil, errors.Wrap(err, "MediumV2Repository: Create: failed to get data;")
	}

	return &result, nil
}