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
	FertId     		uuid.UUID	`json:"FertId,omitempty" gorm:"column:fertilizer_id"`
	FertEN     		string		`json:"FertEN,omitempty" gorm:"column:fertilizer_en"`
	FertTH     		string		`json:"FertTH,omitempty" gorm:"column:fertilizer_th"`
	Nitrogen       	float64		`json:"Nitrogen,omitempty"`
	Phosphorus    	float64		`json:"Phosphorus,omitempty"`
	Potassium      	float64		`json:"Potassium,omitempty"`
	CreateDate		time.Time	`json:"CreateDate,omitempty"`
	ChangeDate	   	time.Time	`json:"ChangeDate,omitempty"`
	FertCatId	 	uuid.UUID	`json:"FertCatId,omitempty" gorm:"column:fertilizer_cat_id"`
	StatusId		uuid.UUID	`json:"StatusId,omitempty"`
}
// New instance fertilizer
func (u *Fertilizer) New() *Fertilizer {
	return &Fertilizer{
		FertId:			u.FertId ,
		FertEN:			u.FertEN ,
		FertTH:			u.FertTH ,
		Nitrogen:		u.Nitrogen ,
		Phosphorus:		u.Phosphorus ,
		Potassium:		u.Potassium ,
		CreateDate:		u.CreateDate ,
		ChangeDate:		u.ChangeDate ,
		FertCatId:		u.FertCatId ,
		StatusId:		u.StatusId ,
	}
}

// Custom table name for GORM
func (Fertilizer) TableName() string {
	return config.DB_FERTILIZER
}

