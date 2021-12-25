package sf_dashboard

import (
	"github.com/jjoykkm/ln-backend/common/models/model_other"
)

type Servicer interface {
	//status, farmId string
	GetFarmAreaDashboardList(status string, reqModel *model_other.ReqModel) (*model_other.RespModel, error)
}

type Service struct {
	repo  Repositorier
}

func NewService(repo Repositorier) Servicer {
	return &Service{
		repo:  repo,
	}
}

func (s *Service) GetFarmAreaDashboardList(status string, reqModel *model_other.ReqModel) (*model_other.RespModel, error) {
	dashboardList, err := s.repo.GetFarmAreaDashboardList(status, reqModel.FarmId)
	if err != nil{
		return nil, err
	}
	return &model_other.RespModel{
		Item: dashboardList,
		Total: len(dashboardList),
	}, nil
}
