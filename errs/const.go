package errs

import "net/http"

const (
	keyCacheThaiHoliday string = "THAIHOLIDAY"
	ERROR_4002005       string = "4002005"
	ERROR_4001001       string = "4001001"
)

const (
	ERROR_BIND_DATA       string = "20000"
	ERROR_UNAUTH          string = "40000"
	ERROR_UNKNOWN         string = "60000"
	ERROR_INTERNAL        string = "80000"
)

var mapErrorCode = map[string]int{
		ERROR_4002005: http.StatusConflict,
		ERROR_4001001: http.StatusConflict,
	}
