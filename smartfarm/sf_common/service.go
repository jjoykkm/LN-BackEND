package sf_common

import (
	"github.com/jjoykkm/ln-backend/common/models/model_other"
)

type Servicer interface {
	// status, uid string
	GetFarmList(status string, ReqModel *model_other.ReqModel) (*model_other.RespModel, error)
	// status, uid string
	GetFarmAndFarmAreaList(status string, ReqModel *model_other.ReqModel) (*model_other.RespModel, error)
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

func (s *Service) GetFarmList(status string, ReqModel *model_other.ReqModel) (*model_other.RespModel, error) {
	myFarm, err := s.repo.FindAllMyFarm(status, ReqModel.Uid)
	if err != nil{
		return nil, err
	}
	return &model_other.RespModel{
		Item: myFarm,
		Total: len(myFarm),
	}, nil
}

func (s *Service) GetFarmAndFarmAreaList(status string, ReqModel *model_other.ReqModel) (*model_other.RespModel, error) {
	myFarm, err := s.repo.FindAllMyFarmAndFarmArea(status, ReqModel.Uid)
	if err != nil{
		return nil, err
	}
	return &model_other.RespModel{
		Item: myFarm,
		Total: len(myFarm),
	}, nil
}

func (s *Service) GetProvinceList(status string) (*model_other.RespModel, error) {
	myFarm, err := s.repo.FindAllProvince(status)
	if err != nil{
		return nil, err
	}
	return &model_other.RespModel{
		Item: myFarm,
		Total: len(myFarm),
	}, nil
}