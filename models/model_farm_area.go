package models

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table farm_area
//-------------------------------------------------------------------------------//
//model farm_area
type FarmArea struct {
	FarmAreaId      	uuid.UUID	 `gorm:"farm_area_id" json:"farm_area_id,omitempty"`
	FarmAreaName    	string		 `gorm:"farm_area_name" json:"farm_area_name,omitempty"`
	CreateDate			time.Time	 `gorm:"create_date" json:"create_date,omitempty"`
	ChangeDate	    	time.Time	 `gorm:"change_date" json:"change_date,omitempty"`
	FarmId				uuid.UUID	 `gorm:"farm_id" json:"farm_id,omitempty"`
	FormulaPlantId		uuid.UUID	 `gorm:"formula_plant_id" json:"formula_plant_id,omitempty"`
	StatusId			uuid.UUID	 `gorm:"status_id" json:"status_id,omitempty"`
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

