package sf_formula_plant

import (
	"github.com/jjoykkm/ln-backend/common/models/model_other"
	"github.com/jjoykkm/ln-backend/config"
	"github.com/jjoykkm/ln-backend/models/model_databases"
	"github.com/mitchellh/mapstructure"
)

type Servicer interface {
	// Service for API
	GetPlantCategoryList(status, language string) *model_other.BodyResp
	GetPlantCategoryItem(status, plantTypeId, language string, offset int) *model_other.BodyRespOffset
	GetPlantOverviewByPlant(status, uid, plantId string, offset int) *model_other.BodyRespOffset
	//GetPlantOverviewFavorite(status, uid, language string, offset int) *model_other.BodyRespOffset

	// Function
	GetFavoriteFormulaPlant(status, uid, resultType string) ([]model_databases.FavoritePlant, []string, map[string]bool)
	GetRateScoreAndPeople(formulaPlant model_databases.FormulaPlant) (float32, int)
	//GetMyPlantOverview(status, uid, language string, offset int) ([]model_services.ForPlantItem, int, int)
	//GetPlantOverviewByPlant(status, uid, plantId, language string, offset int) ([]model_services.ForPlantItem, int, int)
	//GetFertilizerRatioRelate(status, formulaPlantId, language string) ([]model_services.ForPlantFert, int)
	//GetSensorValueRecRelate(status, formulaPlantId, language string) ([]model_services.ForPlantSensor, int)
	//GetFormulaPlantDetail(status, formulaPlantId, language string) model_services.ForPlantFormula
}

type Service struct {
	repo  Repositorier
}

func NewService(repo Repositorier) Servicer {
	return &Service{
		repo:  repo,
	}
}

func (s *Service) GetPlantCategoryList(status, language string) *model_other.BodyResp {
	var ptCat PlantTypeCat
	var ptCatList []PlantTypeCat

	plantTypeList, err := s.repo.FindAllPlantType(status)
	if err != nil{
		return &model_other.BodyResp{
			Item: nil,
			Total: 0,
		}
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
	}
}

