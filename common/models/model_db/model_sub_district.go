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
	SubDistrictId      uuid.UUID	 `json:"SubDistrictId,omitempty"`
	SubDistrictEN      string		 `json:"SubDistrictEN,omitempty"`
	SubDistrictTH      string		 `json:"SubDistrictTH,omitempty"`
	CreateDate		   time.Time	 `json:"CreateDate,omitempty"`
	ChangeDate	       time.Time	 `json:"ChangeDate,omitempty"`
	StatusId		   uuid.UUID	 `json:"StatusId,omitempty"`
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

