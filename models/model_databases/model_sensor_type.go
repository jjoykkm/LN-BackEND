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
	SensorTypeId      	uuid.UUID	 `mapstructure:"sensor_type_id" json:"sensor_type_id,omitempty"`
	SensorTypeName      string	 	 `mapstructure:"sensor_type_name" json:"sensor_type_name,omitempty"`
	CreateDate			time.Time	 `mapstructure:"create_date" json:"create_date,omitempty"`
	ChangeDate	    	time.Time	 `mapstructure:"change_date" json:"change_date,omitempty"`
	StatusId			uuid.UUID	 `mapstructure:"status_id" json:"status_id,omitempty"`
}
// New instance sensor_type
func (u *SensorType) New() *SensorType {
	return &SensorType{
		SensorTypeId:		u.SensorTypeId ,
		SensorTypeName:		u.SensorTypeName ,
		CreateDate:			u.CreateDate ,
		ChangeDate:			u.ChangeDate ,
		StatusId:			u.StatusId ,
	}
}
