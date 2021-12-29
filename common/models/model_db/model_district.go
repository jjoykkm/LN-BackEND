package model_db

import (
	"github.com/jjoykkm/ln-backend/common/config"
)

//-------------------------------------------------------------------------------//
//							Table district
//-------------------------------------------------------------------------------//
//model district
type District struct {
	DBCommonGet
	DistrictId      string	 	 `json:"district_id"`
	DistrictEN      string	 `json:"district_name_en"`
	DistrictTH      string	 `json:"district_name_th"`
}
// New instance district
func (u *District) New() *District {
	return &District{
		DBCommonGet:      	u.DBCommonGet ,
		DistrictId:		u.DistrictId ,
		DistrictEN:		u.DistrictEN ,
		DistrictTH:		u.DistrictTH ,
	}
}

// Custom table name for GORM
func (District) TableName() string {
	return config.DB_DISTRICT
}


