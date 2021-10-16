package sf_my_farm

import (
	"github.com/jjoykkm/ln-backend/common/config"
	"github.com/jjoykkm/ln-backend/common/models/model_other"
)

type Servicer interface {
	GetAuthorizeCheckForManageFarm(uid, farmId string) (bool, error)
	// FarmId
	GetOverviewFarm(status string, ReqModel *model_other.ReqModel) (*model_other.RespModel, error)
	// FarmId
	GetManageRole(status string, ReqModel *model_other.ReqModel) (*model_other.RespModel, error)
	// FarmId
	GetManageFarmArea(status string, ReqModel *model_other.ReqModel) (*model_other.RespModel, error)
	// FarmId
	GetManageMainbox(status string, ReqModel *model_other.ReqModel) (*model_other.RespModel, error)
}

type Service struct {
	repo  Repositorier
}

func NewService(repo Repositorier) Servicer {
	return &Service{
		repo:  repo,
	}
}

func (s *Service) GetAuthorizeCheckForManageFarm(uid, farmId string) (bool, error) {
	authManage := false
	trans, err := s.repo.FindOneTransManagement(uid, farmId)
	if err != nil{
		return false, err
	}
	if trans.RoleId.UUID.String() != config.GetRole().View {
		authManage = true
	}
	return authManage, err
}

func (s *Service) GetOverviewFarm(status string, ReqModel *model_other.ReqModel) (*model_other.RespModel, error) {
	farm, err := s.repo.FindOneFarm(status, ReqModel.FarmId)
	if err != nil{
		return nil, err
	}
	// Get Mainbox count
	farm.MainboxCount, err = s.repo.GetCountMainbox(status, ReqModel.FarmId)
	if err != nil{
		return nil, err
	}
	// Get Farm area count
	farm.FarmAreaCount, err = s.repo.GetCountFarmArea(status, ReqModel.FarmId)
	if err != nil{
		return nil, err
	}
	return &model_other.RespModel{
		Item: farm,
		Total: 1,
	}, nil
}

func (s *Service) GetManageRole(status string, ReqModel *model_other.ReqModel) (*model_other.RespModel, error) {
	// Check auth for edit
	//isAuth, err := Servicer.GetAuthorizeCheckForManageFarm(s, ReqModel.Uid, ReqModel.FarmId)
	//if err != nil{
	//	return nil, err
	//}
	//// No Auth
	//if isAuth != true {
	//	return nil, &errs.ErrContext{
	//		Code: ERROR_4002005,
	//		Err:  err,
	//		Msg:  MSG_NO_AUTH,
	//	}
	//}

	manageRole, err := s.repo.FindAllManageRole(status, ReqModel.FarmId)
	if err != nil{
		return nil, err
	}
	return &model_other.RespModel{
		Item: manageRole,
		Total: len(manageRole),
	}, nil
}

func (s *Service) GetManageFarmArea(status string, ReqModel *model_other.ReqModel) (*model_other.RespModel, error) {
	manageFarmArea, err := s.repo.FindAllManageFarmArea(status, ReqModel.FarmId)
	if err != nil{
		return nil, err
	}
	return &model_other.RespModel{
		Item: manageFarmArea,
		Total: len(manageFarmArea),
	}, nil
}

func (s *Service) GetManageMainbox(status string, ReqModel *model_other.ReqModel) (*model_other.RespModel, error) {
	manageFarmArea, err := s.repo.FindAllManageMainbox(status, ReqModel.FarmId)
	if err != nil{
		return nil, err
	}
	return &model_other.RespModel{
		Item: manageFarmArea,
		Total: len(manageFarmArea),
	}, nil
}
