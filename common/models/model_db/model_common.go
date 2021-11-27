package model_db

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table Common Database
//-------------------------------------------------------------------------------//
//model Common Database
type DBCommon struct {
	CreateDate		time.Time	 `json:"create_date"`
	//CreateBy		string	 	 `json:"create_by"`
	ChangeDate	   	time.Time	 `json:"change_date"`
	//ChangeBy	    string	 	 `json:"change_by"`
	StatusId		uuid.UUID	 `json:"status_id"`
}
// New instance Common Database
func (u *DBCommon) New() *DBCommon {
	return &DBCommon{
		CreateDate:		u.CreateDate ,
		//CreateBy:		u.CreateBy ,
		ChangeDate:		u.ChangeDate ,
		//ChangeBy:		u.ChangeBy ,
		StatusId:		u.StatusId ,
	}
}
