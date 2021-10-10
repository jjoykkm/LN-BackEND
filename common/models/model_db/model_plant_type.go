package model_db

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
	PlantTypeId      uuid.UUID	 `json:"PlantTypeId,omitempty"`
	PlantTypeEN      string		 `json:"PlantTypeEN,omitempty"`
	PlantTypeTH      string		 `json:"PlantTypeTH,omitempty"`
	CreateDate		 time.Time	 `json:"CreateDate,omitempty"`
	ChangeDate	     time.Time	 `json:"ChangeDate,omitempty"`
	StatusId		 uuid.UUID	 `json:"StatusId,omitempty"`
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
