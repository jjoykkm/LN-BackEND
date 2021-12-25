package sf_dashboard

import (
	"errors"
	"github.com/jjoykkm/ln-backend/common/config"
	"gorm.io/gorm"
)

type Repositorier interface {
	GetFarmAreaDashboardList(status, farmId string) ([]FarmDashboard, error)
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repositorier {
	return &Repository{db: db}
}

func (r *Repository) GetFarmAreaDashboardList(status, farmId string) ([]FarmDashboard, error) {
	var result []FarmDashboard

	// Get Mainbox Detail
	mainboxDet := r.db.Debug().Where("status_id = ?", config.GetStatus().Active)
	//Get Sensor Detail
	sensorDet := r.db.Debug().Where("status_id = ?", config.GetStatus().Active).Preload("SensorType",
		"status_id = ?", config.GetStatus().Active)
	// Get Socket Detail
	socketDet := r.db.Debug().Where("status_id = ?", config.GetStatus().Active).Preload("StatusSensor",
		"status_id = ?", config.GetStatus().Active).Preload("Mainbox",
		func(db *gorm.DB) *gorm.DB {
			return mainboxDet
		}).Preload("Sensor",
		func(db *gorm.DB) *gorm.DB {
			return sensorDet
		})
	// Get FormulaPlant Detail
	forPlantDet := r.db.Debug().Where("status_id = ?", config.GetStatus().Active)

	resp := r.db.Debug().Where("status_id = ? AND farm_id = ?", status, farmId).Preload("FormulaPlant",
		func(db *gorm.DB) *gorm.DB {
			return forPlantDet
		}).Preload("SocSenDetail",
		func(db *gorm.DB) *gorm.DB {
			return socketDet
		}).Find(&result)

	if resp.Error != nil && !errors.Is(resp.Error, gorm.ErrRecordNotFound) {
		return nil, resp.Error
	}
	return result, nil
}
