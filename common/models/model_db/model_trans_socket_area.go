package model_db

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/common/config"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table trans_socket_area
//-------------------------------------------------------------------------------//
//model trans_socket_area
type TransSocketArea struct {
	FarmAreaId      uuid.UUID	 `json:"farm_area_id"`
	SocketId     	uuid.UUID	 `json:"socket_id"`
	CreateDate		time.Time	 `json:"create_date"`
	ChangeDate	    time.Time	 `json:"change_date"`
	StatusId		uuid.UUID	 `json:"status_id"`
}
// New instance trans_socket_area
func (u *TransSocketArea) New() *TransSocketArea {
	return &TransSocketArea{
		FarmAreaId:		u.FarmAreaId ,
		SocketId:		u.SocketId ,
		CreateDate:		u.CreateDate ,
		ChangeDate:		u.ChangeDate ,
		StatusId:		u.StatusId ,
	}
}

// Custom table name for GORM
func (TransSocketArea) TableName() string {
	return config.DB_TRANS_SOCKET_AREA
}
