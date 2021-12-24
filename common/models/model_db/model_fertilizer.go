package model_db

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/common/config"
)

//-------------------------------------------------------------------------------//
//							Table fertilizer
//-------------------------------------------------------------------------------//
//model fertilizer
type Fertilizer struct {
	DBCommonGet
	FertId     	 		 uuid.UUID	 `json:"fert_id" gorm:"column:fertilizer_id"`
	FertEN     	 		 string		 `json:"fert_name_en" gorm:"column:fertilizer_en"`
	FertTH     	 		 string		 `json:"fert_name_th" gorm:"column:fertilizer_th"`
	Nitrogen       	 	 float64	 `json:"nitrogen"`
	Phosphorus    	 	 float64	 `json:"phosphorus"`
	Potassium      	 	 float64	 `json:"potassium"`
	FertCatId		 	 uuid.UUID	 `json:"fert_cat_id" gorm:"column:fertilizer_cat_id"`
}
// New instance fertilizer
func (u *Fertilizer) New() *Fertilizer {
	return &Fertilizer{
		DBCommonGet:      	u.DBCommonGet ,
		FertId:				u.FertId ,
		FertEN:				u.FertEN ,
		FertTH:				u.FertTH ,
		Nitrogen:			u.Nitrogen ,
		Phosphorus:			u.Phosphorus ,
		Potassium:			u.Potassium ,
		FertCatId:			u.FertCatId ,
	}
}

// Custom table name for GORM
func (Fertilizer) TableName() string {
	return config.DB_FERTILIZER
}

//-------------------------------------------------------------------------------//
//									Upsert
//-------------------------------------------------------------------------------//
type FertilizerUS struct {
	DBCommonCreateUpdate
	FertId     	 		 string	 	 `json:"fert_id" gorm:"column:fertilizer_id"`
	FertEN     	 		 string		 `json:"fert_name_en" gorm:"column:fertilizer_en"`
	FertTH     	 		 string		 `json:"fert_name_th" gorm:"column:fertilizer_th"`
	Nitrogen       	 	 float64	 `json:"nitrogen"`
	Phosphorus    	 	 float64	 `json:"phosphorus"`
	Potassium      	 	 float64	 `json:"potassium"`
	FertCatId		 	 string	 	 `json:"fert_cat_id" gorm:"column:fertilizer_cat_id"`
	StatusId		 	 string		 `json:"status_id"`
}
func (FertilizerUS) TableName() string {
	return config.DB_FERTILIZER
}

