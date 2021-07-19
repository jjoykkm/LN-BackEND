package model_databases

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table plant_type
//-------------------------------------------------------------------------------//
//model plant_type
type PlantType struct {
	PlantTypeId      uuid.UUID	 `gorm:"plant_type_id" json:"plant_type_id,omitempty"`
	PlantTypeEN      string		 `gorm:"plant_type_en" json:"plant_type_en,omitempty"`
	PlantTypeTH      string		 `gorm:"plant_type_th" json:"plant_type_th,omitempty"`
	CreateDate		 time.Time	 `gorm:"create_date" json:"create_date,omitempty"`
	ChangeDate	     time.Time	 `gorm:"change_date" json:"change_date,omitempty"`
	StatusId		 uuid.UUID	 `gorm:"status_id" json:"status_id,omitempty"`
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

