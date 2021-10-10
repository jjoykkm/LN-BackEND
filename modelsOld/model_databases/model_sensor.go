package model_databases

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
	SensorId      	uuid.UUID	 `json:"sensor_id,omitempty"`
	SensorModel     string	 	 `json:"sensor_model,omitempty"`
	SensorLots      string	 	 `json:"sensor_lots,omitempty"`
	BitTransfer	    pgtype.Bit	 `json:"bit_transfer,omitempty"`
	CreateDate		time.Time	 `json:"create_date,omitempty"`
	ChangeDate	    time.Time	 `json:"change_date,omitempty"`
	SensorTypeId	uuid.UUID	 `json:"sensor_type_id,omitempty"`
	StatusId		uuid.UUID	 `json:"status_id,omitempty"`
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