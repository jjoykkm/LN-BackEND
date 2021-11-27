package model_db

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/common/config"
)

//-------------------------------------------------------------------------------//
//							Table gender
//-------------------------------------------------------------------------------//
//model gender
type Gender struct {
	DBCommon
	GenderId      	uuid.UUID	 `json:"gender_id"`
	GenderName      string		 `json:"gender_name"`
}
// New instance gender
func (u *Gender) New() *Gender {
	return &Gender{
		DBCommon:      	u.DBCommon ,
		GenderId:		u.GenderId ,
		GenderName:		u.GenderName ,
	}
}

// Custom table name for GORM
func (Gender) TableName() string {
	return config.DB_GENDER
}
