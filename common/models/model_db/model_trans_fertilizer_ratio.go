package model_db

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/config"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table trans_fertilizer_Ratio
//-------------------------------------------------------------------------------//
//model trans_fertilizer_Ratio
type TransFertRatio struct {
	FertId       		uuid.UUID	 `json:"fert_id,omitempty" gorm:"column:fertilizer_id"`
	FormulaPlantId		uuid.UUID	 `json:"formula_plant_id,omitempty"`
	StatusId		 	uuid.UUID	 `json:"status_id,omitempty"`
	Ratio		     	float64		 `json:"ratio,omitempty"`
	CreateDate		 	time.Time	 `json:"create_date,omitempty"`
	ChangeDate	     	time.Time	 `json:"change_date,omitempty"`
}
// New instance trans_fertilizer_Ratio
func (u *TransFertRatio) New() *TransFertRatio {
	return &TransFertRatio{
		FertId:				u.FertId ,
		FormulaPlantId:		u.FormulaPlantId ,
		StatusId:			u.StatusId ,
		Ratio:				u.Ratio ,
		CreateDate:			u.CreateDate ,
		ChangeDate:			u.ChangeDate ,
	}
}

// Custom table name for GORM
func (TransFertRatio) TableName() string {
	return config.DB_TRANS_FERTILIZER_RATIO
}
