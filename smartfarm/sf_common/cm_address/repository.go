package cm_address

import (
	"errors"
	"github.com/jjoykkm/ln-backend/modelsOld/model_databases"
	"gorm.io/gorm"
)

type Repositorier interface {
	FindAllProvince(status string) ([]model_databases.Province, error)
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repositorier {
	return &Repository{db: db}
}

func (r *Repository) FindAllProvince(status string) ([]model_databases.Province, error) {
	var result []model_databases.Province

	err := r.db.Where("status_id = ?", status).Find(&result).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return result, nil
}