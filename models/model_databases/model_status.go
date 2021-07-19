package model_databases

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table status
//-------------------------------------------------------------------------------//
//model status
type Status struct {
	StatusId      	uuid.UUID	 `gorm:"status_id" json:"status_id,omitempty"`
	StatusName      string		 `gorm:"status_name" json:"status_name,omitempty"`
	CreateDate		time.Time	 `gorm:"create_date" json:"create_date,omitempty"`
	ChangeDate	    time.Time	 `gorm:"change_date" json:"change_date,omitempty"`
}
// New instance status
func (u *Status) New() *Status {
	return &Status{
		StatusId:			u.StatusId ,
		StatusName:			u.StatusName ,
		CreateDate:			u.CreateDate ,
		ChangeDate:			u.ChangeDate ,
	}
}
