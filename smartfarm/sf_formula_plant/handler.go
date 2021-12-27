package sf_formula_plant

type Handler struct {
	service Servicer
}

func NewHandler(service Servicer) *Handler {
	return &Handler{service: service}
}

//func (h *Handler) GetPlantCategoryList(c *gin.Context) {
//	var reqModel model_other.ReqModel
//	reqModel.Language = c.DefaultQuery("lang", config.GetLanguage().Th)
//	if err := c.Bind(&reqModel); err != nil {
//		c.JSON(http.StatusBadRequest, &errs.ErrContext{
//			Code: "20000",
//			Err:  err,
//			Msg:  err.Error(),
//		})
//		return
//	}
//	respModel,err := h.service.GetPlantCategoryList(config.GetStatus().Active, &reqModel)
//	if err != nil {
//		if errx, ok := err.(*errs.ErrContext); ok {
//			if httpCode, ok := mapErrorCode[errx.Code]; ok {
//				c.JSON(httpCode, err)
//				return
//			}
//		}
//		c.JSON(http.StatusInternalServerError, &errs.ErrContext{
//			Code: "80000",
//			Err:  err,
//			Msg:  err.Error(),
//		})
//		return
//	}
//	c.JSON(http.StatusOK, respModel)
//}
//
//func (h *Handler) GetPlantCategoryItem(c *gin.Context) {
//	var reqModel model_other.ReqModel
//	reqModel.Language = c.DefaultQuery("lang", config.GetLanguage().Th)
//	if err := c.Bind(&reqModel); err != nil {
//		c.JSON(http.StatusBadRequest, &errs.ErrContext{
//			Code: "20000",
//			Err:  err,
//			Msg:  err.Error(),
//		})
//		return
//	}
//	respModel,err := h.service.GetPlantCategoryItem(config.GetStatus().Active, &reqModel)
//	if err != nil {
//		if errx, ok := err.(*errs.ErrContext); ok {
//			if httpCode, ok := mapErrorCode[errx.Code]; ok {
//				c.JSON(httpCode, err)
//				return
//			}
//		}
//		c.JSON(http.StatusInternalServerError, &errs.ErrContext{
//			Code: "80000",
//			Err:  err,
//			Msg:  err.Error(),
//		})
//		return
//	}
//	c.JSON(http.StatusOK, respModel)
//}
//
//func (h *Handler) GetPlantOverviewByPlant(c *gin.Context) {
//	var reqModel model_other.ReqModel
//	reqModel.Language = c.DefaultQuery("lang", config.GetLanguage().Th)
//	if err := c.Bind(&reqModel); err != nil {
//		c.JSON(http.StatusBadRequest, &errs.ErrContext{
//			Code: "20000",
//			Err:  err,
//			Msg:  err.Error(),
//		})
//		return
//	}
//	respModel,err := h.service.GetPlantOverviewByPlant(config.GetStatus().Active, &reqModel)
//	if err != nil {
//		if errx, ok := err.(*errs.ErrContext); ok {
//			if httpCode, ok := mapErrorCode[errx.Code]; ok {
//				c.JSON(httpCode, err)
//				return
//			}
//		}
//		c.JSON(http.StatusInternalServerError, &errs.ErrContext{
//			Code: "80000",
//			Err:  err,
//			Msg:  err.Error(),
//		})
//		return
//	}
//	c.JSON(http.StatusOK, respModel)
//}
//
//func (h *Handler) GetPlantOverviewFavorite(c *gin.Context) {
//	var reqModel model_other.ReqModel
//	reqModel.Language = c.DefaultQuery("lang", config.GetLanguage().Th)
//	if err := c.Bind(&reqModel); err != nil {
//		c.JSON(http.StatusBadRequest, &errs.ErrContext{
//			Code: "20000",
//			Err:  err,
//			Msg:  err.Error(),
//		})
//		return
//	}
//	respModel,err := h.service.GetPlantOverviewFavorite(config.GetStatus().Active, &reqModel)
//	if err != nil {
//		if errx, ok := err.(*errs.ErrContext); ok {
//			if httpCode, ok := mapErrorCode[errx.Code]; ok {
//				c.JSON(httpCode, err)
//				return
//			}
//		}
//		c.JSON(http.StatusInternalServerError, &errs.ErrContext{
//			Code: "80000",
//			Err:  err,
//			Msg:  err.Error(),
//		})
//		return
//	}
//	c.JSON(http.StatusOK, respModel)
//}
//
//func (h *Handler) GetPlantOfMine(c *gin.Context) {
//	var reqModel model_other.ReqModel
//	reqModel.Language = c.DefaultQuery("lang", config.GetLanguage().Th)
//	if err := c.Bind(&reqModel); err != nil {
//		c.JSON(http.StatusBadRequest, &errs.ErrContext{
//			Code: "20000",
//			Err:  err,
//			Msg:  err.Error(),
//		})
//		return
//	}
//	respModel,err := h.service.GetPlantOfMine(config.GetStatus().Active, &reqModel)
//	if err != nil {
//		if errx, ok := err.(*errs.ErrContext); ok {
//			if httpCode, ok := mapErrorCode[errx.Code]; ok {
//				c.JSON(httpCode, err)
//				return
//			}
//		}
//		c.JSON(http.StatusInternalServerError, &errs.ErrContext{
//			Code: "80000",
//			Err:  err,
//			Msg:  err.Error(),
//		})
//		return
//	}
//	c.JSON(http.StatusOK, respModel)
//}
//
//func (h *Handler) GetFormulaPlantDetail(c *gin.Context) {
//	var reqModel model_other.ReqModel
//	reqModel.Language = c.DefaultQuery("lang", config.GetLanguage().Th)
//	if err := c.Bind(&reqModel); err != nil {
//		c.JSON(http.StatusBadRequest, &errs.ErrContext{
//			Code: "20000",
//			Err:  err,
//			Msg:  err.Error(),
//		})
//		return
//	}
//	respModel,err := h.service.GetFormulaPlantDetail(config.GetStatus().Active, &reqModel)
//	if err != nil {
//		if errx, ok := err.(*errs.ErrContext); ok {
//			if httpCode, ok := mapErrorCode[errx.Code]; ok {
//				c.JSON(httpCode, err)
//				return
//			}
//		}
//		c.JSON(http.StatusInternalServerError, &errs.ErrContext{
//			Code: "80000",
//			Err:  err,
//			Msg:  err.Error(),
//		})
//		return
//	}
//	c.JSON(http.StatusOK, respModel)
//}
//
//
//func (h *Handler) AddChangeFormulaPlant(c *gin.Context) {
//	var reqModel ForPlantUS
//	//reqModel.Language = c.DefaultQuery("lang", config.GetLanguage().Th)
//
//	if err := c.Bind(&reqModel); err != nil {
//		c.JSON(http.StatusBadRequest, &errs.ErrContext{
//			Code: "20000",
//			Err:  err,
//			Msg:  err.Error(),
//		})
//		return
//	}
//	err := h.service.AddChangeFormulaPlant(&reqModel)
//	if err != nil {
//		if errx, ok := err.(*errs.ErrContext); ok {
//			if httpCode, ok := mapErrorCode[errx.Code]; ok {
//				c.JSON(httpCode, err)
//				return
//			}
//		}
//		c.JSON(http.StatusInternalServerError, &errs.ErrContext{
//			Code: "80000",
//			Err:  err,
//			Msg:  err.Error(),
//		})
//		return
//	}
//	c.JSON(http.StatusOK, &model_other.RespSuccessModel{
//		MsgTh: config.MSG_SUC_TH,
//		MsgEn: config.MSG_SUC_EN,
//	})
//}
//
//func (h *Handler) AddFavoriteForPlant(c *gin.Context) {
//	var reqModel model_db.FavForPlantUS
//	//reqModel.Language = c.DefaultQuery("lang", config.GetLanguage().Th)
//
//	if err := c.Bind(&reqModel); err != nil {
//		c.JSON(http.StatusBadRequest, &errs.ErrContext{
//			Code: "20000",
//			Err:  err,
//			Msg:  err.Error(),
//		})
//		return
//	}
//	err := h.service.AddFavoriteForPlant(&reqModel)
//	if err != nil {
//		if errx, ok := err.(*errs.ErrContext); ok {
//			if httpCode, ok := mapErrorCode[errx.Code]; ok {
//				c.JSON(httpCode, err)
//				return
//			}
//		}
//		c.JSON(http.StatusInternalServerError, &errs.ErrContext{
//			Code: "80000",
//			Err:  err,
//			Msg:  err.Error(),
//		})
//		return
//	}
//	c.JSON(http.StatusOK, &model_other.RespSuccessModel{
//		MsgTh: config.MSG_SUC_TH,
//		MsgEn: config.MSG_SUC_EN,
//	})
//}
//
//func (h *Handler) RemoveFavoriteForPlant(c *gin.Context) {
//	var reqModel model_db.FavForPlantUS
//	//reqModel.Language = c.DefaultQuery("lang", config.GetLanguage().Th)
//
//	if err := c.Bind(&reqModel); err != nil {
//		c.JSON(http.StatusBadRequest, &errs.ErrContext{
//			Code: "20000",
//			Err:  err,
//			Msg:  err.Error(),
//		})
//		return
//	}
//	err := h.service.RemoveFavoriteForPlant(&reqModel)
//	if err != nil {
//		if errx, ok := err.(*errs.ErrContext); ok {
//			if httpCode, ok := mapErrorCode[errx.Code]; ok {
//				c.JSON(httpCode, err)
//				return
//			}
//		}
//		c.JSON(http.StatusInternalServerError, &errs.ErrContext{
//			Code: "80000",
//			Err:  err,
//			Msg:  err.Error(),
//		})
//		return
//	}
//	c.JSON(http.StatusOK, &model_other.RespSuccessModel{
//		MsgTh: config.MSG_SUC_TH,
//		MsgEn: config.MSG_SUC_EN,
//	})
//}
//
//func (h *Handler) AddChangeFertilizer(c *gin.Context) {
//	var reqModel model_db.FertilizerUS
//	//reqModel.Language = c.DefaultQuery("lang", config.GetLanguage().Th)
//
//	if err := c.Bind(&reqModel); err != nil {
//		c.JSON(http.StatusBadRequest, &errs.ErrContext{
//			Code: "20000",
//			Err:  err,
//			Msg:  err.Error(),
//		})
//		return
//	}
//	err := h.service.AddChangeFertilizer(&reqModel)
//	if err != nil {
//		if errx, ok := err.(*errs.ErrContext); ok {
//			if httpCode, ok := mapErrorCode[errx.Code]; ok {
//				c.JSON(httpCode, err)
//				return
//			}
//		}
//		c.JSON(http.StatusInternalServerError, &errs.ErrContext{
//			Code: "80000",
//			Err:  err,
//			Msg:  err.Error(),
//		})
//		return
//	}
//	c.JSON(http.StatusOK, &model_other.RespSuccessModel{
//		MsgTh: config.MSG_SUC_TH,
//		MsgEn: config.MSG_SUC_EN,
//	})
//}