package model_databases

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
	SubDistrictId      uuid.UUID	 `json:"sub_district_id,omitempty"`
	SubDistrictEN      string		 `json:"sub_district_en,omitempty"`
	SubDistrictTH      string		 `json:"sub_district_th,omitempty"`
	CreateDate		   time.Time	 `json:"create_date,omitempty"`
	ChangeDate	       time.Time	 `json:"change_date,omitempty"`
	StatusId		   uuid.UUID	 `json:"status_id,omitempty"`
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

