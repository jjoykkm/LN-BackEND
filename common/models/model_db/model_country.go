package model_db

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/common/config"
)

//-------------------------------------------------------------------------------//
//							Table country
//-------------------------------------------------------------------------------//
//model country
type Country struct {
	DBCommonGet
	CountryId       uuid.UUID	 `json:"country_id"`
	CountryEN       string		 `json:"country_name_en"`
	CountryTH       string		 `json:"country_name_th"`
}
// New instance country
func (u *Country) New() *Country {
	return &Country{
		DBCommonGet:      	u.DBCommonGet ,
		CountryId:		u.CountryId ,
		CountryEN:		u.CountryEN ,
		CountryTH:		u.CountryTH ,
	}
}

// Custom table name for GORM
func (Country) TableName() string {
	return config.DB_COUNTRY
}
