package model_db

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/common/config"
)

//-------------------------------------------------------------------------------//
//							Table fertilizer_cat
//-------------------------------------------------------------------------------//
//model fertilizer_cat
type FertilizerCat struct {
	DBCommon
	FertCatId      uuid.UUID	 `json:"fert_cat_id" gorm:"column:fertilizer_cat_id"`
	FertCatEN      string		 `json:"fert_cat_en" gorm:"column:fertilizer_cat_en"`
	FertCatTH      string		 `json:"fert_cat_th" gorm:"column:fertilizer_cat_th"`
}
// New instance fertilizer_cat
func (u *FertilizerCat) New() *FertilizerCat {
	return &FertilizerCat{
		DBCommon:      	u.DBCommon ,
		FertCatId:		u.FertCatId ,
		FertCatEN:		u.FertCatEN ,
		FertCatTH:		u.FertCatTH ,
	}
}

// Custom table name for GORM
func (FertilizerCat) TableName() string {
	return config.DB_FERTILIZER_CAT
}

