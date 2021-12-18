package model_db

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/common/config"
)

//-------------------------------------------------------------------------------//
//							Table socket_action
//-------------------------------------------------------------------------------//
//model socket_action
type SocketAction struct {
	DBCommonGet
	ScheduleId      uuid.UUID	 `json:"schedule_id"`
	SocketId      	uuid.UUID	 `json:"socket_id"`
	Uid	      		uuid.UUID	 `json:"uid"`
	StatusSensorId	uuid.UUID	 `json:"status_sensor_id"`
	IsManual    	bool		 `json:"is_manual"`
}
// New instance socket_action
func (u *SocketAction) New() *SocketAction {
	return &SocketAction{
		DBCommonGet:      		u.DBCommonGet ,
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

