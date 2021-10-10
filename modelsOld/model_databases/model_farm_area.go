package model_databases

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/config"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table farm_area
//-------------------------------------------------------------------------------//
//model farm_area
type FarmArea struct {
	FarmAreaId      	uuid.UUID	 `mapstructure:"farm_area_id" json:"farm_area_id,omitempty"`
	FarmAreaName    	string		 `mapstructure:"farm_area_name" json:"farm_area_name,omitempty"`
	CreateDate			time.Time	 `mapstructure:"create_date" json:"create_date,omitempty"`
	ChangeDate	    	time.Time	 `mapstructure:"change_date" json:"change_date,omitempty"`
	FarmId				uuid.UUID	 `mapstructure:"farm_id" json:"farm_id,omitempty"`
	FormulaPlantId		uuid.UUID	 `mapstructure:"formula_plant_id" json:"formula_plant_id,omitempty"`
	StatusId			uuid.UUID	 `mapstructure:"status_id" json:"status_id,omitempty"`
}
// New instance farm_area
func (u *FarmArea) New() *FarmArea {
	return &FarmArea{
		FarmAreaId:			u.FarmAreaId ,
		FarmAreaName:		u.FarmAreaName ,
		CreateDate:			u.CreateDate ,
		ChangeDate:			u.ChangeDate ,
		FarmId:				u.FarmId ,
		FormulaPlantId:		u.FormulaPlantId ,
		StatusId:			u.StatusId ,
	}
}

// Custom table name for GORM
func (FarmArea) TableName() string {
	return config.DB_FARM_AREA
}


