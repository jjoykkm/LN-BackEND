package model_db

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/common/config"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table sensor_type
//-------------------------------------------------------------------------------//
//model sensor_type
type SensorType struct {
	SensorTypeId      	uuid.UUID	 `json:"TypeId,omitempty"`
	SensorTypeNameEN    string	 	 `json:"TypeNameEN,omitempty"`
	CreateDate			time.Time	 `json:"CreateDate,omitempty"`
	ChangeDate	    	time.Time	 `json:"ChangeDate,omitempty"`
	StatusId			uuid.UUID	 `json:"StatusId,omitempty"`
	SensorTypeNameTH    string	 	 `json:"TypeNameTH,omitempty"`
}
// New instance sensor_type
func (u *SensorType) New() *SensorType {
	return &SensorType{
		SensorTypeId:		u.SensorTypeId ,
		SensorTypeNameEN:	u.SensorTypeNameEN ,
		CreateDate:			u.CreateDate ,
		ChangeDate:			u.ChangeDate ,
		StatusId:			u.StatusId ,
		SensorTypeNameTH:	u.SensorTypeNameTH ,
	}
}

// Custom table name for GORM
func (SensorType) TableName() string {
	return config.DB_SENSOR_TYPE
}
