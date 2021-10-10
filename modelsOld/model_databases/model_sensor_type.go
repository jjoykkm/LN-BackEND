package model_databases

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/common/config"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table sensor_type
//-------------------------------------------------------------------------------//
//model sensor_type
type SensorType struct {
	SensorTypeId      	uuid.UUID	 `mapstructure:"sensor_type_id" json:"sensor_type_id"`
	SensorTypeNameEN    string	 	 `mapstructure:"sensor_type_name_en" json:"sensor_type_name_en"`
	CreateDate			time.Time	 `mapstructure:"create_date" json:"create_date"`
	ChangeDate	    	time.Time	 `mapstructure:"change_date" json:"change_date"`
	StatusId			uuid.UUID	 `mapstructure:"status_id" json:"status_id"`
	SensorTypeNameTH    string	 	 `mapstructure:"sensor_type_name_th" json:"sensor_type_name_th"`
}
// New instance sensor_type
func (u *SensorType) New() *SensorType {
	return &SensorType{
		SensorTypeId:		u.SensorTypeId ,
		SensorTypeNameEN:	u.SensorTypeNameEN ,
		CreateDate:			u.CreateDate ,
		ChangeDate:			u.ChangeDate ,
		StatusId:			u.StatusId ,
		SensorTypeNameTH:	u.SensorTypeNameTH ,
	}
}

// Custom table name for GORM
func (SensorType) TableName() string {
	return config.DB_SENSOR_TYPE
}
