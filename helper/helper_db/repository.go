package helper_db

import (
	"errors"
	"github.com/jjoykkm/ln-backend/models/model_databases"
	"gorm.io/gorm"
)

type Repositorier interface {
	FindOneCountry(countryId, status string) (*model_databases.Country, error)
	FindOneProvince(provinceId, status string) (*model_databases.Province, error)
	FindOnePlantType(plantTypeId, status string) (*model_databases.PlantType, error)
	FindOneFertCat(fertCatId, status string) (*model_databases.FertilizerCat, error)
	FindOneUser(uid, status string) (*model_databases.Users, error)
	FindOneSensorType(sensorTypeId, status string) (*model_databases.SensorType, error)
	FindOneRole(roleId, status string) (*model_databases.Role, error)
	FindOneFarmArea(farmAreaId, status string) (*model_databases.FarmArea, error)
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repositorier {
	return &Repository{db: db}
}

func (r *Repository) FindOneCountry(countryId, status string) (*model_databases.Country, error) {
	var result model_databases.Country

	err := r.db.Where("status_id = ? AND country_id = ?", status, countryId).First(&result).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &result, nil
}

func (r *Repository) FindOneProvince(provinceId, status string) (*model_databases.Province, error) {
	var result model_databases.Province

	err := r.db.Where("status_id = ? AND province_id = ?", status, provinceId).First(&result).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &result, nil
}

func (r *Repository) FindOnePlantType(plantTypeId, status string) (*model_databases.PlantType, error) {
	var result model_databases.PlantType

	err := r.db.Where("status_id = ? AND plant_type_id = ?", status, plantTypeId).First(&result).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &result, nil
}

func (r *Repository) FindOneFertCat(fertCatId, status string) (*model_databases.FertilizerCat, error) {
	var result model_databases.FertilizerCat

	err := r.db.Where("status_id = ? AND fertilizer_cat_id = ?", status, fertCatId).First(&result).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &result, nil
}

func (r *Repository) FindOneUser(uid, status string) (*model_databases.Users, error) {
	var result model_databases.Users

	err := r.db.Where("status_id = ? AND uid = ?", status, uid).First(&result).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &result, nil
}

func (r *Repository) FindOneSensorType(sensorTypeId, status string) (*model_databases.SensorType, error) {
	var result model_databases.SensorType

	err := r.db.Where("status_id = ? AND sensor_type_id = ?", status, sensorTypeId).First(&result).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &result, nil
}

func (r *Repository) FindOneRole(roleId, status string) (*model_databases.Role, error) {
	var result model_databases.Role

	err := r.db.Where("status_id = ? AND role_id = ?", status, roleId).First(&result).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &result, nil
}

func (r *Repository) FindOneFarmArea(farmAreaId, status string) (*model_databases.FarmArea, error) {
	var result model_databases.FarmArea

	err := r.db.Where("status_id = ? AND farm_area_id = ?", status, farmAreaId).First(&result).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &result, nil
}