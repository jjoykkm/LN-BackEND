package model_db

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
	FarmId      	uuid.UUID	 `json:"farm_id"`
	FarmName    	string		 `json:"farm_name"`
	FarmDesc    	string		 `json:"farm_desc"`
	CreateDate		time.Time	 `json:"create_date"`
	ChangeDate	    time.Time	 `json:"change_date"`
	StatusId		uuid.UUID	 `json:"status_id"`
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

//-------------------------------------------------------------------------------//
//							Upsert Farm
//-------------------------------------------------------------------------------//
type FarmUS struct {
	FarmId      	string	 `json:"farm_id" gorm:"default:uuid_generate_v4()"`
	FarmName    	string	 `json:"farm_name"`
	FarmDesc    	string	 `json:"farm_desc"`
	StatusId		string	 `json:"status_id"`
}
func (FarmUS) TableName() string {
	return config.DB_FARM
}

