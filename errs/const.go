package errs

import "net/http"

const (
	keyCacheThaiHoliday string = "THAIHOLIDAY"
	ERROR_4002005       string = "4002005"
	ERROR_4001001       string = "4001001"
)

var (
	mapErrorCode = map[string]int{
		ERROR_4002005: http.StatusConflict,
		ERROR_4001001: http.StatusConflict,
	}
)
