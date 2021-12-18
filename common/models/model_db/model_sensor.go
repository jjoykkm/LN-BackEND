package model_db

import (
	"github.com/jackc/pgtype"
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/common/config"
)

//-------------------------------------------------------------------------------//
//							Table sensor
//-------------------------------------------------------------------------------//
//model sensor
type Sensor struct {
	DBCommonGet
	SensorId      	uuid.UUID	 `json:"sensor_id"`
	SensorModel     string	 	 `json:"sensor_model"`
	SensorLots      string	 	 `json:"sensor_lots"`
	BitTransfer	    pgtype.Bit	 `json:"bit_transfer"`
	SensorTypeId	uuid.UUID	 `json:"sensor_type_id"`
}
// New instance sensor
func (u *Sensor) New() *Sensor {
	return &Sensor{
		DBCommonGet:      	u.DBCommonGet ,
		SensorId:		u.SensorId ,
		SensorModel:	u.SensorModel ,
		SensorLots:		u.SensorLots ,
		BitTransfer:	u.BitTransfer ,
		SensorTypeId:	u.SensorTypeId ,
	}
}

// Custom table name for GORM
func (Sensor) TableName() string {
	return config.DB_SENSOR
}

//-------------------------------------------------------------------------------//
//							Upsert Socket
//-------------------------------------------------------------------------------//
type SensorUS struct {
	SensorId      	string	 `json:"sensor_id"  gorm:"default:uuid_generate_v4()"`
	SensorModel     string	 `json:"sensor_model"`
	SensorLots      string	 `json:"sensor_lots"`
	SensorTypeId	string	 `json:"sensor_type_id"`
	StatusId		string	 `json:"status_id"`
}
func (SensorUS) TableName() string {
	return config.DB_SENSOR
}