package helper_gorm

import (
	"github.com/jjoykkm/ln-backend/common/config"
	"github.com/jjoykkm/ln-backend/modelsOld/model_databases"
)

type Servicer interface {
	GetFarmAreaByFarm(status, farmId, resultType string) ([]model_databases.FarmArea, []string)
}

type Service struct {
	repo Repositorier
}

func NewService(repo Repositorier) Servicer {
	return &Service{
		repo:  repo,
	}
}
func (s *Service) GetFarmAreaByFarm(status, farmId, resultType string) ([]model_databases.FarmArea, []string) {
	var farmAreaIdList []string
	farmModel, err := s.repo.FindAllFarmAreaByFarm(status, farmId)
	if err != nil {
		return nil, nil
	}
	if resultType != config.GetResType().Struct {
		for _, array := range farmModel {
			farmAreaIdList = append(farmAreaIdList, array.FarmAreaId.UUID.String())
		}
	}
	return farmModel, farmAreaIdList
}