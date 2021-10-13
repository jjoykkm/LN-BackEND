package model_db

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/common/config"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table trans_fertilizer_Ratio
//-------------------------------------------------------------------------------//
//model trans_fertilizer_Ratio
type TransFertRatio struct {
	FertId       		uuid.UUID	 `json:"fert_id" gorm:"column:fertilizer_id"`
	FormulaPlantId		uuid.UUID	 `json:"formula_plant_id"`
	StatusId		 	uuid.UUID	 `json:"status_id"`
	Ratio		     	float64		 `json:"ratio"`
	CreateDate		 	time.Time	 `json:"create_date"`
	ChangeDate	     	time.Time	 `json:"change_date"`
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
