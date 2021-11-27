package model_db

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/common/config"
)

//-------------------------------------------------------------------------------//
//							Table sub_district
//-------------------------------------------------------------------------------//
//model sub_district
type SubDistrict struct {
	DBCommon
	SubDistrictId      uuid.UUID	 `json:"sub_district_id"`
	SubDistrictEN      string		 `json:"sub_district_name_en"`
	SubDistrictTH      string		 `json:"sub_district_name_th"`
}
// New instance sub_district
func (u *SubDistrict) New() *SubDistrict {
	return &SubDistrict{
		DBCommon:      		u.DBCommon ,
		SubDistrictId:		u.SubDistrictId ,
		SubDistrictEN:		u.SubDistrictEN ,
		SubDistrictTH:		u.SubDistrictTH ,
	}
}

// Custom table name for GORM
func (SubDistrict) TableName() string {
	return config.DB_SUB_DISTRICT
}

