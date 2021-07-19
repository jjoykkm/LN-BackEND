package model_databases

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table status_sensor
//-------------------------------------------------------------------------------//
//model status_sensor
type StatusSensor struct {
	StatusSensorId      uuid.UUID	 `gorm:"status_sensor_id" json:"status_sensor_id,omitempty"`
	StatusName      	string		 `gorm:"status_name" json:"status_name,omitempty"`
	CreateDate			time.Time	 `gorm:"create_date" json:"create_date,omitempty"`
	ChangeDate	    	time.Time	 `gorm:"change_date" json:"change_date,omitempty"`
	StatusId			uuid.UUID	 `gorm:"status_id" json:"status_id,omitempty"`
}
// New instance status_sensor
func (u *StatusSensor) New() *StatusSensor {
	return &StatusSensor{
		StatusSensorId:		u.StatusSensorId ,
		StatusName:			u.StatusName ,
		CreateDate:			u.CreateDate ,
		ChangeDate:			u.ChangeDate ,
		StatusId:			u.StatusId ,
	}
}
