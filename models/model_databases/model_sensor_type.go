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
	SensorTypeId      	uuid.UUID	 `gorm:"sensor_type_id" json:"sensor_type_id,omitempty"`
	SensorTypeName      string	 	 `gorm:"sensor_type_name" json:"sensor_type_name,omitempty"`
	CreateDate			time.Time	 `gorm:"create_date" json:"create_date,omitempty"`
	ChangeDate	    	time.Time	 `gorm:"change_date" json:"change_date,omitempty"`
	StatusId			uuid.UUID	 `gorm:"status_id" json:"status_id,omitempty"`
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
