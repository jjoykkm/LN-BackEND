package helper

type Servicer interface {
	//GetModelFromBody(c *gin.Context) PostBodyReq
}

type Service struct {
	repo  Repositorier
}

func NewService(repo Repositorier) Servicer {
	return &Service{
		repo:  repo,
	}
}

//func (s *Service) GetModelFromBody(c *gin.Context) helper.PostBodyReq {
//	var bodyReq helper.PostBodyReq
//
//	jsonData, err := c.GetRawData()
//	if err != nil {
//		c.JSON(http.StatusBadRequest, err)
//	}
//	json.Unmarshal([]byte(jsonData), &bodyReq)
//	//fmt.Printf("%+v/n", bodyReq)
//	return bodyReq
//}
