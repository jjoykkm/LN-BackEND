package sf_formula_plant

import (
	"github.com/jjoykkm/ln-backend/common/config"
	"github.com/jjoykkm/ln-backend/common/models/model_db"
)

////-------------------------------------------------------------------------------//
////				 	    Plant And PlantType
////-------------------------------------------------------------------------------//
////Model
//type PlantAndPlantType struct {
//	PlantType   model_db.PlantType	`json:"plant_type" gorm:"foreignkey:PlantTypeId; references:PlantTypeId"`
//	Plant     	model_db.Plant	 	`json:"plant" gorm:"embedded"`
//}
//func (PlantAndPlantType) TableName() string {
//	return config.DB_PLANT
//}

//-------------------------------------------------------------------------------//
//				 	   	Formula Plant Item
//-------------------------------------------------------------------------------//
type FormulaPlant struct {
	FormulaPlant    model_db.FormulaPlant	`json:"formula_plant" gorm:"embedded"`
	Owner     		model_db.UsersShort		`json:"owner" gorm:"foreignkey:Uid; references:Uid;"`
}
func (FormulaPlant) TableName() string {
	return config.DB_FORMULA_PLANT
}

type FormulaPlantItem struct {
	FormulaPlant    FormulaPlant		`json:"formula_plant" gorm:"embedded"`
	Plant     		PlantAndPlantType	`json:"plant" gorm:"foreignkey:PlantId; references:PlantId;"`
	Province   		model_db.Province	`json:"province" gorm:"foreignkey:ProvinceId; references:ProvinceId"`
	Country   		model_db.Country	`json:"country" gorm:"foreignkey:CountryId; references:CountryId"`
	IsPlanted		bool				`json:"is_planted"`
	IsFavorite		bool				`json:"is_favorite"`
}
func (FormulaPlantItem) TableName() string {
	return config.DB_FORMULA_PLANT
}

//-------------------------------------------------------------------------------//
//				 	    	Sensor
//-------------------------------------------------------------------------------//
type ForPlantSensor struct {
	SensorValueRec	model_db.TransSensorValueRec	`json:"sensor_value_rec" gorm:"embedded"`
	SensorType		model_db.SensorType			 	`json:"sensor_type" gorm:"foreignkey:SensorTypeId; references:SensorTypeId;"`
}
func (ForPlantSensor) TableName() string {
	return config.DB_TRANS_SENSOR_VALUE_REC
}

//-------------------------------------------------------------------------------//
//				 	    	Fertilizer
//-------------------------------------------------------------------------------//
type Fertilizer struct {
	Fertilizer		model_db.Fertilizer		`json:"fertilizer" gorm:"embedded"`
	FertilizerCat	model_db.FertilizerCat	`json:"fertilizer_cat" gorm:"foreignkey:FertCatId; references:FertCatId;"`
}
func (Fertilizer) TableName() string {
	return config.DB_FERTILIZER
}

type ForPlantFert struct {
	TransFertRatio	model_db.TransFertRatio		`json:"trans_fert_ratio" gorm:"embedded"`
	Fertilizer		Fertilizer					`json:"fertilizer" gorm:"foreignkey:FertId; references:FertId;"`
}
func (ForPlantFert) TableName() string {
	return config.DB_TRANS_FERTILIZER_RATIO
}

//-------------------------------------------------------------------------------//
//				 	   	Formula Plant Detail
//-------------------------------------------------------------------------------//
type ForPlantFormula struct {
	FormulaPlant 	 	 FormulaPlant		`json:"formula_plant" gorm:"embedded"`
	ForPlantSensor	 	 []ForPlantSensor	`json:"sensor_list" gorm:"foreignkey:FormulaPlantId; references:FormulaPlantId"`
	ForPlantFert		 []ForPlantFert		`json:"fert_list" gorm:"foreignkey:FormulaPlantId; references:FormulaPlantId"`
	IsPlanted			 bool				`json:"is_planted"`
	IsFavorite			 bool				`json:"is_favorite"`
}
func (ForPlantFormula) TableName() string {
	return config.DB_FORMULA_PLANT
}


//-------------------------------------------------------------------------------//
//				 	    Plant And PlantType
//-------------------------------------------------------------------------------//
//Model
type PlantAndPlantType struct {
	PlantType   model_db.PlantType	 `json:"plant_type" gorm:"embedded"`
	Plant     	[]model_db.Plant	 `json:"plant" gorm:"foreignkey:PlantTypeId; references:PlantTypeId"`
}
func (PlantAndPlantType) TableName() string {
	return config.DB_PLANT_TYPE
}

//-------------------------------------------------------------------------------//
//									Upsert
//-------------------------------------------------------------------------------//
type ForPlantUS struct {
	FormulaPlant	model_db.FormulaPlantUS 		`json:"formula_plant" gorm:"embedded"`
	SensorValue		[]model_db.TransSenValueRecUS	`json:"sensor_value"`
	FertRatio		[]model_db.TransFertRatioUS		`json:"fertilizer_ratio"`
}
func (ForPlantUS) TableName() string {
	return config.DB_FORMULA_PLANT
}
