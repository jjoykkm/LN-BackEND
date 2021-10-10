package model_databases

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/common/config"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table farm
//-------------------------------------------------------------------------------//
//model farm
type Farm struct {
	FarmId      	uuid.UUID	 `json:"farm_id,omitempty"`
	FarmName    	string		 `json:"farm_name,omitempty"`
	FarmDesc    	string		 `json:"farm_desc,omitempty"`
	CreateDate		time.Time	 `json:"create_date,omitempty"`
	ChangeDate	    time.Time	 `json:"change_date,omitempty"`
	StatusId		uuid.UUID	 `json:"status_id,omitempty"`
}
// New instance farm
func (u *Farm) New() *Farm {
	return &Farm{
		FarmId:			u.FarmId ,
		FarmName:		u.FarmName ,
		FarmDesc:		u.FarmDesc ,
		CreateDate:		u.CreateDate ,
		ChangeDate:		u.ChangeDate ,
		StatusId:		u.StatusId ,
	}
}

// Custom table name for GORM
func (Farm) TableName() string {
	return config.DB_FARM
}

