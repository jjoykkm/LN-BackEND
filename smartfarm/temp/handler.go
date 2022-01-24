package sf_dashboard

import (
	"github.com/gin-gonic/gin"
	"github.com/jjoykkm/ln-backend/common/config"
	"github.com/jjoykkm/ln-backend/common/models/model_other"
	"github.com/jjoykkm/ln-backend/errs"
	"net/http"
)

type Handler struct {
	service Servicer
}

func NewHandler(service Servicer) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetFarmList(c *gin.Context) {
	reqModel := (&cm_auth.Service{}).PrepareDataGet(c, c.Request.Header.Get("Bearer"))
	if reqModel == nil {
		return
	}
	respModel,err := h.service.GetFarmList(config.GetStatus().Active, reqModel)
	if err != nil {
		if errx, ok := err.(*errs.ErrContext); ok {
			if httpCode, ok := mapErrorCode[errx.Code]; ok {
				c.JSON(httpCode, err)
				return
			}
		}
		(&errs.Service{}).ErrMsgInternal(c, err)
		return
	}
	c.JSON(http.StatusOK, respModel)
}


func (h *Handler) GetFarmAreaDashboardList(c *gin.Context) {
	reqModel := (&cm_auth.Service{}).PrepareDataGet(c, c.Request.Header.Get("Bearer"))
	if reqModel == nil {
		return
	}
	respModel,err := h.service.GetFarmAreaDashboardList(config.GetStatus().Active, reqModel)
	if err != nil {
		if errx, ok := err.(*errs.ErrContext); ok {
			if httpCode, ok := mapErrorCode[errx.Code]; ok {
				c.JSON(httpCode, err)
				return
			}
		}
		(&errs.Service{}).ErrMsgInternal(c, err)
		return
	}
	c.JSON(http.StatusOK, respModel)
}