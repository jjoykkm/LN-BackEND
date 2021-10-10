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
	SensorId      	uuid.UUID	 `json:"SensorId,omitempty"`
	SensorModel     string	 	 `json:"SensorModel,omitempty"`
	SensorLots      string	 	 `json:"SensorLots,omitempty"`
	BitTransfer	    pgtype.Bit	 `json:"BitTransfer,omitempty"`
	CreateDate		time.Time	 `json:"CreateDate,omitempty"`
	ChangeDate	    time.Time	 `json:"ChangeDate,omitempty"`
	SensorTypeId	uuid.UUID	 `json:"SensorTypeId,omitempty"`
	StatusId		uuid.UUID	 `json:"StatusId,omitempty"`
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