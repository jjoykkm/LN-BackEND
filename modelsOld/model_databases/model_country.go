package model_databases

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
	CountryId       uuid.UUID	 `json:"country_id,omitempty"`
	CountryEN       string		 `json:"country_en,omitempty"`
	CountryTH       string		 `json:"country_th,omitempty"`
	CreateDate		time.Time	 `json:"create_date,omitempty"`
	ChangeDate	   	time.Time	 `json:"change_date,omitempty"`
	StatusId		uuid.UUID	 `json:"status_id,omitempty"`
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
