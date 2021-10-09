package errs

//-------------------------------------------------------------------------------//
//				 		    Error Context
//-------------------------------------------------------------------------------//
//Model post body
type ErrContext struct {
	Code	string	`mapstructure:"code" json:"code"`
	Err		error	`mapstructure:"err" json:"err"`
}

// New instance
func (u *ErrContext) New() *ErrContext {
	return &ErrContext{
		Code:	u.Code ,
		Err:	u.Err ,
	}
}

// For Assertions
func (r *ErrContext) Error() string {
	return r.Err.Error()
}
