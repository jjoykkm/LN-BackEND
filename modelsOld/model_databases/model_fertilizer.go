package model_databases

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/config"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table fertilizer
//-------------------------------------------------------------------------------//
//model fertilizer
type Fertilizer struct {
	FertilizerId     	 uuid.UUID	 `mapstructure:"fertilizer_id" json:"fertilizer_id,omitempty"`
	FertilizerEN     	 string		 `mapstructure:"fertilizer_en" json:"fertilizer_en,omitempty"`
	FertilizerTH     	 string		 `mapstructure:"fertilizer_th" json:"fertilizer_th,omitempty"`
	Nitrogen       	 	 float64	 `mapstructure:"nitrogen" json:"nitrogen,omitempty"`
	Phosphorus    	 	 float64	 `mapstructure:"phosphorus" json:"phosphorus,omitempty"`
	Potassium      	 	 float64	 `mapstructure:"potassium" json:"potassium,omitempty"`
	CreateDate		 	 time.Time	 `mapstructure:"create_date" json:"create_date,omitempty"`
	ChangeDate	     	 time.Time	 `mapstructure:"change_date" json:"change_date,omitempty"`
	FertilizerCatId		 uuid.UUID	 `mapstructure:"fertilizer_cat_id" json:"fertilizer_cat_id,omitempty"`
	StatusId		 	 uuid.UUID	 `mapstructure:"status_id" json:"status_id,omitempty"`
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