func (s *Service) GetPlantCategoryItem(status, plantTypeId, language string, offset int) *model_other.BodyRespOffset {
	var currentOffset int
	var total int

	joinList, err := s.repo.FindJoinPlantWithPlantType(status, plantTypeId, offset)
	if err != nil{
		return &model_other.BodyRespOffset{
			Item: nil,
			Offset: offset,
			Total: 0,
		}
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

	total = len(joinList)
	currentOffset = offset + total

	return &model_other.BodyRespOffset{
		Item: joinList,
		Offset: currentOffset,
		Total: total,
	}
}

func (s *Service) GetFavoriteFormulaPlant(status, uid, resultType string) ([]model_databases.FavoritePlant, []string, map[string]bool) {
	var favFPKey []string
	favFPMap := make(map[string]bool)

	if uid == "" {
		return nil, nil, nil
	}
	favFPList, err := s.repo.FindAllFavoriteFormulaPlant(status, uid)
	if err != nil{
		return nil, nil, nil
	}
	if resultType != config.GetResType().Struct {
		for _, favFP := range favFPList {
			favFPKey = append(favFPKey, favFP.FormulaPlantId.UUID.String())
			favFPMap[favFP.FormulaPlantId.UUID.String()] = true
		}
	}
	return favFPList, favFPKey, favFPMap
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

func (s *Service) GetPlantOverviewFavorite(status, uid, language string, offset int) *model_other.BodyRespOffset {
	//var joinArray []JoinPlantAndFormulaPlant
//	var formulaPlant model_databases.FormulaPlant
//	var plantOverview model_services.ForPlantItem
//	var plantOverviewArray []model_services.ForPlantItem
//	var currentOffset int
//	var total int
//	var found bool
//
//	if uid == "" {
//		return &model_other.BodyRespOffset{
//			Item: nil,
//			Offset: offset,
//			Total: 0,
//		}
//	}
//
//	countryMap := make(map[string]string)
//	provinceMap := make(map[string]string)
//	plantTypeMap := make(map[string]string)
//	userMap := make(map[string]string)
//
//	_, formulaPlantList, _ := IntFormulaPlant.GetFavoriteFormulaPlanter(ln, config.GetStatus().Active, uid)
//	sqlIn := utility.ConvertListToStringIn(formulaPlantList)
//	sql := fmt.Sprintf("SELECT * FROM %s INNER JOIN %s ON %s.plant_id = %s.plant_id WHERE %s.status_id = '%s' AND %s.formula_plant_id IN %s OFFSET %d LIMIT 100",
//		config.DB_FORMULA_PLANT, config.DB_PLANT, config.DB_FORMULA_PLANT, config.DB_PLANT, config.DB_FORMULA_PLANT, status, config.DB_FORMULA_PLANT, sqlIn, offset)
//	fmt.Println(sql)
//	err := ln.Db.Raw(sql).Scan(&joinArray).Error
//	if err != nil {
//		log.Print(err)
//	}
//
//	for _, join := range joinArray {
//		mapstructure.Decode(join, &plantOverview)
//		mapstructure.Decode(join, &formulaPlant)
//		plantOverview.RateScore, plantOverview.RatePeople = IntFormulaPlant.GetRateScoreAndPeopleer(ln, formulaPlant)
//		plantOverview.IsFavorite = true
//
//		//Get Country name
//		plantOverview.CountryName, found = countryMap[plantOverview.CountryId.UUID.String()]
//		if !found {
//			_, plantOverview.CountryName = IntCommon.GetCountryNameer(ln, plantOverview.CountryId.UUID.String(), language)
//			countryMap[plantOverview.CountryId.UUID.String()] = plantOverview.CountryName
//		}
//
//		//Get Country name
//		plantOverview.ProvinceName, found = provinceMap[plantOverview.ProvinceId.UUID.String()]
//		if !found {
//			_, plantOverview.ProvinceName = IntCommon.GetProvinceNameer(ln, plantOverview.ProvinceId.UUID.String(), language)
//			provinceMap[plantOverview.ProvinceId.UUID.String()] = plantOverview.ProvinceName
//		}
//
//		//Get Plant Type name
//		plantOverview.PlantTypeName, found = plantTypeMap[plantOverview.PlantTypeId.UUID.String()]
//		if !found {
//			_, plantOverview.PlantTypeName = IntCommon.GetPlantTypeNameer(ln, plantOverview.PlantTypeId.UUID.String(), language)
//			plantTypeMap[plantOverview.PlantTypeId.UUID.String()] = plantOverview.PlantTypeName
//		}
//
//		//Get User name
//		plantOverview.Username, found = userMap[plantOverview.Uid.UUID.String()]
//		if !found {
//			_, plantOverview.Username = IntCommon.GetUserNameer(ln, plantOverview.Uid.UUID.String())
//			plantTypeMap[plantOverview.Uid.UUID.String()] = plantOverview.Username
//		}
//
//		plantOverviewArray = append(plantOverviewArray, plantOverview)
//	}
//	total = len(plantOverviewArray)
//	currentOffset = offset + total
//
	return &model_other.BodyRespOffset{
		Item: nil,
		Offset: offset,
		Total: 0,
	}
}

//func (s *Service) GetMyPlantOverviewer(status, uid, language string, offset int) *model_other.BodyRespOffset {
//	var formulaPlant model_databases.FormulaPlant
//	var plantOverview model_services.ForPlantItem
//	var joinArray []model_services.JoinFormulaPlantAndPlant
//	var plantOverviewArray []model_services.ForPlantItem
//	var currentOffset int
//	var total int
//	var found bool
//	var countryMap map[string]string
//	var provinceMap map[string]string
//	var plantTypeMap map[string]string
//	var formulaPlantFavMap map[string]bool
//	var userMap map[string]string
//
//	if uid == "" {
//		return nil, offset, 0
//	}
//
//	countryMap = make(map[string]string)
//	provinceMap = make(map[string]string)
//	plantTypeMap = make(map[string]string)
//	userMap = make(map[string]string)
//
//	_, _, formulaPlantFavMap = IntFormulaPlant.GetFavoriteFormulaPlanter(ln, config.GetStatus().Active, uid)
//	sql := fmt.Sprintf("SELECT * FROM %s INNER JOIN %s ON %s.plant_id = %s.plant_id WHERE %s.status_id = '%s' AND %s.uid = '%s' OFFSET %d LIMIT 100",
//		config.DB_FORMULA_PLANT, config.DB_PLANT, config.DB_FORMULA_PLANT, config.DB_PLANT, config.DB_FORMULA_PLANT, status, config.DB_FORMULA_PLANT, uid, offset)
//	fmt.Println(sql)
//	err := ln.Db.Raw(sql).Scan(&joinArray).Error
//	if err != nil {
//		log.Print(err)
//	}
//
//	for _, join := range joinArray {
//		mapstructure.Decode(join, &formulaPlant)
//		mapstructure.Decode(join, &plantOverview)
//		plantOverview.RateScore, plantOverview.RatePeople = IntFormulaPlant.GetRateScoreAndPeopleer(ln, formulaPlant)
//
//		//Get Country name
//		plantOverview.CountryName, found = countryMap[plantOverview.CountryId.UUID.String()]
//		if !found {
//			_, plantOverview.CountryName = IntCommon.GetCountryNameer(ln, plantOverview.CountryId.UUID.String(), language)
//			countryMap[plantOverview.CountryId.UUID.String()] = plantOverview.CountryName
//		}
//
//		//Get Province name
//		plantOverview.ProvinceName, found = provinceMap[plantOverview.ProvinceId.UUID.String()]
//		if !found {
//			_, plantOverview.ProvinceName = IntCommon.GetProvinceNameer(ln, plantOverview.ProvinceId.UUID.String(), language)
//			provinceMap[plantOverview.ProvinceId.UUID.String()] = plantOverview.ProvinceName
//		}
//
//		//Get Plant Type name
//		plantOverview.PlantTypeName, found = plantTypeMap[plantOverview.PlantTypeId.UUID.String()]
//		if !found {
//			_, plantOverview.PlantTypeName = IntCommon.GetPlantTypeNameer(ln, plantOverview.PlantTypeId.UUID.String(), language)
//			plantTypeMap[plantOverview.PlantTypeId.UUID.String()] = plantOverview.PlantTypeName
//		}
//
//		//Check Favorite
//		plantOverview.IsFavorite, found = formulaPlantFavMap[plantOverview.Uid.UUID.String()]
//		if !found {
//			plantOverview.IsFavorite = false
//		}
//
//		//Get User name
//		plantOverview.Username, found = userMap[plantOverview.Uid.UUID.String()]
//		if !found {
//			_, plantOverview.Username = IntCommon.GetUserNameer(ln, plantOverview.Uid.UUID.String())
//			plantTypeMap[plantOverview.Uid.UUID.String()] = plantOverview.Username
//		}
//
//		plantOverviewArray = append(plantOverviewArray, plantOverview)
//	}
//
//	total = len(plantOverviewArray)
//	currentOffset = offset + total
//
//	return plantOverviewArray, currentOffset, total
//}

func (s *Service) GetPlantOverviewByPlant(status, uid, plantId string, offset int) *model_other.BodyRespOffset {

	if uid == "" || plantId == "" {
		return &model_other.BodyRespOffset{
			Item: nil,
			Offset: offset,
			Total: 0,
		}
	}

	forPlant, err := s.repo.FindPlantWithFormulaPlant(status, plantId, offset)
	if err != nil{
		return &model_other.BodyRespOffset{
			Item: nil,
			Offset: offset,
			Total: 0,
		}
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
	}
}

//func (s *Service) GetFertilizerRatioRelateer(status, formulaPlantId, language string) ([]model_services.ForPlantFert, int) {
//	var joinArray []model_services.JoinFertilizerAndPlant
//	var plantFert model_services.ForPlantFert
//	var plantFertArray []model_services.ForPlantFert
//	var total int
//	var found bool
//	var fertCatMap map[string]string
//
//	if formulaPlantId == "" {
//		return nil, 0
//	}
//
//	fertCatMap = make(map[string]string)
//	sql := fmt.Sprintf("SELECT * FROM %s INNER JOIN %s ON %s.fertilizer_id = %s.fertilizer_id WHERE %s.status_id = '%s' AND %s.formula_plant_id = '%s'",
//		config.DB_FERTILIZER, config.DB_TRANS_FERTILIZER_RATIO, config.DB_FERTILIZER, config.DB_TRANS_FERTILIZER_RATIO, config.DB_FERTILIZER, status, config.DB_TRANS_FERTILIZER_RATIO, formulaPlantId)
//	fmt.Println(sql)
//	err := ln.Db.Raw(sql).Scan(&joinArray).Error
//	if err != nil {
//		log.Print(err)
//	}
//
//	for _, join := range joinArray {
//		mapstructure.Decode(join, &plantFert)
//		switch language {
//		case config.LANGUAGE_EN:
//			plantFert.FertilizerName = join.FertilizerEN
//		case config.LANGUAGE_TH:
//			plantFert.FertilizerName = join.FertilizerTH
//		}
//
//		//Get Fertilizer category name
//		plantFert.FertCatName, found = fertCatMap[plantFert.FertCatId.UUID.String()]
//		if !found {
//			_, plantFert.FertCatName = IntCommon.GetFertCatNameer(ln, plantFert.FertCatId.UUID.String(), language)
//			fertCatMap[plantFert.FertCatId.UUID.String()] = plantFert.FertCatName
//		}
//
//		plantFertArray = append(plantFertArray, plantFert)
//	}
//
//	total = len(plantFertArray)
//
//	return plantFertArray, total
//}

//func (s *Service) GetSensorValueRecRelateer(status, formulaPlantId, language string) ([]model_services.ForPlantSensor, int) {
//	var joinArray []model_services.JoinSensorTypeAndTrans
//	var plantSensor model_services.ForPlantSensor
//	var plantSensorArray []model_services.ForPlantSensor
//	var total int
//
//	if formulaPlantId == "" {
//		return nil, 0
//	}
//
//	sql := fmt.Sprintf("SELECT * FROM %s INNER JOIN %s ON %s.sensor_type_id= %s.sensor_type_id WHERE %s.status_id = '%s' AND %s.formula_plant_id = '%s'",
//		config.DB_SENSOR_TYPE, config.DB_TRANS_SENSOR_VALUE_REC, config.DB_SENSOR_TYPE, config.DB_TRANS_SENSOR_VALUE_REC, config.DB_SENSOR_TYPE, status, config.DB_TRANS_SENSOR_VALUE_REC, formulaPlantId)
//	fmt.Println(sql)
//	err := ln.Db.Raw(sql).Scan(&joinArray).Error
//	if err != nil {
//		log.Print(err)
//	}
//	for _, join := range joinArray {
//		mapstructure.Decode(join, &plantSensor)
//		switch language {
//		case config.LANGUAGE_EN:
//			plantSensor.SensorTypeName = join.SensorTypeNameEN
//		case config.LANGUAGE_TH:
//			plantSensor.SensorTypeName = join.SensorTypeNameTH
//		}
//		plantSensorArray = append(plantSensorArray, plantSensor)
//	}
//	total = len(plantSensorArray)
//
//	return plantSensorArray, total
//}

//func (s *Service) GetFormulaPlantDetailer(status, formulaPlantId, language string) model_services.ForPlantFormula {
//	var formulaPlantModel model_databases.FormulaPlant
//	var formula model_services.ForPlantFormula
//
//	if formulaPlantId == "" {
//		return formula
//	}
//
//	condition := fmt.Sprintf("SELECT * FROM %s WHERE status_id = '%s' AND formula_plant_id = '%s'",
//		config.DB_FORMULA_PLANT, status, formulaPlantId)
//	fmt.Println(condition)
//	err := ln.Db.Raw(condition).Scan(&formulaPlantModel).Error
//	if err != nil {
//		log.Print(err)
//	}
//	mapstructure.Decode(formulaPlantModel, &formula)
//
//	_, formula.Username = IntCommon.GetUserNameer(ln, formula.Uid.UUID.String())
//
//	formula.SensorList, _ = IntFormulaPlant.GetSensorValueRecRelateer(ln, config.GetStatus().Active, formula.FormulaPlantId.UUID.String(), language)
//
//	formula.FertList, _ = IntFormulaPlant.GetFertilizerRatioRelateer(ln, config.GetStatus().Active, formula.FormulaPlantId.UUID.String(), language)
//
//	return formula
//}
