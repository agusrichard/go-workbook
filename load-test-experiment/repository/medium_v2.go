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
	GetLarge(filterQuery string, skip, take int) (*[]model.MediumV2Model, error)
	GetSmall(largeKey int) (*[]model.MediumV2SmallModel, error)
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

	key, err := insertMediumLarge(tx, m)
	if err != nil {
		tx.Rollback()
		return errors.Wrap(err, "MediumV2Repository: Create: failed to insert medium large;")
	}

	for _, val := range m.MediumSmallModelList {
		val.SmallLargeKey = key
		err = insertMediumSmall(tx, &val)
		if err != nil {
			tx.Rollback()
			return errors.Wrap(err, "MediumV2Repository: Create: failed to insert m in repository;")
		}
	}

	tx.Commit()

	return nil
}

func insertMediumSmall(tx *sqlx.Tx, m *model.MediumV2SmallModel) error {
	_, err := tx.NamedExec(`
		INSERT INTO medium_small_table(field_one, field_two, field_three, field_four, small_large_key)
		VALUES (:field_one, :field_two, :field_three, :field_four, :small_large_key);
;	`, m)

	return err
}

func insertMediumLarge(tx *sqlx.Tx, m *model.MediumV2Model) (int, error) {
	var key int

	err := tx.QueryRowx(`
		INSERT INTO medium_large_table(
			field_one,
			field_two,
			field_three,
			field_five,
			field_six,
			field_seven,
			field_eight,
			field_ten,
			field_eleven,
			field_twelve,
			field_thirteen
		)
		VALUES (
			$1,
			$2,
			$3,
			$4,
			$5,
			$6,
			$7,
			$8,
			$9,
			$10,
			$11
		)
		RETURNING id;
;	`,
		m.FieldOne,
		m.FieldTwo,
		m.FieldThree,
		m.FieldFive,
		m.FieldSix,
		m.FieldSeven,
		m.FieldEight,
		m.FieldTen,
		m.FieldEleven,
		m.FieldTwelve,
		m.FieldThirteen,
	).Scan(&key)

	return key, err
}

func (r *mediumV2Repository) GetLarge(filterQuery string, skip, take int) (*[]model.MediumV2Model, error) {
	var result []model.MediumV2Model

	query := fmt.Sprintf(`
		SELECT
			id,
			field_one,
			field_two,
			field_three,
			field_four,
			field_five,
			field_six,
			field_seven,
			field_eight,
			field_nine,
			field_ten,
			field_eleven,
			field_twelve,
			field_thirteen,
			field_fourteen
		FROM medium_large_table
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

func (r *mediumV2Repository) GetSmall(largeKey int) (*[]model.MediumV2SmallModel, error) {
	var result []model.MediumV2SmallModel

	query := `
		SELECT
		    id,
			field_one,
			field_two,
			field_three,
			field_four,
		    small_large_key
		FROM medium_small_table
		WHERE small_large_key=$1;
	`

	err := r.db.Select(&result, query, largeKey)
	if err != nil {
		return nil, errors.Wrap(err, "MediumV2Repository: Create: failed to get data;")
	}

	return &result, nil
}