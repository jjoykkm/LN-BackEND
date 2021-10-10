package model_databases

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
	ProvinceId      uuid.UUID	 `mapstructure:"province_id" json:"province_id,omitempty"`
	ProvinceEN      string		 `mapstructure:"province_en" json:"province_en,omitempty"`
	ProvinceTH      string		 `mapstructure:"province_th" json:"province_th,omitempty"`
	CreateDate		time.Time	 `mapstructure:"create_date" json:"create_date,omitempty"`
	ChangeDate	    time.Time	 `mapstructure:"change_date" json:"change_date,omitempty"`
	StatusId		uuid.UUID	 `mapstructure:"status_id" json:"status_id,omitempty"`
	CountryId       uuid.UUID	 `mapstructure:"country_id" json:"country_id,omitempty"`
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

