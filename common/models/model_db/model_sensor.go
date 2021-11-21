package model_db

import (
	"github.com/jackc/pgtype"
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/common/config"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table sensor
//-------------------------------------------------------------------------------//
//model sensor
type Sensor struct {
	SensorId      	uuid.UUID	 `json:"sensor_id"`
	SensorModel     string	 	 `json:"sensor_model"`
	SensorLots      string	 	 `json:"sensor_lots"`
	BitTransfer	    pgtype.Bit	 `json:"bit_transfer"`
	CreateDate		time.Time	 `json:"create_date"`
	ChangeDate	    time.Time	 `json:"change_date"`
	SensorTypeId	uuid.UUID	 `json:"sensor_type_id"`
	StatusId		uuid.UUID	 `json:"status_id"`
}
// New instance sensor
func (u *Sensor) New() *Sensor {
	return &Sensor{
		SensorId:		u.SensorId ,
		SensorModel:	u.SensorModel ,
		SensorLots:		u.SensorLots ,
		BitTransfer:	u.BitTransfer ,
		CreateDate:		u.CreateDate ,
		ChangeDate:		u.ChangeDate ,
		SensorTypeId:	u.SensorTypeId ,
		StatusId:		u.StatusId ,
	}
}

// Custom table name for GORM
func (Sensor) TableName() string {
	return config.DB_SENSOR
}


//-------------------------------------------------------------------------------//
//							Request Socket
//-------------------------------------------------------------------------------//
type SensorUS struct {
	SensorId      	string	 `json:"sensor_id"  gorm:"type:uuid;default:uuid_generate_v4()"`
	SensorModel     string	 `json:"sensor_model"`
	SensorLots      string	 `json:"sensor_lots"`
	SensorTypeId	string	 `json:"sensor_type_id"`
	StatusId		string	 `json:"status_id"`
}
func (SensorUS) TableName() string {
	return config.DB_SENSOR
}