package model_databases

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/config"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table gender
//-------------------------------------------------------------------------------//
//model gender
type Gender struct {
	GenderId      	uuid.UUID	 `mapstructure:"role_id" json:"role_id,omitempty"`
	GenderName      string		 `mapstructure:"role_name" json:"role_name,omitempty"`
	CreateDate		time.Time	 `mapstructure:"create_date" json:"create_date,omitempty"`
	ChangeDate	    time.Time	 `mapstructure:"change_date" json:"change_date,omitempty"`
	StatusId		uuid.UUID	 `mapstructure:"status_id" json:"status_id,omitempty"`
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
