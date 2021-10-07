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

//-------------------------------------------------------------------------------//
//				 	   	Join Plant And PlantType (View/Add/Edit)
//-------------------------------------------------------------------------------//
//Model
type Plant struct {
	Plant  model_databases.Plant `gorm:"embedded"`
	PlantType  model_databases.PlantType `gorm:"foreignkey:PlantTypeId; references:PlantTypeId"`
}
func (Plant) TableName() string {
	return "plant"
}

type JoinPlantAndFormulaPlant struct {
	FormulaPlant    model_databases.FormulaPlant	`mapstructure:"formula_plant" json:"formula_plant" gorm:"embedded"`
	Plant     		Plant			 		`mapstructure:"plant" json:"plant" gorm:"foreignkey:PlantId; references:PlantId;"`
	//Plant     		model_databases.Plant	 		`mapstructure:"plant" json:"plant" gorm:"foreignkey:PlantId; references:PlantId;"`
	//PlantType   	model_databases.PlantType		`mapstructure:"plant_type" json:"plant_type" gorm:"-"` // gorm:"foreignkey:PlantTypeId; references:PlantTypeId;"
	Province   		model_databases.Province		`mapstructure:"province" json:"province" gorm:"foreignkey:ProvinceId; references:ProvinceId"` // association_foreignkey:ProvinceId
	Country   		model_databases.Country			`mapstructure:"country" json:"country" gorm:"foreignkey:CountryId; references:CountryId"`

	//FormulaPlant    model_databases.FormulaPlant	`mapstructure:"formula_plant" json:"formula_plant" gorm:"embedded"`
	//Plant     		model_databases.Plant	 		`mapstructure:"plant" json:"plant" gorm:"foreignkey:PlantId; references:PlantId"`

}

// New instance
//func (u *JoinPlantAndFormulaPlant) New() *JoinPlantAndFormulaPlant {
//	return &JoinPlantAndFormulaPlant{
//		FormulaPlant:   u.FormulaPlant ,
//		Plant:  		u.Plant ,
//		PlantType:  	u.PlantType ,
//	}
//}

func (JoinPlantAndFormulaPlant) TableName() string {
	//return "plant"
	return "formula_plant"
}

//-------------------------------------------------------------------------------//
//				  Button Favorite and My formula (Card Item)
//-------------------------------------------------------------------------------//
//Model
type ForPlantItem struct {
	PlantType		 model_databases.PlantType
	FormulaPlant	 model_databases.FormulaPlant
	RateScore 		 float32	 `mapstructure:"rate_score" json:"rate_score"`
	RatePeople 		 int		 `mapstructure:"rate_people" json:"rate_people"`
	IsPublic		 bool	 	 `mapstructure:"is_public" json:"is_public"`
	IsPlanted		 bool	 	 `mapstructure:"is_planted" json:"is_planted"`
	IsFavorite		 bool	 	 `mapstructure:"is_favorite" json:"is_favorite"`
}

// New instance
//func (u *ForPlantItem) New() *ForPlantItem {
//	return &ForPlantItem{
//		PlantTypeId:		u.PlantTypeId ,
//		PlantTypeName:  	u.PlantTypeName ,
//		FormulaPlantId:		u.FormulaPlantId ,
//		FormulaName:		u.FormulaName ,
//		FormulaDesc:		u.FormulaDesc ,
//		PeopleUsed:			u.PeopleUsed ,
//		RateScore:			u.RateScore ,
//		RatePeople:			u.RatePeople ,
//		CountryId:			u.CountryId ,
//		CountryName:		u.CountryName ,
//		ProvinceId:			u.ProvinceId ,
//		ProvinceName:		u.ProvinceName ,
//		IsPublic:			u.IsPublic ,
//		IsPlanted:			u.IsPlanted ,
//		IsFavorite:			u.IsFavorite ,
//		Uid:				u.Uid ,
//		Username:			u.Username ,
//	}
//}