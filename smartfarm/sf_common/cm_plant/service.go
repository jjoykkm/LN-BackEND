package cm_plant

import (
	"github.com/jjoykkm/ln-backend/common/models/model_other"
)

type Servicer interface {
	// status
	GetFertAndCatList(status string) (*model_other.RespModel, error)
}

type Service struct {
	repo  Repositorier
}

func NewService(repo Repositorier) Servicer {
	return &Service{
		repo:  repo,
	}
}

func (s *Service) GetFertAndCatList(status string) (*model_other.RespModel, error) {
	fert, err := s.repo.FindAllFertAndCatList(status)
	if err != nil{
		return nil, err
	}
	return &model_other.RespModel{
		Item: fert,
		Total: len(fert),
	}, nil
}
