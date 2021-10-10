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
	AddressFarmId		 uuid.UUID	 `json:"address_farm_id,omitempty"`
	HouseNo				 string  	 `json:"house_no,omitempty"`
	Alley				 string  	 `json:"alley,omitempty"`
	Road				 string  	 `json:"road,omitempty"`
	LocationX			 float64	 `json:"change_date,omitempty"`
	LocationY			 float64	 `json:"change_date,omitempty"`
	CreateDate		 	 time.Time	 `json:"create_date,omitempty"`
	ChangeDate	     	 time.Time	 `json:"change_date,omitempty"`
	StatusId			 uuid.UUID	 `json:"status_id,omitempty"`
	SubDistrictId		 uuid.UUID	 `json:"sub_district_id,omitempty"`
	DistrictId			 uuid.UUID	 `json:"district_id,omitempty"`
	ProvinceId			 uuid.UUID	 `json:"province_id,omitempty"`
	CountryId			 uuid.UUID	 `json:"country_id,omitempty"`
	FarmId				 uuid.UUID	 `json:"farm_id,omitempty"`
	Moo					 string  	 `json:"moo,omitempty"`
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


