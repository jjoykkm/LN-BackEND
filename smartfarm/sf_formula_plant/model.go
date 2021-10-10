package sf_formula_plant

import (
	"github.com/jjoykkm/ln-backend/common/config"
	"github.com/jjoykkm/ln-backend/common/models/model_db"
)


//-------------------------------------------------------------------------------//
//				 	    Plant And PlantType
//-------------------------------------------------------------------------------//
//Model
type PlantAndPlantType struct {
	PlantType   model_db.PlantType	`gorm:"foreignkey:PlantTypeId; references:PlantTypeId"`
	Plant     	model_db.Plant	 	`gorm:"embedded"`
}
func (PlantAndPlantType) TableName() string {
	return "plant"
}

//-------------------------------------------------------------------------------//
//				 	   	Formula Plant Item
//-------------------------------------------------------------------------------//
type FormulaPlant struct {
	FormulaPlant    model_db.FormulaPlant	`gorm:"embedded"`
	Owner     		model_db.UsersShort		`gorm:"foreignkey:Uid; references:Uid;"`
}
func (FormulaPlant) TableName() string {
	return "formula_plant"
}

type FormulaPlantItem struct {
	FormulaPlant    FormulaPlant		`gorm:"embedded"`
	Plant     		PlantAndPlantType	`gorm:"foreignkey:PlantId; references:PlantId;"`
	Province   		model_db.Province	`gorm:"foreignkey:ProvinceId; references:ProvinceId"`
	Country   		model_db.Country	`gorm:"foreignkey:CountryId; references:CountryId"`
	IsPlanted		bool
	IsFavorite		bool
}
func (FormulaPlantItem) TableName() string {
	return "formula_plant"
}

//-------------------------------------------------------------------------------//
//				 	    	Sensor
//-------------------------------------------------------------------------------//
type ForPlantSensor struct {
	SensorValueRec	model_db.TransSensorValueRec	`gorm:"embedded"`
	SensorType		model_db.SensorType			 	`gorm:"foreignkey:SensorTypeId; references:SensorTypeId;"`
}
func (ForPlantSensor) TableName() string {
	return config.DB_TRANS_SENSOR_VALUE_REC
}

//-------------------------------------------------------------------------------//
//				 	    	Fertilizer
//-------------------------------------------------------------------------------//
type Fertilizer struct {
	Fertilizer		model_db.Fertilizer		`gorm:"embedded"`
	FertilizerCat	model_db.FertilizerCat	`gorm:"foreignkey:FertCatId; references:FertCatId;"`
}
func (Fertilizer) TableName() string {
	return config.DB_FERTILIZER
}

type ForPlantFert struct {
	TransFertRatio	model_db.TransFertRatio		`gorm:"embedded"`
	Fertilizer		Fertilizer					`gorm:"foreignkey:FertId; references:FertId;"`
}
func (ForPlantFert) TableName() string {
	return config.DB_TRANS_FERTILIZER_RATIO
}

//-------------------------------------------------------------------------------//
//				 	   	Formula Plant Detail
//-------------------------------------------------------------------------------//
type ForPlantFormula struct {
	FormulaPlant 	 	 FormulaPlant		`gorm:"embedded"`
	ForPlantSensor	 	 []ForPlantSensor	`gorm:"foreignkey:FormulaPlantId; references:FormulaPlantId"`
	ForPlantFert		 []ForPlantFert		`gorm:"foreignkey:FormulaPlantId; references:FormulaPlantId"`
}
func (ForPlantFormula) TableName() string {
	return config.DB_FORMULA_PLANT
}