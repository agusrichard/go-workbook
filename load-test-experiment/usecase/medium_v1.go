package usecase

import (
	"database/sql"
	"github.com/pkg/errors"
	"load-test-experiment/model"
	"load-test-experiment/repository"
	util "load-test-experiment/utils"
)

type mediumV1Usecase struct {
	mediumV1Repository repository.MediumV1Repository
}

type MediumV1Usecase interface {
	Create(s *model.MediumV1Shape) error
	Get(filterQuery string, skip, take int) (*[]model.MediumV1Shape, error)
}

func NewMediumV1Usecase(mediumV1Repository repository.MediumV1Repository) MediumV1Usecase {
	return &mediumV1Usecase{mediumV1Repository}
}

func (u *mediumV1Usecase) Create(s *model.MediumV1Shape) error {
	var fieldThree, fieldEight, fieldThirteen sql.NullString
	var m model.MediumV1Model

	if s == nil {
		return errors.New("MediumV1Usecase: Create: m is nil;")
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

	if s.FieldEight == "" {
		fieldEight = sql.NullString{
			Valid: false,
			String: s.FieldEight,
		}
	} else {
		fieldEight = sql.NullString{
			Valid: true,
			String: s.FieldEight,
		}
	}

	if s.FieldThirteen == "" {
		fieldThirteen = sql.NullString{
			Valid: false,
			String: s.FieldThirteen,
		}
	} else {
		fieldThirteen = sql.NullString{
			Valid: true,
			String: s.FieldThirteen,
		}
	}

	m = model.MediumV1Model{
		FieldOne: s.FieldOne,
		FieldTwo: s.FieldTwo,
		FieldThree: fieldThree,
		FieldFour: s.FieldFour,
		FieldFive: s.FieldFive,
		FieldSix: s.FieldSix,
		FieldSeven: s.FieldSeven,
		FieldEight: fieldEight,
		FieldNine: s.FieldNine,
		FieldTen: s.FieldTen,
		FieldEleven: s.FieldEleven,
		FieldTwelve: s.FieldTwelve,
		FieldThirteen: fieldThirteen,
		FieldFourteen: s.FieldFourteen,
	}

	for _, val := range s.MediumSmallModelList {
		var nullStr sql.NullString
		var timeFour sql.NullTime

		if val.FieldThree == "" {
			nullStr = sql.NullString{
				Valid: false,
				String: s.FieldThree,
			}
		} else {
			nullStr = sql.NullString{
				Valid: true,
				String: s.FieldThree,
			}
		}

		if val.FieldFour == "" {
			timeFour = sql.NullTime{
				Valid: false,
			}
		} else {
			t, err := util.ParseTime(val.FieldFour)

			if err != nil {
				return errors.Wrap(err, "MediumV1Usecase: Create: error parse time")
			}
			timeFour = sql.NullTime{
				Valid: true,
				Time: t,
			}
		}

		mdl := model.MediumV1SmallModel{
			FieldOne: val.FieldOne,
			FieldTwo: val.FieldTwo,
			FieldThree: nullStr,
			FieldFour: timeFour,
		}
		m.MediumSmallModelList = append(m.MediumSmallModelList, mdl)
	}

	err := u.mediumV1Repository.Create(&m)
	if err != nil {
		return errors.Wrap(err, "MediumV1Usecase: Create: error create")
	}

	return nil
}

func (u *mediumV1Usecase) Get(filterQuery string, skip, take int) (*[]model.MediumV1Shape, error) {
	var largeListShape []model.MediumV1Shape

	largeListModel, err := u.mediumV1Repository.GetLarge(filterQuery, skip, take)
	if err != nil {
		return nil, errors.Wrap(err, "MediumV1Usecase: Get: failed get large list models;")
	}

	for _, val := range *largeListModel {
		var smallListShape []model.MediumV1SmallShape

		smallListModel, err := u.mediumV1Repository.GetSmall(val.ID)
		if err != nil {
			return nil, errors.Wrap(err, "MediumV1Usecase: Get: failed get small list models;")
		}

		for _, v := range *smallListModel {
			var fieldFour string
			if v.FieldFour.Valid {
				fieldFour = util.TimeToString(v.FieldFour.Time)
			}

			smallShape := model.MediumV1SmallShape{
				ID: val.ID,
				FieldOne: v.FieldOne,
				FieldTwo: v.FieldTwo,
				FieldThree: v.FieldThree.String,
				FieldFour: fieldFour,
				SmallLargeKey: v.SmallLargeKey,
			}
			smallListShape = append(smallListShape, smallShape)
		}

		largeShape :=  model.MediumV1Shape{
			ID: val.ID,
			FieldOne: val.FieldOne,
			FieldTwo: val.FieldTwo,
			FieldThree: val.FieldThree.String,
			FieldFour: val.FieldFour,
			FieldFive: val.FieldFive,
			FieldSix: val.FieldSix,
			FieldSeven: val.FieldSeven,
			FieldEight: val.FieldEight.String,
			FieldNine: val.FieldNine,
			FieldTen: val.FieldTen,
			FieldEleven: val.FieldEleven,
			FieldTwelve: val.FieldTwelve,
			FieldThirteen: val.FieldThirteen.String,
			FieldFourteen: val.FieldFourteen,
			MediumSmallModelList: smallListShape,
		}

		largeListShape = append(largeListShape, largeShape)
	}

	return &largeListShape, nil
}