package models

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table formula_plant
//-------------------------------------------------------------------------------//
//model formula_plant
type FormulaPlant struct {
	FormulaPlantId 	 uuid.UUID	 `gorm:"formula_plant_id" json:"formula_plant_id,omitempty"`
	FormulaName		 string		 `gorm:"formula_name" json:"formula_name,omitempty"`
	FormulaDesc		 string		 `gorm:"formula_desc" json:"formula_desc,omitempty"`
	PeopleUsed 		 int		 `gorm:"people_used" json:"people_used,omitempty"`
	Recommend1 		 int		 `gorm:"recommend1" json:"recommend1,omitempty"`
	Recommend2		 int		 `gorm:"recommend2" json:"recommend2,omitempty"`
	Recommend3		 int		 `gorm:"recommend3" json:"recommend3,omitempty"`
	Recommend4		 int		 `gorm:"recommend4" json:"recommend4,omitempty"`
	Recommend5		 int		 `gorm:"recommend5" json:"recommend5,omitempty"`
	CreateDate		 time.Time	 `gorm:"create_date" json:"create_date,omitempty"`
	ChangeDate		 time.Time	 `gorm:"change_date" json:"change_date,omitempty"`
	PlantId		 	 uuid.UUID	 `gorm:"plant_id" json:"plant_id,omitempty"`
	StatusId		 uuid.UUID	 `gorm:"status_id" json:"status_id,omitempty"`
	ProvinceId		 uuid.UUID	 `gorm:"province_id" json:"province_id,omitempty"`
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
	}
}












