package sf_my_farm

import "net/http"

const LIMIT_GET_DATA = 100

const MSG_NO_AUTH = "You have unauthorized to change this setting."

const (
	ERROR_4002005       string = "4002005"
	ERROR_4001001       string = "4001001"
)

var (
	mapErrorCode = map[string]int{
		ERROR_4002005: http.StatusUnauthorized,
		//ERROR_4001001: http.StatusConflict,
	}
)