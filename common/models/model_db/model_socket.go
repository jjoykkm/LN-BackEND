package model_db

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/common/config"
)

//-------------------------------------------------------------------------------//
//							Table socket
//-------------------------------------------------------------------------------//
//model socket
type Socket struct {
	DBCommon
	SocketId      	uuid.UUID	 `json:"socket_id gorm:"default:uuid_generate_v4()"`
	SensorId      	uuid.UUID	 `json:"sensor_id"`
	SocketName      string	 	 `json:"socket_name"`
	SocketNumber	int64		 `json:"socket_number"`
	StatusSensorId	uuid.UUID	 `json:"status_sensor_id"`
	MainboxId     	uuid.UUID	 `json:"mainbox_id"`
	FarmAreaId      uuid.UUID	 `json:"farm_area_id"`
}
// New instance socket
func (u *Socket) New() *Socket {
	return &Socket{
		DBCommon:      		u.DBCommon ,
		SocketId:			u.SocketId ,
		SensorId:			u.SensorId ,
		SocketName:			u.SocketName ,
		SocketNumber:		u.SocketNumber ,
		StatusSensorId:		u.StatusSensorId ,
		MainboxId:			u.MainboxId ,
		FarmAreaId:			u.FarmAreaId ,
	}
}

// Custom table name for GORM
func (Socket) TableName() string {
	return config.DB_SOCKET
}

//-------------------------------------------------------------------------------//
//							Upsert Socket
//-------------------------------------------------------------------------------//
type SocketUS struct {
	SocketId		string		 `json:"socket_id" gorm:"default:uuid_generate_v4()"`
	SensorId      	string	 	 `json:"sensor_id"`
	StatusId		string	 	 `json:"status_id"`
	SocketName      string	 	 `json:"socket_name"`
	SocketNumber	int64	 	 `json:"socket_number"`
	StatusSensorId	string	 	 `json:"status_sensor_id"`
	MainboxId     	string	 	 `json:"mainbox_id"`
	FarmAreaId      uuid.UUID	 `json:"farm_area_id"`
}
func (SocketUS) TableName() string {
	return config.DB_SOCKET
}