package model_db

import (
	"github.com/jjoykkm/ln-backend/common/config"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table schedule
//-------------------------------------------------------------------------------//
//model schedule
type Schedule struct {
	DBCommonGet
	ScheduleId     		string	 	 `json:"schedule_id"`
	ScheduleName      	string	 	 `json:"schedule_name"`
	ScheduleDesc      	string	 	 `json:"schedule_desc"`
	StartDateTime		time.Time	 `json:"start_date_time"`
	EndDateTime			time.Time	 `json:"end_date_time"`
	FreqInterval      	int		 	 `json:"frequency_interval"`
	IsAlarm		      	bool		 `json:"is_alarm"`
	FrequencyTypeId		string	 	 `json:"frequency_type_id"`
	IndicateTypeId		string	 	 `json:"indicate_type_id"`
	IsAllDay	      	bool		 `json:"is_all_day"`
	IsReminder	      	bool		 `json:"is_reminder"`
}
// New instance schedule
func (u *Schedule) New() *Schedule {
	return &Schedule{
		DBCommonGet:      		u.DBCommonGet ,
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
