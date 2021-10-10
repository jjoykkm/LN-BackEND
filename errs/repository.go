package errs

import (
	"gorm.io/gorm"
)

type Repositorier interface {
	//FindAllFarmAreaByFarm(farmId, status string) ([]model_databases.FarmArea, error)
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repositorier {
	return &Repository{db: db}
}

//func (r *Repository) FindAllFarmAreaByFarm(farmId, status string) ([]model_databases.FarmArea, error) {
//	var result []model_databases.FarmArea
//
//	err := r.db.Where("status_id = ? AND farm_id = ?", status, farmId).Find(&result).Error
//	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
//		return nil, err
//	}
//	return result, nil
//}
