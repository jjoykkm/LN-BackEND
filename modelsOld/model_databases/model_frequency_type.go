package model_databases

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/common/config"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table frequency_type
//-------------------------------------------------------------------------------//
//model frequency_type
type FrequencyType struct {
	FrequencyTypeId		uuid.UUID	 `json:"frequency_type_id"`
	FrequencyName   	string		 `json:"frequency_name"`
	IntervalRange   	string		 `json:"interval_range"`
	IsForCustom     	string		 `json:"is_for_custom"`
	CreateDate			time.Time	 `json:"create_date"`
	ChangeDate	    	time.Time	 `json:"change_date"`
	StatusId			uuid.UUID	 `json:"status_id"`
}
// New instance frequency_type
func (u *FrequencyType) New() *FrequencyType {
	return &FrequencyType{
		FrequencyTypeId:  u.FrequencyTypeId ,
		FrequencyName:	  u.FrequencyName ,
		IntervalRange:	  u.IntervalRange ,
		IsForCustom:	  u.IsForCustom ,
		CreateDate:		  u.CreateDate ,
		ChangeDate:		  u.ChangeDate ,
		StatusId:		  u.StatusId ,
	}
}

// Custom table name for GORM
func (FrequencyType) TableName() string {
	return config.DB_FREQUENCY_TYPE
}
