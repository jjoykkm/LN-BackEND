package model_db

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/common/config"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table trans_schedule_farm
//-------------------------------------------------------------------------------//
//model trans_schedule_farm
type TransScheduleFarm struct {
	FarmAreaId      uuid.UUID	 `json:"FarmAreaId,omitempty"`
	ScheduleId      uuid.UUID	 `json:"ScheduleId,omitempty"`
	StatusId		uuid.UUID	 `json:"StatusId,omitempty"`
	CreateDate		time.Time	 `json:"CreateDate,omitempty"`
	ChangeDate	    time.Time	 `json:"ChangeDate,omitempty"`
}
// New instance trans_schedule_farm
func (u *TransScheduleFarm) New() *TransScheduleFarm {
	return &TransScheduleFarm{
		FarmAreaId:		u.FarmAreaId ,
		ScheduleId:		u.ScheduleId ,
		StatusId:		u.StatusId ,
		CreateDate:		u.CreateDate ,
		ChangeDate:		u.ChangeDate ,
	}
}

// Custom table name for GORM
func (TransScheduleFarm) TableName() string {
	return config.DB_TRANS_SCHEDULE_FARM
}
