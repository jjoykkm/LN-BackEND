package sf_formula_plant

import (
	"github.com/jjoykkm/ln-backend/common/models/model_other"
	"github.com/jjoykkm/ln-backend/config"
	"github.com/jjoykkm/ln-backend/modelsOld/model_databases"
)

type Servicer interface {
	// Service for API
	// status, language string
	GetPlantCategoryList(status string, ReqModel *model_other.ReqModel) (*model_other.RespModel, error)
	// status, plantTypeId, language string, ReqModel.Offset int
	GetPlantCategoryItem(status string, ReqModel *model_other.ReqModel) (*model_other.RespOffsetModel, error)
	// status, ReqModel.Uid, plantId string, ReqModel.Offset int
	GetPlantOverviewByPlant(status string, ReqModel *model_other.ReqModel) (*model_other.RespOffsetModel, error)
	// status, ReqModel.Uid, language string, ReqModel.Offset int
	GetPlantOverviewFavorite(status string, ReqModel *model_other.ReqModel) (*model_other.RespOffsetModel, error)
	// status, ReqModel.Uid, language string, ReqModel.Offset int
	GetMyPlantOverview(status string, ReqModel *model_other.ReqModel) (*model_other.RespOffsetModel, error)
	// status, formulaPlasntId, language string
	GetFormulaPlantDetail(status string, ReqModel *model_other.ReqModel) (*model_other.RespModel, error)

	// Function
	GetRateScoreAndPeople(formulaPlant model_databases.FormulaPlant) (float32, int)
}

type Service struct {
	repo  Repositorier
}

func NewService(repo Repositorier) Servicer {
	return &Service{
		repo:  repo,
	}
}

func (s *Service) GetPlantCategoryList(status string, ReqModel *model_other.ReqModel) (*model_other.RespModel, error) {
	plantTypeList, err := s.repo.FindAllPlantType(status)
	if err != nil{
		return nil, err
	}

	return &model_other.RespModel{
		Item: plantTypeList,
		Total: len(plantTypeList),
	}, nil
}

func (s *Service) GetPlantCategoryItem(status string, ReqModel *model_other.ReqModel) (*model_other.RespOffsetModel, error) {
	joinList, err := s.repo.FindAllPlantWithPlantType(status, ReqModel.PlantTypeId, ReqModel.Offset)
	if err != nil{
		return nil, err
	}

	for idx, join := range joinList {
		switch ReqModel.Language {
		case config.GetLanguage().En:
			join.PlantType.PlantTypeTH = ""
			join.Plant.PlantNameTH = ""
			join.Plant.PlantDescTH = ""
		case config.GetLanguage().Th:
			join.PlantType.PlantTypeEN = ""
			join.Plant.PlantNameEN = ""
			join.Plant.PlantDescEN = ""
		}
		join.Plant.TotalItem = int(s.repo.GetCountFormulaPlant(config.GetStatus().Active, join.Plant.PlantId.UUID.String()))
		// Delete PlantTypeId in struct Plant
		join.Plant.PlantTypeId = nil
		// Modify list
		joinList[idx] = join
	}

	total := len(joinList)
	currentOffset := ReqModel.Offset + total

	return &model_other.RespOffsetModel{
		Item: joinList,
		Offset: currentOffset,
		Total: total,
	}, nil
}

func (s *Service) GetPlantOverviewByPlant(status string, ReqModel *model_other.ReqModel) (*model_other.RespOffsetModel, error) {
	forPlant, err := s.repo.FindAllFormulaPlantByPlant(status, ReqModel.PlantId, ReqModel.Offset)
	if err != nil{
		return nil, err
	}
	// Get favorite formula plant
	_, favMap, _ := s.repo.FindAllFavForPlantId(status, config.GetResType().Map, ReqModel.Uid)
	// Get planted formula plant
	_, plantedMap, _ := s.repo.FindAllPlantedForPlantId(status, config.GetResType().Map, ReqModel.Uid)
	for idx, wa := range forPlant {
		// Check is favorite
		wa.IsFavorite = favMap[wa.FormulaPlant.FormulaPlant.FormulaPlantId.UUID.String()]
		// Check planted
		wa.IsPlanted = plantedMap[wa.FormulaPlant.FormulaPlant.FormulaPlantId.UUID.String()]
		forPlant[idx] = wa
	}
	total := len(forPlant)
	currentOffset := ReqModel.Offset + total
	return &model_other.RespOffsetModel{
		Item: forPlant,
		Offset: currentOffset,
		Total: total,
	},nil
}

