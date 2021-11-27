package model_db

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/common/config"
)

//-------------------------------------------------------------------------------//
//							Table address_farm
//-------------------------------------------------------------------------------//
//model address_farm
type AddressFarm struct {
	DBCommon
	AddressFarmId		 uuid.UUID	 `json:"address_farm_id"`
	HouseNo				 string  	 `json:"house_no"`
	Alley				 string  	 `json:"alley"`
	Road				 string  	 `json:"road"`
	LocationX			 float64	 `json:"location_x"`
	LocationY			 float64	 `json:"location_y"`
	StatusId			 uuid.UUID	 `json:"status_id"`
	SubDistrictId		 uuid.UUID	 `json:"sub_district_id"`
	DistrictId			 uuid.UUID	 `json:"district_id"`
	ProvinceId			 uuid.UUID	 `json:"province_id"`
	CountryId			 uuid.UUID	 `json:"country_id"`
	FarmId				 uuid.UUID	 `json:"farm_id"`
	Moo					 string  	 `json:"moo"`
}
// New instance address_farm
func (u *AddressFarm) New() *AddressFarm {
	return &AddressFarm{
		DBCommon:      		u.DBCommon ,
		AddressFarmId:      u.AddressFarmId ,
		HouseNo:            u.HouseNo ,
		Alley:              u.Alley ,
		Road:               u.Road ,
		LocationX:          u.LocationX ,
		LocationY:          u.LocationY ,
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


