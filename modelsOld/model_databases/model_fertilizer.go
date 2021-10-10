package model_databases

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
	FertilizerId     	 uuid.UUID	 `json:"fertilizer_id,omitempty"`
	FertilizerEN     	 string		 `json:"fertilizer_en,omitempty"`
	FertilizerTH     	 string		 `json:"fertilizer_th,omitempty"`
	Nitrogen       	 	 float64	 `json:"nitrogen,omitempty"`
	Phosphorus    	 	 float64	 `json:"phosphorus,omitempty"`
	Potassium      	 	 float64	 `json:"potassium,omitempty"`
	CreateDate		 	 time.Time	 `json:"create_date,omitempty"`
	ChangeDate	     	 time.Time	 `json:"change_date,omitempty"`
	FertilizerCatId		 uuid.UUID	 `json:"fertilizer_cat_id,omitempty"`
	StatusId		 	 uuid.UUID	 `json:"status_id,omitempty"`
}
// New instance fertilizer
func (u *Fertilizer) New() *Fertilizer {
	return &Fertilizer{
		FertilizerId:		u.FertilizerId ,
		FertilizerEN:		u.FertilizerEN ,
		FertilizerTH:		u.FertilizerTH ,
		Nitrogen:			u.Nitrogen ,
		Phosphorus:			u.Phosphorus ,
		Potassium:			u.Potassium ,
		CreateDate:			u.CreateDate ,
		ChangeDate:			u.ChangeDate ,
		FertilizerCatId:	u.FertilizerCatId ,
		StatusId:			u.StatusId ,
	}
}

// Custom table name for GORM
func (Fertilizer) TableName() string {
	return config.DB_FERTILIZER
}

