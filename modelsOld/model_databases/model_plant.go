package model_databases

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
	PlantId          uuid.UUID	 `json:"id,omitempty"`
	PlantNameEN      string		 `json:"name_en,omitempty"`
	PlantNameTH      string		 `json:"name_th,omitempty"`
	PlantDescEN      string		 `json:"desc_en,omitempty"`
	PlantDescTH      string		 `json:"desc_th,omitempty"`
	CreateDate		 time.Time	 `json:"-"`
	ChangeDate	     time.Time	 `json:"-"`
	StatusId		 uuid.UUID	 `json:"-"`
	PlantTypeId      *uuid.UUID	 `json:"type_id,omitempty"`
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
