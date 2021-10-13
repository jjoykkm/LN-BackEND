package model_db

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/common/config"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table trans_management
//-------------------------------------------------------------------------------//
//model trans_management
type TransManagement struct {
	Uid      		uuid.UUID	 `json:"uid"`
	FarmId     		uuid.UUID	 `json:"farm_id"`
	RoleId      	uuid.UUID	 `json:"role_id"`
	StatusId		uuid.UUID	 `json:"status_id"`
	CreateDate		time.Time	 `json:"create_date"`
	ChangeDate	    time.Time	 `json:"change_date"`
}
// New instance trans_management
func (u *TransManagement) New() *TransManagement {
	return &TransManagement{
		Uid:			u.Uid ,
		FarmId:			u.FarmId ,
		RoleId:			u.RoleId ,
		StatusId:		u.StatusId ,
		CreateDate:		u.CreateDate ,
		ChangeDate:		u.ChangeDate ,
	}
}

// Custom table name for GORM
func (TransManagement) TableName() string {
	return config.DB_TRANS_MANAGEMENT
}
