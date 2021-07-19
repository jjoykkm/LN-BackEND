package model_databases

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table gender
//-------------------------------------------------------------------------------//
//model gender
type Gender struct {
	GenderId      	uuid.UUID	 `gorm:"role_id" json:"role_id,omitempty"`
	GenderName      string		 `gorm:"role_name" json:"role_name,omitempty"`
	CreateDate		time.Time	 `gorm:"create_date" json:"create_date,omitempty"`
	ChangeDate	    time.Time	 `gorm:"change_date" json:"change_date,omitempty"`
	StatusId		uuid.UUID	 `gorm:"status_id" json:"status_id,omitempty"`
}
// New instance gender
func (u *Gender) New() *Gender {
	return &Gender{
		GenderId:		u.GenderId ,
		GenderName:		u.GenderName ,
		CreateDate:		u.CreateDate ,
		ChangeDate:		u.ChangeDate ,
		StatusId:		u.StatusId ,
	}
}
