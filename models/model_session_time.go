package models

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table session_time
//-------------------------------------------------------------------------------//
//model session_time
type SessionTime struct {
	SessionTimeId      uuid.UUID	 `gorm:"session_time_id" json:"session_time_id,omitempty"`
	PublicToken        string	 	 `gorm:"public_token" json:"public_token,omitempty"`
	PrivateToken       string	 	 `gorm:"private_token" json:"private_token,omitempty"`
	BarrierToken       string	 	 `gorm:"barrier_token" json:"barrier_token,omitempty"`
	CreateDate		   time.Time	 `gorm:"create_date" json:"create_date,omitempty"`
	ChangeDate	       time.Time	 `gorm:"change_date" json:"change_date,omitempty"`
	StatusId		   uuid.UUID	 `gorm:"status_id" json:"status_id,omitempty"`
	PlatformId	   	   uuid.UUID	 `gorm:"platform_id" json:"platform_id,omitempty"`
	Uid	   			   uuid.UUID	 `gorm:"uid" json:"uid,omitempty"`
	SessionActionId	   uuid.UUID	 `gorm:"session_action_id" json:"session_action_id,omitempty"`
}
// New instance session_time
func (u *SessionTime) New() *SessionTime {
	return &SessionTime{
		SessionTimeId:		u.SessionTimeId ,
		PublicToken:		u.PublicToken ,
		PrivateToken:		u.PrivateToken ,
		BarrierToken:		u.BarrierToken ,
		CreateDate:			u.CreateDate ,
		ChangeDate:			u.ChangeDate ,
		StatusId:			u.StatusId ,
		PlatformId:			u.PlatformId ,
		Uid:				u.Uid ,
		SessionActionId:	u.SessionActionId ,
	}
}