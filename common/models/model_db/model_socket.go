package model_db

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/common/config"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table socket
//-------------------------------------------------------------------------------//
//model socket
type Socket struct {
	SocketId      	uuid.UUID	 `json:"socket_id,omitempty"`
	SensorId      	uuid.UUID	 `json:"sensor_id,omitempty"`
	StatusId		uuid.UUID	 `json:"status_id,omitempty"`
	SocketName      string	 	 `json:"socket_name,omitempty"`
	CreateDate		time.Time	 `json:"create_date,omitempty"`
	ChangeDate	    time.Time	 `json:"change_date,omitempty"`
	SocketNumber	int64		 `json:"socket_number,omitempty"`
	StatusSensorId	uuid.UUID	 `json:"status_sensor_id,omitempty"`
}
// New instance socket
func (u *Socket) New() *Socket {
	return &Socket{
		SocketId:			u.SocketId ,
		SensorId:			u.SensorId ,
		StatusId:			u.StatusId ,
		SocketName:			u.SocketName ,
		CreateDate:			u.CreateDate ,
		ChangeDate:			u.ChangeDate ,
		SocketNumber:		u.SocketNumber ,
		StatusSensorId:		u.StatusSensorId ,
	}
}

// Custom table name for GORM
func (Socket) TableName() string {
	return config.DB_SOCKET
}
