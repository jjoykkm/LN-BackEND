package model_db

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/common/config"
)

//-------------------------------------------------------------------------------//
//							Table trans_schedule_farm
//-------------------------------------------------------------------------------//
//model trans_schedule_farm
type TransScheduleFarm struct {
	DBCommonGet
	FarmAreaId      uuid.UUID	 `json:"farm_area_id"`
	ScheduleId      uuid.UUID	 `json:"schedule_id"`
}
// New instance trans_schedule_farm
func (u *TransScheduleFarm) New() *TransScheduleFarm {
	return &TransScheduleFarm{
		DBCommonGet:      		u.DBCommonGet ,
		FarmAreaId:		u.FarmAreaId ,
		ScheduleId:		u.ScheduleId ,
	}
}

// Custom table name for GORM
func (TransScheduleFarm) TableName() string {
	return config.DB_TRANS_SCHEDULE_FARM
}
