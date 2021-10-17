package cm_plant

import (
	"errors"
	"github.com/jjoykkm/ln-backend/common/config"
	"gorm.io/gorm"
)

type Repositorier interface {
	GetFertAndCatList(status string) ([]FertilizerAndCat, error)
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repositorier {
	return &Repository{db: db}
}

func (r *Repository) GetFertAndCatList(status string) ([]FertilizerAndCat, error) {
	var result []FertilizerAndCat
	fertId := r.db.Debug().Table(config.DB_FERTILIZER_CAT).Select("fertilizer_cat_id").Where("status_id = ?",
		config.GetStatus().Active)
	resp := r.db.Debug().Where("status_id = ?", status).Preload("Fertilizer",
		func(db *gorm.DB) *gorm.DB {
			return r.db.Debug().Or(" fertilizer_cat_id IN (?)",fertId)
		}).Find(&result)

	if resp.Error != nil && !errors.Is(resp.Error, gorm.ErrRecordNotFound) {
		return nil, resp.Error
	}
	return result, nil
}