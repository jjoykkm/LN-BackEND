package sf_formula_plant

import (
	"github.com/jjoykkm/ln-backend/common/models/model_other"
)

type Servicer interface {
	// Service for API
	// status, language string
	GetPlantCategoryList(status string, ReqModel *model_other.ReqModel) (*model_other.RespModel, error)
}

type Service struct {
	repo  Repositorier
}

func NewService(repo Repositorier) Servicer {
	return &Service{
		repo:  repo,
	}
}

func (s *Service) GetPlantCategoryList(status string, ReqModel *model_other.ReqModel) (*model_other.RespModel, error) {
	plantTypeList, err := s.repo.FindAllPlantType(status)
	if err != nil{
		return nil, err
	}

	return &model_other.RespModel{
		Item: plantTypeList,
		Total: len(plantTypeList),
	}, nil
}