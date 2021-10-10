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
	FormulaPlantId 	 uuid.UUID	 `son:"formula_plant_id,omitempty"`
	FormulaName		 string		 `json:"name,omitempty"`
	FormulaDesc		 string		 `json:"desc,omitempty"`
	PeopleUsed 		 int		 `json:"people_used,omitempty"`
	Recommend1 		 int		 `json:"rec_1,omitempty"`
	Recommend2		 int		 `json:"rec_2,omitempty"`
	Recommend3		 int		 `json:"rec_3,omitempty"`
	Recommend4		 int		 `json:"rec_4,omitempty"`
	Recommend5		 int		 `json:"rec_5,omitempty"`
	CreateDate		 time.Time	 `json:"create_date,omitempty"`
	ChangeDate		 time.Time	 `json:"change_date,omitempty"`
	PlantId		 	 uuid.UUID	 `json:"plant_id,omitempty"`
	StatusId		 uuid.UUID	 `json:"status_id,omitempty"`
	ProvinceId		 uuid.UUID	 `json:"province_id,omitempty"`
	CountryId		 uuid.UUID	 `json:"country_id,omitempty"`
	IsPublic		 bool	 	 `json:"is_public,omitempty"`
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











