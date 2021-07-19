package model_databases

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table trans_socket_area
//-------------------------------------------------------------------------------//
//model trans_socket_area
type TransSocketArea struct {
	FarmAreaId      uuid.UUID	 `gorm:"farm_area_id" json:"farm_area_id,omitempty"`
	SocketId     	string	 	 `gorm:"socket_id" json:"socket_id,omitempty"`
	CreateDate		time.Time	 `gorm:"create_date" json:"create_date,omitempty"`
	ChangeDate	    time.Time	 `gorm:"change_date" json:"change_date,omitempty"`
	StatusId		uuid.UUID	 `gorm:"status_id" json:"status_id,omitempty"`
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
