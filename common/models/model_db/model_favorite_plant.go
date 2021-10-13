package model_db

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/common/config"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table favorite_plant
//-------------------------------------------------------------------------------//
//model favorite_plant
type FavoritePlant struct {
	Uid          	 uuid.UUID	 `json:"uid"`
	FormulaPlantId   uuid.UUID	 `json:"formula_plant_id"`
	CreateDate		 time.Time	 `json:"create_date"`
	ChangeDate	     time.Time	 `json:"change_date"`
	StatusId		 uuid.UUID	 `json:"status_id"`
}
// New instance favorite_plant
func (u *FavoritePlant) New() *FavoritePlant {
	return &FavoritePlant{
		Uid:				u.Uid ,
		FormulaPlantId:		u.FormulaPlantId ,
		CreateDate:			u.CreateDate ,
		ChangeDate:			u.ChangeDate ,
		StatusId:			u.StatusId ,
	}
}

// Custom table name for GORM
func (FavoritePlant) TableName() string {
	return config.DB_FAVORITE_PLANT
}

