package model_db

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/common/config"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table farm_area
//-------------------------------------------------------------------------------//
//model farm_area
type FarmArea struct {
	FarmAreaId      	uuid.UUID	 `json:"farm_area_id,omitempty"`
	FarmAreaName    	string		 `json:"farm_area_name,omitempty"`
	CreateDate			time.Time	 `json:"create_date,omitempty"`
	ChangeDate	    	time.Time	 `json:"change_date,omitempty"`
	FarmId				uuid.UUID	 `json:"farm_id,omitempty"`
	FormulaPlantId		uuid.UUID	 `json:"formula_plant_id,omitempty"`
	StatusId			uuid.UUID	 `json:"status_id,omitempty"`
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


