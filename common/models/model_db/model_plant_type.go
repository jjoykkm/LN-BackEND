package model_db

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/common/config"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table plant_type
//-------------------------------------------------------------------------------//
//model plant_type
type PlantType struct {
	PlantTypeId      uuid.UUID	 `mapstructure:"plant_type_id" json:"plant_type_id,omitempty"`
	PlantTypeEN      string		 `mapstructure:"plant_type_en" json:"plant_type_en,omitempty"`
	PlantTypeTH      string		 `mapstructure:"plant_type_th" json:"plant_type_th,omitempty"`
	CreateDate		 time.Time	 `mapstructure:"create_date" json:"create_date,omitempty"`
	ChangeDate	     time.Time	 `mapstructure:"change_date" json:"change_date,omitempty"`
	StatusId		 uuid.UUID	 `mapstructure:"status_id" json:"status_id,omitempty"`
}
// New instance plant_type
func (u *PlantType) New() *PlantType {
	return &PlantType{
		PlantTypeId:	u.PlantTypeId ,
		PlantTypeEN:	u.PlantTypeEN ,
		PlantTypeTH:	u.PlantTypeTH ,
		CreateDate:		u.CreateDate ,
		ChangeDate:		u.ChangeDate ,
		StatusId:		u.StatusId ,
	}
}

// Custom table name for GORM
func (PlantType) TableName() string {
	return config.DB_PLANT_TYPE
}
