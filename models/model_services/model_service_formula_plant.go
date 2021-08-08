package model_services

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"time"
)
//-------------------------------------------------------------------------------//
//							Set Button Catagory
//-------------------------------------------------------------------------------//
//Model
type ForPlantCatList struct {
	PlantTypeId     uuid.UUID	 `mapstructure:"plant_type_id" json:"plant_type_id"`
	PlantTypeName   string		 `mapstructure:"plant_type_name" json:"plant_type_name"`
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
	PlantTypeId     uuid.UUID	 `mapstructure:"plant_type_id" json:"plant_type_id"`
	PlantTypeName   string		 `mapstructure:"plant_type_name" json:"plant_type_name"`
	PlantId     	uuid.UUID	 `mapstructure:"plant_id" json:"plant_id"`
	PlantName       string		 `mapstructure:"plant_name" json:"plant_name"`
	PlantDesc       string		 `mapstructure:"plant_desc" json:"plant_desc"`
	TotalItem       int		 	 `mapstructure:"total_item" json:"total_item"`
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
	PlantTypeId      uuid.UUID	 `mapstructure:"plant_type_id" json:"plant_type_id"`
	PlantTypeName    string		 `mapstructure:"plant_type_name" json:"plant_type_name"`
	FormulaPlantId 	 uuid.UUID	 `mapstructure:"formula_plant_id" json:"formula_plant_id"`
	FormulaName		 string		 `mapstructure:"formula_name" json:"formula_name"`
	FormulaDesc		 string		 `mapstructure:"formula_desc" json:"formula_desc"`
	PeopleUsed 		 int		 `mapstructure:"people_used" json:"people_used"`
	RateScore 		 float32	 `mapstructure:"rate_score" json:"rate_score"`
	RatePeople 		 int		 `mapstructure:"rate_people" json:"rate_people"`
	CountryId		 uuid.UUID	 `mapstructure:"country_id" json:"country_id"`
	CountryName      string		 `mapstructure:"country_name" json:"country_name"`
	ProvinceId		 uuid.UUID	 `mapstructure:"province_id" json:"province_id"`
	ProvinceName     string		 `mapstructure:"province_name" json:"province_name"`
	IsPublic		 bool	 	 `mapstructure:"is_public" json:"is_public"`
	IsPlanted		 bool	 	 `mapstructure:"is_planted" json:"is_planted"`
	IsFavorite		 bool	 	 `mapstructure:"is_favorite" json:"is_favorite"`
	Uid				 uuid.UUID 	 `mapstructure:"uid" json:"uid"`
	Username     	 string	 	 `mapstructure:"username" json:"username"`
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
		Username:			u.Username ,
	}
}

//-------------------------------------------------------------------------------//
//				 			 List item in catagory
//-------------------------------------------------------------------------------//
//Model
type ForPlantCatItem struct {
	PlantId     	uuid.UUID	 		`mapstructure:"plant_id" json:"plant_id"`
	PlantName       string		 		`mapstructure:"plant_name" json:"plant_name"`
	PlantDesc       string		 		`mapstructure:"plant_desc" json:"plant_desc"`
	TotalItem       string		 		`mapstructure:"total_item" json:"total_item"`
	PlantCatItem	[]ForPlantItem		`mapstructure:"plant_cat_item" json:"plant_cat_item"`
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
	FertilizerId     uuid.UUID	 `mapstructure:"fertilizer_id" json:"fertilizer_id"`
	FertilizerName   string		 `mapstructure:"fertilizer_name" json:"fertilizer_name"`
	Nitrogen       	 float64	 `mapstructure:"nitrogen" json:"nitrogen"`
	Phosphorus    	 float64	 `mapstructure:"phosphorus" json:"phosphorus"`
	Potassium      	 float64	 `mapstructure:"potassium" json:"potassium"`
	Ratio	      	 float64	 `mapstructure:"ratio" json:"ratio"`
	FertCatId      	 uuid.UUID	 `mapstructure:"fertilizer_cat_id" json:"fertilizer_cat_id"`
	FertCatName    	 string		 `mapstructure:"fertilizer_cat_name" json:"fertilizer_cat_name"`
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
	SensorTypeId      	uuid.UUID	 `mapstructure:"sensor_type_id" json:"sensor_type_id"`
	SensorTypeName      string	 	 `mapstructure:"sensor_type_name" json:"sensor_type_name"`
	ValueRec      		float64		 `mapstructure:"value_rec" json:"value_rec"`
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
	FormulaPlantId 	 uuid.UUID	 			`mapstructure:"formula_plant_id" json:"formula_plant_id"`
	FormulaName		 string		 			`mapstructure:"formula_name" json:"formula_name"`
	FormulaDesc		 string		 			`mapstructure:"formula_desc" json:"formula_desc"`
	Uid			 	 uuid.UUID	 			`mapstructure:"uid" json:"uid"`
	Username		 string		 			`mapstructure:"username" json:"username"`
	SensorList	 	 []ForPlantSensor		`mapstructure:"sensor_list" json:"sensor_list"`
	FertList		 []ForPlantFert			`mapstructure:"fert_list" json:"fert_list"`
}

