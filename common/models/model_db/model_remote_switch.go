package model_db

import (
	"github.com/jjoykkm/ln-backend/common/config"
)

//-------------------------------------------------------------------------------//
//							Table RemoteSwitch
//-------------------------------------------------------------------------------//
//model
type RemoteSwitch struct {
	DBCommonGet
	RemoteId       string	 	 `json:"remote_id"`
	RemoteName     string	 	 `json:"remote_name"`
}
// New instance
func (u *RemoteSwitch) New() *RemoteSwitch {
	return &RemoteSwitch{
		DBCommonGet:    u.DBCommonGet ,
		RemoteId:		u.RemoteId ,
		RemoteName:		u.RemoteName ,
	}
}

// Custom table name for GORM
func (RemoteSwitch) TableName() string {
	return config.DB_REMOTE_SWITCH
}

//-------------------------------------------------------------------------------//
//									Upsert
//-------------------------------------------------------------------------------//
type RemoteSwitchUS struct {
	DBCommonCreateUpdate
	RemoteId      string	 `json:"remote_id" gorm:"default:uuid_generate_v4()"`
	RemoteName    string	 `json:"remote_name"`
	StatusId	  string	 `json:"status_id"`
}
func (RemoteSwitchUS) TableName() string {
	return config.DB_REMOTE_SWITCH
}
