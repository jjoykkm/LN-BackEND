package model_other

//-------------------------------------------------------------------------------//
//				 		    METHOD POST REQUEST
//-------------------------------------------------------------------------------//
//Model post body
type ReqModel struct {
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
type RespModel struct {
	Item     interface{}	`json:"item"`
	Total    int			`json:"total"`
}

// New instance
func (u *RespModel) New() *RespModel {
	return &RespModel{
		Item:     u.Item ,
		Total:    u.Total ,
	}
}

//Model post body
type RespOffsetModel struct {
	Item     interface{}	`json:"item"`
	Offset   int			`json:"offset"`
	Total    int			`json:"total"`
}

// New instance
func (u *RespOffsetModel) New() *RespOffsetModel {
	return &RespOffsetModel{
		Item:      u.Item ,
		Offset:    u.Offset ,
		Total:     u.Total ,
	}
}

//Model response for success
type RespSuccessModel struct {
	MsgTh   	string		`json:"msg_th"`
	MsgEn    	string		`json:"msg_en"`
}
// New instance
func (u *RespSuccessModel) New() *RespSuccessModel {
	return &RespSuccessModel{
		MsgTh:      u.MsgTh ,
		MsgEn:    	u.MsgEn ,
	}
}
//-------------------------------------------------------------------------------//
//				 		    Error Context
//-------------------------------------------------------------------------------//
//Model post body

////-------------------------------------------------------------------------------//
////				 			 	METHOD POST
////-------------------------------------------------------------------------------//
////Model post body
//type PostBody struct {
//	Uid     		  string	`json:"uid"`
//	PlantId    		  string	`json:"plant_id"`
//	FormulaPlantId    string	`json:"formula_plant_id"`
//	PlantTypeId    	  string	`json:"plant_type_id"`
//	FarmId	    	  string	`json:"farm_id"`
//	FarmAreaId	      string	`json:"farm_area_id"`
//	FarmAreaIdList    []string	`json:"farm_area_id_list"`
//	Language    	  string	`json:"language"`
//	Offset    		  int		`json:"offset"`
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