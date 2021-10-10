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
	SocketId      	uuid.UUID	 `json:"SocketId,omitempty"`
	SensorId      	uuid.UUID	 `json:"SensorId,omitempty"`
	StatusId		uuid.UUID	 `json:"StatusId,omitempty"`
	SocketName      string	 	 `json:"SocketName,omitempty"`
	CreateDate		time.Time	 `json:"CreateDate,omitempty"`
	ChangeDate	    time.Time	 `json:"ChangeDate,omitempty"`
	SocketNumber	int64		 `json:"SocketNumber,omitempty"`
	StatusSensorId	uuid.UUID	 `json:"StatusSensorId,omitempty"`
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
