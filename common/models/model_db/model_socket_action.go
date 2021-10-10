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
	ScheduleId      uuid.UUID	 `mapstructure:"schedule_id" json:"schedule_id,omitempty"`
	SocketId      	uuid.UUID	 `mapstructure:"socket_id" json:"socket_id,omitempty"`
	Uid	      		uuid.UUID	 `mapstructure:"uid" json:"uid,omitempty"`
	StatusSensorId	uuid.UUID	 `mapstructure:"status_sensor_id" json:"status_sensor_id,omitempty"`
	IsManual    	bool		 `mapstructure:"is_manual" json:"is_manual,omitempty"`
	CreateDate		time.Time	 `mapstructure:"create_date" json:"create_date,omitempty"`
	ChangeDate	    time.Time	 `mapstructure:"change_date" json:"change_date,omitempty"`
	StatusId		uuid.UUID	 `mapstructure:"status_id" json:"status_id,omitempty"`
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

