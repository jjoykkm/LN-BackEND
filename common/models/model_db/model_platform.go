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
	PlatformId      uuid.UUID	 `mapstructure:"platform_id" json:"platform_id,omitempty"`
	PlatformName    string		 `mapstructure:"platform_name" json:"platform_name,omitempty"`
	CreateDate		time.Time	 `mapstructure:"create_date" json:"create_date,omitempty"`
	ChangeDate	    time.Time	 `mapstructure:"change_date" json:"change_date,omitempty"`
	StatusId		uuid.UUID	 `mapstructure:"status_id" json:"status_id,omitempty"`
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
