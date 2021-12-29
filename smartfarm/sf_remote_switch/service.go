package sf_remote_switch

import (
	"github.com/jjoykkm/ln-backend/common/config"
	"github.com/jjoykkm/ln-backend/common/models/model_other"
)

type Servicer interface {
	ConfigRemoteSwitch(reqModel *RemoteDetailUS) error
	UnlinkSocketRemote(reqModel *RemoteDetailDel) error
	RemoveRemoteSwitch(reqModel *RemoteDetailDel) error
	GetRemoteSwitch(status string, reqModel *model_other.ReqModel) (*model_other.RespModel, error)
}

type Service struct {
	repo  Repositorier
}

func NewService(repo Repositorier) Servicer {
	return &Service{
		repo:  repo,
	}
}

func (s *Service) GetRemoteSwitch(status string, reqModel *model_other.ReqModel) (*model_other.RespModel, error) {
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

	remoteSwitch, err := s.repo.FindAllRemoteSwitch(status, reqModel.Uid)
	if err != nil{
		return nil, err
	}
	return &model_other.RespModel{
		Item: remoteSwitch,
		Total: len(remoteSwitch),
	}, nil
}

func (s *Service) ConfigRemoteSwitch(reqModel *RemoteDetailUS) error {
	// Prepare model before upsert data
	reqModel.RemoteSwitch.StatusId = config.GetStatus().Active //Assign status active

	tx := s.repo.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	// Upsert Remote Switch detail
	err := tx.UpsertRemoteSwitch(&reqModel.RemoteSwitch)
	if err != nil{
		return err
		tx.Rollback()
	}

	// Check array is not empty
	if len(reqModel.SocketList) > 0 {
		// Update Socket
		err = tx.UpdateSocketFieldRemote(reqModel)
		if err != nil{
			return err
			tx.Rollback()
		}
	}
	return tx.db.Commit().Error
}

func (s *Service) UnlinkSocketRemote(reqModel *RemoteDetailDel) error {
	// Update NULL to remote_id at socket
	if len(reqModel.SocketList) > 0{
		err := s.repo.UpdateNullSocketFieldRemote(reqModel.SocketList)
		if err != nil{
			return err
		}
	}
	return nil
}

func (s *Service) RemoveRemoteSwitch(reqModel *RemoteDetailDel) error {
	// Update Socket
	if reqModel.RemoteId != "" {
		err := s.repo.DeleteRemoteSwitch(&reqModel.RemoteId)
		if err != nil{
			return err
		}
	}
	return nil
}