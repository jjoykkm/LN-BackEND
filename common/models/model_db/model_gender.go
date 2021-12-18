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
	DBCommonGet
	GenderId      	uuid.UUID	 `json:"gender_id"`
	GenderName      string		 `json:"gender_name"`
}
// New instance gender
func (u *Gender) New() *Gender {
	return &Gender{
		DBCommonGet:      	u.DBCommonGet ,
		GenderId:		u.GenderId ,
		GenderName:		u.GenderName ,
	}
}

// Custom table name for GORM
func (Gender) TableName() string {
	return config.DB_GENDER
}
