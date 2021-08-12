package controllers

import (
	"LN-BackEND/config"
	"LN-BackEND/models/model_databases"
	"LN-BackEND/models/model_services"
	"LN-BackEND/utility"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"log"
	"strings"
)

/*-------------------------------------------------------------------------------------------*/
//                                 INTERFACE
/*-------------------------------------------------------------------------------------------*/
type IntFormulaPlant interface {
	GetPlantCategoryLister(status, language string) ([]model_services.ForPlantCatList, int)
	GetPlantCategoryItemer(status, plantTypeId, language string, offset int) ([]model_services.ForPlantCat, int, int)
	GetFavoriteFormulaPlanter(status, uid string) ([]model_databases.FavoritePlant, []string, map[string]bool)
	GetRateScoreAndPeopleer(formulaPlant model_databases.FormulaPlant) (float32, int)
	GetCountryNameer(countryId, language string) (model_databases.Country, string)
	GetProvinceNameer(provinceId, language string) (model_databases.Province, string)
	GetPlantTypeNameer(plantTypeId, language string) (model_databases.PlantType, string)
	GetFertCatNameer(fertCatId, language string) (model_databases.FertilizerCat, string)
	GetUserNameer(uid string) (model_databases.Users, string)
	GetPlantOverviewFavoriteer(status, uid, language string, offset int) ([]model_services.ForPlantItem, int, int)
	GetMyPlantOverviewer(status, uid, language string, offset int) ([]model_services.ForPlantItem, int, int)
	GetPlantOverviewByPlanter(status, uid, plantId, language string, offset int) ([]model_services.ForPlantItem, int, int)
	GetFertilizerRatioRelateer(status, formulaPlantId, language string) ([]model_services.ForPlantFert, int)
	GetSensorValueRecRelateer(status, formulaPlantId, language string) ([]model_services.ForPlantSensor, int)
	GetFormulaPlantDetailer(status, formulaPlantId, language string) model_services.ForPlantFormula
}

/*-------------------------------------------------------------------------------------------*/
//                                   METHOD
/*-------------------------------------------------------------------------------------------*/
func (ln Ln) GetPlantCategoryLister(status, language string) ([]model_services.ForPlantCatList, int) {
	var plantTypeArray []model_databases.PlantType
	var catList model_services.ForPlantCatList
	var catListArray []model_services.ForPlantCatList
	var total int

	sql := fmt.Sprintf("SELECT * FROM %s WHERE status_id = '%s' ORDER BY plant_type_en ASC",
		config.DB_PLANT_TYPE, status)
	fmt.Println(sql)
	err := ln.Db.Raw(sql).Scan(&plantTypeArray).Error
if err != nil {
		log.Print(err)
	}
	

	for _, plantType := range plantTypeArray {
		mapstructure.Decode(plantType, &catList)
		switch language {
		case config.LANGUAGE_EN:
			catList.PlantTypeName = plantType.PlantTypeEN
		case config.LANGUAGE_TH:
			catList.PlantTypeName = plantType.PlantTypeTH
		}
		catListArray = append(catListArray, catList)
	}
	total = len(catListArray)
	return catListArray, total
}

func (ln Ln) GetPlantCategoryItemer(status, plantTypeId, language string, offset int) ([]model_services.ForPlantCat, int, int) {
	var plantCat model_services.ForPlantCat
	var plantCatArray []model_services.ForPlantCat
	var currentOffset int
	var total int
	var sqlScopePT string

	if plantTypeId != "" {
		sqlScopePT = fmt.Sprintf("AND %s.plant_type_id = '%s'",config.DB_PLANT, plantTypeId)
	}

	var joinPlantAndPlantTypeArray []model_services.JoinPlantAndPlantType

	sql := fmt.Sprintf("SELECT * FROM %s INNER JOIN %s ON %s.plant_type_id = %s.plant_type_id WHERE %s.status_id = '%s' %s OFFSET %d LIMIT 100",
		config.DB_PLANT, config.DB_PLANT_TYPE, config.DB_PLANT, config.DB_PLANT_TYPE, config.DB_PLANT, status, sqlScopePT, offset)
	fmt.Println(sql)
	err := ln.Db.Raw(sql).Scan(&joinPlantAndPlantTypeArray).Error
	if err != nil {
		log.Print(err)
	}

	for _, joinPlantAndPlantType := range joinPlantAndPlantTypeArray {
		mapstructure.Decode(joinPlantAndPlantType, &plantCat)
		switch language {
		case config.LANGUAGE_EN:
			plantCat.PlantTypeName = joinPlantAndPlantType.PlantTypeEN
			plantCat.PlantName = joinPlantAndPlantType.PlantNameEN
			plantCat.PlantDesc = joinPlantAndPlantType.PlantDescEN
		case config.LANGUAGE_TH:
			plantCat.PlantTypeName = joinPlantAndPlantType.PlantTypeTH
			plantCat.PlantName = joinPlantAndPlantType.PlantNameTH
			plantCat.PlantDesc = joinPlantAndPlantType.PlantDescTH
		}
		cond := fmt.Sprintf("plant_id = '%s'", joinPlantAndPlantType.PlantId.UUID.String())
		plantCat.TotalItem = utility.GetCountTable(ln.Db, config.STATUS_ACTIVE, config.DB_FORMULA_PLANT, "formula_plant_id", cond)
		plantCatArray = append(plantCatArray, plantCat)
	}
	total = len(plantCatArray)
	currentOffset = offset + total
	return plantCatArray, currentOffset, total
}

