package model_db

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/common/config"
)

//-------------------------------------------------------------------------------//
//							Table district
//-------------------------------------------------------------------------------//
//model district
type District struct {
	DBCommon
	DistrictId      uuid.UUID	 `json:"district_id"`
	DistrictEN      string		 `json:"district_name_en"`
	DistrictTH      string		 `json:"district_name_th"`
}
// New instance district
func (u *District) New() *District {
	return &District{
		DBCommon:      	u.DBCommon ,
		DistrictId:		u.DistrictId ,
		DistrictEN:		u.DistrictEN ,
		DistrictTH:		u.DistrictTH ,
	}
}

// Custom table name for GORM
func (District) TableName() string {
	return config.DB_DISTRICT
}


