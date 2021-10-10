package model_db

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/config"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table trans_schedule_farm
//-------------------------------------------------------------------------------//
//model trans_schedule_farm
type TransScheduleFarm struct {
	FarmAreaId      uuid.UUID	 `mapstructure:"farm_area_id" json:"farm_area_id,omitempty"`
	ScheduleId      uuid.UUID	 `mapstructure:"schedule_id" json:"schedule_id,omitempty"`
	StatusId		uuid.UUID	 `mapstructure:"status_id" json:"status_id,omitempty"`
	CreateDate		time.Time	 `mapstructure:"create_date" json:"create_date,omitempty"`
	ChangeDate	    time.Time	 `mapstructure:"change_date" json:"change_date,omitempty"`
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
