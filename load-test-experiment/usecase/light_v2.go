package usecase

import (
	"github.com/pkg/errors"
	"load-test-experiment/model"
	"load-test-experiment/repository"
)

type lightV2Usecase struct {
	lightV2Repository repository.LightV2Repository
}

type LightV2Usecase interface {
	Create(m *model.LightV2Model) error
	Get(filterQuery string, skip, take int) (*[]model.LightV2Model, error)
}

func NewLightV2Usecase(lightV2Repository repository.LightV2Repository) LightV2Usecase {
	return &lightV2Usecase{lightV2Repository}
}

func (u *lightV2Usecase) Create(m *model.LightV2Model) error {
	if m == nil {
		return errors.New("LightV2Usecase: Create: m is nil;")
	}

	err := u.lightV2Repository.Create(m)
	if err != nil {
		return errors.New("LightV2Usecase: Create: failed create m;")
	}

	return nil
}

func (u *lightV2Usecase) Get(filterQuery string, skip, take int) (*[]model.LightV2Model, error) {
	result, err := u.lightV2Repository.Get(filterQuery, skip, take)
	if err != nil {
		return nil, errors.Wrap(err, "todo usecase: filter todos: error get data")
	}

	return result, nil
}