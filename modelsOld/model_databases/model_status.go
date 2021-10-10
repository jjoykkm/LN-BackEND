package model_databases

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/common/config"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table status
//-------------------------------------------------------------------------------//
//model status
type Status struct {
	StatusId      	uuid.UUID	 `json:"status_id,omitempty"`
	StatusName      string		 `json:"status_name,omitempty"`
	CreateDate		time.Time	 `json:"create_date,omitempty"`
	ChangeDate	    time.Time	 `json:"change_date,omitempty"`
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

// Custom table name for GORM
func (Status) TableName() string {
	return config.DB_STATUS
}
