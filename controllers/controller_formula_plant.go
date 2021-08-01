package controllers

import (
	"LN-BackEND/config"
	"LN-BackEND/models/model_databases"
	"LN-BackEND/models/model_services"
	"database/sql"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"log"
	"strings"
)
/*-------------------------------------------------------------------------------------------*/
//                                 STRUCTURE
/*-------------------------------------------------------------------------------------------*/
type Ln struct {
	Db			*sql.DB
}

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
	GetCountTableer(status, tableName, field, condition string) int
	GetPlantOverviewFavoriteer(status, uid, language string, offset int) ([]model_services.ForPlantItem, int, int)
	GetMyPlantOverviewer(status, uid, language string, offset int) ([]model_services.ForPlantItem, int, int)
	GetPlantOverviewByPlanter(status, uid, plantId, language string, offset int) ([]model_services.ForPlantItem, int, int)
	GetFertilizerRatioRelate(status, formulaPlantId, language string) ([]model_services.ForPlantFert, int)
	GetSensorValueRecRelate(status, formulaPlantId, language string) ([]model_services.ForPlantSensor, int)
	GetFormulaPlantDetailer(status, formulaPlantId, language string) model_services.ForPlantFormula
}

func (ln Ln) GetPlantCategoryLister(status, language string) ([]model_services.ForPlantCatList, int) {
	var plantType model_databases.PlantType
	var catList model_services.ForPlantCatList
	var catListArray []model_services.ForPlantCatList
	var total int

	sql := fmt.Sprintf("SELECT * FROM %s WHERE status_id = '%s' ORDER BY plant_type_en ASC", config.DB_PLANT_TYPE, status)
	rows, err := ln.Db.Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next(){
		rows.Scan(
			&plantType.PlantTypeId ,
			&plantType.PlantTypeEN ,
			&plantType.PlantTypeTH ,
			&plantType.CreateDate ,
			&plantType.ChangeDate ,
			&plantType.StatusId ,
		)
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
	var plantType model_databases.PlantType
	var plant model_databases.Plant
	var plantCat model_services.ForPlantCat
	var plantCatArray []model_services.ForPlantCat
	var currentOffset int
	var total int
	var sqlScopePT string

	if plantTypeId != "" {
		sqlScopePT = fmt.Sprintf("AND %s.plant_type_id = '%s'",config.DB_PLANT, plantTypeId)
	}

	sql := fmt.Sprintf("SELECT * FROM %s INNER JOIN %s ON %s.plant_type_id = %s.plant_type_id WHERE %s.status_id = '%s' %s OFFSET %d LIMIT 100",
		config.DB_PLANT, config.DB_PLANT_TYPE, config.DB_PLANT, config.DB_PLANT_TYPE, config.DB_PLANT, status, sqlScopePT, offset)
	fmt.Println(sql)
	rows, err := ln.Db.Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next(){
		rows.Scan(
			&plant.PlantId ,
			&plant.PlantNameEN ,
			&plant.PlantNameTH ,
			&plant.PlantDescEN ,
			&plant.PlantDescTH ,
			&plant.CreateDate ,
			&plant.ChangeDate ,
			&plant.StatusId ,
			&plant.PlantTypeId ,
			&plant.TotalItem ,
			&plantType.PlantTypeId ,
			&plantType.PlantTypeEN ,
			&plantType.PlantTypeTH ,
			&plantType.CreateDate ,
			&plantType.ChangeDate ,
			&plantType.StatusId ,
		)
		mapstructure.Decode(plantType, &plantCat)
		mapstructure.Decode(plant, &plantCat)
		switch language {
		case config.LANGUAGE_EN:
			plantCat.PlantTypeName = plantType.PlantTypeEN
			plantCat.PlantName = plant.PlantNameEN
			plantCat.PlantDesc = plant.PlantDescEN
		case config.LANGUAGE_TH:
			plantCat.PlantTypeName = plantType.PlantTypeTH
			plantCat.PlantName = plant.PlantNameTH
			plantCat.PlantDesc = plant.PlantDescTH
		}
		cond := fmt.Sprintf("plant_id = '%s'", plantCat.PlantId.UUID.String())
		plantCat.TotalItem = IntFormulaPlant.GetCountTableer(ln, config.STATUS_ACTIVE, config.DB_FORMULA_PLANT, "formula_plant_id", cond)
		plantCatArray = append(plantCatArray, plantCat)
	}
	total = len(plantCatArray)
	currentOffset = offset + total
	return plantCatArray, currentOffset, total
}

func (ln Ln) GetFavoriteFormulaPlanter(status, uid string) ([]model_databases.FavoritePlant, []string, map[string]bool) {
	var favPlant model_databases.FavoritePlant
	var favPlantArray []model_databases.FavoritePlant
	var formulaPlantList []string

	if uid == "" {
		return nil, nil, nil
	}

	formulaPlantMap := make(map[string]bool)

	sql := fmt.Sprintf("SELECT * FROM %s WHERE status_id = '%s' AND uid = '%s' ORDER BY change_date ASC",
		config.DB_FAVORITE_PLANT, status, uid)
	fmt.Println(sql)
	rows, err := ln.Db.Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next(){
		rows.Scan(
			&favPlant.Uid ,
			&favPlant.FormulaPlantId ,
			&favPlant.CreateDate ,
			&favPlant.ChangeDate ,
			&favPlant.StatusId ,
		)
		favPlantArray = append(favPlantArray, favPlant)
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

	condition := fmt.Sprintf("SELECT * FROM %s WHERE status_id = $1 AND country_id = $2", config.DB_COUNTRY)
	err := ln.Db.QueryRow(condition, config.STATUS_ACTIVE, countryId).Scan(
		&countryModel.CountryId ,
		&countryModel.CountryEN ,
		&countryModel.CountryTH ,
		&countryModel.CreateDate ,
		&countryModel.ChangeDate ,
		&countryModel.StatusId ,
		)
	if err != nil {
		log.Fatal(err)
		return countryModel, ""
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

	condition := fmt.Sprintf("SELECT * FROM %s WHERE status_id = $1 AND province_id = $2", config.DB_PROVINCE)
	err := ln.Db.QueryRow(condition, config.STATUS_ACTIVE, provinceId).Scan(
		&provinceModel.ProvinceId ,
		&provinceModel.ProvinceEN ,
		&provinceModel.ProvinceTH ,
		&provinceModel.CreateDate ,
		&provinceModel.ChangeDate ,
		&provinceModel.StatusId ,
		&provinceModel.CountryId ,
	)
	if err != nil {
		log.Fatal(err)
		return provinceModel, ""
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

	condition := fmt.Sprintf("SELECT * FROM %s WHERE status_id = $1 AND plant_type_id = $2", config.DB_PLANT_TYPE)
	err := ln.Db.QueryRow(condition, config.STATUS_ACTIVE, plantTypeId).Scan(
		&plantTypeModel.PlantTypeId ,
		&plantTypeModel.PlantTypeEN ,
		&plantTypeModel.PlantTypeTH ,
		&plantTypeModel.CreateDate ,
		&plantTypeModel.ChangeDate ,
		&plantTypeModel.StatusId  ,
	)
	if err != nil {
		log.Fatal(err)
		return plantTypeModel, ""
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

	condition := fmt.Sprintf("SELECT * FROM %s WHERE status_id = $1 AND fertilizer_cat_id = $2", config.DB_FERTILIZER_CAT)
	err := ln.Db.QueryRow(condition, config.STATUS_ACTIVE, fertCatId).Scan(
		&fertCatModel.FertCatId ,
		&fertCatModel.FertCatEN ,
		&fertCatModel.FertCatTH ,
		&fertCatModel.CreateDate ,
		&fertCatModel.ChangeDate ,
		&fertCatModel.StatusId ,
)
	if err != nil {
		log.Fatal(err)
		return fertCatModel, ""
	}
	switch language {
	case config.LANGUAGE_EN:
		fertCatName = fertCatModel.FertCatEN
	case config.LANGUAGE_TH:
		fertCatName = fertCatModel.FertCatTH
	}
	return fertCatModel, fertCatName
}

func (ln Ln) GetUserNameer(uid string) (model_databases.Users, string) {
	var userModel model_databases.Users
	var userName string

	condition := fmt.Sprintf("SELECT username FROM %s WHERE status_id = $1 AND uid = $2", config.DB_USERS)
	fmt.Println(condition)
	err := ln.Db.QueryRow(condition, config.STATUS_ACTIVE, uid).Scan(
		&userModel.Uid ,
		&userModel.Username ,
		&userModel.Password ,
		&userModel.FullName ,
		&userModel.SurName ,
		&userModel.NickName ,
		&userModel.Gender ,
		&userModel.BirthDate ,
		&userModel.MobilePhone ,
		&userModel.Telephone ,
		&userModel.Mail ,
		&userModel.Image ,
		&userModel.CreateDate ,
		&userModel.ChangeDate ,
		&userModel.StatusId ,
		&userModel.UserNo ,
	)
	if err != nil {
		return userModel, ""
	}
	userName = userModel.Username
	return userModel, userName
}

func (ln Ln) GetCountTableer(status, tableName, field, condition string) int {
	var count int

	sql := fmt.Sprintf("SELECT COUNT(%s) FROM %s WHERE status_id = $1 AND %s ", field, tableName, condition)
	err := ln.Db.QueryRow(sql, status).Scan(
		&count ,
	)
	if err != nil {
		log.Fatal(err)
	}
	return count
}

func (ln Ln) GetPlantOverviewFavoriteer(status, uid, language string, offset int) ([]model_services.ForPlantItem, int, int) {
	var plantType model_databases.PlantType
	var formulaPlant model_databases.FormulaPlant
	var plantOverview model_services.ForPlantItem
	var plant model_databases.Plant
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
	rows, err := ln.Db.Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(
			&formulaPlant.FormulaPlantId ,
			&formulaPlant.FormulaName ,
			&formulaPlant.FormulaDesc ,
			&formulaPlant.PeopleUsed ,
			&formulaPlant.Recommend1 ,
			&formulaPlant.Recommend2 ,
			&formulaPlant.Recommend3 ,
			&formulaPlant.Recommend4 ,
			&formulaPlant.Recommend5 ,
			&formulaPlant.CreateDate ,
			&formulaPlant.ChangeDate ,
			&formulaPlant.PlantId ,
			&formulaPlant.StatusId ,
			&formulaPlant.ProvinceId ,
			&formulaPlant.CountryId ,
			&formulaPlant.IsPublic ,
			&formulaPlant.Uid ,
			&plant.PlantId ,
			&plant.PlantNameEN ,
			&plant.PlantNameTH ,
			&plant.PlantDescEN ,
			&plant.PlantDescTH ,
			&plant.CreateDate ,
			&plant.ChangeDate ,
			&plant.StatusId ,
			&plant.PlantTypeId ,
			&plant.TotalItem ,
		)
		mapstructure.Decode(plant, &plantOverview)
		mapstructure.Decode(formulaPlant, &plantOverview)
		plantOverview.RateScore, plantOverview.RatePeople = IntFormulaPlant.GetRateScoreAndPeopleer(ln, formulaPlant)
		switch language {
		case config.LANGUAGE_EN:
			plantOverview.PlantTypeName = plantType.PlantTypeEN
		case config.LANGUAGE_TH:
			plantOverview.PlantTypeName = plantType.PlantTypeTH
		}
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
	var plantType model_databases.PlantType
	var formulaPlant model_databases.FormulaPlant
	var plantOverview model_services.ForPlantItem
	var plant model_databases.Plant
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
	rows, err := ln.Db.Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(
			&formulaPlant.FormulaPlantId ,
			&formulaPlant.FormulaName ,
			&formulaPlant.FormulaDesc ,
			&formulaPlant.PeopleUsed ,
			&formulaPlant.Recommend1 ,
			&formulaPlant.Recommend2 ,
			&formulaPlant.Recommend3 ,
			&formulaPlant.Recommend4 ,
			&formulaPlant.Recommend5 ,
			&formulaPlant.CreateDate ,
			&formulaPlant.ChangeDate ,
			&formulaPlant.PlantId ,
			&formulaPlant.StatusId ,
			&formulaPlant.ProvinceId ,
			&formulaPlant.CountryId ,
			&formulaPlant.IsPublic ,
			&formulaPlant.Uid ,
			&plant.PlantId ,
			&plant.PlantNameEN ,
			&plant.PlantNameTH ,
			&plant.PlantDescEN ,
			&plant.PlantDescTH ,
			&plant.CreateDate ,
			&plant.ChangeDate ,
			&plant.StatusId ,
			&plant.PlantTypeId ,
			&plant.TotalItem ,
		)
		mapstructure.Decode(plant, &plantOverview)
		mapstructure.Decode(formulaPlant, &plantOverview)
		plantOverview.RateScore, plantOverview.RatePeople = IntFormulaPlant.GetRateScoreAndPeopleer(ln, formulaPlant)
		switch language {
		case config.LANGUAGE_EN:
			plantOverview.PlantTypeName = plantType.PlantTypeEN
		case config.LANGUAGE_TH:
			plantOverview.PlantTypeName = plantType.PlantTypeTH
		}

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
	var plantType model_databases.PlantType
	var formulaPlant model_databases.FormulaPlant
	var plantOverview model_services.ForPlantItem
	var plant model_databases.Plant
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
	rows, err := ln.Db.Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(
			&formulaPlant.FormulaPlantId ,
			&formulaPlant.FormulaName ,
			&formulaPlant.FormulaDesc ,
			&formulaPlant.PeopleUsed ,
			&formulaPlant.Recommend1 ,
			&formulaPlant.Recommend2 ,
			&formulaPlant.Recommend3 ,
			&formulaPlant.Recommend4 ,
			&formulaPlant.Recommend5 ,
			&formulaPlant.CreateDate ,
			&formulaPlant.ChangeDate ,
			&formulaPlant.PlantId ,
			&formulaPlant.StatusId ,
			&formulaPlant.ProvinceId ,
			&formulaPlant.CountryId ,
			&formulaPlant.IsPublic ,
			&formulaPlant.Uid ,
			&plant.PlantId ,
			&plant.PlantNameEN ,
			&plant.PlantNameTH ,
			&plant.PlantDescEN ,
			&plant.PlantDescTH ,
			&plant.CreateDate ,
			&plant.ChangeDate ,
			&plant.StatusId ,
			&plant.PlantTypeId ,
			&plant.TotalItem ,
		)
		mapstructure.Decode(plant, &plantOverview)
		mapstructure.Decode(formulaPlant, &plantOverview)
		plantOverview.RateScore, plantOverview.RatePeople = IntFormulaPlant.GetRateScoreAndPeopleer(ln, formulaPlant)
		switch language {
		case config.LANGUAGE_EN:
			plantOverview.PlantTypeName = plantType.PlantTypeEN
		case config.LANGUAGE_TH:
			plantOverview.PlantTypeName = plantType.PlantTypeTH
		}

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

func (ln Ln) GetFertilizerRatioRelate(status, formulaPlantId, language string) ([]model_services.ForPlantFert, int) {
	var fertRatioModel model_databases.TransFertRatio
	var fertModel model_databases.Fertilizer
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
	rows, err := ln.Db.Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(
			&fertModel.FertilizerId ,
			&fertModel.FertilizerEN ,
			&fertModel.FertilizerTH ,
			&fertModel.Nitrogen ,
			&fertModel.Phosphorus ,
			&fertModel.Potassium ,
			&fertModel.CreateDate ,
			&fertModel.ChangeDate ,
			&fertModel.FertCatId ,
			&fertModel.StatusId ,
			&fertRatioModel.FertilizerId ,
			&fertRatioModel.FormulaPlantId ,
			&fertRatioModel.StatusId ,
			&fertRatioModel.Ratio ,
			&fertRatioModel.CreateDate ,
			&fertRatioModel.ChangeDate ,
		)
		mapstructure.Decode(fertModel, &plantFert)
		mapstructure.Decode(fertRatioModel, &plantFert)
		switch language {
		case config.LANGUAGE_EN:
			plantFert.FertilizerName = fertModel.FertilizerEN
		case config.LANGUAGE_TH:
			plantFert.FertilizerName = fertModel.FertilizerTH
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

func (ln Ln) GetSensorValueRecRelate(status, formulaPlantId, language string) ([]model_services.ForPlantSensor, int) {
	var sensorValueModel model_databases.TransSensorValueRec
	var sensorTypeModel model_databases.SensorType
	var plantSensor model_services.ForPlantSensor
	var plantSensorArray []model_services.ForPlantSensor
	var total int

	if formulaPlantId == "" {
		return nil, 0
	}

	sql := fmt.Sprintf("SELECT * FROM %s INNER JOIN %s ON %s.sensor_type_id= %s.sensor_type_id WHERE %s.status_id = '%s' AND %s.formula_plant_id = '%s'",
		config.DB_SENSOR_TYPE, config.DB_TRANS_SENSOR_VALUE_REC, config.DB_SENSOR_TYPE, config.DB_TRANS_SENSOR_VALUE_REC, config.DB_SENSOR_TYPE, status, config.DB_TRANS_SENSOR_VALUE_REC, formulaPlantId)

	fmt.Println(sql)
	rows, err := ln.Db.Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(
			&sensorTypeModel.SensorTypeId ,
			&sensorTypeModel.SensorTypeNameEN ,
			&sensorTypeModel.CreateDate ,
			&sensorTypeModel.ChangeDate ,
			&sensorTypeModel.StatusId ,
			&sensorTypeModel.SensorTypeNameTH ,
			&sensorValueModel.FormulaPlantId ,
			&sensorValueModel.SensorTypeId ,
			&sensorValueModel.StatusId ,
			&sensorValueModel.ValueRec ,
			&sensorValueModel.CreateDate ,
			&sensorValueModel.ChangeDate ,
		)
		mapstructure.Decode(sensorTypeModel, &plantSensor)
		mapstructure.Decode(sensorValueModel, &plantSensor)
		switch language {
		case config.LANGUAGE_EN:
			plantSensor.SensorTypeName = sensorTypeModel.SensorTypeNameEN
		case config.LANGUAGE_TH:
			plantSensor.SensorTypeName = sensorTypeModel.SensorTypeNameTH
		}

		plantSensorArray = append(plantSensorArray, plantSensor)
	}

	total = len(plantSensorArray)

	return plantSensorArray, total
}

func (ln Ln) GetFormulaPlantDetailer(status, formulaPlantId, language string) model_services.ForPlantFormula {
	var formula model_services.ForPlantFormula

	if formulaPlantId == ""{
		return formula
	}

	condition := fmt.Sprintf("SELECT formula_plant_id, formula_name, formula_desc FROM %s WHERE status_id = $1 AND formula_plant_id = $2", config.DB_FORMULA_PLANT)
	fmt.Println(condition)
	err := ln.Db.QueryRow(condition, status, formulaPlantId).Scan(
		&formula.FormulaPlantId ,
		&formula.FormulaName ,
		&formula.FormulaDesc ,
	)
	if err != nil {
		panic(err)
		return formula
	}
	formula.SensorList, _ = IntFormulaPlant.GetSensorValueRecRelate(ln, config.STATUS_ACTIVE, formula.FormulaPlantId.UUID.String(), language)
	formula.FertList, _ = IntFormulaPlant.GetFertilizerRatioRelate(ln, config.STATUS_ACTIVE, formula.FormulaPlantId.UUID.String(), language)
	return formula
}