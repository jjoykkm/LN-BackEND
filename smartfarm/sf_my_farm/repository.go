package SF_FormulaPlant

import (
	"errors"
	"github.com/jjoykkm/ln-backend/modelsOld/model_databases"
	"gorm.io/gorm"
)

type Repositorier interface {
	FindAllPlantType(status string) ([]model_databases.PlantType, error)
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repositorier {
	return &Repository{db: db}
}

func (r *Repository) FindAllPlantType(status string) ([]model_databases.PlantType, error) {
	var result []model_databases.PlantType

	//SELECT * FROM %s WHERE status_id = '%s'
	err := r.db.Where("status_id = ? ", status).Find(&result).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return result, nil
}