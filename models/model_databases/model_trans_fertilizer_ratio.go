package model_databases

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table trans_fertilizer_Ratio
//-------------------------------------------------------------------------------//
//model trans_fertilizer_Ratio
type TransFertRatio struct {
	FertilizerId       	uuid.UUID	 `mapstructure:"fertilizer_id" json:"fertilizer_id,omitempty"`
	FormulaPlantId		uuid.UUID	 `mapstructure:"formula_plant_id" json:"formula_plant_id,omitempty"`
	StatusId		 	uuid.UUID	 `mapstructure:"status_id" json:"status_id,omitempty"`
	Ratio		     	float64		 `mapstructure:"ratio" json:"ratio,omitempty"`
	CreateDate		 	time.Time	 `mapstructure:"create_date" json:"create_date,omitempty"`
	ChangeDate	     	time.Time	 `mapstructure:"change_date" json:"change_date,omitempty"`
}
// New instance trans_fertilizer_Ratio
func (u *TransFertRatio) New() *TransFertRatio {
	return &TransFertRatio{
		FertilizerId:		u.FertilizerId ,
		FormulaPlantId:		u.FormulaPlantId ,
		StatusId:			u.StatusId ,
		Ratio:				u.Ratio ,
		CreateDate:			u.CreateDate ,
		ChangeDate:			u.ChangeDate ,
	}
}
