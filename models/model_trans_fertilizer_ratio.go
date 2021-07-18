package models

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table trans_fertilizer_Ratio
//-------------------------------------------------------------------------------//
//model trans_fertilizer_Ratio
type TransFertRatio struct {
	PlantId          	uuid.UUID	 `gorm:"fertilizer_id" json:"fertilizer_id,omitempty"`
	FormulaPlantId		uuid.UUID	 `gorm:"formula_plant_id" json:"formula_plant_id,omitempty"`
	StatusId		 	uuid.UUID	 `gorm:"status_id" json:"status_id,omitempty"`
	Ratio		     	float64		 `gorm:"Ratio" json:"Ratio,omitempty"`
	CreateDate		 	time.Time	 `gorm:"create_date" json:"create_date,omitempty"`
	ChangeDate	     	time.Time	 `gorm:"change_date" json:"change_date,omitempty"`
}
// New instance trans_fertilizer_Ratio
func (u *TransFertRatio) New() *TransFertRatio {
	return &TransFertRatio{
		PlantId:			u.PlantId ,
		FormulaPlantId:		u.FormulaPlantId ,
		StatusId:			u.StatusId ,
		Ratio:				u.Ratio ,
		CreateDate:			u.CreateDate ,
		ChangeDate:			u.ChangeDate ,
	}
}
