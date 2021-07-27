package model_databases

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table role
//-------------------------------------------------------------------------------//
//model role
type Role struct {
	RoleId      	uuid.UUID	 `mapstructure:"role_id" json:"role_id,omitempty"`
	RoleName      	string		 `mapstructure:"role_name" json:"role_name,omitempty"`
	RoleDesc      	string		 `mapstructure:"role_desc" json:"role_desc,omitempty"`
	CreateDate		time.Time	 `mapstructure:"create_date" json:"create_date,omitempty"`
	ChangeDate	    time.Time	 `mapstructure:"change_date" json:"change_date,omitempty"`
	StatusId		uuid.UUID	 `mapstructure:"status_id" json:"status_id,omitempty"`
}
// New instance role
func (u *Role) New() *Role {
	return &Role{
		RoleId:			u.RoleId ,
		RoleName:		u.RoleName ,
		RoleDesc:		u.RoleDesc ,
		CreateDate:		u.CreateDate ,
		ChangeDate:		u.ChangeDate ,
		StatusId:		u.StatusId ,
	}
}

