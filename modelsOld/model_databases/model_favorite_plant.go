package model_databases

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/config"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table favorite_plant
//-------------------------------------------------------------------------------//
//model favorite_plant
type FavoritePlant struct {
	Uid          	 uuid.UUID	 `mapstructure:"uid" json:"uid,omitempty"`
	FormulaPlantId   uuid.UUID	 `mapstructure:"formula_plant_id" json:"formula_plant_id,omitempty"`
	CreateDate		 time.Time	 `mapstructure:"create_date" json:"create_date,omitempty"`
	ChangeDate	     time.Time	 `mapstructure:"change_date" json:"change_date,omitempty"`
	StatusId		 uuid.UUID	 `mapstructure:"status_id" json:"status_id,omitempty"`
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

