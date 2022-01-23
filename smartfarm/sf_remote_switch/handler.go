package sf_remote_switch

import (
	"github.com/gin-gonic/gin"
	"github.com/jjoykkm/ln-backend/common/config"
	"github.com/jjoykkm/ln-backend/common/models/model_other"
	"github.com/jjoykkm/ln-backend/errs"
	"github.com/jjoykkm/ln-backend/smartfarm/sf_common/cm_auth"
	"net/http"
)

type Handler struct {
	service Servicer
}

func NewHandler(service Servicer) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetRemoteSwitch(c *gin.Context) {
	reqModel := (&cm_auth.Service{}).PrepareData(c, c.Request.Header.Get("Bearer"))
	if reqModel == nil {
		return
	}
	respModel,err := h.service.GetRemoteSwitch(config.GetStatus().Active, reqModel)
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

func (h *Handler) ConfigRemoteSwitch(c *gin.Context) {
	var reqModel *RemoteDetailUS
	//reqModel.Language = c.DefaultQuery("lang", config.GetLanguage().Th)

	if err := c.Bind(&reqModel); err != nil {
		c.JSON(http.StatusBadRequest, &errs.ErrContext{
			Code: "20000",
			Err:  err,
			Msg:  err.Error(),
		})
		return
	}

	err := h.service.ConfigRemoteSwitch(reqModel)
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
	c.JSON(http.StatusOK, &model_other.RespSuccessModel{
		MsgTh: config.MSG_SUC_TH,
		MsgEn: config.MSG_SUC_EN,
	})
}

func (h *Handler) UnlinkSocketRemote(c *gin.Context) {
	var reqModel *RemoteDetailDel
	//reqModel.Language = c.DefaultQuery("lang", config.GetLanguage().Th)

	if err := c.Bind(&reqModel); err != nil {
		c.JSON(http.StatusBadRequest, &errs.ErrContext{
			Code: "20000",
			Err:  err,
			Msg:  err.Error(),
		})
		return
	}

	err := h.service.UnlinkSocketRemote(reqModel)
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
	c.JSON(http.StatusOK, &model_other.RespSuccessModel{
		MsgTh: config.MSG_SUC_TH,
		MsgEn: config.MSG_SUC_EN,
	})
}

func (h *Handler) RemoveRemoteSwitch(c *gin.Context) {
	var reqModel *RemoteDetailDel
	//reqModel.Language = c.DefaultQuery("lang", config.GetLanguage().Th)

	if err := c.Bind(&reqModel); err != nil {
		c.JSON(http.StatusBadRequest, &errs.ErrContext{
			Code: "20000",
			Err:  err,
			Msg:  err.Error(),
		})
		return
	}

	err := h.service.RemoveRemoteSwitch(reqModel)
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
	c.JSON(http.StatusOK, &model_other.RespSuccessModel{
		MsgTh: config.MSG_SUC_TH,
		MsgEn: config.MSG_SUC_EN,
	})
}

func (h *Handler) ChangeStatusSensor(c *gin.Context) {
	var reqModel *ControlSwitch
	//reqModel.Language = c.DefaultQuery("lang", config.GetLanguage().Th)
	if err := c.Bind(&reqModel); err != nil {
		c.JSON(http.StatusBadRequest, &errs.ErrContext{
			Code: "20000",
			Err:  err,
			Msg:  err.Error(),
		})
		return
	}
	respModel, err := h.service.ChangeStatusSensor(reqModel)
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