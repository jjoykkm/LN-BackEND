package cm_other

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"reflect"
	"strings"
)

type Servicer interface {
	//// status, uid string
	//GetFarmList(status string, reqModel *model_other.ReqModel) (*model_other.RespModel, error)
	//// status, uid string
	//GetFarmAndFarmAreaList(status string, reqModel *model_other.ReqModel) (*model_other.RespModel, error)
}

type Service struct {
	repo Repositorier
}

func NewService(repo Repositorier) Servicer {
	return &Service{
		repo:  repo,
	}
}

//
//func (s *Service) GetFarmList(status string, reqModel *model_other.ReqModel) (*model_other.RespModel, error) {
//	myFarm, err := s.repo.FindAllMyFarm(status, reqModel.Uid)
//	if err != nil{
//		return nil, err
//	}
//	return &model_other.RespModel{
//		Item: myFarm,
//		Total: len(myFarm),
//	}, nil
//}
//
//func (s *Service) GetFarmAndFarmAreaList(status string, reqModel *model_other.ReqModel) (*model_other.RespModel, error) {
//	myFarm, err := s.repo.FindAllMyFarmAndFarmArea(status, reqModel.Uid)
//	if err != nil{
//		return nil, err
//	}
//	return &model_other.RespModel{
//		Item: myFarm,
//		Total: len(myFarm),
//	}, nil
//}

func RemoveDuplicateValues(listData interface{}) interface{} {
	source := reflect.ValueOf(listData)
	list := reflect.MakeSlice(source.Type(), 0, 0)
	visited := make(map[interface{}]struct{})
	for i := 0; i < source.Len(); i++ {
		value := source.Index(i)
		if _, ok := visited[value.Interface()]; ok {
			continue
		}
		visited[value.Interface()] = struct{}{}
		list = reflect.Append(list, value)
	}
	return list.Interface()
}

func ConvertListToStringIn(listData []string) string {
	result := "('" + strings.Join(listData, "','") + "')"
	return result
}
func ConvertUUIDtoStringMap(uuidList []uuid.UUID) map[string]bool {
	strMap := make(map[string]bool)
	for _, wa := range uuidList {
		strMap[wa.UUID.String()] = true
	}
	return strMap
}