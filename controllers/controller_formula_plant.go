package controllers

import (
	"LN-BackEND/config"
	"LN-BackEND/models/model_databases"
	"LN-BackEND/models/model_services"
	"LN-BackEND/utility"
	"database/sql"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"strings"
)

func GetPlantCategoryList(db *sql.DB, status string, language string) []model_services.ForPlantCatList {
	var plantType model_databases.PlantType
	var catList model_services.ForPlantCatList
	var catListArray []model_services.ForPlantCatList
	condition := fmt.Sprintf(" status_id = '%s'", status)
	rows := utility.SelectData(db, "*", config.DB_PLANT_TYPE, condition, "", "", "plant_type_en ASC", 0, 0 )
	defer rows.Close()
	for rows.Next(){
		rows.Scan(&plantType.PlantTypeId ,
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

func GetPlantCategoryItem(db *sql.DB, status string, plantTypeId string, language string, offset int) ([]model_services.ForPlantCat, int) {
	var plantType model_databases.PlantType
	var plant model_databases.Plant
	var plantCat model_services.ForPlantCat
	var plantCatArray []model_services.ForPlantCat
	var currentOffset int
	condition := fmt.Sprintf(" %s.status_id = '%s'", config.DB_PLANT, status)
	if plantTypeId != "" {
		condPlantType := fmt.Sprintf(" AND plant_type_id = '%s'", plantTypeId)
		condition = condition + condPlantType
	}
	joinKey := fmt.Sprintf(" %s.plant_type_id = %s.plant_type_id", config.DB_PLANT, config.DB_PLANT_TYPE)
	rows := utility.SelectData(db, "*", config.DB_PLANT, condition, config.DB_PLANT_TYPE, joinKey, "", offset, 100 )
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
			plantCat.PlantTypeName = plant.PlantNameEN
			plantCat.PlantTypeName = plant.PlantDescEN
		case config.LANGUAGE_TH:
			plantCat.PlantTypeName = plantType.PlantTypeTH
			plantCat.PlantTypeName = plant.PlantNameTH
			plantCat.PlantTypeName = plant.PlantDescTH
		}
		plantCatArray = append(plantCatArray, plantCat)
	}
	currentOffset = offset + len(plantCatArray)
	return plantCatArray, currentOffset
}

func GetFavoriteFormulaPlant(db *sql.DB, status string, uid string, language string) ([]model_databases.FavoritePlant, []string) {
	var favPlant model_databases.FavoritePlant
	var favPlantArray []model_databases.FavoritePlant
	var formulaPlantList []string

	if uid == "" {
		return nil, nil
	}
	condition := fmt.Sprintf(" status_id = '%s' AND uid = '%s' ", status, uid)
	rows := utility.SelectData(db, "*", config.DB_FAVORITE_PLANT, condition, "", "", "change_date ASC", 0, 0 )
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

func GetRateScoreAndPeople(formulaPlant model_databases.FormulaPlant) (float32, int) {
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

func GetCountryName(db *sql.DB, status string, countryId string, language string) string {
	var country *model_databases.Country
	var countryName string

	condition := fmt.Sprintf(" status_id = '%s' AND country_id = '%s' ", status, countryId)
	utility.SelectSingleData(db, "*", config.DB_COUNTRY, condition, "", "", country)

	switch language {
	case config.LANGUAGE_EN:
		countryName = country.CountryEN
	case config.LANGUAGE_TH:
		countryName = country.CountryTH
	}
	return countryName
}

func GetPlantOverviewFavorite(db *sql.DB, status string, uid string, language string, offset int) ([]model_services.ForPlantItem, int) {
	var plantType model_databases.PlantType
	var formulaPlant model_databases.FormulaPlant
	var plantOverview model_services.ForPlantItem
	var plant model_databases.Plant
	var plantOverviewArray []model_services.ForPlantItem
	var currentOffset int

	if uid == "" {
		return nil, offset
	}

	_, formulaPlantList := GetFavoriteFormulaPlant(db, config.STATUS_ACTIVE, uid, config.LANGUAGE_EN)
	sqlIn := "('" + strings.Join(formulaPlantList, "','") + "')"
	condition := fmt.Sprintf(" %s.status_id = '%s' AND %s.formula_plant_id IN %s", config.DB_FORMULA_PLANT, status, config.DB_FORMULA_PLANT, sqlIn)
	joinKey := fmt.Sprintf(" %s.plant_id = %s.plant_id", config.DB_FORMULA_PLANT, config.DB_PLANT)
	rows := utility.SelectData(db, "*", config.DB_FORMULA_PLANT, condition, config.DB_PLANT, joinKey, "", offset, 100 )

	defer rows.Close()
	for rows.Next() {
		rows.Scan(
			&formulaPlant.FormulaPlantId,
			&formulaPlant.FormulaName,
			&formulaPlant.FormulaDesc,
			&formulaPlant.PeopleUsed,
			&formulaPlant.Recommend1,
			&formulaPlant.Recommend2,
			&formulaPlant.Recommend3,
			&formulaPlant.Recommend4,
			&formulaPlant.Recommend5,
			&formulaPlant.CreateDate,
			&formulaPlant.ChangeDate,
			&formulaPlant.PlantId,
			&formulaPlant.StatusId,
			&formulaPlant.ProvinceId,
			&formulaPlant.CountryId,
			&formulaPlant.IsPublic,
			&formulaPlant.Uid,
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
			//&plantType.PlantTypeId,
			//&plantType.PlantTypeEN,
			//&plantType.PlantTypeTH,
			//&plantType.CreateDate,
			//&plantType.ChangeDate,
			//&plantType.StatusId ,
		)
		mapstructure.Decode(plant, &plantOverview)
		mapstructure.Decode(formulaPlant, &plantOverview)
		fmt.Println(formulaPlant.Uid)
		plantOverview.RateScore, plantOverview.RatePeople = GetRateScoreAndPeople(formulaPlant)
		switch language {
		case config.LANGUAGE_EN:
			plantOverview.PlantTypeName = plantType.PlantTypeEN
		case config.LANGUAGE_TH:
			plantOverview.PlantTypeName = plantType.PlantTypeTH
		}
		plantOverview.IsFavorite = true
		fmt.Printf("%+v\n", plantOverview)
		plantOverviewArray = append(plantOverviewArray, plantOverview)
	}
	currentOffset = offset + len(plantOverviewArray)
	return plantOverviewArray, currentOffset
}