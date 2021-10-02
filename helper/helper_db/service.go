package helper_db

import (
	"github.com/jjoykkm/ln-backend/config"
	"github.com/jjoykkm/ln-backend/models/model_databases"
)

type Servicer interface {
	GetFarmAreaByFarm(status, farmId, resultType string) ([]model_databases.FarmArea, []string)
	GetCountryAndName(countryId, language string) (*model_databases.Country, string)
	GetProvinceAndName(provinceId, language string) (*model_databases.Province, string)
	GetPlantTypeAndName(plantTypeId, language string) (*model_databases.PlantType, string)
	GetFertCatAndName(fertCatId, language string) (*model_databases.FertilizerCat, string)
	GetUserAndName(uid string) (*model_databases.Users, string)
	GetSensorTypeAndName(sensorTypeId, language string) (*model_databases.SensorType, string)
	GetRoleAndName(roleId, language string) (*model_databases.Role, string, string)
	GetFarmAreaAndName(farmAreaId string) (*model_databases.FarmArea, string)
}

type Service struct {
	repo Repositorier
}

func NewService(repo Repositorier) Servicer {
	return &Service{
		repo:  repo,
	}
}
func (s *Service) GetFarmAreaByFarm(status, farmId, resultType string) ([]model_databases.FarmArea, []string) {
	var farmAreaIdList []string
	farmModel, err := s.repo.FindAllFarmAreaByFarm(status, farmId)
	if err != nil {
		return nil, nil
	}
	if resultType != config.RES_TYPE_STRUCT {
		for _, array := range farmModel {
			farmAreaIdList = append(farmAreaIdList, array.FarmAreaId.UUID.String())
		}
	}
	return farmModel, farmAreaIdList
}

func (s *Service) GetCountryAndName(countryId, language string) (*model_databases.Country, string) {
	var countryName string

	countryModel, err := s.repo.FindOneCountry(config.STATUS_ACTIVE, countryId)
	if err != nil {
		return nil, ""
	}
	switch language {
	case config.LANGUAGE_EN:
		countryName = countryModel.CountryEN
	case config.LANGUAGE_TH:
		countryName = countryModel.CountryTH
	}
	return countryModel, countryName
}

func (s *Service) GetProvinceAndName(provinceId, language string) (*model_databases.Province, string) {
	var provinceName string

	provinceModel, err := s.repo.FindOneProvince(config.STATUS_ACTIVE, provinceId)
	if err != nil {
		return nil, ""
	}
	switch language {
	case config.LANGUAGE_EN:
		provinceName = provinceModel.ProvinceEN
	case config.LANGUAGE_TH:
		provinceName = provinceModel.ProvinceTH
	}
	return provinceModel, provinceName
}

func (s *Service) GetPlantTypeAndName(plantTypeId, language string) (*model_databases.PlantType, string) {
	var plantTypeName string

	plantTypeModel, err := s.repo.FindOnePlantType(config.STATUS_ACTIVE, plantTypeId)
	if err != nil {
		return nil, ""
	}
	switch language {
	case config.LANGUAGE_EN:
		plantTypeName = plantTypeModel.PlantTypeEN
	case config.LANGUAGE_TH:
		plantTypeName = plantTypeModel.PlantTypeTH
	}
	return plantTypeModel, plantTypeName
}

func (s *Service) GetFertCatAndName(fertCatId, language string) (*model_databases.FertilizerCat, string) {
	var fertCatName string

	fertCatModel, err := s.repo.FindOneFertCat(config.STATUS_ACTIVE, fertCatId)
	if err != nil {
		return nil, ""
	}
	switch language {
	case config.LANGUAGE_EN:
		fertCatName = fertCatModel.FertilizerCatEN
	case config.LANGUAGE_TH:
		fertCatName = fertCatModel.FertilizerCatTH
	}
	return fertCatModel, fertCatName
}

func (s *Service) GetUserAndName(uid string) (*model_databases.Users, string) {
	var userName string

	userModel, err := s.repo.FindOneUser(config.STATUS_ACTIVE, uid)
	if err != nil {
		return nil, ""
	}
	userName = userModel.Username
	return userModel, userName
}

func (s *Service) GetSensorTypeAndName(sensorTypeId, language string) (*model_databases.SensorType, string) {
	var sensorTypeName string
	
	sensorTypeModel, err := s.repo.FindOneSensorType(config.STATUS_ACTIVE, sensorTypeId)
	if err != nil {
		return nil, ""
	}
	switch language {
	case config.LANGUAGE_EN:
		sensorTypeName = sensorTypeModel.SensorTypeNameEN
	case config.LANGUAGE_TH:
		sensorTypeName = sensorTypeModel.SensorTypeNameTH
	}
	return sensorTypeModel, sensorTypeName
}

func (s *Service) GetRoleAndName(roleId, language string) (*model_databases.Role, string, string) {
	var name string
	var desc string

	roleModel, err := s.repo.FindOneRole(config.STATUS_ACTIVE, roleId)
	if err != nil {
		return nil, "", ""
	}
	switch language {
	case config.LANGUAGE_EN:
		name = roleModel.RoleNameEN
		desc = roleModel.RoleDescEN
	case config.LANGUAGE_TH:
		name = roleModel.RoleNameTH
		desc = roleModel.RoleDescTH
	}
	return roleModel, name, desc
}

func (s *Service) GetFarmAreaAndName(farmAreaId string) (*model_databases.FarmArea, string) {
	farmAreaModel, err := s.repo.FindOneFarmArea(config.STATUS_ACTIVE, farmAreaId)
	if err != nil {
		return nil, ""
	}
	return farmAreaModel, farmAreaModel.FarmAreaName
}
