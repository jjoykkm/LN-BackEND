package model_db

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/common/config"
)

//-------------------------------------------------------------------------------//
//							Table plant_type
//-------------------------------------------------------------------------------//
//model plant_type
type PlantType struct {
	DBCommon
	PlantTypeId      uuid.UUID	 `json:"plant_type_id"`
	PlantTypeEN      string		 `json:"plant_type_name_en"`
	PlantTypeTH      string		 `json:"plant_type_name_th"`
}
// New instance plant_type
func (u *PlantType) New() *PlantType {
	return &PlantType{
		DBCommon:      	u.DBCommon ,
		PlantTypeId:	u.PlantTypeId ,
		PlantTypeEN:	u.PlantTypeEN ,
		PlantTypeTH:	u.PlantTypeTH ,
	}
}

// Custom table name for GORM
func (PlantType) TableName() string {
	return config.DB_PLANT_TYPE
}
