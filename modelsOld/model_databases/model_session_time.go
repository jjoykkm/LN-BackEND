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
	SessionTimeId      uuid.UUID	 `json:"session_time_id"`
	PublicToken        string	 	 `json:"public_token"`
	PrivateToken       string	 	 `json:"private_token"`
	BarrierToken       string	 	 `json:"barrier_token"`
	CreateDate		   time.Time	 `json:"create_date"`
	ChangeDate	       time.Time	 `json:"change_date"`
	StatusId		   uuid.UUID	 `json:"status_id"`
	PlatformId	   	   uuid.UUID	 `json:"platform_id"`
	Uid	   			   uuid.UUID	 `json:"uid"`
	SessionActionId	   uuid.UUID	 `json:"session_action_id"`
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