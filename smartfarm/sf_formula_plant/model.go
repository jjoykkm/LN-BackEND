package sf_formula_plant

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/models/model_databases"
)

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
	PlantType   model_databases.PlantType	`mapstructure:"plant_type" json:"plant_type" gorm:"foreignkey:PlantTypeId; references:PlantTypeId"`
	Plant     	model_databases.Plant	 	`mapstructure:"plant" json:"plant" gorm:"embedded"`
}

// New instance
func (u *JoinPlantAndPlantType) New() *JoinPlantAndPlantType {
	return &JoinPlantAndPlantType{
		PlantType:  u.PlantType ,
		Plant:  	u.Plant ,
	}
}

func (JoinPlantAndPlantType) TableName() string {
	return "plant"
}


//-------------------------------------------------------------------------------//
//							Button Catagory
//-------------------------------------------------------------------------------//
//Model
type PlantCat struct {
	PlantTypeId     uuid.UUID	 `mapstructure:"plant_type_id" json:"plant_type_id"`
	PlantTypeName   string		 `mapstructure:"plant_type_name" json:"plant_type_name"`
	PlantId     	uuid.UUID	 `mapstructure:"plant_id" json:"plant_id"`
	PlantName       string		 `mapstructure:"plant_name" json:"plant_name"`
	PlantDesc       string		 `mapstructure:"plant_desc" json:"plant_desc"`
	TotalItem       int64	 	 `mapstructure:"total_item" json:"total_item"`
}

// New instance
func (u *PlantCat) New() *PlantCat {
	return &PlantCat{
		PlantTypeId:  	u.PlantTypeId ,
		PlantTypeName:  u.PlantTypeName ,
		PlantId:      	u.PlantId ,
		PlantName:  	u.PlantName ,
		PlantDesc:  	u.PlantDesc ,
		TotalItem:    	u.TotalItem ,
	}
}