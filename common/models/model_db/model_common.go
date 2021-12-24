package model_db

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"time"
)

//-------------------------------------------------------------------------------//
//						Table Common Database (Get)
//-------------------------------------------------------------------------------//
//model Common Database (Get)
type DBCommonGet struct {
	CreateDate		time.Time	 `json:"create_date"`
	//CreateBy		string	 	 `json:"create_by"`
	ChangeDate	   	time.Time	 `json:"change_date"`
	//ChangeBy	    string	 	 `json:"change_by"`
	StatusId		uuid.UUID	 `json:"status_id"`
}
// New instance Common Database
func (u *DBCommonGet) New() *DBCommonGet {
	return &DBCommonGet{
		CreateDate:		u.CreateDate ,
		//CreateBy:		u.CreateBy ,
		ChangeDate:		u.ChangeDate ,
		//ChangeBy:		u.ChangeBy ,
		StatusId:		u.StatusId ,
	}
}

//-------------------------------------------------------------------------------//
//						Table Common Database (Create, Update)
//-------------------------------------------------------------------------------//
//model Common Database (Create, Update)
type DBCommonCreateUpdate struct {
	CreateDate		time.Time	 `json:"create_date" gorm:"autoCreateTime"`
	//CreateBy		*string	 	 `json:"create_by"`
	ChangeDate	   	time.Time	 `json:"change_date" gorm:"autoUpdateTime"`
	//ChangeBy	    string	 	 `json:"change_by"`
	//StatusId		string		 `json:"status_id"`
}
