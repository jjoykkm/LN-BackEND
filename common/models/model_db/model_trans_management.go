package model_db

import (
	"github.com/jjoykkm/ln-backend/common/config"
)

//-------------------------------------------------------------------------------//
//							Table trans_management
//-------------------------------------------------------------------------------//
//model trans_management
type TransManagement struct {
	DBCommonGet
	Uid      		string	 `json:"uid"`
	FarmId     		string	 `json:"farm_id"`
	RoleId      	string	 `json:"role_id"`
}
// New instance trans_management
func (u *TransManagement) New() *TransManagement {
	return &TransManagement{
		DBCommonGet:    u.DBCommonGet ,
		Uid:			u.Uid ,
		FarmId:			u.FarmId ,
		RoleId:			u.RoleId ,
	}
}

// Custom table name for GORM
func (TransManagement) TableName() string {
	return config.DB_TRANS_MANAGEMENT
}
