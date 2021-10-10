package model_databases

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/common/config"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table address_farm
//-------------------------------------------------------------------------------//
//model address_farm
type AddressFarm struct {
	AddressFarmId		 uuid.UUID	 `mapstructure:"address_farm_id" json:"address_farm_id,omitempty"`
	HouseNo				 string  	 `mapstructure:"house_no" json:"house_no,omitempty"`
	Alley				 string  	 `mapstructure:"alley" json:"alley,omitempty"`
	Road				 string  	 `mapstructure:"road" json:"road,omitempty"`
	LocationX			 float64	 `mapstructure:"change_date" json:"change_date,omitempty"`
	LocationY			 float64	 `mapstructure:"change_date" json:"change_date,omitempty"`
	CreateDate		 	 time.Time	 `mapstructure:"create_date" json:"create_date,omitempty"`
	ChangeDate	     	 time.Time	 `mapstructure:"change_date" json:"change_date,omitempty"`
	StatusId			 uuid.UUID	 `mapstructure:"status_id" json:"status_id,omitempty"`
	SubDistrictId		 uuid.UUID	 `mapstructure:"sub_district_id" json:"sub_district_id,omitempty"`
	DistrictId			 uuid.UUID	 `mapstructure:"district_id" json:"district_id,omitempty"`
	ProvinceId			 uuid.UUID	 `mapstructure:"province_id" json:"province_id,omitempty"`
	CountryId			 uuid.UUID	 `mapstructure:"country_id" json:"country_id,omitempty"`
	FarmId				 uuid.UUID	 `mapstructure:"farm_id" json:"farm_id,omitempty"`
	Moo					 string  	 `mapstructure:"moo" json:"moo,omitempty"`
}
// New instance address_farm
func (u *AddressFarm) New() *AddressFarm {
	return &AddressFarm{
		AddressFarmId:      u.AddressFarmId ,
		HouseNo:            u.HouseNo ,
		Alley:              u.Alley ,
		Road:               u.Road ,
		LocationX:          u.LocationX ,
		LocationY:          u.LocationY ,
		CreateDate:         u.CreateDate ,
		ChangeDate:         u.ChangeDate ,
		StatusId:           u.StatusId ,
		SubDistrictId:      u.SubDistrictId ,
		DistrictId:         u.DistrictId ,
		ProvinceId:         u.ProvinceId ,
		CountryId:          u.CountryId ,
		FarmId:             u.FarmId ,
		Moo:                u.Moo ,
	}
}

// Custom table name for GORM
func (AddressFarm) TableName() string {
	return config.DB_ADDRESS_FARM
}


