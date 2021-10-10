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
	PlantId          uuid.UUID	 `json:"PlantId,omitempty"`
	PlantNameEN      string		 `json:"PlantNameEN,omitempty"`
	PlantNameTH      string		 `json:"PlantNameTH,omitempty"`
	PlantDescEN      string		 `json:"PlantDescEN,omitempty"`
	PlantDescTH      string		 `json:"PlantDescTH,omitempty"`
	CreateDate		 time.Time	 `json:"CreateDate,omitempty"`
	ChangeDate	     time.Time	 `json:"ChangeDate,omitempty"`
	StatusId		 uuid.UUID	 `json:"StatusId,omitempty"`
	PlantTypeId      *uuid.UUID	 `json:"PlantTypeId,omitempty"`
	TotalItem      	 int		 `json:"TotalItem,omitempty"`
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
