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
	ProvinceId      uuid.UUID	 `json:"province_id"`
	ProvinceEN      string		 `json:"province_name_en"`
	ProvinceTH      string		 `json:"province_name_th"`
	CreateDate		time.Time	 `json:"create_date"`
	ChangeDate	    time.Time	 `json:"change_date"`
	StatusId		uuid.UUID	 `json:"status_id"`
	CountryId       uuid.UUID	 `json:"country_id"`
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

