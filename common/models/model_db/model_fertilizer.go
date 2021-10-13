package model_db

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/common/config"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table fertilizer
//-------------------------------------------------------------------------------//
//model fertilizer
type Fertilizer struct {
	FertId     	 		 uuid.UUID	 `json:"fert_id" gorm:"column:fertilizer_id"`
	FertEN     	 		 string		 `json:"fert_name_en" gorm:"column:fertilizer_en"`
	FertTH     	 		 string		 `json:"fert_name_th" gorm:"column:fertilizer_th"`
	Nitrogen       	 	 float64	 `json:"nitrogen"`
	Phosphorus    	 	 float64	 `json:"phosphorus"`
	Potassium      	 	 float64	 `json:"potassium"`
	CreateDate		 	 time.Time	 `json:"create_date"`
	ChangeDate	     	 time.Time	 `json:"change_date"`
	FertCatId		 	 uuid.UUID	 `json:"fert_cat_id" gorm:"column:fertilizer_cat_id"`
	StatusId		 	 uuid.UUID	 `json:"status_id"`
}
// New instance fertilizer
func (u *Fertilizer) New() *Fertilizer {
	return &Fertilizer{
		FertId:				u.FertId ,
		FertEN:				u.FertEN ,
		FertTH:				u.FertTH ,
		Nitrogen:			u.Nitrogen ,
		Phosphorus:			u.Phosphorus ,
		Potassium:			u.Potassium ,
		CreateDate:			u.CreateDate ,
		ChangeDate:			u.ChangeDate ,
		FertCatId:			u.FertCatId ,
		StatusId:			u.StatusId ,
	}
}

// Custom table name for GORM
func (Fertilizer) TableName() string {
	return config.DB_FERTILIZER
}

