package model_databases

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/config"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table plant
//-------------------------------------------------------------------------------//
//model plant
type Plant struct {
	PlantId          uuid.UUID	 `mapstructure:"plant_id" json:"id,omitempty"`
	PlantNameEN      string		 `mapstructure:"plant_name_en" json:"name_en,omitempty"`
	PlantNameTH      string		 `mapstructure:"plant_name_th" json:"name_th,omitempty"`
	PlantDescEN      string		 `mapstructure:"plant_desc_en" json:"desc_en,omitempty"`
	PlantDescTH      string		 `mapstructure:"plant_desc_th" json:"desc_th,omitempty"`
	CreateDate		 time.Time	 `mapstructure:"create_date" json:"-"`
	ChangeDate	     time.Time	 `mapstructure:"change_date" json:"-"`
	StatusId		 uuid.UUID	 `mapstructure:"status_id" json:"-"`
	PlantTypeId      *uuid.UUID	 `mapstructure:"plant_type_id" json:"type_id,omitempty"`
	TotalItem      	 int		 `mapstructure:"total_item" json:"total_item,omitempty"`
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
