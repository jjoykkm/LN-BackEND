package model_databases

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table schedule
//-------------------------------------------------------------------------------//
//model schedule
type Schedule struct {
	ScheduleId     		uuid.UUID	 `mapstructure:"schedule_id" json:"schedule_id,omitempty"`
	ScheduleName      	string	 	 `mapstructure:"schedule_name" json:"schedule_name,omitempty"`
	ScheduleDesc      	string	 	 `mapstructure:"schedule_desc" json:"schedule_desc,omitempty"`
	StartDateTime		time.Time	 `mapstructure:"start_date_time" json:"start_date_time,omitempty"`
	EndDateTime			time.Time	 `mapstructure:"end_date_time" json:"end_date_time,omitempty"`
	FreqInterval      	int		 	 `mapstructure:"frequency_interval" json:"frequency_interval,omitempty"`
	IsAlarm		      	bool		 `mapstructure:"is_alarm" json:"is_alarm,omitempty"`
	CreateDate			time.Time	 `mapstructure:"create_date" json:"create_date,omitempty"`
	ChangeDate	    	time.Time	 `mapstructure:"change_date" json:"change_date,omitempty"`
	FreqTypeId			uuid.UUID	 `mapstructure:"frequency_type_id" json:"frequency_type_id,omitempty"`
	IndicateTypeId		uuid.UUID	 `mapstructure:"indicate_type_id" json:"indicate_type_id,omitempty"`
	StatusId			uuid.UUID	 `mapstructure:"status_id" json:"status_id,omitempty"`
	IsAllDay	      	bool		 `mapstructure:"is_all_day" json:"is_all_day,omitempty"`
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
		FreqTypeId:			u.FreqTypeId ,
		IndicateTypeId:		u.IndicateTypeId ,
		StatusId:			u.StatusId ,
		IsAllDay:			u.IsAllDay ,
	}
}
