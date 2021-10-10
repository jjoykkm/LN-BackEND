package model_db

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/common/config"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table district
//-------------------------------------------------------------------------------//
//model district
type District struct {
	DistrictId      uuid.UUID	 `json:"district_id,omitempty"`
	DistrictEN      string		 `json:"district_en,omitempty"`
	DistrictTH      string		 `json:"district_th,omitempty"`
	CreateDate		 time.Time	 `json:"create_date,omitempty"`
	ChangeDate	     time.Time	 `json:"change_date,omitempty"`
	StatusId		 uuid.UUID	 `json:"status_id,omitempty"`
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

// Custom table name for GORM
func (District) TableName() string {
	return config.DB_DISTRICT
}


