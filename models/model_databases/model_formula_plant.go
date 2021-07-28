package model_databases

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table formula_plant
//-------------------------------------------------------------------------------//
//model formula_plant
type FormulaPlant struct {
	FormulaPlantId 	 uuid.UUID	 `mapstructure:"formula_plant_id" json:"formula_plant_id,omitempty"`
	FormulaName		 string		 `mapstructure:"formula_name" json:"formula_name,omitempty"`
	FormulaDesc		 string		 `mapstructure:"formula_desc" json:"formula_desc,omitempty"`
	PeopleUsed 		 int		 `mapstructure:"people_used" json:"people_used,omitempty"`
	Recommend1 		 int		 `mapstructure:"recommend1" json:"recommend1,omitempty"`
	Recommend2		 int		 `mapstructure:"recommend2" json:"recommend2,omitempty"`
	Recommend3		 int		 `mapstructure:"recommend3" json:"recommend3,omitempty"`
	Recommend4		 int		 `mapstructure:"recommend4" json:"recommend4,omitempty"`
	Recommend5		 int		 `mapstructure:"recommend5" json:"recommend5,omitempty"`
	CreateDate		 time.Time	 `mapstructure:"create_date" json:"create_date,omitempty"`
	ChangeDate		 time.Time	 `mapstructure:"change_date" json:"change_date,omitempty"`
	PlantId		 	 uuid.UUID	 `mapstructure:"plant_id" json:"plant_id,omitempty"`
	StatusId		 uuid.UUID	 `mapstructure:"status_id" json:"status_id,omitempty"`
	ProvinceId		 uuid.UUID	 `mapstructure:"province_id" json:"province_id,omitempty"`
	CountryId		 uuid.UUID	 `mapstructure:"country_id" json:"country_id,omitempty"`
	IsPublic		 bool	 	 `mapstructure:"is_public" json:"is_public,omitempty"`
	Uid				 bool	 	 `mapstructure:"uid" json:"uid,omitempty"`
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












