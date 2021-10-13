package model_db

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/common/config"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table socket_action
//-------------------------------------------------------------------------------//
//model socket_action
type SocketAction struct {
	ScheduleId      uuid.UUID	 `json:"schedule_id"`
	SocketId      	uuid.UUID	 `json:"socket_id"`
	Uid	      		uuid.UUID	 `json:"uid"`
	StatusSensorId	uuid.UUID	 `json:"status_sensor_id"`
	IsManual    	bool		 `json:"is_manual"`
	CreateDate		time.Time	 `json:"create_date"`
	ChangeDate	    time.Time	 `json:"change_date"`
	StatusId		uuid.UUID	 `json:"status_id"`
}
// New instance socket_action
func (u *SocketAction) New() *SocketAction {
	return &SocketAction{
		ScheduleId:			u.ScheduleId ,
		SocketId:			u.SocketId ,
		Uid:				u.Uid ,
		StatusSensorId:		u.StatusSensorId ,
		IsManual:			u.IsManual ,
		CreateDate:			u.CreateDate ,
		ChangeDate:			u.ChangeDate ,
		StatusId:			u.StatusId ,
	}
}

// Custom table name for GORM
func (SocketAction) TableName() string {
	return config.DB_SOCKET_ACTION
}

