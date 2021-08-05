package model_databases

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table sensor_type
//-------------------------------------------------------------------------------//
//model sensor_type
type SensorType struct {
	SensorTypeId      	uuid.UUID	 `gorm:"sensor_type_id" mapstructure:"sensor_type_id" json:"sensor_type_id"`
	SensorTypeNameEN    string	 	 `gorm:"sensor_type_name_en" mapstructure:"sensor_type_name_en" json:"sensor_type_name_en"`
	CreateDate			time.Time	 `gorm:"create_date" mapstructure:"create_date" json:"create_date"`
	ChangeDate	    	time.Time	 `gorm:"change_date" mapstructure:"change_date" json:"change_date"`
	StatusId			uuid.UUID	 `gorm:"status_id" mapstructure:"status_id" json:"status_id"`
	SensorTypeNameTH    string	 	 `gorm:"sensor_type_name_th" mapstructure:"sensor_type_name_th" json:"sensor_type_name_th"`
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
