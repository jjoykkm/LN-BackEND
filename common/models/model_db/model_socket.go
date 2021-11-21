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
	SocketId      	uuid.UUID	 `json:"socket_id gorm:"type:uuid;default:uuid_generate_v4()"`
	SensorId      	uuid.UUID	 `json:"sensor_id"`
	StatusId		uuid.UUID	 `json:"status_id"`
	SocketName      string	 	 `json:"socket_name"`
	CreateDate		time.Time	 `json:"create_date"`
	ChangeDate	    time.Time	 `json:"change_date"`
	SocketNumber	int64		 `json:"socket_number"`
	StatusSensorId	uuid.UUID	 `json:"status_sensor_id"`
	MainboxId     	uuid.UUID	 `json:"mainbox_id"`
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
		MainboxId:			u.MainboxId ,
	}
}

// Custom table name for GORM
func (Socket) TableName() string {
	return config.DB_SOCKET
}

//-------------------------------------------------------------------------------//
//							Request Socket
//-------------------------------------------------------------------------------//
type SocketUS struct {
	SensorId      	string	 `json:"sensor_id"`
	StatusId		string	 `json:"status_id"`
	SocketName      string	 `json:"socket_name"`
	SocketNumber	int64	 `json:"socket_number"`
	StatusSensorId	string	 `json:"status_sensor_id"`
	MainboxId     	string	 `json:"mainbox_id"`
}
func (SocketUS) TableName() string {
	return config.DB_SOCKET
}