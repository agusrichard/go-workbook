package repository

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"load-test-experiment/model"
	"time"
)

type mediumV1Repository struct {
	db *sqlx.DB
}

type MediumV1Repository interface {
	Create(m *model.MediumV1Model) error
	GetLarge(filterQuery string, skip, take int) (*[]model.MediumV1Model, error)
	GetSmall(largeKey int) (*[]model.MediumV1SmallModel, error)
}

func NewMediumV1Repository(db *sqlx.DB) MediumV1Repository {
	return &mediumV1Repository{db}
}

func (r *mediumV1Repository) Create(m *model.MediumV1Model) error {
	if m == nil {
		return errors.New("MediumV1Repository: Create: m is nil;")
	}

	tx, err := r.db.Beginx()
	if err != nil {
		return errors.Wrap(err, "MediumV1Repository: Create: failed to initiate transaction;")
	}

	key, err := insertMediumLarge(tx, m)
	if err != nil {
		tx.Rollback()
		return errors.Wrap(err, "MediumV1Repository: Create: failed to insert medium large;")
	}

	for _, val := range m.MediumSmallModelList {
		val.SmallLargeKey = key
		err = insertMediumSmall(tx, &val)
		if err != nil {
			tx.Rollback()
			return errors.Wrap(err, "MediumV1Repository: Create: failed to insert m in repository;")
		}
	}

	tx.Commit()

	return nil
}

func insertMediumSmall(tx *sqlx.Tx, m *model.MediumV1SmallModel) error {
	_, err := tx.Exec(`
		INSERT INTO medium_small_table(field_one, field_two, field_three, field_four, small_large_key)
		VALUES ($1, $2, $3, $4, $5);
;	`,
		m.FieldOne,
		m.FieldTwo,
		m.FieldThree,
		m.FieldFour,
		m.SmallLargeKey,
	)

	return err
}

func insertMediumLarge(tx *sqlx.Tx, m *model.MediumV1Model) (int, error) {
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

func (r *mediumV1Repository) GetLarge(filterQuery string, skip, take int) (*[]model.MediumV1Model, error) {
	var result []model.MediumV1Model

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

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, errors.Wrap(err, "MediumV1Repository: GetLarge: failed to query the data")
	}
	defer rows.Close()

	for rows.Next() {
		var id, fieldFive, fieldTen int
		var fieldOne, fieldSix, fieldEleven string
		var fieldTwo, fieldSeven, fieldTwelve float64
		var fieldThree, fieldEight, fieldThirteen sql.NullString
		var fieldFour, fieldNine, fieldFourteen time.Time

		err = rows.Scan(
			&id,
			&fieldOne,
			&fieldTwo,
			&fieldThree,
			&fieldFour,
			&fieldFive,
			&fieldSix,
			&fieldSeven,
			&fieldEight,
			&fieldNine,
			&fieldTen,
			&fieldEleven,
			&fieldTwelve,
			&fieldThirteen,
			&fieldFourteen,
		)
		if err != nil {
			return nil, errors.Wrap(err, "MediumV1Repository: GetLarge: failed to scan rows")
		}

		result = append(result, model.MediumV1Model{
			ID: id,
			FieldOne: fieldOne,
			FieldTwo: fieldTwo,
			FieldThree: fieldThree,
			FieldFour: fieldFour,
			FieldFive: fieldFive,
			FieldSix: fieldSix,
			FieldSeven: fieldSeven,
			FieldEight: fieldEight,
			FieldNine: fieldNine,
			FieldTen: fieldTen,
			FieldEleven: fieldEleven,
			FieldTwelve: fieldTwelve,
			FieldThirteen: fieldThirteen,
			FieldFourteen: fieldFourteen,
		})
	}

	return &result, nil
}

func (r *mediumV1Repository) GetSmall(largeKey int) (*[]model.MediumV1SmallModel, error) {
	var result []model.MediumV1SmallModel

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

	rows, err := r.db.Query(query, largeKey)
	if err != nil {
		return nil, errors.Wrap(err, "MediumV1Repository: GetSmall: failed to query the data")
	}
	defer rows.Close()

	for rows.Next() {
		var id, smallLargeKey int
		var fieldOne string
		var fieldTwo float64
		var fieldThree sql.NullString
		var fieldFour sql.NullTime

		err = rows.Scan(
			&id,
			&fieldOne,
			&fieldTwo,
			&fieldThree,
			&fieldFour,
			&smallLargeKey,
		)
		if err != nil {
			return nil, errors.Wrap(err, "MediumV1Repository: GetSmall: failed to scan rows")
		}

		result = append(result, model.MediumV1SmallModel{
			ID: id,
			FieldOne: fieldOne,
			FieldTwo: fieldTwo,
			FieldThree: fieldThree,
			FieldFour: fieldFour,
			SmallLargeKey: smallLargeKey,
		})
	}

	return &result, nil
}