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
	ScheduleId     		uuid.UUID	 `json:"schedule_id,omitempty"`
	ScheduleName      	string	 	 `json:"schedule_name,omitempty"`
	ScheduleDesc      	string	 	 `json:"schedule_desc,omitempty"`
	StartDateTime		time.Time	 `json:"start_date_time,omitempty"`
	EndDateTime			time.Time	 `json:"end_date_time,omitempty"`
	FreqInterval      	int		 	 `json:"frequency_interval,omitempty"`
	IsAlarm		      	bool		 `json:"is_alarm,omitempty"`
	CreateDate			time.Time	 `json:"create_date,omitempty"`
	ChangeDate	    	time.Time	 `json:"change_date,omitempty"`
	FrequencyTypeId		uuid.UUID	 `json:"frequency_type_id,omitempty"`
	IndicateTypeId		uuid.UUID	 `json:"indicate_type_id,omitempty"`
	StatusId			uuid.UUID	 `json:"status_id,omitempty"`
	IsAllDay	      	bool		 `json:"is_all_day,omitempty"`
	IsReminder	      	bool		 `json:"is_reminder,omitempty"`
}
// New instance schedule
func (u *Schedule) New() *Schedule {
	return &Schedule{
		ScheduleId:			u.ScheduleId ,
		ScheduleName:		u.ScheduleName ,
		ScheduleDesc:		u.ScheduleDesc ,
		StartDateTime:		u.StartDateTime ,
		EndDateTime:		u.EndDateTime ,
		FreqInterval:		u.FreqInterval ,
		IsAlarm:			u.IsAlarm ,
		CreateDate:			u.CreateDate ,
		ChangeDate:			u.ChangeDate ,
		FrequencyTypeId:	u.FrequencyTypeId ,
		IndicateTypeId:		u.IndicateTypeId ,
		StatusId:			u.StatusId ,
		IsAllDay:			u.IsAllDay ,
		IsReminder:			u.IsReminder ,
	}
}

// Custom table name for GORM
func (Schedule) TableName() string {
	return config.DB_SCHEDULE
}
