package cm_auth

import (
	"github.com/jjoykkm/ln-backend/common/config"
	"github.com/jjoykkm/ln-backend/common/models/model_db"
	"gorm.io/gorm"
)

type Repositorier interface {
	FindOneTransManagement(uid, farmId string) (*model_db.TransManagement, error)
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repositorier {
	return &Repository{db: db}
}

func (r *Repository) FindOneTransManagement(uid, farmId string) (*model_db.TransManagement, error) {
	var result model_db.TransManagement

	resp := r.db.Debug().Where("status_id = ? AND uid = ? AND farm_id = ?",
		config.GetStatus().Active, uid, farmId).First(&result)
	if resp.Error != nil {
		return nil, resp.Error
	}
	return &result, nil
}