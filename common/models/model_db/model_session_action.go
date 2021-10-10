package model_db

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/config"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table session_action
//-------------------------------------------------------------------------------//
//model session_action
type SessionAction struct {
	SessionActionId     uuid.UUID	 `mapstructure:"session_action_id" json:"session_action_id,omitempty"`
	SessionAction     	string	 	 `mapstructure:"session_action" json:"session_action,omitempty"`
	CreateDate			time.Time	 `mapstructure:"create_date" json:"create_date,omitempty"`
	ChangeDate	    	time.Time	 `mapstructure:"change_date" json:"change_date,omitempty"`
	StatusId			uuid.UUID	 `mapstructure:"status_id" json:"status_id,omitempty"`
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
