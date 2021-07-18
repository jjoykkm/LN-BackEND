package models

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table frequency_type
//-------------------------------------------------------------------------------//
//model frequency_type
type FrequencyType struct {
	FreqTypeId     	uuid.UUID	 `gorm:"frequency_type_id" json:"frequency_type_id,omitempty"`
	FreqName      	string		 `gorm:"frequency_name" json:"frequency_name,omitempty"`
	IntervalRange   string		 `gorm:"interval_range" json:"interval_range,omitempty"`
	IsForCustom     string		 `gorm:"is_for_custom" json:"is_for_custom,omitempty"`
	CreateDate		time.Time	 `gorm:"create_date" json:"create_date,omitempty"`
	ChangeDate	    time.Time	 `gorm:"change_date" json:"change_date,omitempty"`
	StatusId		uuid.UUID	 `gorm:"status_id" json:"status_id,omitempty"`
}
// New instance frequency_type
func (u *FrequencyType) New() *FrequencyType {
	return &FrequencyType{
		FreqTypeId:		  u.FreqTypeId ,
		FreqName:		  u.FreqName ,
		IntervalRange:	  u.IntervalRange ,
		IsForCustom:	  u.IsForCustom ,
		CreateDate:		  u.CreateDate ,
		ChangeDate:		  u.ChangeDate ,
		StatusId:		  u.StatusId ,
	}
}
