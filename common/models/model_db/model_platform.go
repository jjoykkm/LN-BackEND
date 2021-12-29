package model_db

import (
	"github.com/jjoykkm/ln-backend/common/config"
)

//-------------------------------------------------------------------------------//
//							Table platform
//-------------------------------------------------------------------------------//
//model platform
type Platform struct {
	DBCommonGet
	PlatformId      string	 `json:"platform_id"`
	PlatformName    string	 `json:"platform_name"`
}
// New instance platform
func (u *Platform) New() *Platform {
	return &Platform{
		DBCommonGet:      	u.DBCommonGet ,
		PlatformId:		u.PlatformId ,
		PlatformName:	u.PlatformName ,
	}
}

// Custom table name for GORM
func (Platform) TableName() string {
	return config.DB_PLATFORM
}
