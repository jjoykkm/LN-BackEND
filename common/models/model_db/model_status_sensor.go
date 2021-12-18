package model_db

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/common/config"
)

//-------------------------------------------------------------------------------//
//							Table status_sensor
//-------------------------------------------------------------------------------//
//model status_sensor
type StatusSensor struct {
	DBCommonGet
	StatusSensorId      uuid.UUID	 `json:"status_sensor_id"`
	StatusName      	string		 `json:"status_name"`
}
// New instance status_sensor
func (u *StatusSensor) New() *StatusSensor {
	return &StatusSensor{
		DBCommonGet:      		u.DBCommonGet ,
		StatusSensorId:		u.StatusSensorId ,
		StatusName:			u.StatusName ,
	}
}

// Custom table name for GORM
func (StatusSensor) TableName() string {
	return config.DB_STATUS_SENSOR
}
