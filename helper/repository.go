package helper

import (
	"gorm.io/gorm"
)

type Repositorier interface {
	//FindAllPlantType(status string) ([]model_db.PlantType, error)
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repositorier {
	return &Repository{db: db}
}
