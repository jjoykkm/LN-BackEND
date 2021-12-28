package sf_formula_plant

import (
	"github.com/jjoykkm/ln-backend/common/config"
	"github.com/jjoykkm/ln-backend/common/models/model_db"
)

//-------------------------------------------------------------------------------//
//				 	   	Formula Plant Item
//-------------------------------------------------------------------------------//
type ForPlantDetail struct {
	FormulaPlant    model_db.FormulaPlant	`json:"formula_plant" gorm:"embedded"`
	Owner     		model_db.UsersShort		`json:"owner" gorm:"foreignkey:Uid; references:Uid;"`
	Province   		model_db.Province		`json:"province" gorm:"foreignkey:ProvinceId; references:ProvinceId"`
	Country   		model_db.Country		`json:"country" gorm:"foreignkey:CountryId; references:CountryId"`
	IsPlanted		bool					`json:"is_planted"`
	IsFavorite		bool					`json:"is_favorite"`
}
func (ForPlantDetail) TableName() string {
	return config.DB_FORMULA_PLANT
}

type PlantDetail struct {
	Plant     	model_db.Plant	 	`json:"plant" gorm:"embedded"`
	PlantType   model_db.PlantType	`json:"plant_type" gorm:"foreignkey:PlantTypeId; references:PlantTypeId"`
}
func (PlantDetail) TableName() string {
	return config.DB_PLANT
}

type FormulaPlant struct {
	FormulaPlant    model_db.FormulaPlant	`json:"formula_plant" gorm:"embedded"`
	Owner     		model_db.UsersShort		`json:"owner" gorm:"foreignkey:Uid; references:Uid;"`
	Province   		model_db.Province		`json:"province" gorm:"foreignkey:ProvinceId; references:ProvinceId"`
	Country   		model_db.Country		`json:"country" gorm:"foreignkey:CountryId; references:CountryId"`
}
func (FormulaPlant) TableName() string {
	return config.DB_FORMULA_PLANT
}

type PlantTypeAndPlantList struct {
	PlantType   model_db.PlantType	 `json:"plant_type" gorm:"embedded"`
	Plant     	[]model_db.Plant	 `json:"plant" gorm:"foreignkey:PlantTypeId; references:PlantTypeId"`
}
func (PlantTypeAndPlantList) TableName() string {
	return config.DB_PLANT_TYPE
}

type ForPlantlist struct {
	Plant     		PlantDetail			`json:"plant" gorm:"embedded"`
	ForPlantDetail  []ForPlantDetail	`json:"formula_plant_detail" gorm:"foreignkey:PlantId; references:PlantId;"`
}
func (ForPlantlist) TableName() string {
	return config.DB_PLANT
}

type FormulaPlantItem struct {
	Plant     		PlantDetail			`json:"plant" gorm:"foreignkey:PlantId; references:PlantId;"`
	FormulaPlant    FormulaPlant		`json:"formula_plant" gorm:"embedded"`
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
	FormulaPlant 	 	 FormulaPlantItem	`json:"formula_plant" gorm:"embedded"`
	ForPlantSensor	 	 []ForPlantSensor	`json:"sensor_list" gorm:"foreignkey:FormulaPlantId; references:FormulaPlantId"`
	ForPlantFert		 []ForPlantFert		`json:"fert_list" gorm:"foreignkey:FormulaPlantId; references:FormulaPlantId"`
}
func (ForPlantFormula) TableName() string {
	return config.DB_FORMULA_PLANT
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
