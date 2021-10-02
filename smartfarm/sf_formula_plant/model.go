package sf_formula_plant

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/helper"
	"time"
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
