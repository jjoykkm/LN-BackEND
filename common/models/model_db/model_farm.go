package model_db

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/config"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table farm
//-------------------------------------------------------------------------------//
//model farm
type Farm struct {
<<<<<<< HEAD
	FarmId      	uuid.UUID	 `json:"FarmId,omitempty"`
	FarmName    	string		 `json:"FarmName,omitempty"`
	FarmDesc    	string		 `json:"FarmDesc,omitempty"`
	CreateDate		time.Time	 `json:"CreateDate,omitempty"`
	ChangeDate	    time.Time	 `json:"ChangeDate,omitempty"`
	StatusId		uuid.UUID	 `json:"StatusId,omitempty"`
=======
	FarmId      	uuid.UUID	 `mapstructure:"farm_id" json:"farm_id,omitempty"`
	FarmName    	string		 `mapstructure:"farm_name" json:"farm_name,omitempty"`
	FarmDesc    	string		 `mapstructure:"farm_desc" json:"farm_desc,omitempty"`
	CreateDate		time.Time	 `mapstructure:"create_date" json:"create_date,omitempty"`
	ChangeDate	    time.Time	 `mapstructure:"change_date" json:"change_date,omitempty"`
	StatusId		uuid.UUID	 `mapstructure:"status_id" json:"status_id,omitempty"`
>>>>>>> parent of b32535e (change model name)
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

