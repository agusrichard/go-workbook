package usecase

import (
	"github.com/pkg/errors"
	"load-test-experiment/model"
	"load-test-experiment/repository"
)

type mediumV2Usecase struct {
	mediumV2Repository repository.MediumV2Repository
}

type MediumV2Usecase interface {
	Create(m *model.MediumV2Model) error
	Get(filterQuery string, skip, take int) (*[]model.MediumV2Model, error)
}

func NewMediumV2Usecase(mediumV2Repository repository.MediumV2Repository) MediumV2Usecase {
	return &mediumV2Usecase{mediumV2Repository}
}

func (u *mediumV2Usecase) Create(m *model.MediumV2Model) error {
	if m == nil {
		return errors.New("MediumV2Usecase: Create: m is nil;")
	}

	err := u.mediumV2Repository.Create(m)
	if err != nil {
		return errors.Wrap(err, "MediumV2Usecase: Create: failed create m;")
	}

	return nil
}

func (u *mediumV2Usecase) Get(filterQuery string, skip, take int) (*[]model.MediumV2Model, error) {
	largeList, err := u.mediumV2Repository.GetLarge(filterQuery, skip, take)
	if err != nil {
		return nil, errors.Wrap(err, "MediumV2Usecase: Get: failed get large list;")
	}

	for i, val := range *largeList {
		smallList, err := u.mediumV2Repository.GetSmall(val.ID)
		if err != nil {
			return nil, errors.Wrap(err, "MediumV2Usecase: Get: failed get small list;")
		}
		(*largeList)[i].MediumSmallModelList = *smallList
	}

	return largeList, nil
}