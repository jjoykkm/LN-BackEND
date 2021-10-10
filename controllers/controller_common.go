package controllers

import (
	"fmt"
	"github.com/jjoykkm/ln-backend/common/config"
	"github.com/jjoykkm/ln-backend/modelsOld/model_databases"
	"log"
)

/*-------------------------------------------------------------------------------------------*/
//                                 INTERFACE
/*-------------------------------------------------------------------------------------------*/
type IntCommon interface {
	GetFarmAreaByFarmId(status, farmId string) ([]model_databases.FarmArea, []string, int)
	GetCountryNameer(countryId, language string) (model_databases.Country, string)
	GetProvinceNameer(provinceId, language string) (model_databases.Province, string)
	GetPlantTypeNameer(plantTypeId, language string) (model_databases.PlantType, string)
	GetFertCatNameer(fertCatId, language string) (model_databases.FertilizerCat, string)
	GetUserNameer(uid string) (model_databases.Users, string)
	GetSensorTypeNameer(sensorTypeId, language string) (model_databases.SensorType, string)
	GetRoleNameer(roleId, language string) (model_databases.Role, string, string)
	GetFarmAreaNameer(farmAreaId string) (model_databases.FarmArea, string)
}

/*-------------------------------------------------------------------------------------------*/
//                                   METHOD
/*-------------------------------------------------------------------------------------------*/
func (ln Ln) GetFarmAreaByFarmId(status, farmId string) ([]model_databases.FarmArea, []string, int) {
	var farmModel []model_databases.FarmArea
	var farmAreaList []string
	var total int

	sql := fmt.Sprintf("SELECT * FROM %s WHERE status_id = '%s' AND farm_id = '%s'",
		config.DB_FARM_AREA, status, farmId)
	fmt.Println(sql)
	err := ln.Db.Raw(sql).Scan(&farmModel).Error
	if err != nil {
		log.Print(err)
	}
	for _, array := range farmModel {
		farmAreaList = append(farmAreaList, array.FarmAreaId.UUID.String())
	}

	total = len(farmModel)
	return farmModel, farmAreaList, total
}

func (ln Ln) GetCountryNameer(countryId, language string) (model_databases.Country, string) {
	var countryModel model_databases.Country
	var countryName string

	sql := fmt.Sprintf("SELECT * FROM %s WHERE status_id = '%s' AND country_id = '%s'",
		config.DB_COUNTRY, config.GetStatus().Active, countryId)
	fmt.Println(sql)
	err := ln.Db.Raw(sql).Scan(&countryModel).Error
	if err != nil {
		log.Print(err)
	}

	switch language {
	case config.GetLanguage().En:
		countryName = countryModel.CountryEN
	case config.GetLanguage().Th:
		countryName = countryModel.CountryTH
	}
	return countryModel, countryName
}

func (ln Ln) GetProvinceNameer(provinceId, language string) (model_databases.Province, string) {
	var provinceModel model_databases.Province
	var provinceName string

	sql := fmt.Sprintf("SELECT * FROM %s WHERE status_id = '%s' AND province_id = '%s'",
		config.DB_PROVINCE, config.GetStatus().Active, provinceId)
	fmt.Println(sql)
	err := ln.Db.Raw(sql).Scan(&provinceModel).Error
	if err != nil {
		log.Print(err)
	}

	switch language {
	case config.GetLanguage().En:
		provinceName = provinceModel.ProvinceEN
	case config.GetLanguage().Th:
		provinceName = provinceModel.ProvinceTH
	}
	return provinceModel, provinceName
}

func (ln Ln) GetPlantTypeNameer(plantTypeId, language string) (model_databases.PlantType, string) {
	var plantTypeModel model_databases.PlantType
	var plantTypeName string

	sql := fmt.Sprintf("SELECT * FROM %s WHERE status_id = '%s' AND plant_type_id = '%s'",
		config.DB_PLANT_TYPE, config.GetStatus().Active, plantTypeId)
	fmt.Println(sql)
	err := ln.Db.Raw(sql).Scan(&plantTypeModel).Error
	if err != nil {
		log.Print(err)
	}
	switch language {
	case config.GetLanguage().En:
		plantTypeName = plantTypeModel.PlantTypeEN
	case config.GetLanguage().Th:
		plantTypeName = plantTypeModel.PlantTypeTH
	}
	return plantTypeModel, plantTypeName
}

