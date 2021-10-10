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
	SensorId      	uuid.UUID	 `mapstructure:"sensor_id" json:"sensor_id,omitempty"`
	SensorModel     string	 	 `mapstructure:"sensor_model" json:"sensor_model,omitempty"`
	SensorLots      string	 	 `mapstructure:"sensor_lots" json:"sensor_lots,omitempty"`
	BitTransfer	    pgtype.Bit	 `mapstructure:"bit_transfer" json:"bit_transfer,omitempty"`
	CreateDate		time.Time	 `mapstructure:"create_date" json:"create_date,omitempty"`
	ChangeDate	    time.Time	 `mapstructure:"change_date" json:"change_date,omitempty"`
	SensorTypeId	uuid.UUID	 `mapstructure:"sensor_type_id" json:"sensor_type_id,omitempty"`
	StatusId		uuid.UUID	 `mapstructure:"status_id" json:"status_id,omitempty"`
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