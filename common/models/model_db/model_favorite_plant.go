package model_db

import (
	"github.com/jjoykkm/ln-backend/common/config"
)

//-------------------------------------------------------------------------------//
//							Table favorite_plant
//-------------------------------------------------------------------------------//
//model favorite_plant
type FavForPlant struct {
	DBCommonGet
	Uid          	 string	 `json:"uid"`
	FormulaPlantId   string	 `json:"formula_plant_id"`
}
// New instance favorite_plant
func (u *FavForPlant) New() *FavForPlant {
	return &FavForPlant{
		DBCommonGet:      	u.DBCommonGet ,
		Uid:				u.Uid ,
		FormulaPlantId:		u.FormulaPlantId ,
	}
}

// Custom table name for GORM
func (FavForPlant) TableName() string {
	return config.DB_FAVORITE_PLANT
}

//-------------------------------------------------------------------------------//
//									Upsert
//-------------------------------------------------------------------------------//
type FavForPlantUS struct {
	DBCommonCreateUpdate
	Uid          	 string	 `json:"uid"`
	FormulaPlantId   string	 `json:"formula_plant_id"`
	StatusId		 string	 `json:"status_id"`
}
func (FavForPlantUS) TableName() string {
	return config.DB_FAVORITE_PLANT
}
