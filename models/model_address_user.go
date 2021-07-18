package models

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table address_user
//-------------------------------------------------------------------------------//
//model address_user
type AddressUser struct {
	AddressUserId	 	 uuid.UUID	 `gorm:"address_user_id" json:"address_user_id,omitempty"`
	HouseNo				 string  	 `gorm:"house_no" json:"house_no,omitempty"`
	Alley				 string  	 `gorm:"alley" json:"alley,omitempty"`
	Road				 string  	 `gorm:"road" json:"road,omitempty"`
	LocationX			 float64	 `gorm:"change_date" json:"change_date,omitempty"`
	LocationY			 float64	 `gorm:"change_date" json:"change_date,omitempty"`
	CreateDate		 	 time.Time	 `gorm:"create_date" json:"create_date,omitempty"`
	ChangeDate	     	 time.Time	 `gorm:"change_date" json:"change_date,omitempty"`
	StatusId			 uuid.UUID	 `gorm:"status_id" json:"status_id,omitempty"`
	SubDistrictId		 uuid.UUID	 `gorm:"sub_district_id" json:"sub_district_id,omitempty"`
	DistrictId			 uuid.UUID	 `gorm:"district_id" json:"district_id,omitempty"`
	ProvinceId			 uuid.UUID	 `gorm:"province_id" json:"province_id,omitempty"`
	CountryId			 uuid.UUID	 `gorm:"country_id" json:"country_id,omitempty"`
	Uid					 uuid.UUID	 `gorm:"uid" json:"uid,omitempty"`
	Moo					 string  	 `gorm:"moo" json:"moo,omitempty"`
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
