package model_db

import (
	"github.com/jjoykkm/ln-backend/common/config"
)

//-------------------------------------------------------------------------------//
//							Table plant_type
//-------------------------------------------------------------------------------//
//model plant_type
type PlantType struct {
	DBCommonGet
	PlantTypeId      string	 `json:"plant_type_id"`
	PlantTypeEN      string	 `json:"plant_type_name_en"`
	PlantTypeTH      string	 `json:"plant_type_name_th"`
}
// New instance plant_type
func (u *PlantType) New() *PlantType {
	return &PlantType{
		DBCommonGet:      	u.DBCommonGet ,
		PlantTypeId:	u.PlantTypeId ,
		PlantTypeEN:	u.PlantTypeEN ,
		PlantTypeTH:	u.PlantTypeTH ,
	}
}

// Custom table name for GORM
func (PlantType) TableName() string {
	return config.DB_PLANT_TYPE
}
