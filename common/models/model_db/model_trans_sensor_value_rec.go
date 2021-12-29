package model_db

import (
	"github.com/jjoykkm/ln-backend/common/config"
)

//-------------------------------------------------------------------------------//
//							Table trans_sensor_value_rec
//-------------------------------------------------------------------------------//
//model trans_sensor_value_rec
type TransSensorValueRec struct {
	DBCommonGet
	FormulaPlantId 		string	 	 `json:"formula_plant_id"`
	SensorTypeId   		string	 	 `json:"sensor_type_id"`
	ValueRec      		float64		 `json:"value_rec"`
}
// New instance trans_sensor_value_rec
func (u *TransSensorValueRec) New() *TransSensorValueRec {
	return &TransSensorValueRec{
		DBCommonGet:      	u.DBCommonGet ,
		FormulaPlantId:		u.FormulaPlantId ,
		SensorTypeId:		u.SensorTypeId ,
		ValueRec:			u.ValueRec ,
	}
}

// Custom table name for GORM
func (TransSensorValueRec) TableName() string {
	return config.DB_TRANS_SENSOR_VALUE_REC
}

//-------------------------------------------------------------------------------//
//									Upsert
//-------------------------------------------------------------------------------//
type TransSenValueRecUS struct {
	DBCommonCreateUpdate
	FormulaPlantId 		string	 	`json:"formula_plant_id" gorm:"primaryKey"`
	SensorTypeId   		string	 	`json:"sensor_type_id" gorm:"primaryKey"`
	ValueRec      		float64		`json:"value_rec"`
	StatusId		 	string		`json:"status_id"`
}
func (TransSenValueRecUS) TableName() string {
	return config.DB_TRANS_SENSOR_VALUE_REC
}
