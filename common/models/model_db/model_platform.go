package model_db

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/common/config"
)

//-------------------------------------------------------------------------------//
//							Table platform
//-------------------------------------------------------------------------------//
//model platform
type Platform struct {
	DBCommon
	PlatformId      uuid.UUID	 `json:"platform_id"`
	PlatformName    string		 `json:"platform_name"`
}
// New instance platform
func (u *Platform) New() *Platform {
	return &Platform{
		DBCommon:      	u.DBCommon ,
		PlatformId:		u.PlatformId ,
		PlatformName:	u.PlatformName ,
	}
}

// Custom table name for GORM
func (Platform) TableName() string {
	return config.DB_PLATFORM
}
