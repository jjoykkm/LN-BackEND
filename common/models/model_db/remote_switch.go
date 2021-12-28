package model_db

import (
	"github.com/jjoykkm/ln-backend/common/config"
)

//-------------------------------------------------------------------------------//
//							Table country
//-------------------------------------------------------------------------------//
//model country
type RemoteSwitch struct {
	DBCommonGet
	RemoteId       string	 	 `json:"remote_id" `
	RemoteName     string		 `json:"remote_name"`
}
// New instance country
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
