package model_db

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/common/config"
)

//-------------------------------------------------------------------------------//
//							Table favorite_plant
//-------------------------------------------------------------------------------//
//model favorite_plant
type FavoritePlant struct {
	DBCommonGet
	Uid          	 uuid.UUID	 `json:"uid"`
	FormulaPlantId   uuid.UUID	 `json:"formula_plant_id"`
}
// New instance favorite_plant
func (u *FavoritePlant) New() *FavoritePlant {
	return &FavoritePlant{
		DBCommonGet:      		u.DBCommonGet ,
		Uid:				u.Uid ,
		FormulaPlantId:		u.FormulaPlantId ,
	}
}

// Custom table name for GORM
func (FavoritePlant) TableName() string {
	return config.DB_FAVORITE_PLANT
}

