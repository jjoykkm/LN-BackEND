package sf_formula_plant

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/helper"
)
//-------------------------------------------------------------------------------//
//							Body Request
//-------------------------------------------------------------------------------//
type BodyReq helper.PostBodyReq

//-------------------------------------------------------------------------------//
//							Body Response
//-------------------------------------------------------------------------------//
type BodyResp helper.PostBodyResp

//-------------------------------------------------------------------------------//
//							Set Button Catagory
//-------------------------------------------------------------------------------//
//Model
type PlantTypeCat struct {
	PlantTypeId     uuid.UUID	 `mapstructure:"plant_type_id" json:"plant_type_id"`
	PlantTypeName   string		 `mapstructure:"plant_type_name" json:"plant_type_name"`
}

// New instance
func (u *PlantTypeCat) New() *PlantTypeCat {
	return &PlantTypeCat{
		PlantTypeId:  	u.PlantTypeId ,
		PlantTypeName:  u.PlantTypeName ,
	}
}