package helper_db

import (
	"github.com/jjoykkm/ln-backend/common/config"
	"github.com/jjoykkm/ln-backend/modelsOld/model_databases"
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
	if resultType != config.GetResType().Struct {
		for _, array := range farmModel {
			farmAreaIdList = append(farmAreaIdList, array.FarmAreaId.UUID.String())
		}
	}
	return farmModel, farmAreaIdList
}

func (s *Service) GetCountryAndName(countryId, language string) (*model_databases.Country, string) {
	var countryName string

	countryModel, err := s.repo.FindOneCountry(config.GetStatus().Active, countryId)
	if err != nil {
		return nil, ""
	}
	switch language {
	case config.GetLanguage().En:
		countryName = countryModel.CountryEN
	case config.GetLanguage().Th:
		countryName = countryModel.CountryTH
	}
	return countryModel, countryName
}

func (s *Service) GetProvinceAndName(provinceId, language string) (*model_databases.Province, string) {
	var provinceName string

	provinceModel, err := s.repo.FindOneProvince(config.GetStatus().Active, provinceId)
	if err != nil {
		return nil, ""
	}
	switch language {
	case config.GetLanguage().En:
		provinceName = provinceModel.ProvinceEN
	case config.GetLanguage().Th:
		provinceName = provinceModel.ProvinceTH
	}
	return provinceModel, provinceName
}

func (s *Service) GetPlantTypeAndName(plantTypeId, language string) (*model_databases.PlantType, string) {
	var plantTypeName string

	plantTypeModel, err := s.repo.FindOnePlantType(config.GetStatus().Active, plantTypeId)
	if err != nil {
		return nil, ""
	}
	switch language {
	case config.GetLanguage().En:
		plantTypeName = plantTypeModel.PlantTypeEN
	case config.GetLanguage().Th:
		plantTypeName = plantTypeModel.PlantTypeTH
	}
	return plantTypeModel, plantTypeName
}

func (s *Service) GetFertCatAndName(fertCatId, language string) (*model_databases.FertilizerCat, string) {
	var fertCatName string

	fertCatModel, err := s.repo.FindOneFertCat(config.GetStatus().Active, fertCatId)
	if err != nil {
		return nil, ""
	}
	switch language {
	case config.GetLanguage().En:
		fertCatName = fertCatModel.FertilizerCatEN
	case config.GetLanguage().Th:
		fertCatName = fertCatModel.FertilizerCatTH
	}
	return fertCatModel, fertCatName
}

func (s *Service) GetUserAndName(uid string) (*model_databases.Users, string) {
	var userName string

	userModel, err := s.repo.FindOneUser(config.GetStatus().Active, uid)
	if err != nil {
		return nil, ""
	}
	userName = userModel.Username
	return userModel, userName
}

func (s *Service) GetSensorTypeAndName(sensorTypeId, language string) (*model_databases.SensorType, string) {
	var sensorTypeName string
	
	sensorTypeModel, err := s.repo.FindOneSensorType(config.GetStatus().Active, sensorTypeId)
	if err != nil {
		return nil, ""
	}
	switch language {
	case config.GetLanguage().En:
		sensorTypeName = sensorTypeModel.SensorTypeNameEN
	case config.GetLanguage().Th:
		sensorTypeName = sensorTypeModel.SensorTypeNameTH
	}
	return sensorTypeModel, sensorTypeName
}

func (s *Service) GetRoleAndName(roleId, language string) (*model_databases.Role, string, string) {
	var name string
	var desc string

	roleModel, err := s.repo.FindOneRole(config.GetStatus().Active, roleId)
	if err != nil {
		return nil, "", ""
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

func (s *Service) GetFarmAreaAndName(farmAreaId string) (*model_databases.FarmArea, string) {
	farmAreaModel, err := s.repo.FindOneFarmArea(config.GetStatus().Active, farmAreaId)
	if err != nil {
		return nil, ""
	}
	return farmAreaModel, farmAreaModel.FarmAreaName
}
