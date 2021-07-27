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
	StatusId      	uuid.UUID	 `mapstructure:"status_id" json:"status_id,omitempty"`
	StatusName      string		 `mapstructure:"status_name" json:"status_name,omitempty"`
	CreateDate		time.Time	 `mapstructure:"create_date" json:"create_date,omitempty"`
	ChangeDate	    time.Time	 `mapstructure:"change_date" json:"change_date,omitempty"`
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
