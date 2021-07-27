package model_databases

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table fertilizer_cat
//-------------------------------------------------------------------------------//
//model fertilizer_cat
type FertilizerCat struct {
	FertCatId      uuid.UUID	 `mapstructure:"fertilizer_cat_id" json:"fertilizer_cat_id,omitempty"`
	FertCatEN      string		 `mapstructure:"fertilizer_cat_en" json:"fertilizer_cat_en,omitempty"`
	FertCatTH      string		 `mapstructure:"fertilizer_cat_th" json:"fertilizer_cat_th,omitempty"`
	CreateDate	   time.Time	 `mapstructure:"create_date" json:"create_date,omitempty"`
	ChangeDate	   time.Time	 `mapstructure:"change_date" json:"change_date,omitempty"`
	StatusId	   uuid.UUID	 `mapstructure:"status_id" json:"status_id,omitempty"`
}
// New instance fertilizer_cat
func (u *FertilizerCat) New() *FertilizerCat {
	return &FertilizerCat{
		FertCatId:		u.FertCatId ,
		FertCatEN:		u.FertCatEN ,
		FertCatTH:		u.FertCatTH ,
		CreateDate:		u.CreateDate ,
		ChangeDate:		u.ChangeDate ,
		StatusId:		u.StatusId ,
	}
}

