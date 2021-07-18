package models

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table district
//-------------------------------------------------------------------------------//
//model district
type District struct {
	DistrictId      uuid.UUID	 `gorm:"district_id" json:"district_id,omitempty"`
	DistrictEN      string		 `gorm:"district_en" json:"district_en,omitempty"`
	DistrictTH      string		 `gorm:"district_th" json:"district_th,omitempty"`
	CreateDate		 time.Time	 `gorm:"create_date" json:"create_date,omitempty"`
	ChangeDate	     time.Time	 `gorm:"change_date" json:"change_date,omitempty"`
	StatusId		 uuid.UUID	 `gorm:"status_id" json:"status_id,omitempty"`
}
// New instance district
func (u *District) New() *District {
	return &District{
		DistrictId:		u.DistrictId ,
		DistrictEN:		u.DistrictEN ,
		DistrictTH:		u.DistrictTH ,
		CreateDate:		u.CreateDate ,
		ChangeDate:		u.ChangeDate ,
		StatusId:		u.StatusId ,
	}
}

