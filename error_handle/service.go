package error_handle

type Servicer interface {
	//GetFarmAreaAndName(farmAreaId string) (*model_databases.FarmArea, string)
}

type Service struct {
	repo Repositorier
}

func NewService(repo Repositorier) Servicer {
	return &Service{
		repo:  repo,
	}
}

//func (s *Service) GetFarmAreaAndName(farmAreaId string) (*model_databases.FarmArea, string) {
//	farmAreaModel, err := s.repo.FindOneFarmArea(config.STATUS_ACTIVE, farmAreaId)
//	if err != nil {
//		return nil, ""
//	}
//	return farmAreaModel, farmAreaModel.FarmAreaName
//}
