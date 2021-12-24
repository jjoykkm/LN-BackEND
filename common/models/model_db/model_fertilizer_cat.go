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
	DBCommonGet
	FertCatId      uuid.UUID	 `json:"fert_cat_id" gorm:"column:fertilizer_cat_id"`
	FertCatEN      string		 `json:"fert_cat_en" gorm:"column:fertilizer_cat_en"`
	FertCatTH      string		 `json:"fert_cat_th" gorm:"column:fertilizer_cat_th"`
}
// New instance fertilizer_cat
func (u *FertilizerCat) New() *FertilizerCat {
	return &FertilizerCat{
		DBCommonGet:      	u.DBCommonGet ,
		FertCatId:		u.FertCatId ,
		FertCatEN:		u.FertCatEN ,
		FertCatTH:		u.FertCatTH ,
	}
}

// Custom table name for GORM
func (FertilizerCat) TableName() string {
	return config.DB_FERTILIZER_CAT
}

//-------------------------------------------------------------------------------//
//									Upsert
//-------------------------------------------------------------------------------//
type FertilizerCatUS struct {
	DBCommonCreateUpdate
	FertCatId      string	 `json:"fert_cat_id" gorm:"column:fertilizer_cat_id;default:uuid_generate_v4()"`
	FertCatEN      string	 `json:"fert_cat_en" gorm:"column:fertilizer_cat_en"`
	FertCatTH      string	 `json:"fert_cat_th" gorm:"column:fertilizer_cat_th"`
	StatusId	   string	 `json:"status_id"`
}
func (FertilizerCatUS) TableName() string {
	return config.DB_FERTILIZER_CAT
}

