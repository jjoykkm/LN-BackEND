package model_db

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/common/config"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table province
//-------------------------------------------------------------------------------//
//model province
type Province struct {
	ProvinceId      uuid.UUID	 `json:"ProvinceId,omitempty"`
	ProvinceEN      string		 `json:"ProvinceEN,omitempty"`
	ProvinceTH      string		 `json:"ProvinceTH,omitempty"`
	CreateDate		time.Time	 `json:"CreateDate,omitempty"`
	ChangeDate	    time.Time	 `json:"ChangeDate,omitempty"`
	StatusId		uuid.UUID	 `json:"StatusId,omitempty"`
	CountryId       uuid.UUID	 `json:"CountryId,omitempty"`
}
// New instance province
func (u *Province) New() *Province {
	return &Province{
		ProvinceId:		u.ProvinceId ,
		ProvinceEN:		u.ProvinceEN ,
		ProvinceTH:		u.ProvinceTH ,
		CreateDate:		u.CreateDate ,
		ChangeDate:		u.ChangeDate ,
		StatusId:		u.StatusId ,
		CountryId:		u.CountryId ,
	}
}

// Custom table name for GORM
func (Province) TableName() string {
	return config.DB_PROVINCE
}

