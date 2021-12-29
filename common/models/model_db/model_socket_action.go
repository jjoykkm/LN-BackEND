package model_db

import (
	"github.com/jjoykkm/ln-backend/common/config"
)

//-------------------------------------------------------------------------------//
//							Table socket_action
//-------------------------------------------------------------------------------//
//model socket_action
type SocketAction struct {
	DBCommonGet
	ScheduleId      string	 `json:"schedule_id"`
	SocketId      	string	 `json:"socket_id"`
	Uid	      		string	 `json:"uid"`
	StatusSensorId	string	 `json:"status_sensor_id"`
	IsManual    	bool	 `json:"is_manual"`
}
// New instance socket_action
func (u *SocketAction) New() *SocketAction {
	return &SocketAction{
		DBCommonGet:      	u.DBCommonGet ,
		ScheduleId:			u.ScheduleId ,
		SocketId:			u.SocketId ,
		Uid:				u.Uid ,
		StatusSensorId:		u.StatusSensorId ,
		IsManual:			u.IsManual ,
	}
}

// Custom table name for GORM
func (SocketAction) TableName() string {
	return config.DB_SOCKET_ACTION
}

