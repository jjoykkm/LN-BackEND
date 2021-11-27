package model_db

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/common/config"
)

//-------------------------------------------------------------------------------//
//							Table frequency_type
//-------------------------------------------------------------------------------//
//model frequency_type
type FrequencyType struct {
	DBCommon
	FrequencyTypeId		uuid.UUID	 `json:"frequency_type_id"`
	FrequencyName   	string		 `json:"frequency_name"`
	IntervalRange   	string		 `json:"interval_range"`
	IsForCustom     	string		 `json:"is_for_custom"`
}
// New instance frequency_type
func (u *FrequencyType) New() *FrequencyType {
	return &FrequencyType{
		DBCommon:      	  u.DBCommon ,
		FrequencyTypeId:  u.FrequencyTypeId ,
		FrequencyName:	  u.FrequencyName ,
		IntervalRange:	  u.IntervalRange ,
		IsForCustom:	  u.IsForCustom ,
	}
}

// Custom table name for GORM
func (FrequencyType) TableName() string {
	return config.DB_FREQUENCY_TYPE
}
