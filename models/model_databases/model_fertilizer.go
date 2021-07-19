package model_databases

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table fertilizer
//-------------------------------------------------------------------------------//
//model fertilizer
type Fertilizer struct {
	FertilizerId     uuid.UUID	 `gorm:"fertilizer_id" json:"fertilizer_id,omitempty"`
	FertilizerEN     string		 `gorm:"fertilizer_en" json:"fertilizer_en,omitempty"`
	FertilizerTH     string		 `gorm:"fertilizer_th" json:"fertilizer_th,omitempty"`
	Nitrogen       	 float64	 `gorm:"nitrogen" json:"nitrogen,omitempty"`
	Phosphorus    	 float64	 `gorm:"phosphorus" json:"phosphorus,omitempty"`
	Potassium      	 float64	 `gorm:"potassium" json:"potassium,omitempty"`
	CreateDate		 time.Time	 `gorm:"create_date" json:"create_date,omitempty"`
	ChangeDate	     time.Time	 `gorm:"change_date" json:"change_date,omitempty"`
	FertCatId		 uuid.UUID	 `gorm:"fertilizer_cat_id" json:"fertilizer_cat_id,omitempty"`
	StatusId		 uuid.UUID	 `gorm:"status_id" json:"status_id,omitempty"`
}
// New instance fertilizer
func (u *Fertilizer) New() *Fertilizer {
	return &Fertilizer{
		FertilizerId:	u.FertilizerId ,
		FertilizerEN:	u.FertilizerEN ,
		FertilizerTH:	u.FertilizerTH ,
		Nitrogen:		u.Nitrogen ,
		Phosphorus:		u.Phosphorus ,
		Potassium:		u.Potassium ,
		CreateDate:		u.CreateDate ,
		ChangeDate:		u.ChangeDate ,
		FertCatId:		u.FertCatId ,
		StatusId:		u.StatusId ,
	}
}
