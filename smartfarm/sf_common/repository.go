package sf_common

import (
	"errors"
	"github.com/jjoykkm/ln-backend/common/config"
	"github.com/jjoykkm/ln-backend/common/models/model_db"
	"github.com/jjoykkm/ln-backend/modelsOld/model_databases"
	"gorm.io/gorm"
)

type Repositorier interface {
	FindAllMyFarm(status, uid string) ([]model_db.Farm, error)
	FindAllMyFarmAndFarmArea(status, uid string) ([]FarmFarmArea, error)
	FindAllProvince(status string) ([]model_databases.Province, error)
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repositorier {
	return &Repository{db: db}
}

func (r *Repository) FindAllMyFarm(status, uid string) ([]model_db.Farm, error) {
	var result []model_db.Farm
	// Get farm_id
	farmId := r.db.Debug().Select("farm_id").Where("status_id = ? AND uid = ?",
		config.GetStatus().Active, uid).Table(config.DB_TRANS_MANAGEMENT)

	resp := r.db.Debug().Where("status_id = ? AND farm_id IN (?)", status, farmId).Find(&result)
	if resp.Error != nil && !errors.Is(resp.Error, gorm.ErrRecordNotFound) {
		return nil, resp.Error
	}
	return result, nil
}

func (r *Repository) FindAllMyFarmAndFarmArea(status, uid string) ([]FarmFarmArea, error) {
	var result []FarmFarmArea
	// Get farm_id
	farmId := r.db.Debug().Select("farm_id").Where("status_id = ? AND uid = ?",
		config.GetStatus().Active, uid).Table(config.DB_TRANS_MANAGEMENT)

	resp := r.db.Debug().Where("status_id = ? AND farm_id IN (?)",
		status, farmId).Preload("FarmArea","status_id = ?",
			config.GetStatus().Active).Find(&result)

	if resp.Error != nil && !errors.Is(resp.Error, gorm.ErrRecordNotFound) {
		return nil, resp.Error
	}
	return result, nil
}

func (r *Repository) FindAllProvince(status string) ([]model_databases.Province, error) {
	var result []model_databases.Province

	err := r.db.Where("status_id = ?", status).Find(&result).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return result, nil
}