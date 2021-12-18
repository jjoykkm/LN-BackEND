package model_db

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/common/config"
)

//-------------------------------------------------------------------------------//
//							Table session_time
//-------------------------------------------------------------------------------//
//model session_time
type SessionTime struct {
	DBCommonGet
	SessionTimeId      uuid.UUID	 `json:"session_time_id"`
	PublicToken        string	 	 `json:"public_token"`
	PrivateToken       string	 	 `json:"private_token"`
	BarrierToken       string	 	 `json:"barrier_token"`
	PlatformId	   	   uuid.UUID	 `json:"platform_id"`
	Uid	   			   uuid.UUID	 `json:"uid"`
	SessionActionId	   uuid.UUID	 `json:"session_action_id"`
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