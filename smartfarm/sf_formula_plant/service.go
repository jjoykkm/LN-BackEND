package sf_formula_plant

import (
	"github.com/jjoykkm/ln-backend/common/config"
	"github.com/jjoykkm/ln-backend/common/models/model_db"
	"github.com/jjoykkm/ln-backend/common/models/model_other"
)

type Servicer interface {
	// Service for API
	// status, language string
	GetPlantCategoryList(status string) (*model_other.RespModel, error)
	// status, plantTypeId, language string, reqModel.Offset int
	GetPlantCategoryItem(status string, reqModel *model_other.ReqModel) (*model_other.RespOffsetModel, error)
	// status, reqModel.User.Uid, plantId string, reqModel.Offset int
	GetPlantOverviewByPlant(status string, reqModel *model_other.ReqModel) (*model_other.RespOffsetModel, error)
	// status, reqModel.User.Uid, language string, reqModel.Offset int
	GetPlantOverviewFavorite(status string, reqModel *model_other.ReqModel) (*model_other.RespOffsetModel, error)
	// status, reqModel.User.Uid, language string, reqModel.Offset int
	GetPlantOfMine(status string, reqModel *model_other.ReqModel) (*model_other.RespOffsetModel, error)
	// status, formulaPlasntId, language string
	GetFormulaPlantDetail(status string, reqModel *model_other.ReqModel) (*model_other.RespModel, error)

	//Upsert
	AddChangeFormulaPlant(reqModel *ForPlantUS) error
	AddChangeFertilizer(req *model_db.FertilizerUS) error

	//Update
	AddFavoriteForPlant(reqModel *model_db.FavForPlantUS) error

	//Delete
	RemoveFavoriteForPlant(reqModel *model_db.FavForPlantUS) error
}

type Service struct {
	repo  Repositorier
}

func NewService(repo Repositorier) Servicer {
	return &Service{
		repo:  repo,
	}
}

func (s *Service) GetPlantCategoryList(status string) (*model_other.RespModel, error) {
	plantTypeList, err := s.repo.FindAllPlantType(status)
	if err != nil{
		return nil, err
	}

	return &model_other.RespModel{
		Item: plantTypeList,
		Total: len(plantTypeList),
	}, nil
}

func (s *Service) GetPlantCategoryItem(status string, reqModel *model_other.ReqModel) (*model_other.RespOffsetModel, error) {
	result, err := s.repo.FindAllPlantWithPlantType(status, reqModel.PlantTypeId, reqModel.Offset)
	if err != nil{
		return nil, err
	}

	total := len(result)
	currentOffset := reqModel.Offset + total

	return &model_other.RespOffsetModel{
		Item: result,
		Offset: currentOffset,
		Total: total,
	}, nil
}

func (s *Service) GetPlantOverviewByPlant(status string, reqModel *model_other.ReqModel) (*model_other.RespOffsetModel, error) {

	forPlant, err := s.repo.FindAllFormulaPlantByPlant(status, reqModel.PlantId, reqModel.Offset)
	if err != nil{
		return nil, err
	}
	// Get favorite formula plant
	favMap, _ := s.repo.FindAllFavForPlantId(status, reqModel.User.Uid)
	// Get planted formula plant
	plantedMap, _ := s.repo.FindAllPlantedForPlantId(status, reqModel.User.Uid)
	// Fill data
	for idx, wa := range forPlant.ForPlantDetail {
		forPlantId := wa.FormulaPlant.FormulaPlantId
		// Check is favorite
		wa.IsFavorite = favMap[forPlantId]
		// Check planted
		wa.IsPlanted = plantedMap[forPlantId]
		forPlant.ForPlantDetail[idx] = wa
	}
	total := len(forPlant.ForPlantDetail)
	currentOffset := reqModel.Offset + total
	return &model_other.RespOffsetModel{
		Item: forPlant,
		Offset: currentOffset,
		Total: total,
	},nil
}

func (s *Service) GetPlantOverviewFavorite(status string, reqModel *model_other.ReqModel) (*model_other.RespOffsetModel, error) {
	forPlant, err := s.repo.FindAllFormulaPlantFavorite(status, reqModel.User.Uid, reqModel.Offset)
	if err != nil{
		return nil, err
	}
	// Get planted formula plant
	plantedMap, _ := s.repo.FindAllPlantedForPlantId(status, reqModel.User.Uid)
	for idx, wa := range forPlant {
		forPlantId := wa.FormulaPlant.FormulaPlant.FormulaPlantId
		wa.IsFavorite = true
		// Check planted
		wa.IsPlanted = plantedMap[forPlantId]
		forPlant[idx] = wa
	}
	total := len(forPlant)
	currentOffset := reqModel.Offset + total
	return &model_other.RespOffsetModel{
		Item: forPlant,
		Offset: currentOffset,
		Total: total,
	}, nil
}

