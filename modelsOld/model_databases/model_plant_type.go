package model_databases

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/common/config"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table plant_type
//-------------------------------------------------------------------------------//
//model plant_type
type PlantType struct {
	PlantTypeId      uuid.UUID	 `json:"id"`
	PlantTypeEN      string		 `json:"name_en"`
	PlantTypeTH      string		 `json:"name_th"`
	CreateDate		 time.Time	 `json:"-"`
	ChangeDate	     time.Time	 `json:"-"`
	StatusId		 uuid.UUID	 `json:"-"`
}
// New instance plant_type
func (u *PlantType) New() *PlantType {
	return &PlantType{
		PlantTypeId:	u.PlantTypeId ,
		PlantTypeEN:	u.PlantTypeEN ,
		PlantTypeTH:	u.PlantTypeTH ,
		CreateDate:		u.CreateDate ,
		ChangeDate:		u.ChangeDate ,
		StatusId:		u.StatusId ,
	}
}

// Custom table name for GORM
func (PlantType) TableName() string {
	return config.DB_PLANT_TYPE
}
