package model_db

import (
	"github.com/jjoykkm/ln-backend/common/config"
)

//-------------------------------------------------------------------------------//
//							Table farm_area
//-------------------------------------------------------------------------------//
//model farm_area
type FarmArea struct {
	DBCommonGet
	FarmAreaId      	string	 `json:"farm_area_id"`
	FarmAreaName    	string	 `json:"farm_area_name"`
	FarmId				string	 `json:"farm_id"`
	FormulaPlantId		string	 `json:"formula_plant_id"`
}
// New instance farm
func (u *FarmArea) New() *FarmArea {
	return &FarmArea{
		DBCommonGet:      		u.DBCommonGet ,
		FarmAreaId:			u.FarmAreaId ,
		FarmAreaName:		u.FarmAreaName ,
		FarmId:				u.FarmId ,
		FormulaPlantId:		u.FormulaPlantId ,
	}
}
// Custom table name for GORM
func (FarmArea) TableName() string {
	return config.DB_FARM_AREA
}

//-------------------------------------------------------------------------------//
//							Upsert FarmArea
//-------------------------------------------------------------------------------//
type FarmAreaUS struct {
	FarmAreaId      	string	 `json:"farm_area_id" gorm:"default:uuid_generate_v4()"`
	FarmAreaName    	string	 `json:"farm_area_name"`
	FarmId				string	 `json:"farm_id"`
	FormulaPlantId		string	 `json:"formula_plant_id"`
	StatusId			string	 `json:"status_id"`
}
func (FarmAreaUS) TableName() string {
	return config.DB_FARM_AREA
}


