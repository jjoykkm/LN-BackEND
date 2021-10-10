package model_db

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/config"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table status_sensor
//-------------------------------------------------------------------------------//
//model status_sensor
type StatusSensor struct {
	StatusSensorId      uuid.UUID	 `mapstructure:"status_sensor_id" json:"status_sensor_id,omitempty"`
	StatusName      	string		 `mapstructure:"status_name" json:"status_name,omitempty"`
	CreateDate			time.Time	 `mapstructure:"create_date" json:"create_date,omitempty"`
	ChangeDate	    	time.Time	 `mapstructure:"change_date" json:"change_date,omitempty"`
	StatusId			uuid.UUID	 `mapstructure:"status_id" json:"status_id,omitempty"`
}
// New instance status_sensor
func (u *StatusSensor) New() *StatusSensor {
	return &StatusSensor{
		StatusSensorId:		u.StatusSensorId ,
		StatusName:			u.StatusName ,
		CreateDate:			u.CreateDate ,
		ChangeDate:			u.ChangeDate ,
		StatusId:			u.StatusId ,
	}
}

// Custom table name for GORM
func (StatusSensor) TableName() string {
	return config.DB_STATUS_SENSOR
}
