package helper

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetModelFromBody(c *gin.Context) PostBodyReq {
	var bodyReq PostBodyReq

	if err := c.Bind(&bodyReq); err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	//fmt.Printf("%+v/n", bodyReq)
	return bodyReq
}
