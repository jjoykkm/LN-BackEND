package model_databases

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table session_action
//-------------------------------------------------------------------------------//
//model session_action
type SessionAction struct {
	SessionActionId     uuid.UUID	 `gorm:"session_action_id" json:"session_action_id,omitempty"`
	SessionAction     	string	 	 `gorm:"session_action" json:"session_action,omitempty"`
	CreateDate			time.Time	 `gorm:"create_date" json:"create_date,omitempty"`
	ChangeDate	    	time.Time	 `gorm:"change_date" json:"change_date,omitempty"`
	StatusId			uuid.UUID	 `gorm:"status_id" json:"status_id,omitempty"`
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
