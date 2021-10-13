package model_db

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
	GenderId      	uuid.UUID	 `json:"gender_id"`
	GenderName      string		 `json:"gender_name"`
	CreateDate		time.Time	 `json:"create_date"`
	ChangeDate	    time.Time	 `json:"change_date"`
	StatusId		uuid.UUID	 `json:"status_id"`
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