func (s *Service) GetPlantOverviewFavorite(status string, ReqModel *model_other.ReqModel) (*model_other.RespOffsetModel, error) {
	forPlant, err := s.repo.FindAllFormulaPlantFavorite(status, ReqModel.Uid, ReqModel.Offset)
	if err != nil{
		return nil, err
	}
	// Get planted formula plant
	_, plantedMap, _ := s.repo.FindAllPlantedForPlantId(status, config.GetResType().Map, ReqModel.Uid)
	for idx, wa := range forPlant {
		wa.IsFavorite = true
		// Check planted
		wa.IsPlanted = plantedMap[wa.FormulaPlant.FormulaPlant.FormulaPlantId.UUID.String()]
		forPlant[idx] = wa
	}
	total := len(forPlant)
	currentOffset := ReqModel.Offset + total
	return &model_other.RespOffsetModel{
		Item: forPlant,
		Offset: currentOffset,
		Total: total,
	}, nil
}

func (s *Service) GetMyPlantOverview(status string, ReqModel *model_other.ReqModel) (*model_other.RespOffsetModel, error) {
	forPlant, err := s.repo.FindAllMyFormulaPlant(status, ReqModel.Uid, ReqModel.Offset)
	if err != nil{
		return nil, err
	}
	// Get favorite formula plant
	_, favMap, _ := s.repo.FindAllFavForPlantId(status, config.GetResType().Map, ReqModel.Uid)
	// Get planted formula plant
	_, plantedMap, _ := s.repo.FindAllPlantedForPlantId(status, config.GetResType().Map, ReqModel.Uid)
	for idx, wa := range forPlant {
		// Check is favorite
		wa.IsFavorite = favMap[wa.FormulaPlant.FormulaPlant.FormulaPlantId.UUID.String()]
		// Check planted
		wa.IsPlanted = plantedMap[wa.FormulaPlant.FormulaPlant.FormulaPlantId.UUID.String()]
		forPlant[idx] = wa
	}
	total := len(forPlant)
	currentOffset := ReqModel.Offset + total
	return &model_other.RespOffsetModel{
		Item: forPlant,
		Offset: currentOffset,
		Total: total,
	}, nil
}

func (s *Service) GetFormulaPlantDetail(status string, ReqModel *model_other.ReqModel) (*model_other.RespModel, error) {
	forPlantDetail, err := s.repo.FindAllFormulaPlantDetail(status, ReqModel.FormulaPlantId)
	if err != nil{
		return nil, err
	}
	return &model_other.RespModel{
		Item: forPlantDetail,
		Total: len(forPlantDetail),
	}, nil
}

func (s *Service) GetRateScoreAndPeople(formulaPlant model_databases.FormulaPlant) (float32, int) {
	var rateScore float32
	var ratePeople int

	ratePeople = formulaPlant.Recommend1 + formulaPlant.Recommend2 + formulaPlant.Recommend3 + formulaPlant.Recommend4 + formulaPlant.Recommend5

	rateScore += float32(formulaPlant.Recommend1)
	rateScore += (float32(formulaPlant.Recommend2) * 2)
	rateScore += (float32(formulaPlant.Recommend3) * 3)
	rateScore += (float32(formulaPlant.Recommend4) * 4)
	rateScore += (float32(formulaPlant.Recommend5) * 5)
	rateScore = rateScore / float32(ratePeople)

	return rateScore, ratePeople
}
