package model_other

//-------------------------------------------------------------------------------//
//				 		    METHOD POST REQUEST
//-------------------------------------------------------------------------------//
//Model post body
type ReqModel struct {
	Uid     		  string	`mapstructure:"uid" json:"uid"`
	PlantId    		  string	`mapstructure:"plant_id" json:"plant_id"`
	FormulaPlantId    string	`mapstructure:"formula_plant_id" json:"formula_plant_id"`
	PlantTypeId    	  string	`mapstructure:"plant_type_id" json:"plant_type_id"`
	FarmId	    	  string	`mapstructure:"farm_id" json:"farm_id"`
	FarmAreaId	      string	`mapstructure:"farm_area_id" json:"farm_area_id"`
	FarmAreaIdList    []string	`mapstructure:"farm_area_id_list" json:"farm_area_id_list"`
	Language    	  string	`mapstructure:"language" json:"language"`
	Offset    		  int		`mapstructure:"offset" json:"offset"`
}

// New instance
func (u *ReqModel) New() *ReqModel {
	return &ReqModel{
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

//-------------------------------------------------------------------------------//
//				 		    METHOD POST RESPONSE
//-------------------------------------------------------------------------------//
//Model post body
type BodyResp struct {
	Item     interface{}	`mapstructure:"item" json:"item"`
	Total    int			`mapstructure:"total" json:"total"`
}

// New instance
func (u *BodyResp) New() *BodyResp {
	return &BodyResp{
		Item:     u.Item ,
		Total:    u.Total ,
	}
}

//Model post body
type BodyRespOffset struct {
	Item     interface{}	`mapstructure:"item" json:"item"`
	Offset   int			`mapstructure:"offset" json:"offset"`
	Total    int			`mapstructure:"total" json:"total"`
}

// New instance
func (u *BodyRespOffset) New() *BodyRespOffset {
	return &BodyRespOffset{
		Item:      u.Item ,
		Offset:    u.Offset ,
		Total:     u.Total ,
	}
}



////-------------------------------------------------------------------------------//
////				 			 	METHOD POST
////-------------------------------------------------------------------------------//
////Model post body
//type PostBody struct {
//	Uid     		  string	`mapstructure:"uid" json:"uid"`
//	PlantId    		  string	`mapstructure:"plant_id" json:"plant_id"`
//	FormulaPlantId    string	`mapstructure:"formula_plant_id" json:"formula_plant_id"`
//	PlantTypeId    	  string	`mapstructure:"plant_type_id" json:"plant_type_id"`
//	FarmId	    	  string	`mapstructure:"farm_id" json:"farm_id"`
//	FarmAreaId	      string	`mapstructure:"farm_area_id" json:"farm_area_id"`
//	FarmAreaIdList    []string	`mapstructure:"farm_area_id_list" json:"farm_area_id_list"`
//	Language    	  string	`mapstructure:"language" json:"language"`
//	Offset    		  int		`mapstructure:"offset" json:"offset"`
//}
//
//// New instance
//func (u *PostBody) New() *PostBody {
//	return &PostBody{
//		Uid:      			u.Uid ,
//		PlantId:    		u.PlantId ,
//		FormulaPlantId:   	u.FormulaPlantId ,
//		PlantTypeId:   		u.PlantTypeId ,
//		FarmId:   			u.FarmId ,
//		FarmAreaId:   		u.FarmAreaId ,
//		FarmAreaIdList:   	u.FarmAreaIdList ,
//		Language:   		u.Language ,
//		Offset:   			u.Offset ,
//	}
//}

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