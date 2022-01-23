package cm_plant

import (
	"github.com/gin-gonic/gin"
	"github.com/jjoykkm/ln-backend/common/config"
	"github.com/jjoykkm/ln-backend/common/models/model_db"
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

func (h *Handler) GetFertAndCatList(c *gin.Context) {
	reqModel := (&cm_auth.Service{}).PrepareData(c, c.Request.Header.Get("Bearer"))
	if reqModel == nil {
		return
	}
	respModel,err := h.service.GetFertAndCatList(config.GetStatus().Active)
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

func (h *Handler) GetSensorTypeList(c *gin.Context) {
	reqModel := (&cm_auth.Service{}).PrepareData(c, c.Request.Header.Get("Bearer"))
	if reqModel == nil {
		return
	}
	respModel,err := h.service.GetSensorTypeList(config.GetStatus().Active)
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

func (h *Handler) GetFertCatList(c *gin.Context) {
	reqModel := (&cm_auth.Service{}).PrepareData(c, c.Request.Header.Get("Bearer"))
	if reqModel == nil {
		return
	}
	respModel,err := h.service.GetFertCatList(config.GetStatus().Active)
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

func (h *Handler) AddChangeFertCat(c *gin.Context) {
	var reqModel []model_db.FertilizerCatUS
	//reqModel.Language = c.DefaultQuery("lang", config.GetLanguage().Th)

	if err := c.Bind(&reqModel); err != nil {
		c.JSON(http.StatusBadRequest, &errs.ErrContext{
			Code: "20000",
			Err:  err,
			Msg:  err.Error(),
		})
		return
	}
	err := h.service.AddChangeFertCat(reqModel)
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