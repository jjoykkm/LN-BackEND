package model_db

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/common/config"
)

//-------------------------------------------------------------------------------//
//							Table trans_management
//-------------------------------------------------------------------------------//
//model trans_management
type TransManagement struct {
	DBCommon
	Uid      		uuid.UUID	 `json:"uid"`
	FarmId     		uuid.UUID	 `json:"farm_id"`
	RoleId      	uuid.UUID	 `json:"role_id"`
}
// New instance trans_management
func (u *TransManagement) New() *TransManagement {
	return &TransManagement{
		DBCommon:      		u.DBCommon ,
		Uid:			u.Uid ,
		FarmId:			u.FarmId ,
		RoleId:			u.RoleId ,
	}
}

// Custom table name for GORM
func (TransManagement) TableName() string {
	return config.DB_TRANS_MANAGEMENT
}
