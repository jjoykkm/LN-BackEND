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
	GetPlantCategoryLister(status, language string) []model_services.ForPlantCatList
	GetPlantCategoryItemer(status, plantTypeId, language string, offset int) ([]model_services.ForPlantCat, int)
	GetFavoriteFormulaPlanter(status, uid string) ([]model_databases.FavoritePlant, []string)
	GetMyFormulaPlanter(status, uid string) ([]model_databases.FavoritePlant, []string)
	GetRateScoreAndPeopleer(formulaPlant model_databases.FormulaPlant) (float32, int)
	GetCountryNameer(countryId, language string) string
	GetProvinceNameer(provinceId, language string) string
	GetPlantTypeNameer(plantTypeId, language string) string
	GetPlantOverviewFavoriteer(status, uid, language string, offset int) ([]model_services.ForPlantItem, int)
	GetMyPlantOverview(status, uid, language string, offset int) ([]model_services.ForPlantItem, int)
}

func (ln Ln) GetPlantCategoryLister(status, language string) []model_services.ForPlantCatList {
	var plantType model_databases.PlantType
	var catList model_services.ForPlantCatList
	var catListArray []model_services.ForPlantCatList

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
	return catListArray
}

func (ln Ln) GetPlantCategoryItemer(status, plantTypeId, language string, offset int) ([]model_services.ForPlantCat, int) {
	var plantType model_databases.PlantType
	var plant model_databases.Plant
	var plantCat model_services.ForPlantCat
	var plantCatArray []model_services.ForPlantCat
	var currentOffset int

	sql := fmt.Sprintf("SELECT * FROM %s INNER JOIN %s ON %s.plant_type_id = %s.plant_type_id WHERE %s.status_id = '%s' OFFSET %d LIMIT 100",
		config.DB_PLANT, config.DB_PLANT_TYPE, config.DB_PLANT, config.DB_PLANT_TYPE, config.DB_PLANT, status, offset)
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
		plantCatArray = append(plantCatArray, plantCat)
	}
	currentOffset = offset + len(plantCatArray)
	return plantCatArray, currentOffset
}

func (ln Ln) GetFavoriteFormulaPlanter(status, uid string) ([]model_databases.FavoritePlant, []string) {
	var favPlant model_databases.FavoritePlant
	var favPlantArray []model_databases.FavoritePlant
	var formulaPlantList []string

	if uid == "" {
		return nil, nil
	}

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
	}
	return favPlantArray, formulaPlantList
}

func (ln Ln) GetMyFormulaPlanter(status, uid string) ([]model_databases.FavoritePlant, []string) {
	var favPlant model_databases.FavoritePlant
	var favPlantArray []model_databases.FavoritePlant
	var formulaPlantList []string

	if uid == "" {
		return nil, nil
	}

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
	}
	return favPlantArray, formulaPlantList
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

func (ln Ln) GetCountryNameer(countryId, language string) string {
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
	}
	switch language {
	case config.LANGUAGE_EN:
		countryName = countryModel.CountryEN
	case config.LANGUAGE_TH:
		countryName = countryModel.CountryTH
	}
	return countryName
}

func (ln Ln) GetProvinceNameer(provinceId, language string) string {
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
	}
	switch language {
	case config.LANGUAGE_EN:
		provinceName = provinceModel.ProvinceEN
	case config.LANGUAGE_TH:
		provinceName = provinceModel.ProvinceTH
	}
	return provinceName
}

func (ln Ln) GetPlantTypeNameer(plantTypeId, language string) string {
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
	}
	switch language {
	case config.LANGUAGE_EN:
		plantTypeName = plantTypeModel.PlantTypeEN
	case config.LANGUAGE_TH:
		plantTypeName = plantTypeModel.PlantTypeTH
	}
	return plantTypeName
}

