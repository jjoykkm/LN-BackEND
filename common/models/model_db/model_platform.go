package model_db

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/common/config"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table platform
//-------------------------------------------------------------------------------//
//model platform
type Platform struct {
	PlatformId      uuid.UUID	 `json:"platform_id"`
	PlatformName    string		 `json:"platform_name"`
	CreateDate		time.Time	 `json:"create_date"`
	ChangeDate	    time.Time	 `json:"change_date"`
	StatusId		uuid.UUID	 `json:"status_id"`
}
// New instance platform
func (u *Platform) New() *Platform {
	return &Platform{
		PlatformId:		u.PlatformId ,
		PlatformName:	u.PlatformName ,
		CreateDate:		u.CreateDate ,
		ChangeDate:		u.ChangeDate ,
		StatusId:		u.StatusId ,
	}
}

// Custom table name for GORM
func (Platform) TableName() string {
	return config.DB_PLATFORM
}
