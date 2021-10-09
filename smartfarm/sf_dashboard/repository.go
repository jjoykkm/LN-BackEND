package sf_dashboard

import (
	"errors"
	"github.com/jjoykkm/ln-backend/models/model_databases"
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