// New instance
func (u *ForPlantFormula) New() *ForPlantFormula {
	return &ForPlantFormula{
		FormulaPlantId:		u.FormulaPlantId ,
		FormulaName:		u.FormulaName ,
		FormulaDesc:		u.FormulaDesc ,
		Uid:				u.Uid ,
		Username:			u.Username ,
		SensorList:   		u.SensorList ,
		FertList:    		u.FertList ,
	}
}

//-------------------------------------------------------------------------------//
//				 	   	Join Plant And PlantType (View/Add/Edit)
//-------------------------------------------------------------------------------//

//Model
type JoinPlantAndPlantType struct {
	PlantId          uuid.UUID	 `mapstructure:"plant_id" json:"plant_id"`
	PlantNameEN      string		 `mapstructure:"plant_name_en" json:"plant_name_en"`
	PlantNameTH      string		 `mapstructure:"plant_name_th" json:"plant_name_th"`
	PlantDescEN      string		 `mapstructure:"plant_desc_en" json:"plant_desc_en"`
	PlantDescTH      string		 `mapstructure:"plant_desc_th" json:"plant_desc_th"`
	CreateDate		 time.Time	 `mapstructure:"create_date" json:"create_date"`
	ChangeDate	     time.Time	 `mapstructure:"change_date" json:"change_date"`
	StatusId		 uuid.UUID	 `mapstructure:"status_id" json:"status_id"`
	TotalItem      	 int		 `mapstructure:"total_item" json:"total_item"`
	PlantTypeId      uuid.UUID	 `mapstructure:"plant_type_id" json:"plant_type_id"`
	PlantTypeEN      string		 `mapstructure:"plant_type_en" json:"plant_type_en"`
	PlantTypeTH      string		 `mapstructure:"plant_type_th" json:"plant_type_th"`
}

// New instance
func (u *JoinPlantAndPlantType) New() *JoinPlantAndPlantType {
	return &JoinPlantAndPlantType{
		PlantId:    	u.PlantId ,
		PlantNameEN:    u.PlantNameEN ,
		PlantNameTH:    u.PlantNameTH ,
		PlantDescEN:    u.PlantDescEN ,
		PlantDescTH:    u.PlantDescTH ,
		CreateDate:    	u.CreateDate ,
		ChangeDate:    	u.ChangeDate ,
		StatusId:    	u.StatusId ,
		TotalItem:    	u.TotalItem   ,
		PlantTypeId:    u.PlantTypeId ,
		PlantTypeEN:    u.PlantTypeEN ,
		PlantTypeTH:    u.PlantTypeTH ,
	}
}

//-------------------------------------------------------------------------------//
//				 	   	Join FormulaPlant And Plant (View/Add/Edit)
//-------------------------------------------------------------------------------//

//Model
type JoinFormulaPlantAndPlant struct {
	FormulaPlantId 	 uuid.UUID	 `mapstructure:"formula_plant_id" json:"formula_plant_id"`
	FormulaName		 string		 `mapstructure:"formula_name" json:"formula_name"`
	FormulaDesc		 string		 `mapstructure:"formula_desc" json:"formula_desc"`
	PeopleUsed 		 int		 `mapstructure:"people_used" json:"people_used"`
	Recommend1 		 int		 `mapstructure:"recommend1" json:"recommend1"`
	Recommend2		 int		 `mapstructure:"recommend2" json:"recommend2"`
	Recommend3		 int		 `mapstructure:"recommend3" json:"recommend3"`
	Recommend4		 int		 `mapstructure:"recommend4" json:"recommend4"`
	Recommend5		 int		 `mapstructure:"recommend5" json:"recommend5"`
	CreateDate		 time.Time	 `mapstructure:"create_date" json:"create_date"`
	ChangeDate		 time.Time	 `mapstructure:"change_date" json:"change_date"`
	StatusId		 uuid.UUID	 `mapstructure:"status_id" json:"status_id"`
	ProvinceId		 uuid.UUID	 `mapstructure:"province_id" json:"province_id"`
	CountryId		 uuid.UUID	 `mapstructure:"country_id" json:"country_id"`
	IsPublic		 bool	 	 `mapstructure:"is_public" json:"is_public"`
	Uid				 uuid.UUID	 `mapstructure:"uid" json:"uid"`
	PlantId          uuid.UUID	 `mapstructure:"plant_id" json:"plant_id"`
	PlantNameEN      string		 `mapstructure:"plant_name_en" json:"plant_name_en"`
	PlantNameTH      string		 `mapstructure:"plant_name_th" json:"plant_name_th"`
	PlantDescEN      string		 `mapstructure:"plant_desc_en" json:"plant_desc_en"`
	PlantDescTH      string		 `mapstructure:"plant_desc_th" json:"plant_desc_th"`
	PlantTypeId      uuid.UUID	 `mapstructure:"plant_type_id" json:"plant_type_id"`
	TotalItem      	 int		 `mapstructure:"total_item" json:"total_item"`
}

