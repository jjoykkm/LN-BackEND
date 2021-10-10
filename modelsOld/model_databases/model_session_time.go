package model_databases

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/common/config"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table session_time
//-------------------------------------------------------------------------------//
//model session_time
type SessionTime struct {
	SessionTimeId      uuid.UUID	 `mapstructure:"session_time_id" json:"session_time_id,omitempty"`
	PublicToken        string	 	 `mapstructure:"public_token" json:"public_token,omitempty"`
	PrivateToken       string	 	 `mapstructure:"private_token" json:"private_token,omitempty"`
	BarrierToken       string	 	 `mapstructure:"barrier_token" json:"barrier_token,omitempty"`
	CreateDate		   time.Time	 `mapstructure:"create_date" json:"create_date,omitempty"`
	ChangeDate	       time.Time	 `mapstructure:"change_date" json:"change_date,omitempty"`
	StatusId		   uuid.UUID	 `mapstructure:"status_id" json:"status_id,omitempty"`
	PlatformId	   	   uuid.UUID	 `mapstructure:"platform_id" json:"platform_id,omitempty"`
	Uid	   			   uuid.UUID	 `mapstructure:"uid" json:"uid,omitempty"`
	SessionActionId	   uuid.UUID	 `mapstructure:"session_action_id" json:"session_action_id,omitempty"`
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

// Custom table name for GORM
func (SessionTime) TableName() string {
	return config.DB_SESSION_TIME
}