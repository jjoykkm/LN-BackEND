package model_db

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/common/config"
)
//-------------------------------------------------------------------------------//
//							Table Recommend
//-------------------------------------------------------------------------------//
//model formula_plant
type Recommend struct {
	Recommend1 		int		`json:"star_1,"`
	Recommend2		int		`json:"star_2,"`
	Recommend3		int		`json:"star_3,"`
	Recommend4		int		`json:"star_4,"`
	Recommend5		int		`json:"star_5,"`
}
// New instance formula_plant
func (u *Recommend) New() *Recommend {
	return &Recommend{
		Recommend1:			u.Recommend1 ,
		Recommend2:			u.Recommend2 ,
		Recommend3:			u.Recommend3 ,
		Recommend4:			u.Recommend4 ,
		Recommend5:			u.Recommend5 ,
	}
}

//-------------------------------------------------------------------------------//
//							Table formula_plant
//-------------------------------------------------------------------------------//
//model formula_plant
type FormulaPlant struct {
	DBCommon
	FormulaPlantId 	 uuid.UUID	 `json:"formula_plant_id"`
	FormulaName		 string		 `json:"formula_plant_name"`
	FormulaDesc		 string		 `json:"formula_plant_desc"`
	PeopleUsed 		 int		 `json:"people_used"`
	Recommend		 Recommend	 `json:"recommend" gorm:"embedded"`
	PlantId		 	 uuid.UUID	 `json:"plant_id"`
	ProvinceId		 uuid.UUID	 `json:"province_id"`
	CountryId		 uuid.UUID	 `json:"country_id"`
	IsPublic		 bool	 	 `json:"is_public"`
	Uid				 uuid.UUID	 `json:"-"`
}
// New instance formula_plant
func (u *FormulaPlant) New() *FormulaPlant {
	return &FormulaPlant{
		DBCommon:      		u.DBCommon ,
		FormulaPlantId:		u.FormulaPlantId ,
		FormulaName:		u.FormulaName ,
		FormulaDesc:		u.FormulaDesc ,
		PeopleUsed:			u.PeopleUsed ,
		Recommend:			u.Recommend ,
		PlantId:			u.PlantId ,
		ProvinceId:			u.ProvinceId ,
		CountryId:			u.CountryId ,
		IsPublic:			u.IsPublic ,
		Uid:				u.Uid ,
	}
}

// Custom table name for GORM
func (FormulaPlant) TableName() string {
	return config.DB_FORMULA_PLANT
}











