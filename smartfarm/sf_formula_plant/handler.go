package sf_formula_plant

import (
	"github.com/gin-gonic/gin"
	"github.com/jjoykkm/ln-backend/config"
	"github.com/jjoykkm/ln-backend/utility"
	"net/http"
)

type Handler struct {
	service Servicer
}

func NewHandler(service Servicer) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetPlantCategoryList(c *gin.Context) {
	bodyModel := utility.GetModelFromBody(c)
	// GetPlantCategoryList(status, language string) ([]model_services.ForPlantCatList, int)
	plantCategoryList, total := h.service.GetPlantCategoryList(config.STATUS_ACTIVE, bodyModel.Language)
	if total == 0 {
		c.JSON(http.StatusNoContent, gin.H{})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"item":  plantCategoryList,
			"total": total,
		})
	}
}
