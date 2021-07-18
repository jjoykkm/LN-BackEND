package models

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table socket
//-------------------------------------------------------------------------------//
//model socket
type Socket struct {
	SocketId      	uuid.UUID	 `gorm:"socket_id" json:"socket_id,omitempty"`
	MainboxId      	uuid.UUID	 `gorm:"mainbox_id" json:"mainbox_id,omitempty"`
	SensorId      	uuid.UUID	 `gorm:"sensor_id" json:"sensor_id,omitempty"`
	StatusId		uuid.UUID	 `gorm:"status_id" json:"status_id,omitempty"`
	SocketName      string	 	 `gorm:"socket_name" json:"socket_name,omitempty"`
	CreateDate		time.Time	 `gorm:"create_date" json:"create_date,omitempty"`
	ChangeDate	    time.Time	 `gorm:"change_date" json:"change_date,omitempty"`
	StatusSensorId	uuid.UUID	 `gorm:"status_sensor_id" json:"status_sensor_id,omitempty"`
}
// New instance socket
func (u *Socket) New() *Socket {
	return &Socket{
		SocketId:		u.SocketId ,
		MainboxId:		u.MainboxId ,
		SensorId:		u.SensorId ,
		StatusId:		u.StatusId ,
		SocketName:		u.SocketName ,
		CreateDate:		u.CreateDate ,
		ChangeDate:		u.ChangeDate ,
		StatusSensorId:	u.StatusSensorId ,
	}
}
