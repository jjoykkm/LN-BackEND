package cm_plant

import (
	"github.com/jjoykkm/ln-backend/common/config"
	"github.com/jjoykkm/ln-backend/common/models/model_db"
)

//-------------------------------------------------------------------------------//
//				 	    Plant And PlantType
//-------------------------------------------------------------------------------//
//Model
type FertilizerAndCat struct {
	FertilizerCat	model_db.FertilizerCat	`json:"fertilizer_cat" gorm:"embedded"`
	Fertilizer    	[]model_db.Fertilizer	`json:"fertilizer" gorm:"foreignkey:FertCatId; references:FertCatId"`
}
func (FertilizerAndCat) TableName() string {
	return config.DB_FERTILIZER_CAT
}

