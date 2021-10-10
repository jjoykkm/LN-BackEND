package model_databases

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/common/config"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table gender
//-------------------------------------------------------------------------------//
//model gender
type Gender struct {
	GenderId      	uuid.UUID	 `json:"role_id,omitempty"`
	GenderName      string		 `json:"role_name,omitempty"`
	CreateDate		time.Time	 `json:"create_date,omitempty"`
	ChangeDate	    time.Time	 `json:"change_date,omitempty"`
	StatusId		uuid.UUID	 `json:"status_id,omitempty"`
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

// Custom table name for GORM
func (Gender) TableName() string {
	return config.DB_GENDER
}
