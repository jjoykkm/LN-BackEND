package model_db

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/common/config"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table plant
//-------------------------------------------------------------------------------//
//model plant
type Plant struct {
	PlantId          uuid.UUID	 `json:"plant_id,omitempty"`
	PlantNameEN      string		 `json:"plant_name_en,omitempty"`
	PlantNameTH      string		 `json:"plant_name_th,omitempty"`
	PlantDescEN      string		 `json:"plant_desc_en,omitempty"`
	PlantDescTH      string		 `json:"plant_desc_th,omitempty"`
	CreateDate		 time.Time	 `json:"create_date,omitempty"`
	ChangeDate	     time.Time	 `json:"change_date,omitempty"`
	StatusId		 uuid.UUID	 `json:"status_id,omitempty"`
	PlantTypeId      *uuid.UUID	 `json:"plant_type_id,omitempty"`
	TotalItem      	 int		 `json:"total_item,omitempty"`
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
		TotalItem:		u.TotalItem ,
	}
}

// Custom table name for GORM
func (Plant) TableName() string {
	return config.DB_PLANT
}
