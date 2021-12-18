package model_db

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/common/config"
)

//-------------------------------------------------------------------------------//
//							Table session_action
//-------------------------------------------------------------------------------//
//model session_action
type SessionAction struct {
	DBCommonGet
	SessionActionId     uuid.UUID	 `json:"session_action_id"`
	SessionAction     	string	 	 `json:"session_action"`
}
// New instance session_action
func (u *SessionAction) New() *SessionAction {
	return &SessionAction{
		DBCommonGet:      		u.DBCommonGet ,
		SessionActionId:	u.SessionActionId ,
		SessionAction:		u.SessionAction ,
	}
}

// Custom table name for GORM
func (SessionAction) TableName() string {
	return config.DB_SESSION_ACTION
}
