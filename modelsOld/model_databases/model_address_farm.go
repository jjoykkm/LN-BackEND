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
	AddressFarmId		 uuid.UUID	 `json:"address_farm_id"`
	HouseNo				 string  	 `json:"house_no"`
	Alley				 string  	 `json:"alley"`
	Road				 string  	 `json:"road"`
	LocationX			 float64	 `json:"change_date"`
	LocationY			 float64	 `json:"change_date"`
	CreateDate		 	 time.Time	 `json:"create_date"`
	ChangeDate	     	 time.Time	 `json:"change_date"`
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


