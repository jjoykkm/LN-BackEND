package model_services

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"time"
)
//-------------------------------------------------------------------------------//
//							Table Manage FarmArea
//-------------------------------------------------------------------------------//
//Model
type ScheduleFarmArea struct {
	FarmId      	uuid.UUID	 			`mapstructure:"farm_id" json:"farm_id"`
	FarmAreaId      uuid.UUID	 			`mapstructure:"farm_area_id" json:"farm_area_id"`
	FarmAreaName    string		 			`mapstructure:"farm_area_name" json:"farm_area_name"`
}

// New instance
func (u *ScheduleFarmArea) New() *ScheduleFarmArea {
	return &ScheduleFarmArea{
		FarmId:				u.FarmId ,
		FarmAreaId:			u.FarmAreaId ,
		FarmAreaName:		u.FarmAreaName ,
	}
}

//-------------------------------------------------------------------------------//
//							Table ScheduleFreqType
//-------------------------------------------------------------------------------//
//model
type ScheduleFreqType struct {
	FrequencyTypeId 	uuid.UUID	 `mapstructure:"frequency_type_id" json:"frequency_type_id"`
	FrequencyName   	string		 `mapstructure:"frequency_name" json:"frequency_name"`
	IntervalRange   	string		 `mapstructure:"interval_range" json:"interval_range"`
	IsForCustom     	string		 `mapstructure:"is_for_custom" json:"is_for_custom"`
}
// New instance
func (u *ScheduleFreqType) New() *ScheduleFreqType {
	return &ScheduleFreqType{
		FrequencyTypeId:  	u.FrequencyTypeId ,
		FrequencyName:	  	u.FrequencyName ,
		IntervalRange:	  	u.IntervalRange ,
		IsForCustom:	  	u.IsForCustom ,
	}
}

//-------------------------------------------------------------------------------//
//							Table ScheduleInType
//-------------------------------------------------------------------------------//
// model
type ScheduleInType struct {
	IndicateTypeId      uuid.UUID	 `mapstructure:"indicate_type_id" json:"indicate_type_id"`
	IndicateName      	string	 	 `mapstructure:"indicate_name" json:"indicate_name"`
	IndicateDesc      	string	 	 `mapstructure:"indicate_desc" json:"indicate_desc"`
	Important	      	int			 `mapstructure:"important" json:"important"`
	IndColorName      	string	 	 `mapstructure:"ind_color_name" json:"ind_color_name"`
	IndColorCode      	string	 	 `mapstructure:"ind_color_code" json:"ind_color_code"`
	IndColorCodeR      	string	 	 `mapstructure:"ind_color_code_r" json:"ind_color_code_r"`
	IndColorCodeG      	string	 	 `mapstructure:"ind_color_code_g" json:"ind_color_code_g"`
	IndColorCodeB      	string	 	 `mapstructure:"ind_color_code_b" json:"ind_color_code_b"`
}
// New instance
func (u *ScheduleInType) New() *ScheduleInType {
	return &ScheduleInType{
		IndicateTypeId:		u.IndicateTypeId ,
		IndicateName:		u.IndicateName ,
		IndicateDesc:		u.IndicateDesc ,
		Important:			u.Important ,
		IndColorName:		u.IndColorName ,
		IndColorCode:		u.IndColorCode ,
		IndColorCodeR:		u.IndColorCodeR ,
		IndColorCodeG:		u.IndColorCodeG ,
		IndColorCodeB:		u.IndColorCodeB ,
	}
}


//-------------------------------------------------------------------------------//
//							Table socket_action
//-------------------------------------------------------------------------------//
//model socket_action
type SocketAction struct {
	ScheduleId      uuid.UUID	 `mapstructure:"schedule_id" json:"schedule_id"`
	SocketId      	uuid.UUID	 `mapstructure:"socket_id" json:"socket_id"`
	Uid	      		uuid.UUID	 `mapstructure:"uid" json:"uid"`
	StatusSensorId	uuid.UUID	 `mapstructure:"status_sensor_id" json:"status_sensor_id"`
	IsManual    	bool		 `mapstructure:"is_manual" json:"is_manual"`
	CreateDate		time.Time	 `mapstructure:"create_date" json:"create_date"`
	ChangeDate	    time.Time	 `mapstructure:"change_date" json:"change_date"`
	StatusId		uuid.UUID	 `mapstructure:"status_id" json:"status_id"`
}
// New instance socket_action
func (u *SocketAction) New() *SocketAction {
	return &SocketAction{
		ScheduleId:			u.ScheduleId ,
		SocketId:			u.SocketId ,
		Uid:				u.Uid ,
		StatusSensorId:		u.StatusSensorId ,
		IsManual:			u.IsManual ,
		CreateDate:			u.CreateDate ,
		ChangeDate:			u.ChangeDate ,
		StatusId:			u.StatusId ,
	}
}

//-------------------------------------------------------------------------------//
//							Table ScheRemind
//-------------------------------------------------------------------------------//
//Model
type ScheduleStruct struct {
	FarmAreaId      	uuid.UUID	 		`mapstructure:"farm_area_id" json:"farm_area_id"`
	FarmAreaName    	string		 		`mapstructure:"farm_area_name" json:"farm_area_name"`
	ScheduleDesc      	string	 	 		`mapstructure:"schedule_desc" json:"schedule_desc"`
	ScheduleId     		uuid.UUID	 		`mapstructure:"schedule_id" json:"schedule_id"`
	ScheduleName      	string	 	 		`mapstructure:"schedule_name" json:"schedule_name"`
	StartDateTime		time.Time	 		`mapstructure:"start_date_time" json:"start_date_time"`
	EndDateTime			time.Time	 		`mapstructure:"end_date_time" json:"end_date_time"`
	FreqInterval      	int		 	 		`mapstructure:"frequency_interval" json:"frequency_interval"`
	IsAlarm		      	bool		 		`mapstructure:"is_alarm" json:"is_alarm"`
	FreqType			ScheduleFreqType	`mapstructure:"freq_type" json:"freq_type" gorm:"embedded"`
	IndicateType		ScheduleInType		`mapstructure:"indicate_type" json:"indicate_type" gorm:"embedded"`
	IsAllDay	      	bool		 		`mapstructure:"is_all_day" json:"is_all_day"`
	IsReminder	      	bool		 		`mapstructure:"is_reminder" json:"is_reminder"`
}

// New instance
func (u *ScheduleStruct) New() *ScheduleStruct {
	return &ScheduleStruct{
		FarmAreaId:			u.FarmAreaId ,
		FarmAreaName:		u.FarmAreaName ,
		ScheduleId:			u.ScheduleId ,
		ScheduleName:		u.ScheduleName ,
		ScheduleDesc:		u.ScheduleDesc ,
		FreqInterval:		u.FreqInterval ,
		IsAlarm:			u.IsAlarm ,
		FreqType:			u.FreqType ,
		IndicateType:		u.IndicateType ,
		IsAllDay:			u.IsAllDay ,
		IsReminder:			u.IsReminder ,
	}
}

//-------------------------------------------------------------------------------//
//							Table ScheRemind
//-------------------------------------------------------------------------------//
//Model
type ScheduleScheRemind struct {
	ScheduleList    []ScheduleStruct	`mapstructure:"schedule_list" json:"schedule_list"`
	ReminderList    []ScheduleStruct	`mapstructure:"reminder_list" json:"reminder_list"`
}

// New instance
func (u *ScheduleScheRemind) New() *ScheduleScheRemind {
	return &ScheduleScheRemind{
		ScheduleList:		u.ScheduleList ,
		ReminderList:		u.ReminderList ,
	}
}
