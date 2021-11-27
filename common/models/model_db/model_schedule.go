package model_db

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/common/config"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table schedule
//-------------------------------------------------------------------------------//
//model schedule
type Schedule struct {
	DBCommon
	ScheduleId     		uuid.UUID	 `json:"schedule_id"`
	ScheduleName      	string	 	 `json:"schedule_name"`
	ScheduleDesc      	string	 	 `json:"schedule_desc"`
	StartDateTime		time.Time	 `json:"start_date_time"`
	EndDateTime			time.Time	 `json:"end_date_time"`
	FreqInterval      	int		 	 `json:"frequency_interval"`
	IsAlarm		      	bool		 `json:"is_alarm"`
	FrequencyTypeId		uuid.UUID	 `json:"frequency_type_id"`
	IndicateTypeId		uuid.UUID	 `json:"indicate_type_id"`
	IsAllDay	      	bool		 `json:"is_all_day"`
	IsReminder	      	bool		 `json:"is_reminder"`
}
// New instance schedule
func (u *Schedule) New() *Schedule {
	return &Schedule{
		DBCommon:      		u.DBCommon ,
		ScheduleId:			u.ScheduleId ,
		ScheduleName:		u.ScheduleName ,
		ScheduleDesc:		u.ScheduleDesc ,
		StartDateTime:		u.StartDateTime ,
		EndDateTime:		u.EndDateTime ,
		FreqInterval:		u.FreqInterval ,
		IsAlarm:			u.IsAlarm ,
		FrequencyTypeId:	u.FrequencyTypeId ,
		IndicateTypeId:		u.IndicateTypeId ,
		IsAllDay:			u.IsAllDay ,
		IsReminder:			u.IsReminder ,
	}
}

// Custom table name for GORM
func (Schedule) TableName() string {
	return config.DB_SCHEDULE
}
