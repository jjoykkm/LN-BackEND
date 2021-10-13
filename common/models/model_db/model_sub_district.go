package model_db

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/common/config"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table sub_district
//-------------------------------------------------------------------------------//
//model sub_district
type SubDistrict struct {
	SubDistrictId      uuid.UUID	 `json:"sub_district_id"`
	SubDistrictEN      string		 `json:"sub_district_name_en"`
	SubDistrictTH      string		 `json:"sub_district_name_th"`
	CreateDate		   time.Time	 `json:"create_date"`
	ChangeDate	       time.Time	 `json:"change_date"`
	StatusId		   uuid.UUID	 `json:"status_id"`
}
// New instance sub_district
func (u *SubDistrict) New() *SubDistrict {
	return &SubDistrict{
		SubDistrictId:		u.SubDistrictId ,
		SubDistrictEN:		u.SubDistrictEN ,
		SubDistrictTH:		u.SubDistrictTH ,
		CreateDate:			u.CreateDate ,
		ChangeDate:			u.ChangeDate ,
		StatusId:			u.StatusId ,
	}
}

// Custom table name for GORM
func (SubDistrict) TableName() string {
	return config.DB_SUB_DISTRICT
}

