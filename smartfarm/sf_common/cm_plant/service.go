package cm_plant

import (
	"github.com/jjoykkm/ln-backend/common/config"
	"github.com/jjoykkm/ln-backend/common/models/model_db"
	"github.com/jjoykkm/ln-backend/common/models/model_other"
)

type Servicer interface {
	// status
	GetFertAndCatList(status string) (*model_other.RespModel, error)
	GetSensorTypeList(status string) (*model_other.RespModel, error)
	GetFertCatList(status string) (*model_other.RespModel, error)
	AddChangeFertCat(req []model_db.FertilizerCatUS) error
}

type Service struct {
	repo  Repositorier
}

func NewService(repo Repositorier) Servicer {
	return &Service{
		repo:  repo,
	}
}

func (s *Service) GetFertAndCatList(status string) (*model_other.RespModel, error) {
	fert, err := s.repo.FindAllFertAndCat(status)
	if err != nil{
		return nil, err
	}
	return &model_other.RespModel{
		Item: fert,
		Total: len(fert),
	}, nil
}

func (s *Service) GetSensorTypeList(status string) (*model_other.RespModel, error) {
	fert, err := s.repo.FindAllSensorType(status)
	if err != nil{
		return nil, err
	}
	return &model_other.RespModel{
		Item: fert,
		Total: len(fert),
	}, nil
}

func (s *Service) GetFertCatList(status string) (*model_other.RespModel, error) {
	fert, err := s.repo.FindAllFertCatList(status)
	if err != nil{
		return nil, err
	}
	return &model_other.RespModel{
		Item: fert,
		Total: len(fert),
	}, nil
}


//-------------------------------------------------------------------------------//
//									Upsert
//-------------------------------------------------------------------------------//
func (s *Service) AddChangeFertCat(req []model_db.FertilizerCatUS) error {
	for idx, wa := range req {
		wa.StatusId = config.GetStatus().Active
		req[idx] = wa
	}
	err := s.repo.UpsertFertCat(req)
	if err != nil{
		return err
	}
	return nil
}