func (s *Service) GetPlantOfMine(status string, reqModel *model_other.ReqModel) (*model_other.RespOffsetModel, error) {
	forPlant, err := s.repo.FindAllMyFormulaPlant(status, reqModel.User.Uid, reqModel.Offset)
	if err != nil{
		return nil, err
	}
	// Get favorite formula plant
	favMap, _ := s.repo.FindAllFavForPlantId(status, reqModel.User.Uid)
	// Get planted formula plant
	plantedMap, _ := s.repo.FindAllPlantedForPlantId(status, reqModel.User.Uid)
	for idx, wa := range forPlant {
		forPlantId := wa.FormulaPlant.FormulaPlant.FormulaPlantId
		// Check is favorite
		wa.IsFavorite = favMap[forPlantId]
		// Check planted
		wa.IsPlanted = plantedMap[forPlantId]
		forPlant[idx] = wa
	}
	total := len(forPlant)
	currentOffset := reqModel.Offset + total
	return &model_other.RespOffsetModel{
		Item: forPlant,
		Offset: currentOffset,
		Total: total,
	}, nil
}

func (s *Service) GetFormulaPlantDetail(status string, reqModel *model_other.ReqModel) (*model_other.RespModel, error) {
	forPlantDetail, err := s.repo.FindAllFormulaPlantDetail(status, reqModel.FormulaPlantId)
	if err != nil{
		return nil, err
	}
	// Get favorite formula plant
	favMap, _ := s.repo.FindAllFavForPlantId(status, reqModel.User.Uid)
	// Get planted formula plant
	plantedMap, _ := s.repo.FindAllPlantedForPlantId(status, reqModel.User.Uid)
	// Fill data
	forPlantId := forPlantDetail.FormulaPlant.FormulaPlant.FormulaPlant.FormulaPlantId
	// Check is favorite
	forPlantDetail.FormulaPlant.IsFavorite = favMap[forPlantId]
	// Check planted
	forPlantDetail.FormulaPlant.IsPlanted = plantedMap[forPlantId]

	return &model_other.RespModel{
		Item: forPlantDetail,
		Total: 1,
	}, nil
}

//-------------------------------------------------------------------------------//
//				 	   				Upsert
//-------------------------------------------------------------------------------//
func (s *Service) AddChangeFormulaPlant(reqModel *ForPlantUS) error {
	var (
		forPlantId *string
		err	error
	)
	//Prepare data before process
	reqModel.FormulaPlant.StatusId = config.GetStatus().Active
	tx := s.repo.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	//Upsert Formula Plant
	if (model_db.FormulaPlantUS{}) != reqModel.FormulaPlant {
		err, forPlantId = tx.UpsertFormulaPlant(&reqModel.FormulaPlant)
		if err != nil{
			tx.Rollback()
			return err
		}
	}
	//Upsert Sensor Value
	if len(reqModel.SensorValue) > 0 {
		if *forPlantId != "" {
			//Assign FormulaPlantId
			for idx, wa := range reqModel.SensorValue {
				wa.FormulaPlantId = reqModel.FormulaPlant.FormulaPlantId
				wa.StatusId = config.GetStatus().Active
				reqModel.SensorValue[idx] = wa
			}
			err := tx.UpsertForPlantSensor(reqModel.SensorValue)
			if err != nil{
				tx.Rollback()
				return err
			}
		}
	}
	//Upsert Fertilizer Ratio
	if len(reqModel.FertRatio) > 0 {
		if *forPlantId != "" {
			//Assign FormulaPlantId
			for idx, wa := range reqModel.FertRatio {
				wa.FormulaPlantId = reqModel.FormulaPlant.FormulaPlantId
				wa.StatusId = config.GetStatus().Active
				reqModel.FertRatio[idx] = wa
			}
			err := tx.UpsertForPlantFert(reqModel.FertRatio)
			if err != nil{
				tx.Rollback()
				return err
			}
		}
	}

	return tx.db.Commit().Error
}

func (s *Service) AddChangeFertilizer(req *model_db.FertilizerUS) error {
	req.StatusId = config.GetStatus().Active
	err := s.repo.UpsertFertilizer(req)
	if err != nil{
		return err
	}
	return nil
}
//-------------------------------------------------------------------------------//
//									Create
//-------------------------------------------------------------------------------//
func (s *Service) AddFavoriteForPlant(reqModel *model_db.FavForPlantUS) error {
	if (model_db.FavForPlantUS{}) == *reqModel  {
		//return err.ErrContext
	}
	err := s.repo.CreateFavFormulaPlant(reqModel)
	if err != nil{
		return err
	}
	return nil
}

//-------------------------------------------------------------------------------//
//									Delete
//-------------------------------------------------------------------------------//
func (s *Service) RemoveFavoriteForPlant(reqModel *model_db.FavForPlantUS) error {
	if (model_db.FavForPlantUS{}) == *reqModel  {
		//return err.ErrContext
	}
	err := s.repo.DeleteFavFormulaPlant(reqModel)
	if err != nil{
		return err
	}
	return nil
}