package model_other


//-------------------------------------------------------------------------------//
//				 			 	METHOD POST
//-------------------------------------------------------------------------------//
//Model post body
type PostBody struct {
	Uid     		  string	`mapstructure:"uid" json:"uid"`
	PlantId    		  string	`mapstructure:"plant_id" json:"plant_id"`
	FormulaPlantId    string	`mapstructure:"formula_plant_id" json:"formula_plant_id"`
	PlantTypeId    	  string	`mapstructure:"plant_type_id" json:"plant_type_id"`
	FarmId	    	  string	`mapstructure:"farm_id" json:"farm_id"`
	Language    	  string	`mapstructure:"language" json:"language"`
	Offset    		  int		`mapstructure:"offset" json:"offset"`
}

// New instance
func (u *PostBody) New() *PostBody {
	return &PostBody{
		Uid:      			u.Uid ,
		PlantId:    		u.PlantId ,
		FormulaPlantId:   	u.FormulaPlantId ,
		PlantTypeId:   		u.PlantTypeId ,
		FarmId:   			u.FarmId ,
		Language:   		u.Language ,
		Offset:   			u.Offset ,
	}
}

//{
//	"uid": "",
//	"plant_id": "",
//	"formula_plant_id": "",
//	"plant_type_id": "",
//	"language": "",
//	"offset": 0
//}