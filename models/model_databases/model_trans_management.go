package model_databases

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table trans_management
//-------------------------------------------------------------------------------//
//model trans_management
type TransManagement struct {
	Uid      		uuid.UUID	 `gorm:"uid" json:"uid,omitempty"`
	FarmId     		uuid.UUID	 `gorm:"farm_id" json:"farm_id,omitempty"`
	RoleId      	uuid.UUID	 `gorm:"role_id" json:"role_id,omitempty"`
	StatusId		uuid.UUID	 `gorm:"status_id" json:"status_id,omitempty"`
	CreateDate		time.Time	 `gorm:"create_date" json:"create_date,omitempty"`
	ChangeDate	    time.Time	 `gorm:"change_date" json:"change_date,omitempty"`
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
