package model_databases

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
	Uid      		uuid.UUID	 `json:"uid,omitempty"`
	FarmId     		uuid.UUID	 `json:"farm_id,omitempty"`
	RoleId      	uuid.UUID	 `json:"role_id,omitempty"`
	StatusId		uuid.UUID	 `json:"status_id,omitempty"`
	CreateDate		time.Time	 `json:"create_date,omitempty"`
	ChangeDate	    time.Time	 `json:"change_date,omitempty"`
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
