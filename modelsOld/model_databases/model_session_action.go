package model_databases

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/common/config"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table session_action
//-------------------------------------------------------------------------------//
//model session_action
type SessionAction struct {
	SessionActionId     uuid.UUID	 `json:"session_action_id"`
	SessionAction     	string	 	 `json:"session_action"`
	CreateDate			time.Time	 `json:"create_date"`
	ChangeDate	    	time.Time	 `json:"change_date"`
	StatusId			uuid.UUID	 `json:"status_id"`
}
// New instance session_action
func (u *SessionAction) New() *SessionAction {
	return &SessionAction{
		SessionActionId:	u.SessionActionId ,
		SessionAction:		u.SessionAction ,
		CreateDate:			u.CreateDate ,
		ChangeDate:			u.ChangeDate ,
		StatusId:			u.StatusId ,
	}
}

// Custom table name for GORM
func (SessionAction) TableName() string {
	return config.DB_SESSION_ACTION
}
