package sf_remote_switch

import "net/http"

const LIMIT_GET_DATA = 100

const (
	MSG_NO_AUTH = "You have unauthorized to change this setting."
	MSG_DUP_MB = "This Serial Number has been Activated."
	MSG_WRONG_MB = "Wrong Serial Number."
)

const (
	ERROR_4002005       string = "4002005"
	ERROR_4001001       string = "4001001"
	ERROR_4001002       string = "4001002"
)

var (
	mapErrorCode = map[string]int{
		ERROR_4002005: http.StatusUnauthorized,
		ERROR_4001001: http.StatusConflict,
		ERROR_4001002: http.StatusConflict,
	}
)