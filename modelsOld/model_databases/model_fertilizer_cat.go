package model_databases

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/common/config"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table fertilizer_cat
//-------------------------------------------------------------------------------//
//model fertilizer_cat
type FertilizerCat struct {
	FertilizerCatId      uuid.UUID	 `json:"fertilizer_cat_id"`
	FertilizerCatEN      string		 `json:"fertilizer_cat_en"`
	FertilizerCatTH      string		 `json:"fertilizer_cat_th"`
	ChangeDate	   		 time.Time	 `json:"change_date"`
	CreateDate	   		 time.Time	 `json:"create_date"`
	StatusId	   		 uuid.UUID	 `json:"status_id"`
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

