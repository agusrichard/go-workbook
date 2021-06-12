package usecase

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"load-test-experiment/model"
	"load-test-experiment/repository"
	util "load-test-experiment/utils"
)

type lightV1Usecase struct {
	lightV1Repository repository.LightV1Repository
}

type LightV1Usecase interface {
	Create(m *model.LightV1Shape) error
	Get(filterQuery string, skip, take int) (*[]model.LightV1Shape, error)
}

func NewLightV1Usecase(lightV1Repository repository.LightV1Repository) LightV1Usecase {
	return &lightV1Usecase{lightV1Repository}
}

func (u *lightV1Usecase) Create(s *model.LightV1Shape) error {
	var fieldThree sql.NullString
	var fieldFour sql.NullTime

	fmt.Printf("s :=> %+v\n", s)
	if s == nil {
		return errors.New("LightV1Usecase: Create: s is nil;")
	}

	if s.FieldThree == "" {
		fieldThree = sql.NullString{
			Valid: false,
			String: s.FieldThree,
		}
	} else {
		fieldThree = sql.NullString{
			Valid: true,
			String: s.FieldThree,
		}
	}

	if s.FieldFour == "" {
		fieldFour = sql.NullTime{
			Valid: false,
		}
	} else {
		t, err := util.ParseTime(s.FieldFour)

		if err != nil {
			return errors.Wrap(err, "LightV1Usecase: Create: error parse time")
		}
		fieldFour = sql.NullTime{
			Valid: true,
			Time: t,
		}
	}


	m := model.LightV1Model{
		FieldOne: s.FieldOne,
		FieldTwo: s.FieldTwo,
		FieldThree: fieldThree,
		FieldFour: fieldFour,
	}

	err := u.lightV1Repository.Create(&m)
	if err != nil {
		return errors.Wrap(err, "LightV1Usecase: Create: failed to create in usecase;")
	}

	return nil
}

func (u *lightV1Usecase) Get(filterQuery string, skip, take int) (*[]model.LightV1Shape, error) {
	var ss []model.LightV1Shape

	result, err := u.lightV1Repository.Get(filterQuery, skip, take)
	if err != nil {
		return nil, errors.Wrap(err, "LightV1Usecase: Get: error get data")
	}

	for _, value := range *result {
		var fieldFour string
		if value.FieldFour.Valid {
			fieldFour = util.TimeToString(value.FieldFour.Time)
		}
		ss = append(ss, model.LightV1Shape{
			ID: value.ID,
			FieldOne: value.FieldOne,
			FieldTwo: value.FieldTwo,
			FieldThree: value.FieldThree.String,
			FieldFour: fieldFour,
		})
	}


	return &ss, nil
}