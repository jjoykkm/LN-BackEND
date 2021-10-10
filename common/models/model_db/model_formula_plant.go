package model_db

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/common/config"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table formula_plant
//-------------------------------------------------------------------------------//
//model formula_plant
type FormulaPlant struct {
	FormulaPlantId 	 uuid.UUID	 `json:"ForPlantId,omitempty"`
	FormulaName		 string		 `json:"ForName,omitempty"`
	FormulaDesc		 string		 `json:"ForDesc,omitempty"`
	PeopleUsed 		 int		 `json:"PeopleUsed,omitempty"`
	Recommend1 		 int		 `json:"Recommend_1,omitempty"`
	Recommend2		 int		 `json:"Recommend_2,omitempty"`
	Recommend3		 int		 `json:"Recommend_3,omitempty"`
	Recommend4		 int		 `json:"Recommend_4,omitempty"`
	Recommend5		 int		 `json:"Recommend_5,omitempty"`
	CreateDate		 time.Time	 `json:"CreateDate,omitempty"`
	ChangeDate		 time.Time	 `json:"ChangeDate,omitempty"`
	PlantId		 	 uuid.UUID	 `json:"PlantId,omitempty"`
	StatusId		 uuid.UUID	 `json:"StatusId,omitempty"`
	ProvinceId		 uuid.UUID	 `json:"ProvinceId,omitempty"`
	CountryId		 uuid.UUID	 `json:"CountryId,omitempty"`
	IsPublic		 bool	 	 `json:"IsPublic,omitempty"`
	Uid				 uuid.UUID	 `json:"-"`
}
// New instance formula_plant
func (u *FormulaPlant) New() *FormulaPlant {
	return &FormulaPlant{
		FormulaPlantId:		u.FormulaPlantId ,
		FormulaName:		u.FormulaName ,
		FormulaDesc:		u.FormulaDesc ,
		PeopleUsed:			u.PeopleUsed ,
		Recommend1:			u.Recommend1 ,
		Recommend2:			u.Recommend2 ,
		Recommend3:			u.Recommend3 ,
		Recommend4:			u.Recommend4 ,
		Recommend5:			u.Recommend5 ,
		CreateDate:			u.CreateDate ,
		ChangeDate:			u.ChangeDate ,
		PlantId:			u.PlantId ,
		StatusId:			u.StatusId ,
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











