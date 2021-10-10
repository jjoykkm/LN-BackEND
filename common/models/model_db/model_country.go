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
	CountryId       uuid.UUID	 `json:"CountryId,omitempty"`
	CountryEN       string		 `json:"CountryEN,omitempty"`
	CountryTH       string		 `json:"CountryTH,omitempty"`
	CreateDate		time.Time	 `json:"CreateDate,omitempty"`
	ChangeDate	   	time.Time	 `json:"ChangeDate,omitempty"`
	StatusId		uuid.UUID	 `json:"StatusId,omitempty"`
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
