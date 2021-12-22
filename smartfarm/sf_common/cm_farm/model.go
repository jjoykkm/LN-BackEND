package cm_farm

import (
	"github.com/jjoykkm/ln-backend/common/config"
	"github.com/jjoykkm/ln-backend/common/models/model_db"
)

//-------------------------------------------------------------------------------//
//				 	    Plant And PlantType
//-------------------------------------------------------------------------------//
//Model
type FarmFarmArea struct {
	Farm    	model_db.Farm			`json:"farm" gorm:"embedded"`
	FarmArea    []model_db.FarmArea		`json:"farm_area" gorm:"foreignkey:FarmId; references:FarmId"`
}
func (FarmFarmArea) TableName() string {
	return config.DB_FARM
}

