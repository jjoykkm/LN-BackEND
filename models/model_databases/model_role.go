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
	RoleId      	uuid.UUID	 `gorm:"role_id" json:"role_id,omitempty"`
	RoleName      	string		 `gorm:"role_name" json:"role_name,omitempty"`
	RoleDesc      	string		 `gorm:"role_desc" json:"role_desc,omitempty"`
	CreateDate		time.Time	 `gorm:"create_date" json:"create_date,omitempty"`
	ChangeDate	    time.Time	 `gorm:"change_date" json:"change_date,omitempty"`
	StatusId		uuid.UUID	 `gorm:"status_id" json:"status_id,omitempty"`
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

