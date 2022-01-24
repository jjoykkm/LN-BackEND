package cm_farm

import (
	"github.com/jjoykkm/ln-backend/common/models/model_other"
)

type Servicer interface {
	// status, uid string
	GetFarmList(status string, reqModel *model_other.ReqModel) (*model_other.RespModel, error)
	// status, uid string
	GetFarmAndFarmAreaList(status string, reqModel *model_other.ReqModel) (*model_other.RespModel, error)
}

type Service struct {
	repo Repositorier
}

func NewService(repo Repositorier) Servicer {
	return &Service{
		repo:  repo,
	}
}

func (s *Service) GetFarmList(status string, reqModel *model_other.ReqModel) (*model_other.RespModel, error) {
	myFarm, err := s.repo.FindAllMyFarm(status, reqModel.User.Uid)
	if err != nil{
		return nil, err
	}
	return &model_other.RespModel{
		Item: myFarm,
		Total: len(myFarm),
	}, nil
}

func (s *Service) GetFarmAndFarmAreaList(status string, reqModel *model_other.ReqModel) (*model_other.RespModel, error) {
	myFarm, err := s.repo.FindAllMyFarmAndFarmArea(status, reqModel.User.Uid)
	if err != nil{
		return nil, err
	}
	return &model_other.RespModel{
		Item: myFarm,
		Total: len(myFarm),
	}, nil
}