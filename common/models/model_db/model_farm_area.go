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
	FarmAreaId      	uuid.UUID	 `json:"FarmAreaId,omitempty"`
	FarmAreaName    	string		 `json:"FarmAreaName,omitempty"`
	CreateDate			time.Time	 `json:"CreateDate,omitempty"`
	ChangeDate	    	time.Time	 `json:"ChangeDate,omitempty"`
	FarmId				uuid.UUID	 `json:"FarmId,omitempty"`
	FormulaPlantId		uuid.UUID	 `json:"FormulaPlantId,omitempty"`
	StatusId			uuid.UUID	 `json:"StatusId,omitempty"`
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


