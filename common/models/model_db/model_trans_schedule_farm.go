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
	FarmAreaId      uuid.UUID	 `json:"farm_area_id"`
	ScheduleId      uuid.UUID	 `json:"schedule_id"`
	StatusId		uuid.UUID	 `json:"status_id"`
	CreateDate		time.Time	 `json:"create_date"`
	ChangeDate	    time.Time	 `json:"change_date"`
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
