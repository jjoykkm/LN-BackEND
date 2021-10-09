package sf_formula_plant

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/config"
	"github.com/jjoykkm/ln-backend/models/model_databases"
)

//-------------------------------------------------------------------------------//
//							Set Button Catagory
//-------------------------------------------------------------------------------//
//Model
type PlantTypeCat struct {
	PlantTypeId     uuid.UUID	 `mapstructure:"plant_type_id" json:"plant_type_id"`
	PlantTypeName   string		 `mapstructure:"plant_type_name" json:"plant_type_name"`
}

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
	//RateScore		float32							`mapstructure:"rate_score" json:"rate_score"`
	//RatePeople		int64							`mapstructure:"rate_people" json:"rate_people"`
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
	PlantType		 model_databases.PlantType
	FormulaPlant	 model_databases.FormulaPlant
	RateScore 		 float32	 `mapstructure:"rate_score" json:"rate_score"`
	RatePeople 		 int		 `mapstructure:"rate_people" json:"rate_people"`
	IsPublic		 bool	 	 `mapstructure:"is_public" json:"is_public"`
	IsPlanted		 bool	 	 `mapstructure:"is_planted" json:"is_planted"`
	IsFavorite		 bool	 	 `mapstructure:"is_favorite" json:"is_favorite"`
}

//-------------------------------------------------------------------------------//
//				 	    	Sensor
//-------------------------------------------------------------------------------//
//Model
type ForPlantSensor struct {
	SensorValueRec	model_databases.TransSensorValueRec	 `mapstructure:"sensor_value_rec" json:"sensor_value_rec" gorm:"embedded"`
	SensorType		model_databases.SensorType			 `mapstructure:"sensor_type" json:"sensor_type" gorm:"foreignkey:SensorTypeId; references:SensorTypeId;"`

//	SensorTypeId      	uuid.UUID	 `mapstructure:"sensor_type_id" json:"sensor_type_id"`
//	SensorTypeName      string	 	 `mapstructure:"sensor_type_name" json:"sensor_type_name"`
//	ValueRec      		float64		 `mapstructure:"value_rec" json:"value_rec"`
}

func (ForPlantSensor) TableName() string {
	return config.DB_TRANS_SENSOR_VALUE_REC
}
//-------------------------------------------------------------------------------//
//				 	    	Fertilizer
//-------------------------------------------------------------------------------//
//Model

type Fertilizer struct {
	Fertilizer		model_databases.Fertilizer		`mapstructure:"fertilizer" json:"fertilizer" gorm:"embedded"`
	FertilizerCat	model_databases.FertilizerCat	`mapstructure:"fertilizer_cat" json:"fertilizer_cat" gorm:"foreignkey:FertilizerCatId; references:FertilizerCatId;"`
}

func (Fertilizer) TableName() string {
	return config.DB_FERTILIZER
}

type ForPlantFert struct {
	TransFertRatio	model_databases.TransFertRatio	`mapstructure:"trans_fert_ratio" json:"trans_fert_ratio" gorm:"embedded"`
	Fertilizer		Fertilizer						`mapstructure:"fertilizer" json:"fertilizer" gorm:"foreignkey:FertilizerId; references:FertilizerId;"`
	//FertilizerCat	model_databases.FertilizerCat	`mapstructure:"fertilizer_cat" json:"fertilizer_cat" gorm:"foreignkey:FertilizerCatId; references:FertilizerCatId;"`

//	FertilizerId     uuid.UUID	 `mapstructure:"fertilizer_id" json:"fertilizer_id"`
//	FertilizerName   string		 `mapstructure:"fertilizer_name" json:"fertilizer_name"`
//	Nitrogen       	 float64	 `mapstructure:"nitrogen" json:"nitrogen"`
//	Phosphorus    	 float64	 `mapstructure:"phosphorus" json:"phosphorus"`
//	Potassium      	 float64	 `mapstructure:"potassium" json:"potassium"`
//	Ratio	      	 float64	 `mapstructure:"ratio" json:"ratio"`
//	FertCatId      	 uuid.UUID	 `mapstructure:"fertilizer_cat_id" json:"fertilizer_cat_id"`
//	FertCatName    	 string		 `mapstructure:"fertilizer_cat_name" json:"fertilizer_cat_name"`
}

func (ForPlantFert) TableName() string {
	return config.DB_TRANS_FERTILIZER_RATIO
}
//-------------------------------------------------------------------------------//
//				 	   	Formula Plant Detail
//-------------------------------------------------------------------------------//

//Model
type ForPlantFormula struct {
	FormulaPlant 	 model_databases.FormulaPlant	`mapstructure:"formula_plant" json:"formula_plant" gorm:"embedded"`
	ForPlantSensor	 	 []ForPlantSensor				`mapstructure:"sensor_list" json:"sensor_list" gorm:"foreignkey:FormulaPlantId; references:FormulaPlantId"`
	ForPlantFert		 []ForPlantFert					`mapstructure:"fert_list" json:"fert_list" gorm:"foreignkey:FormulaPlantId; references:FormulaPlantId"`
	//SensorList	 	 []ForPlantSensor				`mapstructure:"sensor_list" json:"sensor_list" gorm:"foreignkey:FormulaPlantId; references:FormulaPlantId"`
	//FertList		 []ForPlantFert					`mapstructure:"fert_list" json:"fert_list" gorm:"foreignkey:FormulaPlantId; references:FormulaPlantId"`
	//SensorList	 	 []ForPlantSensor		`mapstructure:"sensor_list" json:"sensor_list"`
	//FertList		 []ForPlantFert			`mapstructure:"fert_list" json:"fert_list"`
}
func (ForPlantFormula) TableName() string {
	return config.DB_FORMULA_PLANT
}