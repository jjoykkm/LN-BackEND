package sf_formula_plant

import (
	"errors"
	"fmt"
	"github.com/jjoykkm/ln-backend/config"
	"github.com/jjoykkm/ln-backend/helper"
	"github.com/jjoykkm/ln-backend/models/model_databases"
	"gorm.io/gorm"
)

type Repositorier interface {
	FindAllPlantType(status string) ([]model_databases.PlantType, error)
	FindJoinPlantWithPlantType(status, plantTypeId string, offset int) ([]JoinPlantAndPlantType, error)
	GetCountFormulaPlant(status, plantId string) int64
	FindAllFavoriteFormulaPlant(status, uid string) ([]model_databases.FavoritePlant, error)
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repositorier {
	return &Repository{db: db}
}

func (r *Repository) FindAllPlantType(status string) ([]model_databases.PlantType, error) {
	var result []model_databases.PlantType

	err := r.db.Where("status_id = ? ", status).Find(&result).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return result, nil
}

func (r *Repository) FindJoinPlantWithPlantType(status, plantTypeId string, offset int) ([]JoinPlantAndPlantType, error) {
	var result []JoinPlantAndPlantType
	var sqlWhere string
	var err error

	sqlJoin := helper.ConcatJoin(config.JOIN_TYPE_INNER, config.DB_PLANT, config.DB_PLANT_TYPE,"plant_type_id")
	if plantTypeId == ""{
		sqlWhere = fmt.Sprintf("%s.status_id = ?",config.DB_PLANT)
		err = r.db.Table(config.DB_PLANT).Joins(sqlJoin).Limit(LIMIT_GET_DATA).Offset(offset).Where(sqlWhere, status).Scan(&result).Error
	}else {
		sqlWhere = fmt.Sprintf("%s.status_id = ? AND %s.plant_type_id = ?", config.DB_PLANT, config.DB_PLANT)
		err = r.db.Table(config.DB_PLANT).Joins(sqlJoin).Limit(LIMIT_GET_DATA).Offset(offset).Where(sqlWhere, status, plantTypeId).Scan(&result).Error
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return result, nil
}

func (r *Repository) GetCountFormulaPlant(status, plantId string) int64 {
	var forPlant []model_databases.FormulaPlant
	var count int64

	err := r.db.Model(&forPlant).Where("status_id = ? AND plant_id = ?", status, plantId).Count(&count).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0
	}
	return count
}

func (r *Repository) FindAllFavoriteFormulaPlant(status, uid string) ([]model_databases.FavoritePlant, error) {
	var result []model_databases.FavoritePlant

	err := r.db.Where("status_id = ? AND uid = ?", status, uid).Order("change_date desc").Find(&result).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return result, nil
}