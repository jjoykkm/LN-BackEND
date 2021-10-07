package sf_formula_plant

import (
	"errors"
	"fmt"
	"github.com/jjoykkm/ln-backend/config"
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

	resp := r.db.Debug().Where("status_id = ? ", status).Find(&result)
	if resp.Error != nil && !errors.Is(resp.Error, gorm.ErrRecordNotFound) {
		return nil, resp.Error
	}
	return result, nil
}

func (r *Repository) FindJoinPlantWithPlantType(status, plantTypeId string, offset int) ([]JoinPlantAndPlantType, error) {
	var result []JoinPlantAndPlantType
	var sqlWhere string
	// Generate condition when get plant
	sqlWhere = fmt.Sprintf("%s.status_id = ?",config.DB_PLANT)
	if plantTypeId != "" {
		sqlWhere = sqlWhere + fmt.Sprintf(" AND %s.plant_type_id = ?", config.DB_PLANT)
	}
	resp := r.db.Debug().Where(sqlWhere, status, plantTypeId).Preload("PlantType", "status_id = ?", config.STATUS_ACTIVE).Limit(LIMIT_GET_DATA).Offset(offset).Find(&result)
	if resp.Error != nil && !errors.Is(resp.Error, gorm.ErrRecordNotFound) {
		return nil, resp.Error
	}
	return result, nil
}

func (r *Repository) GetCountFormulaPlant(status, plantId string) int64 {
	var forPlant []model_databases.FormulaPlant
	var count int64

	resp := r.db.Debug().Model(&forPlant).Where("status_id = ? AND plant_id = ?", status, plantId).Count(&count)
	if resp.Error != nil && !errors.Is(resp.Error, gorm.ErrRecordNotFound) {
		return 0
	}
	return count
}

func (r *Repository) FindAllFavoriteFormulaPlant(status, uid string) ([]model_databases.FavoritePlant, error) {
	var result []model_databases.FavoritePlant

	resp := r.db.Debug().Where("status_id = ? AND uid = ?", status, uid).Order("change_date desc").Find(&result)
	if resp.Error != nil && !errors.Is(resp.Error, gorm.ErrRecordNotFound) {
		return nil, resp.Error
	}
	return result, nil
}

func (r *Repository) FindJoinPlantWithFormulaPlant(status, plantId string, offset int) ([]JoinPlantAndFormulaPlant, error) {
	var result []JoinPlantAndFormulaPlant
	var sqlWhere string
	//// Generate condition when get plant
	sqlWhere = fmt.Sprintf("status_id = '%s'",config.STATUS_ACTIVE)
	if plantId != "" {
		sqlWhere = sqlWhere + fmt.Sprintf(" AND plant_id = '%s'", plantId)
	}
	resp := r.db.Debug().Limit(5).Offset(0).Where(sqlWhere).Preload("Plant", func(db *gorm.DB) *gorm.DB {
		return db.Where("status_id = ?", config.STATUS_ACTIVE).Preload("PlantType", "status_id = ?", config.STATUS_ACTIVE)
	}).Preload("Province", "status_id = ?", config.STATUS_ACTIVE).Preload("Country", "status_id = ?", config.STATUS_ACTIVE).Find(&result)

	if resp.Error != nil && !errors.Is(resp.Error, gorm.ErrRecordNotFound) {
		return nil, resp.Error
	}
	return result, nil
}