package sf_schedule_reminder

import (
	"github.com/gin-gonic/gin"
	"github.com/jjoykkm/ln-backend/common/config"
	"github.com/jjoykkm/ln-backend/obsolete_utility"
	"net/http"
)

type Handler struct {
	service Servicer
}

func NewHandler(service Servicer) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetPlantCategoryList(c *gin.Context) {
	bodyModel := obsolete_utility.GetModelFromBody(c)
	// GetPlantCategoryList(status, language string) ([]Obsolete_model_services.ForPlantCatList, int)
	plantCategoryList, total := h.service.GetPlantCategoryList(config.GetStatus().Active, bodyModel.Language)
	if total == 0 {
		c.JSON(http.StatusNoContent, gin.H{})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"item":  plantCategoryList,
			"total": total,
		})
	}
}
