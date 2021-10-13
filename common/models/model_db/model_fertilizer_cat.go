package model_db

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
	FertCatId      uuid.UUID	 `json:"fert_cat_id" gorm:"column:fertilizer_cat_id"`
	FertCatEN      string		 `json:"fert_cat_en" gorm:"column:fertilizer_cat_en"`
	FertCatTH      string		 `json:"fert_cat_th" gorm:"column:fertilizer_cat_th"`
	ChangeDate	   time.Time	 `json:"change_date"`
	CreateDate	   time.Time	 `json:"create_date"`
	StatusId	   uuid.UUID	 `json:"status_id"`
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

// Custom table name for GORM
func (FertilizerCat) TableName() string {
	return config.DB_FERTILIZER_CAT
}

