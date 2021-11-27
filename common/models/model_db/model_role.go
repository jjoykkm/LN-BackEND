package model_db

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/common/config"
)

//-------------------------------------------------------------------------------//
//							Table role
//-------------------------------------------------------------------------------//
//model role
type Role struct {
	DBCommon
	RoleId      	uuid.UUID	 `json:"role_id"`
	RoleNameEN      string		 `json:"role_name_en"`
	RoleDescEN      string		 `json:"role_desc_en"`
	RoleNameTH      string		 `json:"role_name_th"`
	RoleDescTH      string		 `json:"role_desc_th"`
}
// New instance role
func (u *Role) New() *Role {
	return &Role{
		DBCommon:      	u.DBCommon ,
		RoleId:			u.RoleId ,
		RoleNameEN:		u.RoleNameEN ,
		RoleDescEN:		u.RoleDescEN ,
		RoleNameTH:		u.RoleNameTH ,
		RoleDescTH:		u.RoleDescTH ,
	}
}

// Custom table name for GORM
func (Role) TableName() string {
	return config.DB_ROLE
}

