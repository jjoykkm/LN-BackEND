package sf_formula_plant

import (
	"github.com/gin-gonic/gin"
	"github.com/jjoykkm/ln-backend/common/models/model_other"
	"github.com/jjoykkm/ln-backend/config"
	"net/http"
)

type Handler struct {
	service Servicer
}

func NewHandler(service Servicer) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetPlantCategoryList(c *gin.Context) {
	var bodyReq model_other.BodyReq
	if err := c.Bind(&bodyReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	// GetPlantCategoryList(status, language string) ([]model_services.ForPlantCatList, int)
	bodyResp := h.service.GetPlantCategoryList(config.STATUS_ACTIVE, bodyReq.Language)
	if bodyResp.Total < 1 {
		c.JSON(http.StatusNoContent, gin.H{})
	} else {
		c.JSON(http.StatusOK, bodyResp)
	}
}

func (h *Handler) GetPlantCategoryItem(c *gin.Context) {
	var bodyReq model_other.BodyReq
	if err := c.Bind(&bodyReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	// GetPlantCategoryItem(status, plantTypeId, language string, offset int)
	bodyResp := h.service.GetPlantCategoryItem(config.STATUS_ACTIVE, bodyReq.PlantTypeId, bodyReq.Language, bodyReq.Offset)
	if bodyResp.Total < 1 {
		c.JSON(http.StatusNoContent, gin.H{})
	} else {
		c.JSON(http.StatusOK, bodyResp)
	}
}
