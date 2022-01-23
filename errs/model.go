package errs

//-------------------------------------------------------------------------------//
//				 		    Error Context
//-------------------------------------------------------------------------------//
//Model post body
type ErrContext struct {
	Code	string	`json:"code"`
	Err		error	`json:"-"`
	Type	string	`json:"type"`
	Msg		string	`json:"error"`
}

// New instance
func (u *ErrContext) New() *ErrContext {
	return &ErrContext{
		Code:	u.Code ,
		Err:	u.Err ,
		Type:	u.Type ,
		Msg:	u.Msg ,
	}
}

// For Assertions
func (u *ErrContext) Error() string {
	return u.Err.Error()
}
