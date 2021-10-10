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
	ScheduleId      uuid.UUID	 `json:"schedule_id,omitempty"`
	SocketId      	uuid.UUID	 `json:"socket_id,omitempty"`
	Uid	      		uuid.UUID	 `json:"uid,omitempty"`
	StatusSensorId	uuid.UUID	 `json:"status_sensor_id,omitempty"`
	IsManual    	bool		 `json:"is_manual,omitempty"`
	CreateDate		time.Time	 `json:"create_date,omitempty"`
	ChangeDate	    time.Time	 `json:"change_date,omitempty"`
	StatusId		uuid.UUID	 `json:"status_id,omitempty"`
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

