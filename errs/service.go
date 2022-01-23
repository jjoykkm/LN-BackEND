package errs

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Servicer interface {
	//20000
	ErrMsgBindData(c *gin.Context, err error)
	//40000
	ErrMsgUnAuth(c *gin.Context, err error)
	//60000
	ErrMsgCustom(c *gin.Context, err error, httpCode int)
	//80000
	ErrMsgInternal(c *gin.Context, err error)
}

type Service struct {
	repo Repositorier
}

func NewService(repo Repositorier) Servicer {
	return &Service{
		repo:  repo,
	}
}

func (s *Service) ErrMsgBindData(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, ErrContext{
		Code: ERROR_BIND_DATA,
		Err:  err,
		Msg:  err.Error(),
	})
}

func (s *Service) ErrMsgUnAuth(c *gin.Context, err error) {
	c.JSON(http.StatusUnauthorized, ErrContext{
		Code: ERROR_UNAUTH,
		Err:  err,
		Msg:  err.Error(),
	})
}

func (s *Service) ErrMsgInternal(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, ErrContext{
		Code: ERROR_INTERNAL,
		Err:  err,
		Msg:  err.Error(),
	})
}

func (s *Service) ErrMsgCustom(c *gin.Context, err error, httpCode int) {
	fmt.Println("ErrMsgCustom")
	fmt.Println(httpCode)
	fmt.Println(err)
	c.JSON(httpCode, ErrContext{
		Code: ERROR_UNKNOWN,
		Err:  err,
		Msg:  err.Error(),
	})
}
