package cm_auth

import (
	"github.com/jjoykkm/ln-backend/common/models/model_other"
)

type Servicer interface {
	// status
	GetProvinceList(status string) (*model_other.RespModel, error)
}

type Service struct {
	repo  Repositorier
}

func NewService(repo Repositorier) Servicer {
	return &Service{
		repo:  repo,
	}
}

func (s *Service) GetProvinceList(status string) (*model_other.RespModel, error) {
	province, err := s.repo.FindAllProvince(status)
	if err != nil{
		return nil, err
	}
	return &model_other.RespModel{
		Item: province,
		Total: len(province),
	}, nil
}