// New instance
func (u *JoinFormulaPlantAndPlant) New() *JoinFormulaPlantAndPlant {
	return &JoinFormulaPlantAndPlant{
		FormulaPlantId:		u.FormulaPlantId ,
		FormulaName:		u.FormulaName ,
		FormulaDesc:		u.FormulaDesc ,
		PeopleUsed:			u.PeopleUsed ,
		Recommend1:			u.Recommend1 ,
		Recommend2:			u.Recommend2 ,
		Recommend3:			u.Recommend3 ,
		Recommend4:			u.Recommend4 ,
		Recommend5:			u.Recommend5 ,
		CreateDate:			u.CreateDate ,
		ChangeDate:			u.ChangeDate ,
		StatusId:			u.StatusId ,
		ProvinceId:			u.ProvinceId ,
		CountryId:			u.CountryId ,
		IsPublic:			u.IsPublic ,
		Uid:				u.Uid ,
		PlantId:			u.PlantId ,
		PlantNameEN:		u.PlantNameEN ,
		PlantNameTH:		u.PlantNameTH ,
		PlantDescEN:		u.PlantDescEN ,
		PlantDescTH:		u.PlantDescTH ,
		PlantTypeId:		u.PlantTypeId ,
		TotalItem:			u.TotalItem ,
	}
}

//-------------------------------------------------------------------------------//
//				 	   	Join Fertilizer And Plant (View/Add/Edit)
//-------------------------------------------------------------------------------//

//Model
type JoinFertilizerAndPlant struct {
	FertilizerId     uuid.UUID	 `mapstructure:"fertilizer_id" json:"fertilizer_id"`
	FertilizerEN     string		 `mapstructure:"fertilizer_en" json:"fertilizer_en"`
	FertilizerTH     string		 `mapstructure:"fertilizer_th" json:"fertilizer_th"`
	Nitrogen       	 float64	 `mapstructure:"nitrogen" json:"nitrogen"`
	Phosphorus    	 float64	 `mapstructure:"phosphorus" json:"phosphorus"`
	Potassium      	 float64	 `mapstructure:"potassium" json:"potassium"`
	CreateDate		 time.Time	 `mapstructure:"create_date" json:"create_date"`
	ChangeDate	     time.Time	 `mapstructure:"change_date" json:"change_date"`
	FertilizerCatId	 uuid.UUID	 `mapstructure:"fertilizer_cat_id" json:"fertilizer_cat_id"`
	StatusId		 uuid.UUID	 `mapstructure:"status_id" json:"status_id"`
	FormulaPlantId   uuid.UUID	 `mapstructure:"formula_plant_id" json:"formula_plant_id"`
	Ratio		     float64	 `mapstructure:"ratio" json:"ratio"`
}

// New instance
func (u *JoinFertilizerAndPlant) New() *JoinFertilizerAndPlant {
	return &JoinFertilizerAndPlant{
		FertilizerId:		u.FertilizerId ,
		FertilizerEN:		u.FertilizerEN ,
		FertilizerTH:		u.FertilizerTH ,
		Nitrogen:			u.Nitrogen ,
		Phosphorus:			u.Phosphorus ,
		Potassium:			u.Potassium ,
		CreateDate:			u.CreateDate ,
		ChangeDate:			u.ChangeDate ,
		FertilizerCatId:	u.FertilizerCatId ,
		StatusId:			u.StatusId ,
		FormulaPlantId:		u.FormulaPlantId ,
		Ratio:				u.Ratio ,
	}
}


//-------------------------------------------------------------------------------//
//			Join SensorType And TransSensorValueRec (View/Add/Edit)
//-------------------------------------------------------------------------------//
//Model
type JoinSensorTypeAndTrans struct {
	SensorTypeId      	uuid.UUID	 `mapstructure:"sensor_type_id" json:"sensor_type_id"`
	SensorTypeNameEN    string	 	 `mapstructure:"sensor_type_name_en" json:"sensor_type_name_en"`
	CreateDate			time.Time	 `mapstructure:"create_date" json:"create_date"`
	ChangeDate	    	time.Time	  `mapstructure:"change_date" json:"change_date"`
	StatusId			uuid.UUID	 `mapstructure:"status_id" json:"status_id"`
	SensorTypeNameTH    string	 	 `mapstructure:"sensor_type_name_th" json:"sensor_type_name_th"`
	FormulaPlantId 		uuid.UUID	 `mapstructure:"formula_plant_id" json:"formula_plant_id"`
	ValueRec      		float64		 `mapstructure:"value_rec" json:"value_rec"`
}

// New instance
func (u *JoinSensorTypeAndTrans) New() *JoinSensorTypeAndTrans {
	return &JoinSensorTypeAndTrans{
		SensorTypeId:		u.SensorTypeId ,
		SensorTypeNameEN:	u.SensorTypeNameEN ,
		CreateDate:			u.CreateDate ,
		ChangeDate:			u.ChangeDate ,
		StatusId:			u.StatusId ,
		SensorTypeNameTH:	u.SensorTypeNameTH ,
		FormulaPlantId:		u.FormulaPlantId ,
		ValueRec:			u.ValueRec ,
	}
}
