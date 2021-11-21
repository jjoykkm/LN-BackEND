package sf_my_farm

import (
	"errors"
	"github.com/jjoykkm/ln-backend/common/config"
	"github.com/jjoykkm/ln-backend/common/models/model_db"
	"github.com/jjoykkm/ln-backend/common/models/model_other"
	"github.com/jjoykkm/ln-backend/errs"
	"gorm.io/gorm"
)

type Servicer interface {
	GetAuthorizeCheckForManageFarm(uid, farmId string) (bool, error)
	// FarmId
	GetOverviewFarm(status string, reqModel *model_other.ReqModel) (*model_other.RespModel, error)
	// FarmId
	GetManageRole(status string, reqModel *model_other.ReqModel) (*model_other.RespModel, error)
	// FarmId
	GetManageFarmArea(status string, reqModel *model_other.ReqModel) (*model_other.RespModel, error)
	// FarmId
	GetManageMainbox(status string, reqModel *model_other.ReqModel) (*model_other.RespModel, error)

	CheckMainboxIsInactivated(serialNo string) (bool, error)
	ActivateMainbox(reqModel *model_db.MainboxSerialUS) error
	ConfigMainbox(reqModel *ReqConfMainbox) error
	ConfigAddSensor(reqModel *ReqConfMainbox) error
	ConfigDeleteSocket(reqModel *ReqDeleteConfig) error
	ConfigDeleteMainbox(reqModel *ReqDeleteConfig) error
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

func (s *Service) GetOverviewFarm(status string, reqModel *model_other.ReqModel) (*model_other.RespModel, error) {
	farm, err := s.repo.FindOneFarm(status, reqModel.FarmId)
	if err != nil{
		return nil, err
	}
	// Get Mainbox count
	farm.MainboxCount, err = s.repo.GetCountMainbox(status, reqModel.FarmId)
	if err != nil{
		return nil, err
	}
	// Get Farm area count
	farm.FarmAreaCount, err = s.repo.GetCountFarmArea(status, reqModel.FarmId)
	if err != nil{
		return nil, err
	}
	return &model_other.RespModel{
		Item: farm,
		Total: 1,
	}, nil
}

func (s *Service) GetManageRole(status string, reqModel *model_other.ReqModel) (*model_other.RespModel, error) {
	// Check auth for edit
	//isAuth, err := Servicer.GetAuthorizeCheckForManageFarm(s, reqModel.Uid, reqModel.FarmId)
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

	manageRole, err := s.repo.FindAllManageRole(status, reqModel.FarmId)
	if err != nil{
		return nil, err
	}
	return &model_other.RespModel{
		Item: manageRole,
		Total: len(manageRole),
	}, nil
}

func (s *Service) GetManageFarmArea(status string, reqModel *model_other.ReqModel) (*model_other.RespModel, error) {
	manageFarmArea, err := s.repo.FindAllManageFarmArea(status, reqModel.FarmId)
	if err != nil{
		return nil, err
	}
	return &model_other.RespModel{
		Item: manageFarmArea,
		Total: len(manageFarmArea),
	}, nil
}

func (s *Service) GetManageMainbox(status string, reqModel *model_other.ReqModel) (*model_other.RespModel, error) {
	manageFarmArea, err := s.repo.FindAllManageMainbox(status, reqModel.FarmId)
	if err != nil{
		return nil, err
	}
	return &model_other.RespModel{
		Item: manageFarmArea,
		Total: len(manageFarmArea),
	}, nil
}

func (s *Service) CheckMainboxIsInactivated(serialNo string) (bool, error) {
	mainbox, err := s.repo.FindOneMainboxBySerialNo(serialNo)
	if err != nil{
		// Check serial no has been found in DB
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, &errs.ErrContext{
					Code: ERROR_4001002,
					Err:  err,
					Msg:  MSG_WRONG_MB,
				}
		}else {
			return false, err
		}
	}
	// Check serial no is inactive
	if mainbox.StatusId.UUID.String() == config.GetStatus().Inactive {
		return true, nil
	}else {
		return false, &errs.ErrContext{
			Code: ERROR_4001001,
			Err:  err,
			Msg:  MSG_DUP_MB,
		}
	}
}
//-------------------------------------------------------------------------------//
//							Update data
//-------------------------------------------------------------------------------//
func (s *Service) ActivateMainbox(reqModel *model_db.MainboxSerialUS) error {
	isInactive, err := s.CheckMainboxIsInactivated(reqModel.MainboxSerialNo)
	if !isInactive || err != nil {
		return err
	}
	err = s.repo.UpdateOneMainboxBySerialNo(reqModel)
	if err != nil{
		return err
	}
	return nil
}

func (s *Service) ConfigMainbox(reqModel *ReqConfMainbox) error {
	// Update Mainbox detail
	if !((model_db.MainboxUS{}) == *reqModel.Mainbox) {	// Check model has value
		err := s.repo.UpdateOneMainbox(reqModel.Mainbox)
		if err != nil{
			return err
		}
	}
	// Upsert Socket
	err := s.repo.UpsertSocket(reqModel.Socket)
	if err != nil{
		return err
	}
	return nil
}

func (s *Service) ConfigAddSensor(reqModel *ReqConfMainbox) error {
	// Create Sensor
	err := s.repo.CreateOneSensor(reqModel.Sensor)
	if err != nil{
		return err
	}
	return nil
}

func (s *Service) ConfigDeleteSocket(reqModel *ReqDeleteConfig) error {
	// Delete Socket
	err := s.repo.DeleteOneSocket(reqModel.SocketId)
	if err != nil{
		return err
	}
	return nil
}

func (s *Service) ConfigDeleteMainbox(reqModel *ReqDeleteConfig) error {
	// Deactivate Mainbox
	err := s.repo.DeactivateOneMainbox(reqModel.MainboxId)
	if err != nil{
		return err
	}
	return nil
}