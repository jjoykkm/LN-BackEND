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
	PlantId          uuid.UUID	 `json:"plant_id"`
	PlantNameEN      string		 `json:"plant_name_en"`
	PlantNameTH      string		 `json:"plant_name_th"`
	PlantDescEN      string		 `json:"plant_desc_en"`
	PlantDescTH      string		 `json:"plant_desc_th"`
	CreateDate		 time.Time	 `json:"create_date"`
	ChangeDate	     time.Time	 `json:"change_date"`
	StatusId		 uuid.UUID	 `json:"status_id"`
	PlantTypeId      *uuid.UUID	 `json:"plant_type_id"`
	TotalItem      	 int		 `json:"total_item"`
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
