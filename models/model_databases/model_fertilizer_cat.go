package model_databases

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/config"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table fertilizer_cat
//-------------------------------------------------------------------------------//
//model fertilizer_cat
type FertilizerCat struct {
	FertilizerCatId      uuid.UUID	 `mapstructure:"fertilizer_cat_id" json:"fertilizer_cat_id,omitempty"`
	FertilizerCatEN      string		 `mapstructure:"fertilizer_cat_en" json:"fertilizer_cat_en,omitempty"`
	FertilizerCatTH      string		 `mapstructure:"fertilizer_cat_th" json:"fertilizer_cat_th,omitempty"`
	ChangeDate	   		 time.Time	 `mapstructure:"change_date" json:"change_date,omitempty"`
	CreateDate	   		 time.Time	 `mapstructure:"create_date" json:"create_date,omitempty"`
	StatusId	   		 uuid.UUID	 `mapstructure:"status_id" json:"status_id,omitempty"`
}
// New instance fertilizer_cat
func (u *FertilizerCat) New() *FertilizerCat {
	return &FertilizerCat{
		FertilizerCatId:		u.FertilizerCatId ,
		FertilizerCatEN:		u.FertilizerCatEN ,
		FertilizerCatTH:		u.FertilizerCatTH ,
		CreateDate:				u.CreateDate ,
		ChangeDate:				u.ChangeDate ,
		StatusId:				u.StatusId ,
	}
}

// Custom table name for GORM
func (FertilizerCat) TableName() string {
	return config.DB_FERTILIZER_CAT
}

