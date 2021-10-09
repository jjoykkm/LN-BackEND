package helper

type Servicer interface {
	//GetModelFromBody(c *gin.Context) PostReqModel
}

type Service struct {
	repo  Repositorier
}

func NewService(repo Repositorier) Servicer {
	return &Service{
		repo:  repo,
	}
}

//func (s *Service) GetModelFromBody(c *gin.Context) helper.PostReqModel {
//	var ReqModel helper.PostReqModel
//
//	jsonData, err := c.GetRawData()
//	if err != nil {
//		c.JSON(http.StatusBadRequest, err)
//	}
//	json.Unmarshal([]byte(jsonData), &ReqModel)
//	//fmt.Printf("%+v/n", ReqModel)
//	return ReqModel
//}
