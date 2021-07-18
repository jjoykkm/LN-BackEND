package models

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table plant
//-------------------------------------------------------------------------------//
//model plant
type Plant struct {
	PlantId          uuid.UUID	 `gorm:"plant_id" json:"plant_id,omitempty"`
	PlantNameEN      string		 `gorm:"plant_name_en" json:"plant_name_en,omitempty"`
	PlantNameTH      string		 `gorm:"plant_name_th" json:"plant_name_th,omitempty"`
	PlantDescEN      string		 `gorm:"plant_desc_en" json:"plant_desc_en,omitempty"`
	PlantDescTH      string		 `gorm:"plant_desc_th" json:"plant_desc_th,omitempty"`
	CreateDate		 time.Time	 `gorm:"create_date" json:"create_date,omitempty"`
	ChangeDate	     time.Time	 `gorm:"change_date" json:"change_date,omitempty"`
	StatusId		 uuid.UUID	 `gorm:"status_id" json:"status_id,omitempty"`
	PlantTypeId      uuid.UUID	 `gorm:"plant_type_id" json:"plant_type_id,omitempty"`
}
// New instance plant
func (u *Plant) New() *Plant {
	return &Plant{
		PlantId:		u.PlantId ,
		PlantNameEN:	u.PlantNameEN ,
		PlantNameTH:	u.PlantNameTH ,
		PlantDescEN:	u.PlantDescEN ,
		PlantDescTH:	u.PlantDescTH ,
		CreateDate:		u.CreateDate ,
		ChangeDate:		u.ChangeDate ,
		StatusId:		u.StatusId ,
		PlantTypeId:	u.PlantTypeId ,
	}
}
