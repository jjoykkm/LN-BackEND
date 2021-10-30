package errs

//-------------------------------------------------------------------------------//
//				 		    Error Context
//-------------------------------------------------------------------------------//
//Model post body
type ErrContext struct {
	Code	string	`json:"code"`
	Err		error	`json:"-"`
	Msg		string	`json:"msg"`
}

// New instance
func (u *ErrContext) New() *ErrContext {
	return &ErrContext{
		Code:	u.Code ,
		Err:	u.Err ,
		Msg:	u.Msg ,
	}
}

// For Assertions
func (u *ErrContext) Error() string {
	return u.Err.Error()
}
