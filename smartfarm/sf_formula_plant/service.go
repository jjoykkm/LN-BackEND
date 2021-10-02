package sf_formula_plant

import (
	"github.com/jjoykkm/ln-backend/config"
	"github.com/jjoykkm/ln-backend/models/model_services"
	"github.com/mitchellh/mapstructure"
)

type Servicer interface {
	GetPlantCategoryList(status, language string) ([]model_services.ForPlantCatList, int)
}

type Service struct {
	repo  Repositorier
}

func NewService(repo Repositorier) Servicer {
	return &Service{
		repo:  repo,
	}
}

func (s *Service) GetPlantCategoryList(status, language string) ([]model_services.ForPlantCatList, int) {
	var catList model_services.ForPlantCatList
	var catListArray []model_services.ForPlantCatList

	plantTypeList, err := s.repo.FindAllPlantType(status)

	if err != nil{
		return nil, 0
	}

	for _, plantType := range plantTypeList {
		mapstructure.Decode(plantType, &catList)
		switch language {
		case config.LANGUAGE_EN:
			catList.PlantTypeName = plantType.PlantTypeEN
		case config.LANGUAGE_TH:
			catList.PlantTypeName = plantType.PlantTypeTH
		}
		catListArray = append(catListArray, catList)
	}

	return catListArray, len(catListArray)
}