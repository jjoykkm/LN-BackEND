package model_db

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/common/config"
)

//-------------------------------------------------------------------------------//
//							Table sensor_type
//-------------------------------------------------------------------------------//
//model sensor_type
type SensorType struct {
	DBCommonGet
	SensorTypeId      	uuid.UUID	 `json:"sensor_type_id"`
	SensorTypeNameEN    string	 	 `json:"sensor_type_name_en"`
	SensorTypeNameTH    string	 	 `json:"sensor_type_name_th"`
}
// New instance sensor_type
func (u *SensorType) New() *SensorType {
	return &SensorType{
		DBCommonGet:      		u.DBCommonGet ,
		SensorTypeId:		u.SensorTypeId ,
		SensorTypeNameEN:	u.SensorTypeNameEN ,
		SensorTypeNameTH:	u.SensorTypeNameTH ,
	}
}

// Custom table name for GORM
func (SensorType) TableName() string {
	return config.DB_SENSOR_TYPE
}
