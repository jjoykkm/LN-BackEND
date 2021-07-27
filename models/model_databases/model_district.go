package model_databases

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table district
//-------------------------------------------------------------------------------//
//model district
type District struct {
	DistrictId      uuid.UUID	 `mapstructure:"district_id" json:"district_id,omitempty"`
	DistrictEN      string		 `mapstructure:"district_en" json:"district_en,omitempty"`
	DistrictTH      string		 `mapstructure:"district_th" json:"district_th,omitempty"`
	CreateDate		 time.Time	 `mapstructure:"create_date" json:"create_date,omitempty"`
	ChangeDate	     time.Time	 `mapstructure:"change_date" json:"change_date,omitempty"`
	StatusId		 uuid.UUID	 `mapstructure:"status_id" json:"status_id,omitempty"`
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

