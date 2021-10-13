package sf_my_farm

type Handler struct {
	service Servicer
}

func NewHandler(service Servicer) *Handler {
	return &Handler{service: service}
}
//
//func (h *Handler) GetPlantCategoryList(c *gin.Context) {
//	bodyModel := utility.GetModelFromBody(c)
//	// GetPlantCategoryList(status, language string) ([]model_services.ForPlantCatList, int)
//	plantCategoryList, total := h.service.GetPlantCategoryList(config.GetStatus().Active, bodyModel.Language)
//	if total == 0 {
//		c.JSON(http.StatusNoContent, gin.H{})
//	} else {
//		c.JSON(http.StatusOK, gin.H{
//			"item":  plantCategoryList,
//			"total": total,
//		})
//	}
//}
