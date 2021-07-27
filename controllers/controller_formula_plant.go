package controllers

import (
	"LN-BackEND/config"
	"LN-BackEND/models/model_databases"
	"LN-BackEND/models/model_services"
	"LN-BackEND/utility"
	"database/sql"
	"fmt"
	"github.com/mitchellh/mapstructure"
)

func GetPlantCategoryList(db *sql.DB, status string, language string) []model_services.ForPlantCatList {
	var plantType model_databases.PlantType
	var catList model_services.ForPlantCatList
	var catListArray []model_services.ForPlantCatList
	condition := fmt.Sprintf(" status_id = '%s'", status)
	rows := utility.SelectData(db, "*", config.DbPlantType, condition, "", "", "plant_type_en ASC", 0, 0 )
	defer rows.Close()
	for rows.Next(){
		rows.Scan(&plantType.PlantTypeId ,
			&plantType.PlantTypeEN ,
			&plantType.PlantTypeTH ,
			&plantType.CreateDate ,
			&plantType.ChangeDate ,
			&plantType.StatusId )
		mapstructure.Decode(plantType, &catList)
		switch language {
		case config.LanguageEN:
			catList.PlantTypeName = plantType.PlantTypeEN
		case config.LanguageTH:
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
	condition := fmt.Sprintf("plant.status_id = '%s'", status)
	if plantTypeId != "" {
		condPlantType := fmt.Sprintf(" AND plant_type_id = '%s'", plantTypeId)
		condition = condition + condPlantType
	}
	rows := utility.SelectData(db, "*", config.DbPlant, condition, config.DbPlantType, "plant.plant_type_id = plant_type.plant_type_id", "", offset, 100 )
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
			&plantType.StatusId )
		mapstructure.Decode(plantType, &plantCat)
		mapstructure.Decode(plant, &plantCat)
		switch language {
		case config.LanguageEN:
			plantCat.PlantTypeName = plantType.PlantTypeEN
			plantCat.PlantTypeName = plant.PlantNameEN
			plantCat.PlantTypeName = plant.PlantDescEN
		case config.LanguageTH:
			plantCat.PlantTypeName = plantType.PlantTypeTH
			plantCat.PlantTypeName = plant.PlantNameTH
			plantCat.PlantTypeName = plant.PlantDescTH
		}
		plantCatArray = append(plantCatArray, plantCat)
	}
	currentOffset = offset + len(plantCatArray)
	return plantCatArray, currentOffset
}