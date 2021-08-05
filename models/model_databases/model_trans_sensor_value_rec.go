package model_databases

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table trans_sensor_value_rec
//-------------------------------------------------------------------------------//
//model trans_sensor_value_rec
type TransSensorValueRec struct {
	FormulaPlantId 		uuid.UUID	 `mapstructure:"formula_plant_id" json:"formula_plant_id"`
	SensorTypeId   		uuid.UUID	 `mapstructure:"sensor_type_id" json:"sensor_type_id"`
	StatusId	   		uuid.UUID	 `mapstructure:"status_id" json:"status_id"`
	ValueRec      		float64		 `mapstructure:"value_rec" json:"value_rec"`
	CreateDate	   		time.Time	 `mapstructure:"create_date" json:"create_date"`
	ChangeDate	   		time.Time	 `mapstructure:"change_date" json:"change_date"`
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
