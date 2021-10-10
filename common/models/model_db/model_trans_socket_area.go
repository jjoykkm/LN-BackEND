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
	FarmAreaId      uuid.UUID	 `mapstructure:"farm_area_id" json:"farm_area_id,omitempty"`
	SocketId     	uuid.UUID	 `mapstructure:"socket_id" json:"socket_id,omitempty"`
	CreateDate		time.Time	 `mapstructure:"create_date" json:"create_date,omitempty"`
	ChangeDate	    time.Time	 `mapstructure:"change_date" json:"change_date,omitempty"`
	StatusId		uuid.UUID	 `mapstructure:"status_id" json:"status_id,omitempty"`
	MainboxId     	uuid.UUID	 `mapstructure:"mainbox_id" json:"mainbox_id,omitempty"`
}
// New instance trans_socket_area
func (u *TransSocketArea) New() *TransSocketArea {
	return &TransSocketArea{
		FarmAreaId:		u.FarmAreaId ,
		SocketId:		u.SocketId ,
		CreateDate:		u.CreateDate ,
		ChangeDate:		u.ChangeDate ,
		StatusId:		u.StatusId ,
		MainboxId:		u.MainboxId ,
	}
}

// Custom table name for GORM
func (TransSocketArea) TableName() string {
	return config.DB_TRANS_SOCKET_AREA
}
