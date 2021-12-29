package model_db

import (
	"github.com/jjoykkm/ln-backend/common/config"
)

//-------------------------------------------------------------------------------//
//							Table address_farm
//-------------------------------------------------------------------------------//
//model address_farm
type AddressFarm struct {
	DBCommonGet
	AddressFarmId		 string	 	 `json:"address_farm_id"`
	HouseNo				 string  	 `json:"house_no"`
	Alley				 string  	 `json:"alley"`
	Road				 string  	 `json:"road"`
	LocationX			 float64	 `json:"location_x"`
	LocationY			 float64	 `json:"location_y"`
	StatusId			 string	 	 `json:"status_id"`
	SubDistrictId		 string	 	 `json:"sub_district_id"`
	DistrictId			 string	 	 `json:"district_id"`
	ProvinceId			 string	 	 `json:"province_id"`
	CountryId			 string	 	 `json:"country_id"`
	FarmId				 string	 	 `json:"farm_id"`
	Moo					 string  	 `json:"moo"`
}
// New instance address_farm
func (u *AddressFarm) New() *AddressFarm {
	return &AddressFarm{
		DBCommonGet:      	u.DBCommonGet ,
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


