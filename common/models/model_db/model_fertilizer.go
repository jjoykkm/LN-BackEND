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
	FertId     	 		 uuid.UUID	 `json:"fert_id,omitempty" gorm:"column:fertilizer_id"`
	FertEN     	 		 string		 `json:"name_en,omitempty" gorm:"column:fertilizer_en"`
	FertTH     	 		 string		 `json:"name_th,omitempty" gorm:"column:fertilizer_th"`
	Nitrogen       	 	 float64	 `json:"nitrogen,omitempty"`
	Phosphorus    	 	 float64	 `json:"phosphorus,omitempty"`
	Potassium      	 	 float64	 `json:"potassium,omitempty"`
	CreateDate		 	 time.Time	 `json:"create_date,omitempty"`
	ChangeDate	     	 time.Time	 `json:"change_date,omitempty"`
	FertCatId		 	 uuid.UUID	 `json:"fert_cat_id,omitempty" gorm:"column:fertilizer_cat_id"`
	StatusId		 	 uuid.UUID	 `json:"status_id,omitempty"`
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

