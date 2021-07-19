package model_databases

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table farm
//-------------------------------------------------------------------------------//
//model farm
type Farm struct {
	FarmId      	uuid.UUID	 `gorm:"farm_id" json:"farm_id,omitempty"`
	FarmName    	string		 `gorm:"farm_name" json:"farm_name,omitempty"`
	FarmDesc    	string		 `gorm:"farm_desc" json:"farm_desc,omitempty"`
	CreateDate		time.Time	 `gorm:"create_date" json:"create_date,omitempty"`
	ChangeDate	    time.Time	 `gorm:"change_date" json:"change_date,omitempty"`
	StatusId		uuid.UUID	 `gorm:"status_id" json:"status_id,omitempty"`
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
