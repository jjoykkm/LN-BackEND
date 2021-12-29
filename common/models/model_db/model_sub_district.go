package model_db

import (
	"github.com/jjoykkm/ln-backend/common/config"
)

//-------------------------------------------------------------------------------//
//							Table sub_district
//-------------------------------------------------------------------------------//
//model sub_district
type SubDistrict struct {
	DBCommonGet
	SubDistrictId      string	 `json:"sub_district_id"`
	SubDistrictEN      string	 `json:"sub_district_name_en"`
	SubDistrictTH      string	 `json:"sub_district_name_th"`
}
// New instance sub_district
func (u *SubDistrict) New() *SubDistrict {
	return &SubDistrict{
		DBCommonGet:      	u.DBCommonGet ,
		SubDistrictId:		u.SubDistrictId ,
		SubDistrictEN:		u.SubDistrictEN ,
		SubDistrictTH:		u.SubDistrictTH ,
	}
}

// Custom table name for GORM
func (SubDistrict) TableName() string {
	return config.DB_SUB_DISTRICT
}

