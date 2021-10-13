package model_db

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/common/config"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table country
//-------------------------------------------------------------------------------//
//model country
type Country struct {
	CountryId       uuid.UUID	 `json:"country_id"`
	CountryEN       string		 `json:"country_name_en"`
	CountryTH       string		 `json:"country_name_th"`
	CreateDate		time.Time	 `json:"create_date"`
	ChangeDate	   	time.Time	 `json:"change_date"`
	StatusId		uuid.UUID	 `json:"status_id"`
}
// New instance country
func (u *Country) New() *Country {
	return &Country{
		CountryId:		u.CountryId ,
		CountryEN:		u.CountryEN ,
		CountryTH:		u.CountryTH ,
		CreateDate:		u.CreateDate ,
		ChangeDate:		u.ChangeDate ,
		StatusId:		u.StatusId ,
	}
}

// Custom table name for GORM
func (Country) TableName() string {
	return config.DB_COUNTRY
}
