package model_services

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
)
//-------------------------------------------------------------------------------//
//							Set Button Catagory
//-------------------------------------------------------------------------------//
//Model
type ForPlantCatList struct {
	PlantTypeId     uuid.UUID	 `mapstructure:"plant_type_id" json:"plant_type_id,omitempty"`
	PlantTypeName   string		 `mapstructure:"plant_type_name" json:"plant_type_name,omitempty"`
}

// New instance
func (u *ForPlantCatList) New() *ForPlantCatList {
	return &ForPlantCatList{
		PlantTypeId:  	u.PlantTypeId ,
		PlantTypeName:  u.PlantTypeName ,
	}
}

//-------------------------------------------------------------------------------//
//							Button Catagory
//-------------------------------------------------------------------------------//
//Model
type ForPlantCat struct {
	PlantTypeId     uuid.UUID	 `mapstructure:"plant_type_id" json:"plant_type_id,omitempty"`
	PlantTypeName   string		 `mapstructure:"plant_type_name" json:"plant_type_name,omitempty"`
	PlantId     	uuid.UUID	 `mapstructure:"plant_id" json:"plant_id,omitempty"`
	PlantName       string		 `mapstructure:"plant_name" json:"plant_name,omitempty"`
	PlantDesc       string		 `mapstructure:"plant_desc" json:"plant_desc,omitempty"`
	TotalItem       string		 `mapstructure:"total_item" json:"total_item,omitempty"`
}

// New instance
func (u *ForPlantCat) New() *ForPlantCat {
	return &ForPlantCat{
		PlantTypeId:  	u.PlantTypeId ,
		PlantTypeName:  u.PlantTypeName ,
		PlantId:      	u.PlantId ,
		PlantName:  	u.PlantName ,
		PlantDesc:  	u.PlantDesc ,
		TotalItem:    	u.TotalItem ,
	}
}


//-------------------------------------------------------------------------------//
//				  Button Favorite and My formula (Card Item)
//-------------------------------------------------------------------------------//
//Model
type ForPlantItem struct {
	PlantTypeId      uuid.UUID	 `mapstructure:"plant_type_id" json:"plant_type_id,omitempty"`
	PlantTypeName    string		 `mapstructure:"plant_type_name" json:"plant_type_name,omitempty"`
	FormulaPlantId 	 uuid.UUID	 `mapstructure:"formula_plant_id" json:"formula_plant_id,omitempty"`
	FormulaName		 string		 `mapstructure:"formula_name" json:"formula_name,omitempty"`
	FormulaDesc		 string		 `mapstructure:"formula_desc" json:"formula_desc,omitempty"`
	PeopleUsed 		 int		 `mapstructure:"people_used" json:"people_used,omitempty"`
	RateScore 		 float32	 `mapstructure:"rate_score" json:"rate_score,omitempty"`
	RatePeople 		 int		 `mapstructure:"rate_people" json:"rate_people,omitempty"`
	CountryId		 uuid.UUID	 `mapstructure:"country_id" json:"country_id,omitempty"`
	CountryName      string		 `mapstructure:"country_name" json:"country_name,omitempty"`
	ProvinceId		 uuid.UUID	 `mapstructure:"province_id" json:"province_id,omitempty"`
	ProvinceName     string		 `mapstructure:"province_name" json:"province_name,omitempty"`
	IsPublic		 bool	 	 `mapstructure:"is_public" json:"is_public,omitempty"`
	IsPlanted		 bool	 	 `mapstructure:"is_planted" json:"is_planted,omitempty"`
	IsFavorite		 bool	 	 `mapstructure:"is_favorite" json:"is_favorite,omitempty"`
	Uid				 uuid.UUID 	 `mapstructure:"uid" json:"uid,omitempty"`
}

// New instance
func (u *ForPlantItem) New() *ForPlantItem {
	return &ForPlantItem{
		PlantTypeId:		u.PlantTypeId ,
		PlantTypeName:  	u.PlantTypeName ,
		FormulaPlantId:		u.FormulaPlantId ,
		FormulaName:		u.FormulaName ,
		FormulaDesc:		u.FormulaDesc ,
		PeopleUsed:			u.PeopleUsed ,
		RateScore:			u.RateScore ,
		RatePeople:			u.RatePeople ,
		CountryId:			u.CountryId ,
		CountryName:		u.CountryName ,
		ProvinceId:			u.ProvinceId ,
		ProvinceName:		u.ProvinceName ,
		IsPublic:			u.IsPublic ,
		IsPlanted:			u.IsPlanted ,
		IsFavorite:			u.IsFavorite ,
		Uid:				u.Uid ,
	}
}

