package model_databases

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/config"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table socket
//-------------------------------------------------------------------------------//
//model socket
type Socket struct {
	SocketId      	uuid.UUID	 `mapstructure:"socket_id" json:"socket_id,omitempty"`
	SensorId      	uuid.UUID	 `mapstructure:"sensor_id" json:"sensor_id,omitempty"`
	StatusId		uuid.UUID	 `mapstructure:"status_id" json:"status_id,omitempty"`
	SocketName      string	 	 `mapstructure:"socket_name" json:"socket_name,omitempty"`
	CreateDate		time.Time	 `mapstructure:"create_date" json:"create_date,omitempty"`
	ChangeDate	    time.Time	 `mapstructure:"change_date" json:"change_date,omitempty"`
	SocketNumber	int64		 `mapstructure:"socket_number" json:"socket_number,omitempty"`
	StatusSensorId	uuid.UUID	 `mapstructure:"status_sensor_id" json:"status_sensor_id,omitempty"`
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
