package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"load-test-experiment/model"
)

type lightV2Repository struct {
	db *sqlx.DB
}

type LightV2Repository interface {
	Create(m *model.LightV2Model) error
	Get(filterQuery string, skip, take int) (*[]model.LightV2Model, error)
}

func NewLightV2Repository(db *sqlx.DB) LightV2Repository {
	return &lightV2Repository{db}
}

func (r *lightV2Repository) Create(m *model.LightV2Model) error {
	if m == nil {
		return errors.New("LightV2Repository: Create: m is nil;")
	}

	tx, err := r.db.Beginx()
	if err != nil {
		return errors.Wrap(err, "LightV2Repository: Create: failed to initiate transaction;")
	}

	err = insertM(tx, m)
	if err != nil {
		tx.Rollback()
		return errors.Wrap(err, "LightV2Repository: Create: failed to insert m in repository;")
	}

	tx.Commit()

	return nil
}

func insertM(tx *sqlx.Tx, m *model.LightV2Model) error {
	_, err := tx.NamedExec(`
		INSERT INTO light_table(field_one, field_two, field_three, field_four)
		VALUES (:field_one, :field_two, :field_three, :field_four);
;	`, m)

	return err
}


func (r *lightV2Repository) Get(filterQuery string, skip, take int) (*[]model.LightV2Model, error) {
	var result []model.LightV2Model

	query := fmt.Sprintf(`
		SELECT id, field_one, field_two, field_three, field_four
		FROM light_table
		%s
		OFFSET %d
		LIMIT %d;
	`, filterQuery, skip, take)

	err := r.db.Select(&result, query)
	if err != nil {
		return nil, errors.Wrap(err, "LightV2Repository: Create: failed to get data;")
	}

	return &result, nil
}