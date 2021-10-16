package sf_dashboard

import (
	"errors"
	"github.com/jjoykkm/ln-backend/common/config"
	"gorm.io/gorm"
)

type Repositorier interface {
	FindAllFarmAreaDashboard(status, farmId string) ([]FarmSensorDetail, error)
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repositorier {
	return &Repository{db: db}
}

func (r *Repository) FindAllFarmAreaDashboard(status, farmId string) ([]FarmSensorDetail, error) {
	var result []FarmSensorDetail

	//Get Sensor Detail
	sensorDet := r.db.Debug().Where("status_id = ?", config.GetStatus().Active).Preload("SensorType",
		"status_id = ?", config.GetStatus().Active)
	// Get Socket Detail
	socketDet := r.db.Debug().Where("status_id = ?", config.GetStatus().Active).Preload("StatusSensor",
		"status_id = ?", config.GetStatus().Active).Preload("Sensor",
		func(db *gorm.DB) *gorm.DB {
			return sensorDet
		})
	// Get Mainbox
	mainbox := r.db.Debug().Where("status_id = ?", config.GetStatus().Active)
	// Get Sensor, Socket and Mainbox
	senSocMain := r.db.Debug().Where("status_id = ?", config.GetStatus().Active).Preload("Socket",
		func(db *gorm.DB) *gorm.DB {
			return socketDet
		}).Preload("Mainbox",
		func(db *gorm.DB) *gorm.DB {
			return mainbox
		})
	// Get FormulaPlant
	forPlant := r.db.Debug().Where("status_id = ?", config.GetStatus().Active)

	resp := r.db.Debug().Where("status_id = ? AND farm_id = ?", status, farmId).Preload("FormulaPlant",
		func(db *gorm.DB) *gorm.DB {
			return forPlant
		}).Preload("SensorDetail",
		func(db *gorm.DB) *gorm.DB {
			return senSocMain
		}).Find(&result)

	if resp.Error != nil && !errors.Is(resp.Error, gorm.ErrRecordNotFound) {
		return nil, resp.Error
	}
	return result, nil
}
