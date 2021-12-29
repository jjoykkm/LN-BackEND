package Obsolete_model_services

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"time"
)
//-------------------------------------------------------------------------------//
//							Set Button Catagory
//-------------------------------------------------------------------------------//
//Model
type ForPlantCatList struct {
	PlantTypeId     string	 `json:"plant_type_id"`
	PlantTypeName   string	 `json:"plant_type_name"`
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
	PlantTypeId     string	 `json:"plant_type_id"`
	PlantTypeName   string	 `json:"plant_type_name"`
	PlantId     	string	 `json:"plant_id"`
	PlantName       string	 `json:"plant_name"`
	PlantDesc       string	 `json:"plant_desc"`
	TotalItem       int		 	 `json:"total_item"`
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
	PlantTypeId      string	 `json:"plant_type_id"`
	PlantTypeName    string	 `json:"plant_type_name"`
	FormulaPlantId 	 string	 `json:"formula_plant_id"`
	FormulaName		 string	 `json:"formula_name"`
	FormulaDesc		 string	 `json:"formula_desc"`
	PeopleUsed 		 int		 `json:"people_used"`
	RateScore 		 float32	 `json:"rate_score"`
	RatePeople 		 int		 `json:"rate_people"`
	CountryId		 string	 `json:"country_id"`
	CountryName      string	 `json:"country_name"`
	ProvinceId		 string	 `json:"province_id"`
	ProvinceName     string	 `json:"province_name"`
	IsPublic		 bool	 	 `json:"is_public"`
	IsPlanted		 bool	 	 `json:"is_planted"`
	IsFavorite		 bool	 	 `json:"is_favorite"`
	Uid				 string 	 `json:"uid"`
	Username     	 string	 	 `json:"username"`
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
	PlantId     	string	 		`json:"plant_id"`
	PlantName       string	 		`json:"plant_name"`
	PlantDesc       string	 		`json:"plant_desc"`
	TotalItem       string	 		`json:"total_item"`
	PlantCatItem	[]ForPlantItem		`json:"plant_cat_item"`
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
	FertilizerId     string	 `json:"fertilizer_id"`
	FertilizerName   string	 `json:"fertilizer_name"`
	Nitrogen       	 float64	 `json:"nitrogen"`
	Phosphorus    	 float64	 `json:"phosphorus"`
	Potassium      	 float64	 `json:"potassium"`
	Ratio	      	 float64	 `json:"ratio"`
	FertCatId      	 string	 `json:"fertilizer_cat_id"`
	FertCatName    	 string	 `json:"fertilizer_cat_name"`
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
	SensorTypeId      	string	 `json:"sensor_type_id"`
	SensorTypeName      string	 	 `json:"sensor_type_name"`
	ValueRec      		float64		 `json:"value_rec"`
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
	FormulaPlantId 	 string	 			`json:"formula_plant_id"`
	FormulaName		 string	 			`json:"formula_name"`
	FormulaDesc		 string	 			`json:"formula_desc"`
	Uid			 	 string	 			`json:"uid"`
	Username		 string	 			`json:"username"`
	SensorList	 	 []ForPlantSensor		`json:"sensor_list"`
	FertList		 []ForPlantFert			`json:"fert_list"`
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
	PlantId          string	 `json:"plant_id"`
	PlantNameEN      string	 `json:"plant_name_en"`
	PlantNameTH      string	 `json:"plant_name_th"`
	PlantDescEN      string	 `json:"plant_desc_en"`
	PlantDescTH      string	 `json:"plant_desc_th"`
	CreateDate		 time.Time	 `json:"create_date"`
	ChangeDate	     time.Time	 `json:"change_date"`
	StatusId		 string	 `json:"status_id"`
	TotalItem      	 int		 `json:"total_item"`
	PlantTypeId      string	 `json:"plant_type_id"`
	PlantTypeEN      string	 `json:"plant_type_en"`
	PlantTypeTH      string	 `json:"plant_type_th"`
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
	FormulaPlantId 	 string	 `json:"formula_plant_id"`
	FormulaName		 string	 `json:"formula_name"`
	FormulaDesc		 string	 `json:"formula_desc"`
	PeopleUsed 		 int		 `json:"people_used"`
	Recommend1 		 int		 `json:"recommend1"`
	Recommend2		 int		 `json:"recommend2"`
	Recommend3		 int		 `json:"recommend3"`
	Recommend4		 int		 `json:"recommend4"`
	Recommend5		 int		 `json:"recommend5"`
	CreateDate		 time.Time	 `json:"create_date"`
	ChangeDate		 time.Time	 `json:"change_date"`
	StatusId		 string	 `json:"status_id"`
	ProvinceId		 string	 `json:"province_id"`
	CountryId		 string	 `json:"country_id"`
	IsPublic		 bool	 	 `json:"is_public"`
	Uid				 string	 `json:"uid"`
	PlantId          string	 `json:"plant_id"`
	PlantNameEN      string	 `json:"plant_name_en"`
	PlantNameTH      string	 `json:"plant_name_th"`
	PlantDescEN      string	 `json:"plant_desc_en"`
	PlantDescTH      string	 `json:"plant_desc_th"`
	PlantTypeId      string	 `json:"plant_type_id"`
	TotalItem      	 int		 `json:"total_item"`
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
	FertilizerId     string	 `json:"fertilizer_id"`
	FertilizerEN     string	 `json:"fertilizer_en"`
	FertilizerTH     string	 `json:"fertilizer_th"`
	Nitrogen       	 float64	 `json:"nitrogen"`
	Phosphorus    	 float64	 `json:"phosphorus"`
	Potassium      	 float64	 `json:"potassium"`
	CreateDate		 time.Time	 `json:"create_date"`
	ChangeDate	     time.Time	 `json:"change_date"`
	FertilizerCatId	 string	 `json:"fertilizer_cat_id"`
	StatusId		 string	 `json:"status_id"`
	FormulaPlantId   string	 `json:"formula_plant_id"`
	Ratio		     float64	 `json:"ratio"`
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
	SensorTypeId      	string	 `json:"sensor_type_id"`
	SensorTypeNameEN    string	 	 `json:"sensor_type_name_en"`
	CreateDate			time.Time	 `json:"create_date"`
	ChangeDate	    	time.Time	 `json:"change_date"`
	StatusId			string	 `json:"status_id"`
	SensorTypeNameTH    string	 	 `json:"sensor_type_name_th"`
	FormulaPlantId 		string	 `json:"formula_plant_id"`
	ValueRec      		float64		 `json:"value_rec"`
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
