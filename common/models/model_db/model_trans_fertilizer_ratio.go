package model_db

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/common/config"
)

//-------------------------------------------------------------------------------//
//							Table trans_fertilizer_Ratio
//-------------------------------------------------------------------------------//
//model trans_fertilizer_Ratio
type TransFertRatio struct {
	DBCommonGet
	FertId       		uuid.UUID	 `json:"fert_id" gorm:"column:fertilizer_id"`
	FormulaPlantId		uuid.UUID	 `json:"formula_plant_id"`
	Ratio		     	float64		 `json:"ratio"`
}
// New instance trans_fertilizer_Ratio
func (u *TransFertRatio) New() *TransFertRatio {
	return &TransFertRatio{
		DBCommonGet:      		u.DBCommonGet ,
		FertId:				u.FertId ,
		FormulaPlantId:		u.FormulaPlantId ,
		Ratio:				u.Ratio ,
	}
}

// Custom table name for GORM
func (TransFertRatio) TableName() string {
	return config.DB_TRANS_FERTILIZER_RATIO
}

//-------------------------------------------------------------------------------//
//									Upsert
//-------------------------------------------------------------------------------//
type TransFertRatioUS struct {
	DBCommonCreateUpdate
	FertId       		string	 	`json:"fert_id" gorm:"column:fertilizer_id"`
	FormulaPlantId		string	 	`json:"formula_plant_id"`
	Ratio		     	float64		`json:"ratio"`
	StatusId		 	string		`json:"status_id"`
}
func (TransFertRatioUS) TableName() string {
	return config.DB_TRANS_FERTILIZER_RATIO
}