func (ln Ln) GetFertCatNameer(fertCatId, language string) (model_databases.FertilizerCat, string) {
	var fertCatModel model_databases.FertilizerCat
	var fertCatName string

	sql := fmt.Sprintf("SELECT * FROM %s WHERE status_id = '%s' AND fertilizer_cat_id = '%s'",
		config.DB_FERTILIZER_CAT, config.GetStatus().Active, fertCatId)
	fmt.Println(sql)
	err := ln.Db.Raw(sql).Scan(&fertCatModel).Error
	if err != nil {
		log.Print(err)
	}

	switch language {
	case config.GetLanguage().En:
		fertCatName = fertCatModel.FertilizerCatEN
	case config.GetLanguage().Th:
		fertCatName = fertCatModel.FertilizerCatTH
	}
	return fertCatModel, fertCatName
}

func (ln Ln) GetUserNameer(uid string) (model_databases.Users, string) {
	var userModel model_databases.Users
	var userName string

	sql := fmt.Sprintf("SELECT * FROM %s WHERE status_id = '%s' AND uid = '%s'",
		config.DB_USERS, config.GetStatus().Active, uid)
	err := ln.Db.Raw(sql).Scan(&userModel).Error
	if err != nil {
		log.Print(err)
	}

	userName = userModel.Username
	return userModel, userName
}

func (ln Ln) GetSensorTypeNameer(sensorTypeId, language string) (model_databases.SensorType, string) {
	var sensorTypeModel model_databases.SensorType
	var sensorTypeName string

	sql := fmt.Sprintf("SELECT * FROM %s WHERE status_id = '%s' AND sensor_type_id = '%s'",
		config.DB_SENSOR_TYPE, config.GetStatus().Active, sensorTypeId)
	fmt.Println(sql)
	err := ln.Db.Raw(sql).Scan(&sensorTypeModel).Error
	if err != nil {
		log.Print(err)
	}
	switch language {
	case config.GetLanguage().En:
		sensorTypeName = sensorTypeModel.SensorTypeNameEN
	case config.GetLanguage().Th:
		sensorTypeName = sensorTypeModel.SensorTypeNameTH
	}

	return sensorTypeModel, sensorTypeName
}

func (ln Ln) GetRoleNameer(roleId, language string) (model_databases.Role, string, string) {
	var roleModel model_databases.Role
	var name string
	var desc string

	sql := fmt.Sprintf("SELECT * FROM %s WHERE status_id = '%s' AND role_id = '%s'",
		config.DB_ROLE, config.GetStatus().Active, roleId)
	fmt.Println(sql)
	err := ln.Db.Raw(sql).Scan(&roleModel).Error
	if err != nil {
		log.Print(err)
	}

	switch language {
	case config.GetLanguage().En:
		name = roleModel.RoleNameEN
		desc = roleModel.RoleDescEN
	case config.GetLanguage().Th:
		name = roleModel.RoleNameTH
		desc = roleModel.RoleDescTH
	}

	return roleModel, name, desc
}

func (ln Ln) GetFarmAreaNameer(farmAreaId string) (model_databases.FarmArea, string) {
	var farmAreaModel model_databases.FarmArea

	sql := fmt.Sprintf("SELECT * FROM %s WHERE status_id = '%s' AND farm_area_id = '%s'",
		config.DB_FARM_AREA, config.GetStatus().Active, farmAreaId)
	fmt.Println(sql)
	err := ln.Db.Raw(sql).Scan(&farmAreaModel).Error
	if err != nil {
		log.Print(err)
	}

	return farmAreaModel, farmAreaModel.FarmAreaName
}
