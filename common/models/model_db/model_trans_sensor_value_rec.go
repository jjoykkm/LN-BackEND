package model_db

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/common/config"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table trans_sensor_value_rec
//-------------------------------------------------------------------------------//
//model trans_sensor_value_rec
type TransSensorValueRec struct {
	FormulaPlantId 		uuid.UUID	 `gorm:"primaryKey" json:"formula_plant_id"`
	SensorTypeId   		uuid.UUID	 `json:"sensor_type_id"`
	StatusId	   		uuid.UUID	 `json:"status_id"`
	ValueRec      		float64		 `json:"value_rec"`
	CreateDate	   		time.Time	 `json:"create_date"`
	ChangeDate	   		time.Time	 `json:"change_date"`
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

// Custom table name for GORM
func (TransSensorValueRec) TableName() string {
	return config.DB_TRANS_SENSOR_VALUE_REC
}
