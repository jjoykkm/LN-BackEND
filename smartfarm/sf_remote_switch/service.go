package sf_remote_switch

import (
	"github.com/jjoykkm/ln-backend/common/config"
)

type Servicer interface {
	ConfigRemoteSwitch(reqModel *RemoteDetailUS) error
	UnlinkSocketRemote(reqModel *RemoteDetailDel) error
	RemoveRemoteSwitch(reqModel *RemoteDetailDel) error
}

type Service struct {
	repo  Repositorier
}

func NewService(repo Repositorier) Servicer {
	return &Service{
		repo:  repo,
	}
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