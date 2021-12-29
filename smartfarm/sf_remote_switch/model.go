package sf_remote_switch

import (
	"github.com/jjoykkm/ln-backend/common/config"
	"github.com/jjoykkm/ln-backend/common/models/model_db"
)

//-------------------------------------------------------------------------------//
//									Upsert
//-------------------------------------------------------------------------------//
type RemoteDetailUS struct {
	RemoteSwitch	model_db.RemoteSwitchUS 	`json:"remote_switch" gorm:"embedded"`
	SocketList		[]string					`json:"socket_list"`
}
func (RemoteDetailUS) TableName() string {
	return config.DB_REMOTE_SWITCH
}

//-------------------------------------------------------------------------------//
//									Delete
//-------------------------------------------------------------------------------//
type RemoteDetailDel struct {
	RemoteId       string	 	 `json:"remote_id"`
	SocketList	   []string		 `json:"socket_list"`
}
func (RemoteDetailDel) TableName() string {
	return config.DB_REMOTE_SWITCH
}