//-------------------------------------------------------------------------------//
//				 			 List item in catagory
//-------------------------------------------------------------------------------//
//Model
type ForPlantCatItem struct {
	PlantId     	uuid.UUID	 		`mapstructure:"plant_id" json:"plant_id,omitempty"`
	PlantName       string		 		`mapstructure:"plant_name" json:"plant_name,omitempty"`
	PlantDesc       string		 		`mapstructure:"plant_desc" json:"plant_desc,omitempty"`
	TotalItem       string		 		`mapstructure:"total_item" json:"total_item,omitempty"`
	PlantCatItem	[]ForPlantItem		`mapstructure:"plant_cat_item" json:"plant_cat_item,omitempty"`
}

// New instance
func (u *ForPlantCatItem) New() *ForPlantCatItem {
	return &ForPlantCatItem{
		PlantId:      	u.PlantId ,
		PlantName:  	u.PlantName ,
		PlantDesc:  	u.PlantDesc ,
		TotalItem:    	u.TotalItem ,
		PlantCatItem:   u.PlantCatItem ,
	}
}


//-------------------------------------------------------------------------------//
//				 	    	Fertilizer
//-------------------------------------------------------------------------------//
//Model
type ForPlantFert struct {
	FertilizerId     uuid.UUID	 `mapstructure:"fertilizer_id" json:"fertilizer_id,omitempty"`
	FertilizerName   string		 `mapstructure:"fertilizer_name" json:"fertilizer_name,omitempty"`
	Nitrogen       	 float64	 `mapstructure:"nitrogen" json:"nitrogen,omitempty"`
	Phosphorus    	 float64	 `mapstructure:"phosphorus" json:"phosphorus,omitempty"`
	Potassium      	 float64	 `mapstructure:"potassium" json:"potassium,omitempty"`
	Ratio	      	 float64	 `mapstructure:"ratio" json:"ratio,omitempty"`
	FertCatId      	 uuid.UUID	 `mapstructure:"fertilizer_cat_id" json:"fertilizer_cat_id,omitempty"`
	FertCatName    	 string		 `mapstructure:"fertilizer_cat_name" json:"fertilizer_cat_name,omitempty"`
}
// New instance
func (u *ForPlantFert) New() *ForPlantFert {
	return &ForPlantFert{
		FertilizerId:		u.FertilizerId ,
		FertilizerName:		u.FertilizerName ,
		Nitrogen:			u.Nitrogen ,
		Phosphorus:			u.Phosphorus ,
		Potassium:			u.Potassium ,
		Ratio:				u.Ratio ,
		FertCatId:			u.FertCatId ,
		FertCatName:		u.FertCatName ,
	}
}

//-------------------------------------------------------------------------------//
//				 	    	Sensor
//-------------------------------------------------------------------------------//
//Model
type ForPlantSensor struct {
	SensorTypeId      	uuid.UUID	 `mapstructure:"sensor_type_id" json:"sensor_type_id,omitempty"`
	SensorTypeName      string	 	 `mapstructure:"sensor_type_name" json:"sensor_type_name,omitempty"`
	ValueRec      		float64		 `mapstructure:"value_rec" json:"value_rec,omitempty"`
}
// New instance
func (u *ForPlantSensor) New() *ForPlantSensor {
	return &ForPlantSensor{
		SensorTypeId:		u.SensorTypeId ,
		SensorTypeName:		u.SensorTypeName ,
		ValueRec:			u.ValueRec ,
	}
}

//-------------------------------------------------------------------------------//
//				 	   	Formula Plant (View/Add/Edit)
//-------------------------------------------------------------------------------//

//Model
type ForPlantFormula struct {
	FormulaPlantId 	 uuid.UUID	 			`mapstructure:"formula_plant_id" json:"formula_plant_id,omitempty"`
	FormulaName		 string		 			`mapstructure:"formula_name" json:"formula_name,omitempty"`
	FormulaDesc		 string		 			`mapstructure:"formula_desc" json:"formula_desc,omitempty"`
	SensorList	 	 []ForPlantSensor		`mapstructure:"sensor_list" json:"sensor_list,omitempty"`
	FertList		 []ForPlantFert			`mapstructure:"fert_list" json:"fert_list,omitempty"`
}

// New instance
func (u *ForPlantFormula) New() *ForPlantFormula {
	return &ForPlantFormula{
		FormulaPlantId:		u.FormulaPlantId ,
		FormulaName:		u.FormulaName ,
		FormulaDesc:		u.FormulaDesc ,
		SensorList:   		u.SensorList ,
		FertList:    		u.FertList ,
	}
}

