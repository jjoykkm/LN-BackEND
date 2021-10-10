package model_other


//-------------------------------------------------------------------------------//
//				 			 	METHOD POST
//-------------------------------------------------------------------------------//
//Model post body
type PostBody struct {
	Uid     		  string	`json:"uid"`
	PlantId    		  string	`json:"plant_id"`
	FormulaPlantId    string	`json:"formula_plant_id"`
	PlantTypeId    	  string	`json:"plant_type_id"`
	FarmId	    	  string	`json:"farm_id"`
	FarmAreaId	      string	`json:"farm_area_id"`
	FarmAreaIdList    []string	`json:"farm_area_id_list"`
	Language    	  string	`json:"language"`
	Offset    		  int		`json:"offset"`
}

// New instance
func (u *PostBody) New() *PostBody {
	return &PostBody{
		Uid:      			u.Uid ,
		PlantId:    		u.PlantId ,
		FormulaPlantId:   	u.FormulaPlantId ,
		PlantTypeId:   		u.PlantTypeId ,
		FarmId:   			u.FarmId ,
		FarmAreaId:   		u.FarmAreaId ,
		FarmAreaIdList:   	u.FarmAreaIdList ,
		Language:   		u.Language ,
		Offset:   			u.Offset ,
	}
}

//{
//	"uid": "",
//	"plant_id": "",
//	"formula_plant_id": "",
//	"plant_type_id": "",
//	"farm_id": "",
//	"farm_area_id": "",
//	"language": "",
//	"offset": 0
//}