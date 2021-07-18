package models

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table sensor
//-------------------------------------------------------------------------------//
//model sensor
type Sensor struct {
	SensorId      	uuid.UUID	 `gorm:"sensor_id" json:"sensor_id,omitempty"`
	SenserModel     string	 	 `gorm:"senser_model" json:"senser_model,omitempty"`
	SenserLots      string	 	 `gorm:"senser_lots" json:"senser_lots,omitempty"`
	BitTransfer	    byte		 `gorm:"bit_transfer" json:"bit_transfer,omitempty"`
	CreateDate		time.Time	 `gorm:"create_date" json:"create_date,omitempty"`
	ChangeDate	    time.Time	 `gorm:"change_date" json:"change_date,omitempty"`
	sensorTypeId	uuid.UUID	 `gorm:"sensor_type_id" json:"sensor_type_id,omitempty"`
	StatusId		uuid.UUID	 `gorm:"status_id" json:"status_id,omitempty"`
}
// New instance sensor
func (u *Sensor) New() *Sensor {
	return &Sensor{
		SensorId:		u.SensorId ,
		SenserModel:	u.SenserModel ,
		SenserLots:		u.SenserLots ,
		BitTransfer:	u.BitTransfer ,
		CreateDate:		u.CreateDate ,
		ChangeDate:		u.ChangeDate ,
		sensorTypeId:	u.sensorTypeId ,
		StatusId:		u.StatusId ,
	}
}
