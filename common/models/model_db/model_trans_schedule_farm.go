package model_db

import (
	"github.com/jjoykkm/ln-backend/common/config"
)

//-------------------------------------------------------------------------------//
//							Table trans_schedule_farm
//-------------------------------------------------------------------------------//
//model trans_schedule_farm
type TransScheduleFarm struct {
	DBCommonGet
	FarmAreaId      string	 `json:"farm_area_id"`
	ScheduleId      string	 `json:"schedule_id"`
}
// New instance trans_schedule_farm
func (u *TransScheduleFarm) New() *TransScheduleFarm {
	return &TransScheduleFarm{
		DBCommonGet:    u.DBCommonGet ,
		FarmAreaId:		u.FarmAreaId ,
		ScheduleId:		u.ScheduleId ,
	}
}

// Custom table name for GORM
func (TransScheduleFarm) TableName() string {
	return config.DB_TRANS_SCHEDULE_FARM
}
