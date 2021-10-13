package model_databases

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/common/config"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table role
//-------------------------------------------------------------------------------//
//model role
type Role struct {
	RoleId      	uuid.UUID	 `json:"role_id"`
	RoleNameEN      string		 `json:"role_name_en"`
	RoleDescEN      string		 `json:"role_desc_en"`
	CreateDate		time.Time	 `json:"create_date"`
	ChangeDate	    time.Time	 `json:"change_date"`
	StatusId		uuid.UUID	 `json:"status_id"`
	RoleNameTH      string		 `json:"role_name_th"`
	RoleDescTH      string		 `json:"role_desc_th"`
}
// New instance role
func (u *Role) New() *Role {
	return &Role{
		RoleId:			u.RoleId ,
		RoleNameEN:		u.RoleNameEN ,
		RoleDescEN:		u.RoleDescEN ,
		CreateDate:		u.CreateDate ,
		ChangeDate:		u.ChangeDate ,
		StatusId:		u.StatusId ,
		RoleNameTH:		u.RoleNameTH ,
		RoleDescTH:		u.RoleDescTH ,
	}
}

// Custom table name for GORM
func (Role) TableName() string {
	return config.DB_ROLE
}

