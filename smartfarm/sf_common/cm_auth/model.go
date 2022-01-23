package cm_auth

import (
	"github.com/jjoykkm/ln-backend/common/models/model_db"
	"github.com/jjoykkm/ln-backend/errs"
)

//-------------------------------------------------------------------------------//
//				 	    Error Message
//-------------------------------------------------------------------------------//
type Error struct {
	Type    	string		`json:"type"`
	Message     string		`json:"message"`
}

type ErrMsg struct {
	StatusCode    	int		`json:"status_code"`
	Error 			Error 	`json:"error"`
}

type ErrMsgWithCode struct {
	StatusCode    	int					`json:"status_code"`
	Message     	errs.ErrContext		`json:"message"`
}

type UserAuth struct {
	User  	model_db.UsersBank	`json:"user"`
}

