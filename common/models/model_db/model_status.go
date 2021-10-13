package model_db

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
	StatusId      	uuid.UUID	 `json:"status_id"`
	StatusName      string		 `json:"status_name"`
	CreateDate		time.Time	 `json:"create_date"`
	ChangeDate	    time.Time	 `json:"change_date"`
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
