package helper

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/jjoykkm/ln-backend/models/model_other"
	"net/http"
)

type Servicer interface {
	GetModelFromBody(c *gin.Context) model_other.PostBody
}

type Service struct {
	repo  Repositorier
}

func NewService(repo Repositorier) Servicer {
	return &Service{
		repo:  repo,
	}
}

func (s *Service) GetModelFromBody(c *gin.Context) model_other.PostBody {
	var bodyModel helper.PostBody

	jsonData, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	json.Unmarshal([]byte(jsonData), &bodyModel)
	//fmt.Printf("%+v/n", bodyModel)
	return bodyModel
}
