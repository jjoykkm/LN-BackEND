package model_db

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/config"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table role
//-------------------------------------------------------------------------------//
//model role
type Role struct {
	RoleId      	uuid.UUID	 `mapstructure:"role_id" json:"role_id,omitempty"`
	RoleNameEN      string		 `mapstructure:"role_name_en" json:"role_name_en,omitempty"`
	RoleDescEN      string		 `mapstructure:"role_desc_en" json:"role_desc_en,omitempty"`
	CreateDate		time.Time	 `mapstructure:"create_date" json:"create_date,omitempty"`
	ChangeDate	    time.Time	 `mapstructure:"change_date" json:"change_date,omitempty"`
	StatusId		uuid.UUID	 `mapstructure:"status_id" json:"status_id,omitempty"`
	RoleNameTH      string		 `mapstructure:"role_name_th" json:"role_name_th,omitempty"`
	RoleDescTH      string		 `mapstructure:"role_desc_th" json:"role_desc_th,omitempty"`
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

