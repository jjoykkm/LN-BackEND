package models

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table socket_action
//-------------------------------------------------------------------------------//
//model socket_action
type SocketAction struct {
	ScheduleId      uuid.UUID	 `gorm:"schedule_id" json:"schedule_id,omitempty"`
	SocketId      	uuid.UUID	 `gorm:"socket_id" json:"socket_id,omitempty"`
	Uid	      		uuid.UUID	 `gorm:"uid" json:"uid,omitempty"`
	StatusSensorId	uuid.UUID	 `gorm:"status_sensor_id" json:"status_sensor_id,omitempty"`
	IsManual    	bool		 `gorm:"is_manual" json:"is_manual,omitempty"`
	CreateDate		time.Time	 `gorm:"create_date" json:"create_date,omitempty"`
	ChangeDate	    time.Time	 `gorm:"change_date" json:"change_date,omitempty"`
	StatusId		uuid.UUID	 `gorm:"status_id" json:"status_id,omitempty"`
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

