package cm_plant

import (
	"errors"
	"github.com/jjoykkm/ln-backend/common/config"
	"gorm.io/gorm"
)

type Repositorier interface {
	FindAllFertAndCatList(status string) ([]FertilizerAndCat, error)
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repositorier {
	return &Repository{db: db}
}

func (r *Repository) FindAllFertAndCatList(status string) ([]FertilizerAndCat, error) {
	var result []FertilizerAndCat
	// Get fertilizer
	fertId := r.db.Debug().Table(config.DB_FERTILIZER_CAT).Select("fertilizer_cat_id").Where("status_id = ?",
		config.GetStatus().Active)
	// Add or condition cause preload will select with primary but i wanna select by foreign key
	resp := r.db.Debug().Where("status_id = ?", status).Preload("Fertilizer",
		func(db *gorm.DB) *gorm.DB {
			return r.db.Debug().Or("fertilizer_cat_id IN (?)",fertId)
		}).Find(&result)

	if resp.Error != nil && !errors.Is(resp.Error, gorm.ErrRecordNotFound) {
		return nil, resp.Error
	}
	return result, nil
}
func (r *Repository) GetFertAndCatList(status string) ([]FertilizerAndCat, error) {
	var result []FertilizerAndCat
	resp := r.db.Debug().Where("status_id = ?", status).Preload("Fertilizer",
		func(db *gorm.DB) *gorm.DB {
			return r.db.Debug().Or("fertilizer_cat_id IN (?)",fertId)
		}).Find(&result)

	if resp.Error != nil && !errors.Is(resp.Error, gorm.ErrRecordNotFound) {
		return nil, resp.Error
	}
	return result, nil
}