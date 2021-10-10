package sf_dashboard

import (
	"errors"
	"github.com/jjoykkm/ln-backend/common/config"
	"github.com/jjoykkm/ln-backend/common/models/model_db"
	"gorm.io/gorm"
)

type Repositorier interface {
	FindAllMyFarm(status, uid string) ([]model_db.Farm, error)
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
