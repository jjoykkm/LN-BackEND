package sf_formula_plant

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/config"
	"github.com/jjoykkm/ln-backend/models/model_databases"
)


//-------------------------------------------------------------------------------//
//				 	   	Join Plant And PlantType (View/Add/Edit)
//-------------------------------------------------------------------------------//
//Model
type JoinPlantAndPlantType struct {
	PlantType   model_databases.PlantType	`mapstructure:"plant_type" json:"plant_type" gorm:"foreignkey:PlantTypeId; references:PlantTypeId"`
	Plant     	model_databases.Plant	 	`mapstructure:"plant" json:"plant" gorm:"embedded"`
}
func (JoinPlantAndPlantType) TableName() string {
	return "plant"
}

//-------------------------------------------------------------------------------//
//							Button Catagory
//-------------------------------------------------------------------------------//
//Model
type PlantCat struct {
	PlantTypeId     uuid.UUID	 `mapstructure:"plant_type_id" json:"plant_type_id"`
	PlantTypeName   string		 `mapstructure:"plant_type_name" json:"plant_type_name"`
	PlantId     	uuid.UUID	 `mapstructure:"plant_id" json:"plant_id"`
	PlantName       string		 `mapstructure:"plant_name" json:"plant_name"`
	PlantDesc       string		 `mapstructure:"plant_desc" json:"plant_desc"`
	TotalItem       int64	 	 `mapstructure:"total_item" json:"total_item"`
}

//-------------------------------------------------------------------------------//
//				 	   	Join Plant And PlantType (View/Add/Edit)
//-------------------------------------------------------------------------------//
//Model
type Plant struct {
	Plant  		model_databases.Plant 		`mapstructure:"plant" json:"plant" gorm:"embedded"`
	PlantType  	model_databases.PlantType 	`mapstructure:"plant_type" json:"plant_type" gorm:"foreignkey:PlantTypeId; references:PlantTypeId"`
}
func (Plant) TableName() string {
	return "plant"
}

type FormulaPlantItem struct {
	FormulaPlant    model_databases.FormulaPlant	`mapstructure:"formula_plant" json:"formula_plant" gorm:"embedded"`
	Plant     		Plant			 				`mapstructure:"plant" json:"plant" gorm:"foreignkey:PlantId; references:PlantId;"`
	Province   		model_databases.Province		`mapstructure:"province" json:"province" gorm:"foreignkey:ProvinceId; references:ProvinceId"` // association_foreignkey:ProvinceId
	Country   		model_databases.Country			`mapstructure:"country" json:"country" gorm:"foreignkey:CountryId; references:CountryId"`
	IsPlanted		bool							`mapstructure:"is_planted" json:"is_planted"`
	IsFavorite		bool							`mapstructure:"is_favorite" json:"is_favorite"`
}
func (FormulaPlantItem) TableName() string {
	return "formula_plant"
}

//-------------------------------------------------------------------------------//
//				  Button Favorite and My formula (Card Item)
//-------------------------------------------------------------------------------//
//Model
type ForPlantItem struct {
	PlantType		 model_databases.PlantType		`json:"plant_type"`
	FormulaPlant	 model_databases.FormulaPlant	`json:"formula_plant"`
	RateScore 		 float32	 					`json:"rate_score"`
	RatePeople 		 int		 					`json:"rate_people"`
	IsPublic		 bool	 	 					`json:"is_public"`
	IsPlanted		 bool	 	 					`json:"is_planted"`
	IsFavorite		 bool	 	 					`json:"is_favorite"`
}

//-------------------------------------------------------------------------------//
//				 	    	Sensor
//-------------------------------------------------------------------------------//
type ForPlantSensor struct {
	SensorValueRec	model_databases.TransSensorValueRec	 `json:"sensor_value_rec" gorm:"embedded"`
	SensorType		model_databases.SensorType			 `json:"sensor_type" gorm:"foreignkey:SensorTypeId; references:SensorTypeId;"`
}
func (ForPlantSensor) TableName() string {
	return config.DB_TRANS_SENSOR_VALUE_REC
}

//-------------------------------------------------------------------------------//
//				 	    	Fertilizer
//-------------------------------------------------------------------------------//
type Fertilizer struct {
	Fertilizer		model_databases.Fertilizer		`json:"fertilizer" gorm:"embedded"`
	FertilizerCat	model_databases.FertilizerCat	`json:"fertilizer_cat" gorm:"foreignkey:FertilizerCatId; references:FertilizerCatId;"`
}
func (Fertilizer) TableName() string {
	return config.DB_FERTILIZER
}

type ForPlantFert struct {
	TransFertRatio	model_databases.TransFertRatio	`json:"trans_fert_ratio" gorm:"embedded"`
	Fertilizer		Fertilizer						`json:"fertilizer" gorm:"foreignkey:FertilizerId; references:FertilizerId;"`
}
func (ForPlantFert) TableName() string {
	return config.DB_TRANS_FERTILIZER_RATIO
}

//-------------------------------------------------------------------------------//
//				 	   	Formula Plant Detail
//-------------------------------------------------------------------------------//
type ForPlantFormula struct {
	FormulaPlant 	 model_databases.FormulaPlant	`json:"formula_plant" gorm:"embedded"`
	ForPlantSensor	 	 []ForPlantSensor			`json:"sensor_list" gorm:"foreignkey:FormulaPlantId; references:FormulaPlantId"`
	ForPlantFert		 []ForPlantFert				`json:"fert_list" gorm:"foreignkey:FormulaPlantId; references:FormulaPlantId"`
}
func (ForPlantFormula) TableName() string {
	return config.DB_FORMULA_PLANT
}