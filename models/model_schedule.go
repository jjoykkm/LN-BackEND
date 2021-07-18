package models

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table schedule
//-------------------------------------------------------------------------------//
//model schedule
type Schedule struct {
	ScheduleId     		uuid.UUID	 `gorm:"schedule_id" json:"schedule_id,omitempty"`
	ScheduleName      	string	 	 `gorm:"schedule_name" json:"schedule_name,omitempty"`
	ScheduleDesc      	string	 	 `gorm:"schedule_desc" json:"schedule_desc,omitempty"`
	StartDateTime		time.Time	 `gorm:"start_date_time" json:"start_date_time,omitempty"`
	EndDateTime			time.Time	 `gorm:"end_date_time" json:"end_date_time,omitempty"`
	FreqInterval      	int		 	 `gorm:"frequency_interval" json:"frequency_interval,omitempty"`
	IsAlarm		      	bool		 `gorm:"is_alarm" json:"is_alarm,omitempty"`
	CreateDate			time.Time	 `gorm:"create_date" json:"create_date,omitempty"`
	ChangeDate	    	time.Time	 `gorm:"change_date" json:"change_date,omitempty"`
	FreqTypeId			uuid.UUID	 `gorm:"frequency_type_id" json:"frequency_type_id,omitempty"`
	IndicateTypeId		uuid.UUID	 `gorm:"indicate_type_id" json:"indicate_type_id,omitempty"`
	StatusId			uuid.UUID	 `gorm:"status_id" json:"status_id,omitempty"`
	IsAllDay	      	bool		 `gorm:"is_all_day" json:"is_all_day,omitempty"`
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
