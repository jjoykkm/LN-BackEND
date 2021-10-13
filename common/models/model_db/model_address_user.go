package model_db

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/common/config"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table address_user
//-------------------------------------------------------------------------------//
//model address_user
type AddressUser struct {
	AddressUserId	 	 uuid.UUID	 `json:"address_user_id"`
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
	Uid					 uuid.UUID	 `json:"uid"`
	Moo					 string  	 `json:"moo"`
}
// New instance address_user
func (u *AddressUser) New() *AddressUser {
	return &AddressUser{
		AddressUserId:      u.AddressUserId ,
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
		Uid:             	u.Uid ,
		Moo:                u.Moo ,
	}
}

// Custom table name for GORM
func (AddressUser) TableName() string {
	return config.DB_ADDRESS_USER
}

