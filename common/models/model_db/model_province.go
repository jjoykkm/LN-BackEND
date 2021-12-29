package model_db

import (
	"github.com/jjoykkm/ln-backend/common/config"
)

//-------------------------------------------------------------------------------//
//							Table province
//-------------------------------------------------------------------------------//
//model province
type Province struct {
	DBCommonGet
	ProvinceId      string	 `json:"province_id"`
	ProvinceEN      string	 `json:"province_name_en"`
	ProvinceTH      string	 `json:"province_name_th"`
	CountryId       string	 `json:"country_id"`
}
// New instance province
func (u *Province) New() *Province {
	return &Province{
		DBCommonGet:      	u.DBCommonGet ,
		ProvinceId:		u.ProvinceId ,
		ProvinceEN:		u.ProvinceEN ,
		ProvinceTH:		u.ProvinceTH ,
		CountryId:		u.CountryId ,
	}
}

// Custom table name for GORM
func (Province) TableName() string {
	return config.DB_PROVINCE
}

