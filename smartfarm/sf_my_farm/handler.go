package sf_my_farm

import (
	"github.com/gin-gonic/gin"
	"github.com/jjoykkm/ln-backend/common/config"
	"github.com/jjoykkm/ln-backend/common/models/model_db"
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

func (h *Handler) GetOverviewFarm(c *gin.Context) {
	var reqModel model_other.ReqModel
	reqModel.Language = c.DefaultQuery("lang", config.GetLanguage().Th)
	if err := c.Bind(&reqModel); err != nil {
		c.JSON(http.StatusBadRequest, &errs.ErrContext{
			Code: "20000",
			Err:  err,
		})
		return
	}
	respModel,err := h.service.GetOverviewFarm(config.GetStatus().Active, &reqModel)
	if err != nil {
		if errx, ok := err.(*errs.ErrContext); ok {
			if httpCode, ok := mapErrorCode[errx.Code]; ok {
				c.JSON(httpCode, err)
				return
			}
		}
		c.JSON(http.StatusInternalServerError, &errs.ErrContext{
			Code: "80000",
			Err:  err,
		})
		return
	}
	c.JSON(http.StatusOK, respModel)
}

func (h *Handler) GetManageRole(c *gin.Context) {
	var reqModel model_other.ReqModel
	reqModel.Language = c.DefaultQuery("lang", config.GetLanguage().Th)
	if err := c.Bind(&reqModel); err != nil {
		c.JSON(http.StatusBadRequest, &errs.ErrContext{
			Code: "20000",
			Err:  err,
		})
		return
	}
	respModel,err := h.service.GetManageRole(config.GetStatus().Active, &reqModel)
	if err != nil {
		if errx, ok := err.(*errs.ErrContext); ok {
			if httpCode, ok := mapErrorCode[errx.Code]; ok {
				c.JSON(httpCode, err)
				return
			}
		}
		c.JSON(http.StatusInternalServerError, &errs.ErrContext{
			Code: "80000",
			Err:  err,
		})
		return
	}
	c.JSON(http.StatusOK, respModel)
}

func (h *Handler) GetManageFarmArea(c *gin.Context) {
	var reqModel model_other.ReqModel
	reqModel.Language = c.DefaultQuery("lang", config.GetLanguage().Th)
	if err := c.Bind(&reqModel); err != nil {
		c.JSON(http.StatusBadRequest, &errs.ErrContext{
			Code: "20000",
			Err:  err,
		})
		return
	}
	respModel,err := h.service.GetManageFarmArea(config.GetStatus().Active, &reqModel)
	if err != nil {
		if errx, ok := err.(*errs.ErrContext); ok {
			if httpCode, ok := mapErrorCode[errx.Code]; ok {
				c.JSON(httpCode, err)
				return
			}
		}
		c.JSON(http.StatusInternalServerError, &errs.ErrContext{
			Code: "80000",
			Err:  err,
		})
		return
	}
	c.JSON(http.StatusOK, respModel)
}

func (h *Handler) GetManageMainbox(c *gin.Context) {
	var reqModel model_other.ReqModel
	reqModel.Language = c.DefaultQuery("lang", config.GetLanguage().Th)
	if err := c.Bind(&reqModel); err != nil {
		c.JSON(http.StatusBadRequest, &errs.ErrContext{
			Code: "20000",
			Err:  err,
		})
		return
	}
	respModel,err := h.service.GetManageMainbox(config.GetStatus().Active, &reqModel)
	if err != nil {
		if errx, ok := err.(*errs.ErrContext); ok {
			if httpCode, ok := mapErrorCode[errx.Code]; ok {
				c.JSON(httpCode, err)
				return
			}
		}
		c.JSON(http.StatusInternalServerError, &errs.ErrContext{
			Code: "80000",
			Err:  err,
		})
		return
	}
	c.JSON(http.StatusOK, respModel)
}

func (h *Handler) ActivateMainbox(c *gin.Context) {
	var reqModel model_db.MainboxSerialUS
	//reqModel.Language = c.DefaultQuery("lang", config.GetLanguage().Th)
	if err := c.Bind(&reqModel); err != nil {
		c.JSON(http.StatusBadRequest, &errs.ErrContext{
			Code: "20000",
			Err:  err,
		})
		return
	}
	err := h.service.ActivateMainbox(&reqModel)
	if err != nil {
		if errx, ok := err.(*errs.ErrContext); ok {
			if httpCode, ok := mapErrorCode[errx.Code]; ok {
				c.JSON(httpCode, err)
				return
			}
		}
		c.JSON(http.StatusInternalServerError, &errs.ErrContext{
			Code: "80000",
			Err:  err,
		})
		return
	}
	c.JSON(http.StatusOK, "success")
}
func (h *Handler) ConfigMainbox(c *gin.Context) {
	var reqModel ReqConfMainbox
	//reqModel.Language = c.DefaultQuery("lang", config.GetLanguage().Th)
	if err := c.Bind(&reqModel); err != nil {
		c.JSON(http.StatusBadRequest, &errs.ErrContext{
			Code: "20000",
			Err:  err,
		})
		return
	}
	err := h.service.ConfigMainbox(&reqModel)
	if err != nil {
		if errx, ok := err.(*errs.ErrContext); ok {
			if httpCode, ok := mapErrorCode[errx.Code]; ok {
				c.JSON(httpCode, err)
				return
			}
		}
		c.JSON(http.StatusInternalServerError, &errs.ErrContext{
			Code: "80000",
			Err:  err,
		})
		return
	}
	c.JSON(http.StatusOK, "success")
}
