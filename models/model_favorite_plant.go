package models

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table favorite_plant
//-------------------------------------------------------------------------------//
//model favorite_plant
type FavoritePlant struct {
	Uid          	 uuid.UUID	 `gorm:"uid" json:"uid,omitempty"`
	FormulaPlantId   uuid.UUID	 `gorm:"formula_plant_id" json:"formula_plant_id,omitempty"`
	CreateDate		 time.Time	 `gorm:"create_date" json:"create_date,omitempty"`
	ChangeDate	     time.Time	 `gorm:"change_date" json:"change_date,omitempty"`
	StatusId		 uuid.UUID	 `gorm:"status_id" json:"status_id,omitempty"`
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
