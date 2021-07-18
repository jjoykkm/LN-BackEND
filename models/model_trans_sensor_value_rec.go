package models

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table trans_sensor_value_rec
//-------------------------------------------------------------------------------//
//model trans_sensor_value_rec
type TransSensorValueRec struct {
	FormulaPlantId 		uuid.UUID	 `gorm:"formula_plant_id" json:"formula_plant_id,omitempty"`
	SensorTypeId   		uuid.UUID	 `gorm:"sensor_type_id" json:"sensor_type_id,omitempty"`
	StatusId	   		uuid.UUID	 `gorm:"status_id" json:"status_id,omitempty"`
	ValueRec      		float64		 `gorm:"value_rec" json:"value_rec,omitempty"`
	CreateDate	   		time.Time	 `gorm:"create_date" json:"create_date,omitempty"`
	ChangeDate	   		time.Time	 `gorm:"change_date" json:"change_date,omitempty"`
}
// New instance trans_sensor_value_rec
func (u *TransSensorValueRec) New() *TransSensorValueRec {
	return &TransSensorValueRec{
		FormulaPlantId:		u.FormulaPlantId ,
		SensorTypeId:		u.SensorTypeId ,
		StatusId:			u.StatusId ,
		ValueRec:			u.ValueRec ,
		CreateDate:			u.CreateDate ,
		ChangeDate:			u.ChangeDate ,
	}
}
