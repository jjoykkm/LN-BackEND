package sf_formula_plant

import (
	"github.com/jjoykkm/ln-backend/common/models/model_other"
	"github.com/jjoykkm/ln-backend/config"
	"github.com/jjoykkm/ln-backend/models/model_databases"
	"github.com/mitchellh/mapstructure"
)

type Servicer interface {
	// Service for API
	GetPlantCategoryList(status, language string) (*model_other.BodyResp, error)
	GetPlantCategoryItem(status, plantTypeId, language string, offset int) (*model_other.BodyRespOffset, error)
	GetPlantOverviewByPlant(status, uid, plantId string, offset int) (*model_other.BodyRespOffset, error)
	GetPlantOverviewFavorite(status, uid, language string, offset int) (*model_other.BodyRespOffset, error)
	GetMyPlantOverview(status, uid, language string, offset int) (*model_other.BodyRespOffset, error)
	GetFormulaPlantDetail(status, formulaPlantId, language string) (*model_other.BodyResp, error)

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

func (s *Service) GetPlantCategoryList(status, language string) (*model_other.BodyResp, error) {
	var ptCat PlantTypeCat
	var ptCatList []PlantTypeCat

	plantTypeList, err := s.repo.FindAllPlantType(status)
	if err != nil{
		return nil, err
	}

	for _, plantType := range plantTypeList {
		mapstructure.Decode(plantType, &ptCat)
		switch language {
		case config.LANGUAGE_EN:
			ptCat.PlantTypeName = plantType.PlantTypeEN
		case config.LANGUAGE_TH:
			ptCat.PlantTypeName = plantType.PlantTypeTH
		}
		ptCatList = append(ptCatList, ptCat)
	}

	return &model_other.BodyResp{
		Item: ptCatList,
		Total: len(ptCatList),
	}, nil
}

func (s *Service) GetPlantCategoryItem(status, plantTypeId, language string, offset int) (*model_other.BodyRespOffset, error) {
	joinList, err := s.repo.FindAllPlantWithPlantType(status, plantTypeId, offset)
	if err != nil{
		return nil, err
	}

	for idx, join := range joinList {
		switch language {
		case config.LANGUAGE_EN:
			join.PlantType.PlantTypeTH = ""
			join.Plant.PlantNameTH = ""
			join.Plant.PlantDescTH = ""
		case config.LANGUAGE_TH:
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
	currentOffset := offset + total

	return &model_other.BodyRespOffset{
		Item: joinList,
		Offset: currentOffset,
		Total: total,
	}, nil
}

func (s *Service) GetPlantOverviewByPlant(status, uid, plantId string, offset int) (*model_other.BodyRespOffset, error) {
	forPlant, err := s.repo.FindAllFormulaPlantByPlant(status, plantId, offset)
	if err != nil{
		return nil, err
	}
	// Get favorite formula plant
	_, favMap, _ := s.repo.FindAllFavForPlantId(status, config.GetResType().Map, uid)
	// Get planted formula plant
	_, plantedMap, _ := s.repo.FindAllPlantedForPlantId(status, config.GetResType().Map, uid)
	for idx, wa := range forPlant {
		// Check is favorite
		wa.IsFavorite = favMap[wa.FormulaPlant.FormulaPlantId.UUID.String()]
		// Check planted
		wa.IsPlanted = plantedMap[wa.FormulaPlant.FormulaPlantId.UUID.String()]
		forPlant[idx] = wa
	}
	total := len(forPlant)
	currentOffset := offset + total
	return &model_other.BodyRespOffset{
		Item: forPlant,
		Offset: currentOffset,
		Total: total,
	},nil
}

func (s *Service) GetPlantOverviewFavorite(status, uid, language string, offset int) (*model_other.BodyRespOffset, error) {
	forPlant, err := s.repo.FindAllFormulaPlantFavorite(status, uid, offset)
	if err != nil{
		return nil, err
	}
	// Get planted formula plant
	_, plantedMap, _ := s.repo.FindAllPlantedForPlantId(status, config.GetResType().Map, uid)
	for idx, wa := range forPlant {
		wa.IsFavorite = true
		// Check planted
		wa.IsPlanted = plantedMap[wa.FormulaPlant.FormulaPlantId.UUID.String()]
		forPlant[idx] = wa
	}
	total := len(forPlant)
	currentOffset := offset + total
	return &model_other.BodyRespOffset{
		Item: forPlant,
		Offset: currentOffset,
		Total: total,
	}, nil
}

func (s *Service) GetMyPlantOverview(status, uid, language string, offset int) (*model_other.BodyRespOffset, error) {
	forPlant, err := s.repo.FindAllMyFormulaPlant(status, uid, offset)
	if err != nil{
		return nil, err
	}
	// Get favorite formula plant
	_, favMap, _ := s.repo.FindAllFavForPlantId(status, config.GetResType().Map, uid)
	// Get planted formula plant
	_, plantedMap, _ := s.repo.FindAllPlantedForPlantId(status, config.GetResType().Map, uid)
	for idx, wa := range forPlant {
		// Check is favorite
		wa.IsFavorite = favMap[wa.FormulaPlant.FormulaPlantId.UUID.String()]
		// Check planted
		wa.IsPlanted = plantedMap[wa.FormulaPlant.FormulaPlantId.UUID.String()]
		forPlant[idx] = wa
	}
	total := len(forPlant)
	currentOffset := offset + total
	return &model_other.BodyRespOffset{
		Item: forPlant,
		Offset: currentOffset,
		Total: total,
	}, nil
}

func (s *Service) GetFormulaPlantDetail(status, formulaPlantId, language string) (*model_other.BodyResp, error) {
	if formulaPlantId == "" {
		return nil, nil
	}
	forPlantDetail, err := s.repo.FindAllFormulaPlantDetail(status, formulaPlantId)
	if err != nil{
		return nil, err
	}
	return &model_other.BodyResp{
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
