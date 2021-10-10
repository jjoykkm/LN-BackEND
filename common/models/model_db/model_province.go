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
	ProvinceId      uuid.UUID	 `json:"province_id,omitempty"`
	ProvinceEN      string		 `json:"province_en,omitempty"`
	ProvinceTH      string		 `json:"province_th,omitempty"`
	CreateDate		time.Time	 `json:"create_date,omitempty"`
	ChangeDate	    time.Time	 `json:"change_date,omitempty"`
	StatusId		uuid.UUID	 `json:"status_id,omitempty"`
	CountryId       uuid.UUID	 `json:"country_id,omitempty"`
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

