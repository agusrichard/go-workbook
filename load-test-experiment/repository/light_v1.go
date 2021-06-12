package repository

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"load-test-experiment/model"
)

type lightV1Repository struct {
	db *sqlx.DB
}

type LightV1Repository interface {
	Create(m *model.LightV1Model) error
	Get(filterQuery string, skip, take int) (*[]model.LightV1Model, error)
}

func NewV1Repository(db *sqlx.DB) LightV1Repository {
	return &lightV1Repository{db}
}

func (r *lightV1Repository) Create(m *model.LightV1Model) error {
	if m == nil {
		return errors.New("LightV1Repository: Create: m is nil;")
	}

	tx, err := r.db.Begin()
	if err != nil {
		return errors.Wrap(err, "LightV1Repository: Create: failed to initiate transaction;")
	}

	err = r.insert(tx, m)
	if err != nil {
		tx.Rollback()
		return errors.Wrap(err, "LightV1Repository: Create: failed to insert m in repository;")
	}

	tx.Commit()

	return nil
}

func (r *lightV1Repository) insert(tx *sql.Tx, m *model.LightV1Model) error {
	fmt.Printf("m :=> %+v\n", m)
	_, err := tx.Exec(`
		INSERT INTO light_table(field_one, field_two, field_three, field_four)
		VALUES ($1, $2, $3, $4);
;	`, m.FieldOne, m.FieldTwo, m.FieldThree, m.FieldFour)

	return err
}

func (r *lightV1Repository) Get(filterQuery string, skip, take int) (*[]model.LightV1Model, error) {
	var ms []model.LightV1Model

	query := fmt.Sprintf(
		`SELECT id, field_one, field_two, field_three, field_four FROM light_table
		%s
		OFFSET %d
		LIMIT %d;
		`, filterQuery, skip, take)

	fmt.Println("query", query)

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, errors.Wrap(err, "LightV1Repository: Get: failed to query the data")
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var fieldOne string
		var fieldTwo float64
		var fieldThree sql.NullString
		var fieldFour sql.NullTime

		err = rows.Scan(&id, &fieldOne, &fieldTwo, &fieldThree, &fieldFour)
		if err != nil {
			return nil, errors.Wrap(err, "LightV1Repository: Get: failed to scan rows")
		}

		ms = append(ms, model.LightV1Model{
			ID: id,
			FieldOne: fieldOne,
			FieldTwo: fieldTwo,
			FieldThree: fieldThree,
			FieldFour: fieldFour,
		})
	}

	fmt.Println("ms", ms)

	return &ms, nil
}
