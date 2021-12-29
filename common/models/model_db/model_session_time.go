package model_db

import (
	"github.com/jjoykkm/ln-backend/common/config"
)

//-------------------------------------------------------------------------------//
//							Table session_time
//-------------------------------------------------------------------------------//
//model session_time
type SessionTime struct {
	DBCommonGet
	SessionTimeId      string	 	 `json:"session_time_id"`
	PublicToken        string	 	 `json:"public_token"`
	PrivateToken       string	 	 `json:"private_token"`
	BarrierToken       string	 	 `json:"barrier_token"`
	PlatformId	   	   string	 	 `json:"platform_id"`
	Uid	   			   string	 	 `json:"uid"`
	SessionActionId	   string	 	 `json:"session_action_id"`
}
// New instance session_time
func (u *SessionTime) New() *SessionTime {
	return &SessionTime{
		DBCommonGet:      		u.DBCommonGet ,
		SessionTimeId:		u.SessionTimeId ,
		PublicToken:		u.PublicToken ,
		PrivateToken:		u.PrivateToken ,
		BarrierToken:		u.BarrierToken ,
		PlatformId:			u.PlatformId ,
		Uid:				u.Uid ,
		SessionActionId:	u.SessionActionId ,
	}
}

// Custom table name for GORM
func (SessionTime) TableName() string {
	return config.DB_SESSION_TIME
}