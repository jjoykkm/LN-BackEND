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
	DistrictId      uuid.UUID	 `json:"DistrictId,omitempty"`
	DistrictEN      string		 `json:"DistrictEN,omitempty"`
	DistrictTH      string		 `json:"DistrictTH,omitempty"`
	CreateDate		time.Time	 `json:"CreateDate,omitempty"`
	ChangeDate	    time.Time	 `json:"ChangeDate,omitempty"`
	StatusId		uuid.UUID	 `json:"StatusId,omitempty"`
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


