package sf_remote_switch

import (
	"gorm.io/gorm"
)

type Repositorier interface {
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repositorier {
	return &Repository{db: db}
}

