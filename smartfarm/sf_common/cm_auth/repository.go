package cm_auth

import (
	"errors"
	"github.com/jjoykkm/ln-backend/common/models/model_db"
	"gorm.io/gorm"
)

type Repositorier interface {
	FindAllProvince(status string) ([]model_db.Province, error)
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repositorier {
	return &Repository{db: db}
}

func (r *Repository) FindAllProvince(status string) ([]model_db.Province, error) {
	var result []model_db.Province

	resp := r.db.Debug().Where("status_id = ?", status).Find(&result)
	if resp.Error != nil && !errors.Is(resp.Error, gorm.ErrRecordNotFound) {
		return nil, resp.Error
	}
	return result, nil
}