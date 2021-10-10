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
	SubDistrictId      uuid.UUID	 `mapstructure:"sub_district_id" json:"sub_district_id,omitempty"`
	SubDistrictEN      string		 `mapstructure:"sub_district_en" json:"sub_district_en,omitempty"`
	SubDistrictTH      string		 `mapstructure:"sub_district_th" json:"sub_district_th,omitempty"`
	CreateDate		   time.Time	 `mapstructure:"create_date" json:"create_date,omitempty"`
	ChangeDate	       time.Time	 `mapstructure:"change_date" json:"change_date,omitempty"`
	StatusId		   uuid.UUID	 `mapstructure:"status_id" json:"status_id,omitempty"`
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

