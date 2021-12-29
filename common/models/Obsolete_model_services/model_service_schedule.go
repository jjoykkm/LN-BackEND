package Obsolete_model_services

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"time"
)
//-------------------------------------------------------------------------------//
//							Table Manage FarmArea
//-------------------------------------------------------------------------------//
//Model
type ScheduleFarmArea struct {
	FarmId      	string	 			`json:"farm_id"`
	FarmAreaId      string	 			`json:"farm_area_id"`
	FarmAreaName    string	 			`json:"farm_area_name"`
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
	FrequencyTypeId 	string	 `json:"frequency_type_id"`
	FrequencyName   	string	 `json:"frequency_name"`
	IntervalRange   	string	 `json:"interval_range"`
	IsForCustom     	string	 `json:"is_for_custom"`
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
	IndicateTypeId      string	 `json:"indicate_type_id"`
	IndicateName      	string	 	 `json:"indicate_name"`
	IndicateDesc      	string	 	 `json:"indicate_desc"`
	Important	      	int			 `json:"important"`
	IndColorName      	string	 	 `json:"ind_color_name"`
	IndColorCode      	string	 	 `json:"ind_color_code"`
	IndColorCodeR      	string	 	 `json:"ind_color_code_r"`
	IndColorCodeG      	string	 	 `json:"ind_color_code_g"`
	IndColorCodeB      	string	 	 `json:"ind_color_code_b"`
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
	ScheduleId      string	 `json:"schedule_id"`
	SocketId      	string	 `json:"socket_id"`
	Uid	      		string	 `json:"uid"`
	StatusSensorId	string	 `json:"status_sensor_id"`
	IsManual    	bool		 `json:"is_manual"`
	CreateDate		time.Time	 `json:"create_date"`
	ChangeDate	    time.Time	 `json:"change_date"`
	StatusId		string	 `json:"status_id"`
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
	FarmAreaId      	string	 		`json:"farm_area_id"`
	FarmAreaName    	string	 		`json:"farm_area_name"`
	ScheduleDesc      	string	 	 		`json:"schedule_desc"`
	ScheduleId     		string	 		`json:"schedule_id"`
	ScheduleName      	string	 	 		`json:"schedule_name"`
	StartDateTime		time.Time	 		`json:"start_date_time"`
	EndDateTime			time.Time	 		`json:"end_date_time"`
	FreqInterval      	int		 	 		`json:"frequency_interval"`
	IsAlarm		      	bool		 		`json:"is_alarm"`
	FreqType			ScheduleFreqType	`json:"freq_type" gorm:"embedded"`
	IndicateType		ScheduleInType		`json:"indicate_type" gorm:"embedded"`
	IsAllDay	      	bool		 		`json:"is_all_day"`
	IsReminder	      	bool		 		`json:"is_reminder"`
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
	ScheduleList    []ScheduleStruct	`json:"schedule_list"`
	ReminderList    []ScheduleStruct	`json:"reminder_list"`
}

// New instance
func (u *ScheduleScheRemind) New() *ScheduleScheRemind {
	return &ScheduleScheRemind{
		ScheduleList:		u.ScheduleList ,
		ReminderList:		u.ReminderList ,
	}
}
