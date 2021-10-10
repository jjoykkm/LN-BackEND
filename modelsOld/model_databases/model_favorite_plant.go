package model_databases

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
	Uid          	 uuid.UUID	 `json:"uid,omitempty"`
	FormulaPlantId   uuid.UUID	 `json:"formula_plant_id,omitempty"`
	CreateDate		 time.Time	 `json:"create_date,omitempty"`
	ChangeDate	     time.Time	 `json:"change_date,omitempty"`
	StatusId		 uuid.UUID	 `json:"status_id,omitempty"`
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