func (ln Ln) GetFavoriteFormulaPlanter(status, uid string) ([]model_databases.FavoritePlant, []string, map[string]bool) {
	var favPlantArray []model_databases.FavoritePlant
	var formulaPlantList []string

	if uid == "" {
		return nil, nil, nil
	}

	formulaPlantMap := make(map[string]bool)

	sql := fmt.Sprintf("SELECT * FROM %s WHERE status_id = '%s' AND uid = '%s' ORDER BY change_date ASC",
		config.DB_FAVORITE_PLANT, status, uid)
	fmt.Println(sql)
	err := ln.Db.Raw(sql).Scan(&favPlantArray).Error
	if err != nil {
		log.Print(err)
	}

	for _, favPlant := range favPlantArray {
		formulaPlantList = append(formulaPlantList, favPlant.FormulaPlantId.UUID.String())
		formulaPlantMap[favPlant.FormulaPlantId.UUID.String()] = true
	}

	return favPlantArray, formulaPlantList, formulaPlantMap
}

func (ln Ln) GetRateScoreAndPeopleer(formulaPlant model_databases.FormulaPlant) (float32, int) {
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

func (ln Ln) GetCountryNameer(countryId, language string) (model_databases.Country, string) {
	var countryModel model_databases.Country
	var countryName string

	sql := fmt.Sprintf("SELECT * FROM %s WHERE status_id = '%s' AND country_id = '%s'",
		config.DB_COUNTRY, config.STATUS_ACTIVE, countryId)
	err := ln.Db.Raw(sql).Scan(&countryModel).Error
	if err != nil {
		log.Print(err)
	}

	switch language {
	case config.LANGUAGE_EN:
		countryName = countryModel.CountryEN
	case config.LANGUAGE_TH:
		countryName = countryModel.CountryTH
	}
	return countryModel, countryName
}

func (ln Ln) GetProvinceNameer(provinceId, language string) (model_databases.Province, string) {
	var provinceModel model_databases.Province
	var provinceName string

	sql := fmt.Sprintf("SELECT * FROM %s WHERE status_id = '%s' AND province_id = '%s'",
		config.DB_PROVINCE, config.STATUS_ACTIVE, provinceId)
	err := ln.Db.Raw(sql).Scan(&provinceModel).Error
	if err != nil {
		log.Print(err)
	}

	switch language {
	case config.LANGUAGE_EN:
		provinceName = provinceModel.ProvinceEN
	case config.LANGUAGE_TH:
		provinceName = provinceModel.ProvinceTH
	}
	return provinceModel, provinceName
}

func (ln Ln) GetPlantTypeNameer(plantTypeId, language string) (model_databases.PlantType, string) {
	var plantTypeModel model_databases.PlantType
	var plantTypeName string

	sql := fmt.Sprintf("SELECT * FROM %s WHERE status_id = '%s' AND plant_type_id = '%s'",
		config.DB_PLANT_TYPE, config.STATUS_ACTIVE, plantTypeId)
	err := ln.Db.Raw(sql).Scan(&plantTypeModel).Error
	if err != nil {
		log.Print(err)
	}
	switch language {
	case config.LANGUAGE_EN:
		plantTypeName = plantTypeModel.PlantTypeEN
	case config.LANGUAGE_TH:
		plantTypeName = plantTypeModel.PlantTypeTH
	}
	return plantTypeModel, plantTypeName
}

func (ln Ln) GetFertCatNameer(fertCatId, language string) (model_databases.FertilizerCat, string) {
	var fertCatModel model_databases.FertilizerCat
	var fertCatName string

	sql := fmt.Sprintf("SELECT * FROM %s WHERE status_id = '%s' AND fertilizer_cat_id = '%s'",
		config.DB_FERTILIZER_CAT, config.STATUS_ACTIVE, fertCatId)
	err := ln.Db.Raw(sql).Scan(&fertCatModel).Error
	if err != nil {
		log.Print(err)
	}

	switch language {
	case config.LANGUAGE_EN:
		fertCatName = fertCatModel.FertilizerCatEN
	case config.LANGUAGE_TH:
		fertCatName = fertCatModel.FertilizerCatTH
	}
	return fertCatModel, fertCatName
}

func (ln Ln) GetUserNameer(uid string) (model_databases.Users, string) {
	var userModel model_databases.Users
	var userName string

	sql := fmt.Sprintf("SELECT * FROM %s WHERE status_id = '%s' AND uid = '%s'",
		config.DB_USERS, config.STATUS_ACTIVE, uid)
	err := ln.Db.Raw(sql).Scan(&userModel).Error
	if err != nil {
		log.Print(err)
	}

	userName = userModel.Username
	return userModel, userName
}

func (ln Ln) GetPlantOverviewFavoriteer(status, uid, language string, offset int) ([]model_services.ForPlantItem, int, int) {
	var joinArray []model_services.JoinFormulaPlantAndPlant
	var formulaPlant model_databases.FormulaPlant
	var plantOverview model_services.ForPlantItem
	var plantOverviewArray []model_services.ForPlantItem
	var currentOffset int
	var total int
	var found bool
	var countryMap map[string]string
	var provinceMap map[string]string
	var plantTypeMap map[string]string
	var userMap map[string]string

	if uid == "" {
		return nil, offset, 0
	}

	countryMap = make(map[string]string)
	provinceMap = make(map[string]string)
	plantTypeMap = make(map[string]string)
	userMap = make(map[string]string)

	_, formulaPlantList, _ := IntFormulaPlant.GetFavoriteFormulaPlanter(ln, config.STATUS_ACTIVE, uid)
	sqlIn := "('" + strings.Join(formulaPlantList, "','") + "')"
	sql := fmt.Sprintf("SELECT * FROM %s INNER JOIN %s ON %s.plant_id = %s.plant_id WHERE %s.status_id = '%s' AND %s.formula_plant_id IN %s OFFSET %d LIMIT 100",
		config.DB_FORMULA_PLANT, config.DB_PLANT, config.DB_FORMULA_PLANT, config.DB_PLANT, config.DB_FORMULA_PLANT, status, config.DB_FORMULA_PLANT, sqlIn, offset)
	fmt.Println(sql)
	err := ln.Db.Raw(sql).Scan(&joinArray).Error
	if err != nil {
		log.Print(err)
	}

	for _, join := range joinArray {
		mapstructure.Decode(join, &plantOverview)
		mapstructure.Decode(join, &formulaPlant)
		plantOverview.RateScore, plantOverview.RatePeople = IntFormulaPlant.GetRateScoreAndPeopleer(ln, formulaPlant)
		plantOverview.IsFavorite = true

		//Get Country name
		plantOverview.CountryName, found = countryMap[plantOverview.CountryId.UUID.String()]
		if !found {
			_, plantOverview.CountryName = IntFormulaPlant.GetCountryNameer(ln, plantOverview.CountryId.UUID.String(), language)
			countryMap[plantOverview.CountryId.UUID.String()] = plantOverview.CountryName
		}

		//Get Country name
		plantOverview.ProvinceName, found = provinceMap[plantOverview.ProvinceId.UUID.String()]
		if !found {
			_, plantOverview.ProvinceName = IntFormulaPlant.GetProvinceNameer(ln, plantOverview.ProvinceId.UUID.String(), language)
			provinceMap[plantOverview.ProvinceId.UUID.String()] = plantOverview.ProvinceName
		}

		//Get Plant Type name
		plantOverview.PlantTypeName, found = plantTypeMap[plantOverview.PlantTypeId.UUID.String()]
		if !found {
			_, plantOverview.PlantTypeName = IntFormulaPlant.GetPlantTypeNameer(ln, plantOverview.PlantTypeId.UUID.String(), language)
			plantTypeMap[plantOverview.PlantTypeId.UUID.String()] = plantOverview.PlantTypeName
		}

		//Get User name
		plantOverview.Username, found = userMap[plantOverview.Uid.UUID.String()]
		if !found {
			_, plantOverview.Username = IntFormulaPlant.GetUserNameer(ln, plantOverview.Uid.UUID.String())
			plantTypeMap[plantOverview.Uid.UUID.String()] = plantOverview.Username
		}

		plantOverviewArray = append(plantOverviewArray, plantOverview)
	}
	total = len(plantOverviewArray)
	currentOffset = offset + total
	return plantOverviewArray, currentOffset, total
}

func (ln Ln) GetMyPlantOverviewer(status, uid, language string, offset int) ([]model_services.ForPlantItem, int, int) {
	var formulaPlant model_databases.FormulaPlant
	var plantOverview model_services.ForPlantItem
	var joinArray []model_services.JoinFormulaPlantAndPlant
	var plantOverviewArray []model_services.ForPlantItem
	var currentOffset int
	var total int
	var found bool
	var countryMap map[string]string
	var provinceMap map[string]string
	var plantTypeMap map[string]string
	var formulaPlantFavMap map[string]bool
	var userMap map[string]string

	if uid == "" {
		return nil, offset, 0
	}

	countryMap = make(map[string]string)
	provinceMap = make(map[string]string)
	plantTypeMap = make(map[string]string)
	userMap = make(map[string]string)

	_, _, formulaPlantFavMap = IntFormulaPlant.GetFavoriteFormulaPlanter(ln, config.STATUS_ACTIVE, uid)
	sql := fmt.Sprintf("SELECT * FROM %s INNER JOIN %s ON %s.plant_id = %s.plant_id WHERE %s.status_id = '%s' AND %s.uid = '%s' OFFSET %d LIMIT 100",
		config.DB_FORMULA_PLANT, config.DB_PLANT, config.DB_FORMULA_PLANT, config.DB_PLANT, config.DB_FORMULA_PLANT, status, config.DB_FORMULA_PLANT, uid, offset)
	fmt.Println(sql)
	err := ln.Db.Raw(sql).Scan(&joinArray).Error
	if err != nil {
		log.Print(err)
	}

	for _, join := range joinArray {
		mapstructure.Decode(join, &formulaPlant)
		mapstructure.Decode(join, &plantOverview)
		plantOverview.RateScore, plantOverview.RatePeople = IntFormulaPlant.GetRateScoreAndPeopleer(ln, formulaPlant)

		//Get Country name
		plantOverview.CountryName, found = countryMap[plantOverview.CountryId.UUID.String()]
		if !found {
			_, plantOverview.CountryName = IntFormulaPlant.GetCountryNameer(ln, plantOverview.CountryId.UUID.String(), language)
			countryMap[plantOverview.CountryId.UUID.String()] = plantOverview.CountryName
		}

		//Get Country name
		plantOverview.ProvinceName, found = provinceMap[plantOverview.ProvinceId.UUID.String()]
		if !found {
			_, plantOverview.ProvinceName = IntFormulaPlant.GetProvinceNameer(ln, plantOverview.ProvinceId.UUID.String(), language)
			provinceMap[plantOverview.ProvinceId.UUID.String()] = plantOverview.ProvinceName
		}

		//Get Plant Type name
		plantOverview.PlantTypeName, found = plantTypeMap[plantOverview.PlantTypeId.UUID.String()]
		if !found {
			_, plantOverview.PlantTypeName = IntFormulaPlant.GetPlantTypeNameer(ln, plantOverview.PlantTypeId.UUID.String(), language)
			plantTypeMap[plantOverview.PlantTypeId.UUID.String()] = plantOverview.PlantTypeName
		}

		//Check Favorite
		plantOverview.IsFavorite, found = formulaPlantFavMap[plantOverview.Uid.UUID.String()]
		if !found {
			plantOverview.IsFavorite = false
		}

		//Get User name
		plantOverview.Username, found = userMap[plantOverview.Uid.UUID.String()]
		if !found {
			_, plantOverview.Username = IntFormulaPlant.GetUserNameer(ln, plantOverview.Uid.UUID.String())
			plantTypeMap[plantOverview.Uid.UUID.String()] = plantOverview.Username
		}

		plantOverviewArray = append(plantOverviewArray, plantOverview)
	}

	total = len(plantOverviewArray)
	currentOffset = offset + total

	return plantOverviewArray, currentOffset, total
}

func (ln Ln) GetPlantOverviewByPlanter(status, uid, plantId, language string, offset int) ([]model_services.ForPlantItem, int, int) {
	var formulaPlant model_databases.FormulaPlant
	var joinArray []model_services.JoinFormulaPlantAndPlant
	var plantOverview model_services.ForPlantItem
	var plantOverviewArray []model_services.ForPlantItem
	var currentOffset int
	var total int
	var found bool
	var countryMap map[string]string
	var provinceMap map[string]string
	var plantTypeMap map[string]string
	var formulaPlantFavMap map[string]bool
	var userMap map[string]string

	if uid == "" && plantId == "" {
		return nil, offset, 0
	}

	countryMap = make(map[string]string)
	provinceMap = make(map[string]string)
	plantTypeMap = make(map[string]string)
	userMap = make(map[string]string)

	_, _, formulaPlantFavMap = IntFormulaPlant.GetFavoriteFormulaPlanter(ln, config.STATUS_ACTIVE, uid)
	sql := fmt.Sprintf("SELECT * FROM %s INNER JOIN %s ON %s.plant_id = %s.plant_id WHERE %s.status_id = '%s' AND %s.plant_id = '%s' OFFSET %d LIMIT 100",
		config.DB_FORMULA_PLANT, config.DB_PLANT, config.DB_FORMULA_PLANT, config.DB_PLANT, config.DB_FORMULA_PLANT, status, config.DB_FORMULA_PLANT, plantId, offset)
	fmt.Println(sql)
	err := ln.Db.Raw(sql).Scan(&joinArray).Error
	if err != nil {
		log.Print(err)
	}

	for _, join := range joinArray {
		mapstructure.Decode(join, &plantOverview)
		mapstructure.Decode(join, &formulaPlant)
		plantOverview.RateScore, plantOverview.RatePeople = IntFormulaPlant.GetRateScoreAndPeopleer(ln, formulaPlant)

		//Get Country name
		plantOverview.CountryName, found = countryMap[plantOverview.CountryId.UUID.String()]
		if !found {
			_, plantOverview.CountryName = IntFormulaPlant.GetCountryNameer(ln, plantOverview.CountryId.UUID.String(), language)
			countryMap[plantOverview.CountryId.UUID.String()] = plantOverview.CountryName
		}

		//Get Country name
		plantOverview.ProvinceName, found = provinceMap[plantOverview.ProvinceId.UUID.String()]
		if !found {
			_, plantOverview.ProvinceName = IntFormulaPlant.GetProvinceNameer(ln, plantOverview.ProvinceId.UUID.String(), language)
			provinceMap[plantOverview.ProvinceId.UUID.String()] = plantOverview.ProvinceName
		}

		//Get Plant Type name
		plantOverview.PlantTypeName, found = plantTypeMap[plantOverview.PlantTypeId.UUID.String()]
		if !found {
			_, plantOverview.PlantTypeName = IntFormulaPlant.GetPlantTypeNameer(ln, plantOverview.PlantTypeId.UUID.String(), language)
			plantTypeMap[plantOverview.PlantTypeId.UUID.String()] = plantOverview.PlantTypeName
		}

		//Check Favorite
		plantOverview.IsFavorite, found = formulaPlantFavMap[plantOverview.Uid.UUID.String()]
		if !found {
			plantOverview.IsFavorite = false
		}

		//Get User name
		plantOverview.Username, found = userMap[plantOverview.Uid.UUID.String()]
		if !found {
			_, plantOverview.Username = IntFormulaPlant.GetUserNameer(ln, plantOverview.Uid.UUID.String())
			plantTypeMap[plantOverview.Uid.UUID.String()] = plantOverview.Username
		}

		plantOverviewArray = append(plantOverviewArray, plantOverview)
	}

	total = len(plantOverviewArray)
	currentOffset = offset + total

	return plantOverviewArray, currentOffset, total
}

func (ln Ln) GetFertilizerRatioRelateer(status, formulaPlantId, language string) ([]model_services.ForPlantFert, int) {
	var joinArray []model_services.JoinFertilizerAndPlant
	var plantFert model_services.ForPlantFert
	var plantFertArray []model_services.ForPlantFert
	var total int
	var found bool
	var fertCatMap map[string]string

	if formulaPlantId == "" {
		return nil, 0
	}

	fertCatMap = make(map[string]string)
	sql := fmt.Sprintf("SELECT * FROM %s INNER JOIN %s ON %s.fertilizer_id = %s.fertilizer_id WHERE %s.status_id = '%s' AND %s.formula_plant_id = '%s'",
		config.DB_FERTILIZER, config.DB_TRANS_FERTILIZER_RATIO, config.DB_FERTILIZER, config.DB_TRANS_FERTILIZER_RATIO, config.DB_FERTILIZER, status, config.DB_TRANS_FERTILIZER_RATIO, formulaPlantId)
	fmt.Println(sql)
	err := ln.Db.Raw(sql).Scan(&joinArray).Error
	if err != nil {
		log.Print(err)
	}

	for _, join := range joinArray {
		mapstructure.Decode(join, &plantFert)
		switch language {
		case config.LANGUAGE_EN:
			plantFert.FertilizerName = join.FertilizerEN
		case config.LANGUAGE_TH:
			plantFert.FertilizerName = join.FertilizerTH
		}

		//Get Fertilizer category name
		plantFert.FertCatName, found = fertCatMap[plantFert.FertCatId.UUID.String()]
		if !found {
			_, plantFert.FertCatName = IntFormulaPlant.GetFertCatNameer(ln, plantFert.FertCatId.UUID.String(), language)
			fertCatMap[plantFert.FertCatId.UUID.String()] = plantFert.FertCatName
		}

		plantFertArray = append(plantFertArray, plantFert)
	}

	total = len(plantFertArray)

	return plantFertArray, total
}

func (ln Ln) GetSensorValueRecRelateer(status, formulaPlantId, language string) ([]model_services.ForPlantSensor, int) {
	var joinArray []model_services.JoinSensorTypeAndTrans
	var plantSensor model_services.ForPlantSensor
	var plantSensorArray []model_services.ForPlantSensor
	var total int

	if formulaPlantId == "" {
		return nil, 0
	}

	sql := fmt.Sprintf("SELECT * FROM %s INNER JOIN %s ON %s.sensor_type_id= %s.sensor_type_id WHERE %s.status_id = '%s' AND %s.formula_plant_id = '%s'",
		config.DB_SENSOR_TYPE, config.DB_TRANS_SENSOR_VALUE_REC, config.DB_SENSOR_TYPE, config.DB_TRANS_SENSOR_VALUE_REC, config.DB_SENSOR_TYPE, status, config.DB_TRANS_SENSOR_VALUE_REC, formulaPlantId)
	fmt.Println(sql)
	err := ln.Db.Raw(sql).Scan(&joinArray).Error
	if err != nil {
		log.Print(err)
	}
	for _, join := range joinArray {
		mapstructure.Decode(join, &plantSensor)
		switch language {
			case config.LANGUAGE_EN:
				plantSensor.SensorTypeName = join.SensorTypeNameEN
			case config.LANGUAGE_TH:
				plantSensor.SensorTypeName = join.SensorTypeNameTH
		}
		plantSensorArray = append(plantSensorArray, plantSensor)
	}
	total = len(plantSensorArray)

	return plantSensorArray, total
}

func (ln Ln) GetFormulaPlantDetailer(status, formulaPlantId, language string) model_services.ForPlantFormula {
	var formulaPlantModel model_databases.FormulaPlant
	var formula model_services.ForPlantFormula

	if formulaPlantId == ""{
		return formula
	}

	condition := fmt.Sprintf("SELECT * FROM %s WHERE status_id = '%s' AND formula_plant_id = '%s'",
		config.DB_FORMULA_PLANT, status, formulaPlantId)
	fmt.Println(condition)
	err := ln.Db.Raw(condition).Scan(&formulaPlantModel).Error
	if err != nil {
		log.Print(err)
	}
	mapstructure.Decode(formulaPlantModel, &formula)

	_, formula.Username = IntFormulaPlant.GetUserNameer(ln, formula.Uid.UUID.String())

	formula.SensorList, _ = IntFormulaPlant.GetSensorValueRecRelateer(ln, config.STATUS_ACTIVE, formula.FormulaPlantId.UUID.String(), language)

	formula.FertList, _ = IntFormulaPlant.GetFertilizerRatioRelateer(ln, config.STATUS_ACTIVE, formula.FormulaPlantId.UUID.String(), language)

	return formula
}