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
	FarmAreaId      	uuid.UUID	 `json:"farm_area_id"`
	FarmAreaName    	string		 `json:"farm_area_name"`
	CreateDate			time.Time	 `json:"create_date"`
	ChangeDate	    	time.Time	 `json:"change_date"`
	FarmId				uuid.UUID	 `json:"farm_id"`
	FormulaPlantId		uuid.UUID	 `json:"formula_plant_id"`
	StatusId			uuid.UUID	 `json:"status_id"`
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


