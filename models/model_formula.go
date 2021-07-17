package models

import (
	"github.com/jackc/pgtype/ext/gofrs-uuid"
	"time"
)

//model formula_plant
type FormulaPlant struct {
	formulaPlantId 	 uuid.UUID	 `gorm:"formula_plant_id" json:"formula_plant_id,omitempty"`
	formulaName		 string		 `gorm:"formula_name" json:"formula_name,omitempty"`
	formulaDesc		 string		 `gorm:"formula_desc" json:"formula_desc,omitempty"`
	peopleUsed 		 int		 `gorm:"people_used" json:"people_used,omitempty"`
	recommend1 		 int		 `gorm:"recommend1" json:"recommend1,omitempty"`
	recommend2		 int		 `gorm:"recommend2" json:"recommend2,omitempty"`
	recommend3		 int		 `gorm:"recommend3" json:"recommend3,omitempty"`
	recommend4		 int		 `gorm:"recommend4" json:"recommend4,omitempty"`
	recommend5		 int		 `gorm:"recommend5" json:"recommend5,omitempty"`
	createDate		 time.Time	 `gorm:"create_date" json:"create_date,omitempty"`
	changeDate		 time.Time	 `gorm:"change_date" json:"change_date,omitempty"`
	plantId		 	 uuid.UUID	 `gorm:"plant_id" json:"plant_id,omitempty"`
	statusId		 uuid.UUID	 `gorm:"status_id" json:"status_id,omitempty"`
	provinceId		 uuid.UUID	 `gorm:"province_id" json:"province_id,omitempty"`
}
// New instance formula_plant
func (u *FormulaPlant) New() *FormulaPlant {
	return &FormulaPlant{
		formulaPlantId:		u.formulaPlantId ,
		formulaName:		u.formulaName ,
		formulaDesc:		u.formulaDesc ,
		peopleUsed:			u.peopleUsed ,
		recommend1:			u.recommend1 ,
		recommend2:			u.recommend2 ,
		recommend3:			u.recommend3 ,
		recommend4:			u.recommend4 ,
		recommend5:			u.recommend5 ,
		createDate:			u.createDate ,
		changeDate:			u.changeDate ,
		plantId:			u.plantId ,
		statusId:			u.statusId ,
		provinceId:			u.provinceId ,
	}
}
