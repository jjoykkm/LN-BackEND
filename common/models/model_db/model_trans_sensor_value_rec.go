package model_db

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/common/config"
)

//-------------------------------------------------------------------------------//
//							Table trans_sensor_value_rec
//-------------------------------------------------------------------------------//
//model trans_sensor_value_rec
type TransSensorValueRec struct {
	DBCommonGet
	FormulaPlantId 		uuid.UUID	 `json:"formula_plant_id"`
	SensorTypeId   		uuid.UUID	 `json:"sensor_type_id"`
	ValueRec      		float64		 `json:"value_rec"`
}
// New instance trans_sensor_value_rec
func (u *TransSensorValueRec) New() *TransSensorValueRec {
	return &TransSensorValueRec{
		DBCommonGet:      		u.DBCommonGet ,
		FormulaPlantId:		u.FormulaPlantId ,
		SensorTypeId:		u.SensorTypeId ,
		ValueRec:			u.ValueRec ,
	}
}

// Custom table name for GORM
func (TransSensorValueRec) TableName() string {
	return config.DB_TRANS_SENSOR_VALUE_REC
}