func (ln Ln) GetPlantOverviewFavoriteer(status, uid, language string, offset int) ([]model_services.ForPlantItem, int) {
	var plantType model_databases.PlantType
	var formulaPlant model_databases.FormulaPlant
	var plantOverview model_services.ForPlantItem
	var plant model_databases.Plant
	var plantOverviewArray []model_services.ForPlantItem
	var currentOffset int
	var found bool
	var countryMap map[string]string
	var provinceMap map[string]string
	var plantTypeMap map[string]string

	if uid == "" {
		return nil, offset
	}

	countryMap = make(map[string]string)
	provinceMap = make(map[string]string)
	plantTypeMap = make(map[string]string)

	_, formulaPlantList := IntFormulaPlant.GetFavoriteFormulaPlanter(ln, config.STATUS_ACTIVE, uid)
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
			plantOverview.CountryName = IntFormulaPlant.GetCountryNameer(ln, plantOverview.CountryId.UUID.String(), language)
			countryMap[plantOverview.CountryId.UUID.String()] = plantOverview.CountryName
		}

		//Get Country name
		plantOverview.ProvinceName, found = provinceMap[plantOverview.ProvinceId.UUID.String()]
		if !found {
			plantOverview.ProvinceName = IntFormulaPlant.GetProvinceNameer(ln, plantOverview.ProvinceId.UUID.String(), language)
			provinceMap[plantOverview.ProvinceId.UUID.String()] = plantOverview.ProvinceName
		}

		//Get Plant Type name
		plantOverview.PlantTypeName, found = plantTypeMap[plantOverview.PlantTypeId.UUID.String()]
		if !found {
			plantOverview.PlantTypeName = IntFormulaPlant.GetPlantTypeNameer(ln, plantOverview.PlantTypeId.UUID.String(), language)
			plantTypeMap[plantOverview.PlantTypeId.UUID.String()] = plantOverview.PlantTypeName
		}

		plantOverviewArray = append(plantOverviewArray, plantOverview)
	}
	currentOffset = offset + len(plantOverviewArray)
	return plantOverviewArray, currentOffset
}

func (ln Ln) GetMyPlantOverview(status, uid, language string, offset int) ([]model_services.ForPlantItem, int) {
	var plantType model_databases.PlantType
	var formulaPlant model_databases.FormulaPlant
	var plantOverview model_services.ForPlantItem
	var plant model_databases.Plant
	var plantOverviewArray []model_services.ForPlantItem
	var currentOffset int
	var found bool
	var countryMap map[string]string
	var provinceMap map[string]string
	var plantTypeMap map[string]string

	if uid == "" {
		return nil, offset
	}

	countryMap = make(map[string]string)
	provinceMap = make(map[string]string)
	plantTypeMap = make(map[string]string)

	_, formulaPlantList := IntFormulaPlant.GetFavoriteFormulaPlanter(ln, config.STATUS_ACTIVE, uid)
	_, myFormulaPlantList := IntFormulaPlant.GetMyFormulaPlanter(ln, config.STATUS_ACTIVE, uid)
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
			plantOverview.CountryName = IntFormulaPlant.GetCountryNameer(ln, plantOverview.CountryId.UUID.String(), language)
			countryMap[plantOverview.CountryId.UUID.String()] = plantOverview.CountryName
		}

		//Get Country name
		plantOverview.ProvinceName, found = provinceMap[plantOverview.ProvinceId.UUID.String()]
		if !found {
			plantOverview.ProvinceName = IntFormulaPlant.GetProvinceNameer(ln, plantOverview.ProvinceId.UUID.String(), language)
			provinceMap[plantOverview.ProvinceId.UUID.String()] = plantOverview.ProvinceName
		}

		//Get Plant Type name
		plantOverview.PlantTypeName, found = plantTypeMap[plantOverview.PlantTypeId.UUID.String()]
		if !found {
			plantOverview.PlantTypeName = IntFormulaPlant.GetPlantTypeNameer(ln, plantOverview.PlantTypeId.UUID.String(), language)
			plantTypeMap[plantOverview.PlantTypeId.UUID.String()] = plantOverview.PlantTypeName
		}

		plantOverviewArray = append(plantOverviewArray, plantOverview)
	}
	currentOffset = offset + len(plantOverviewArray)
	return plantOverviewArray, currentOffset
}