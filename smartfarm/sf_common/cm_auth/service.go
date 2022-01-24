package cm_auth

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jjoykkm/ln-backend/common/config"
	"github.com/jjoykkm/ln-backend/common/models/model_other"
	"github.com/jjoykkm/ln-backend/errs"
	"io/ioutil"
	"net/http"
)

type Servicer interface {
	PrepareDataUser(c *gin.Context, token string) *model_other.ReqUser
	PrepareDataGet(c *gin.Context, token string) *model_other.ReqModel
	GetUserFromToken(token string) (interface{}, error)
	GetAuthorizeCheckForManageFarm(uid, farmId string) (bool, error)
}

type Service struct {
	repo  Repositorier
}

func NewService(repo Repositorier) Servicer {
	return &Service{
		repo:  repo,
	}
}
func (s *Service) PrepareDataUser(c *gin.Context, token string) *model_other.ReqUser {
	var model model_other.ReqUser

	result, err := s.GetUserFromToken(token)
	if err != nil{
		(&errs.Service{}).ErrMsgInternal(c, err)
		return nil
	}
	switch v := result.(type) {
	case ErrMsg:
		(&errs.Service{}).ErrMsgCustom(c, errors.New(v.Error.Message), v.StatusCode)
		return nil
	case UserAuth:
		//Assign uid and user no to model input
		model.Uid = v.User.Uid
		model.UserNo = v.User.UserNo
	}
	return &model
}

func (s *Service) PrepareDataGet(c *gin.Context, token string) *model_other.ReqModel {
	var model model_other.ReqModel
	//Bind data to model
	if err := c.Bind(&model); err != nil {
		(&errs.Service{}).ErrMsgBindData(c, err)
		return nil
	}
	//Get Language
	model.Language = c.DefaultQuery("lang", config.GetLanguage().Th)
	//Get User
	result := s.PrepareDataUser(c, token)
	if result == nil {
		return nil
	}
	model.User = *result
	return &model
}

func (s *Service) GetAuthorizeCheckForManageFarm(uid, farmId string) (bool, error) {
	authManage := false
	trans, err := s.repo.FindOneTransManagement(uid, farmId)
	if err != nil{
		return false, err
	}
	if trans.RoleId != config.GetRole().View {
		authManage = true
	}
	return authManage, err
}

func (s *Service) GetUserFromToken(token string) (interface{}, error) {
	var (
		modelErr  ErrMsg
		modelUser UserAuth
		)
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://103.212.181.187/api/userdata/me", nil)
	if err != nil {
		return nil, err
	}
	//Concat bearer token
	tokenStr := "Bearer" + " " + token
	req.Header.Add("Authorization", tokenStr)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	//If status code is 2XX
	if ( resp.StatusCode >= 200 && resp.StatusCode <= 299 ) {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		json.Unmarshal(bodyBytes, &modelUser)
		return modelUser, nil
	}
	//If status code doesn't 2XX
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(bodyBytes, &modelErr)
	modelErr.StatusCode = resp.StatusCode
	return modelErr, nil
}