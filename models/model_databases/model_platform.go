package model_databases

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table platform
//-------------------------------------------------------------------------------//
//model platform
type Platform struct {
	PlatformId      uuid.UUID	 `gorm:"platform_id" json:"platform_id,omitempty"`
	PlatformName    string		 `gorm:"platform_name" json:"platform_name,omitempty"`
	CreateDate		time.Time	 `gorm:"create_date" json:"create_date,omitempty"`
	ChangeDate	    time.Time	 `gorm:"change_date" json:"change_date,omitempty"`
	StatusId		uuid.UUID	 `gorm:"status_id" json:"status_id,omitempty"`
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
