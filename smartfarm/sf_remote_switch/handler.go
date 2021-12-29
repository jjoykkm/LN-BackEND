package sf_remote_switch

import (
	//"github.com/gin-gonic/gin"
	//"github.com/jjoykkm/ln-backend/common/config"
	//"net/http"
)

type Handler struct {
	service Servicer
}

func NewHandler(service Servicer) *Handler {
	return &Handler{service: service}
}

