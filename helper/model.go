package helper

//-------------------------------------------------------------------------------//
//				 		    METHOD POST REQUEST
//-------------------------------------------------------------------------------//
//Model post body
type PostBodyReq struct {
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
func (u *PostBodyReq) New() *PostBodyReq {
	return &PostBodyReq{
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
type PostBodyResp struct {
	Item     interface{}	`mapstructure:"item" json:"item"`
	Total    int			`mapstructure:"total" json:"total"`
}

// New instance
func (u *PostBodyResp) New() *PostBodyResp {
	return &PostBodyResp{
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
