package helper_db

//-------------------------------------------------------------------------------//
//				 		    METHOD POST REQUEST
//-------------------------------------------------------------------------------//
//Model post body
type PostReqModel struct {
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
func (u *PostReqModel) New() *PostReqModel {
	return &PostReqModel{
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
type PostRespModel struct {
	Item     interface{}	`json:"item"`
	Total    int			`json:"total"`
}

// New instance
func (u *PostRespModel) New() *PostRespModel {
	return &PostRespModel{
		Item:     u.Item ,
		Total:    u.Total ,
	}
}

////-------------------------------------------------------------------------------//
////							Table plant_type
////-------------------------------------------------------------------------------//
//type BatchJobInitialReq struct {
//	BatchName string `json:"batchName" validate:"nonzero"`
//	DataDate  string `json:"dataDate"`
//}
//
////-------------------------------------------------------------------------------//
////							Table plant_type
////-------------------------------------------------------------------------------//
//type BatchJobInitialResp struct {
//	DataDate string `json:"dataDate"`
//}
