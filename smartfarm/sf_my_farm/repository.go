package sf_my_farm

import (
	"database/sql"
	"errors"
	"github.com/jjoykkm/ln-backend/common/config"
	"github.com/jjoykkm/ln-backend/common/models/model_db"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repositorier interface {
	Begin() *Repository
	Commit()
	Rollback()
	FindOneFarm(status, farmId string) (*FarmOverview, error)
	FindOneMainboxBySerialNo (serialNo string) (*model_db.Mainbox, error)
	GetCountMainbox(status, farmId string) (int64, error)
	GetCountFarmArea(status, farmId string) (int64, error)
	FindAllManageRole(status, farmId string) ([]ManageRole, error)
	FindAllManageFarmArea(status, farmId string) ([]ManageFarmArea, error)
	FindAllManageMainbox(status, farmId string) ([]ManageMainbox, error)
	UpdateOneMainboxBySerialNo (req *model_db.MainboxSerialUS) error
	UpdateOneMainbox (req *model_db.MainboxUS) error
	UpdateOneSensor (req *model_db.Sensor) error
	UpdateAllSocket (req *LinkedSocFarmArea) error
	UpsertFarm (req *model_db.FarmUS) error
	UpsertSocket (req []model_db.SocketUS) error
	CreateOneSensor (req *model_db.SensorUS) error
	DeleteOneSocket (socketId string) error
	DeactivateOneMainbox (mainboxId string) error
	DeleteOneFarm (farmId string) error
	DeleteOneFarmArea (farmAreaId string) error
	UpsertFarmArea (req *model_db.FarmAreaUS) (error, *string)
	UpdateAllSocketNullFarmArea (req []string) error
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repositorier {
	return &Repository{db: db}
}

func (r *Repository) Begin() *Repository {
	return &Repository{
		db:	r.db.Debug().Begin(),
	}
}

func (r *Repository) Commit() {
	r.db.Debug().Commit()
}

func (r *Repository) Rollback() {
	r.db.Debug().Rollback()
}

func (r *Repository) FindOneFarm(status, farmId string) (*FarmOverview, error) {
	var result FarmOverview

	resp := r.db.Debug().Where("status_id = ? AND farm_id = ?", status, farmId).First(&result)
	if resp.Error != nil && !errors.Is(resp.Error, gorm.ErrRecordNotFound) {
		return nil, resp.Error
	}
	return &result, nil
}

func (r *Repository) FindOneMainboxBySerialNo (serialNo string) (*model_db.Mainbox, error) {
	var result model_db.Mainbox

	resp := r.db.Debug().Where("mainbox_serial_no = ?", serialNo).First(&result)
	if resp.Error != nil {
		return nil, resp.Error
	}
	return &result, nil
}

func (r *Repository) GetCountMainbox(status, farmId string) (int64, error) {
	var mainbox []model_db.Mainbox
	var count int64

	resp := r.db.Debug().Model(&mainbox).Distinct("mainbox_id").Where("status_id = ? AND farm_id = ?",
		status, farmId).Count(&count)
	if resp.Error != nil && !errors.Is(resp.Error, gorm.ErrRecordNotFound) {
		return 0, resp.Error
	}
	return count, nil
}

func (r *Repository) GetCountFarmArea(status, farmId string) (int64, error) {
	var farmArea []model_db.FarmArea
	var count int64

	resp := r.db.Debug().Model(&farmArea).Where("status_id = ? AND farm_id = ?", status, farmId).Count(&count)
	if resp.Error != nil && !errors.Is(resp.Error, gorm.ErrRecordNotFound) {
		return 0, resp.Error
	}
	return count, nil
}

func (r *Repository) FindAllManageRole(status, farmId string) ([]ManageRole, error) {
	var result []ManageRole

	checkStatus := r.db.Debug().Where("status_id = ?", config.GetStatus().Active)
	resp := r.db.Debug().Where("status_id = ? AND farm_id = ?",
		status, farmId).Preload("User").Preload("Role",
			func(db *gorm.DB) *gorm.DB {
				return checkStatus
			}).Find(&result)
	if resp.Error != nil && !errors.Is(resp.Error, gorm.ErrRecordNotFound) {
		return nil, resp.Error
	}
	return result, nil
}


func (r *Repository) FindAllManageFarmArea(status, farmId string) ([]ManageFarmArea, error) {
	var result []ManageFarmArea

	//Get Sensor Detail
	sensorDet := r.db.Debug().Where("status_id = ?", config.GetStatus().Active).Preload("SensorType",
		"status_id = ?", config.GetStatus().Active)
	// Get Socket Detail
	socketDet := r.db.Debug().Where("status_id = ?", config.GetStatus().Active).Preload("StatusSensor",
		"status_id = ?", config.GetStatus().Active).Preload("Sensor",
		func(db *gorm.DB) *gorm.DB {
			return sensorDet
		})
	resp := r.db.Debug().Where("status_id = ? AND farm_id = ?", status, farmId).Preload("SocSenDetail",
		func(db *gorm.DB) *gorm.DB {
			return socketDet
		}).Find(&result)

	if resp.Error != nil && !errors.Is(resp.Error, gorm.ErrRecordNotFound) {
		return nil, resp.Error
	}
	return result, nil
}

func (r *Repository) FindAllManageMainbox(status, farmId string) ([]ManageMainbox, error) {
	var result []ManageMainbox
	//Get Sensor Detail
	sensorDet := r.db.Debug().Where("status_id = ?", config.GetStatus().Active).Preload("SensorType",
		"status_id = ?", config.GetStatus().Active)
	// Get Socket Detail
	socketDet := r.db.Debug().Where("status_id = ?", config.GetStatus().Active).Preload("StatusSensor",
		"status_id = ?", config.GetStatus().Active).Preload("Sensor",
		func(db *gorm.DB) *gorm.DB {
			return sensorDet
		})
	resp := r.db.Debug().Where("status_id = ? AND farm_id = ?",
	status, farmId).Preload("SocSenDetail",
		func(db *gorm.DB) *gorm.DB {
			return socketDet
		}).Find(&result)

	if resp.Error != nil && !errors.Is(resp.Error, gorm.ErrRecordNotFound) {
		return nil, resp.Error
	}
	return result, nil
}

//-------------------------------------------------------------------------------//
//							Update Data
//-------------------------------------------------------------------------------//
func (r *Repository) UpdateOneMainboxBySerialNo (req *model_db.MainboxSerialUS) error {
	resp := r.db.Debug().Where("mainbox_serial_no = ?",
			req.MainboxSerialNo).Updates(&req).Update("status_id", // Get active status
				r.db.Model(&model_db.Status{}).Select("status_id").Where("status_id = ?",
					config.GetStatus().Active))
	if resp.Error != nil {
		return resp.Error
	}
	if resp.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
func (r *Repository) UpdateOneMainbox (req *model_db.MainboxUS) error {
	resp := r.db.Debug().Where("mainbox_id = ?",
		req.MainboxId).Updates(&req)
	if resp.Error != nil {
		return resp.Error
	}
	if resp.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
func (r *Repository) UpdateOneSensor (req *model_db.Sensor) error {
	resp := r.db.Debug().Where("sensor_id = ?",
		req.SensorId).Updates(&req)
	if resp.Error != nil {
		return resp.Error
	}
	if resp.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
func (r *Repository) UpdateAllSocket (req *LinkedSocFarmArea) error {
	resp := r.db.Debug().Table(config.DB_SOCKET).Select("farm_area_id").Where("socket_id IN ?",
		req.SocketId).Updates(&req)
	if resp.Error != nil {
		return resp.Error
	}
	if resp.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *Repository) UpdateAllSocketNullFarmArea (req []string) error {
	resp := r.db.Debug().Table(config.DB_SOCKET).Where("socket_id IN ?",
		req).Update("farm_area_id", sql.NullString{})
	if resp.Error != nil {
		return resp.Error
	}
	if resp.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

//-------------------------------------------------------------------------------//
//							Upsert Data
//-------------------------------------------------------------------------------//
func (r *Repository) UpsertFarm (req *model_db.FarmUS) error {
	resp := r.db.Debug().Model(model_db.FarmUS{}).Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: "farm_id"},
		},
		UpdateAll: true,
	}).Create(&req)
	if resp.Error != nil {
		return resp.Error
	}
	if resp.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *Repository) UpsertSocket (req []model_db.SocketUS) error {
	resp := r.db.Debug().Model(model_db.SocketUS{}).Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: "socket_id"},
		},
		UpdateAll: true,
		}).Create(&req)
	if resp.Error != nil {
		return resp.Error
	}
	if resp.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *Repository) CreateOneSensor (req *model_db.SensorUS) error {
	resp := r.db.Debug().Create(&req)
	if resp.Error != nil {
		return resp.Error
	}
	if resp.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *Repository) UpsertFarmArea (req *model_db.FarmAreaUS) (error, *string) {
	resp := r.db.Debug().Model(model_db.FarmAreaUS{}).Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: "farm_area_id"},
		},
		UpdateAll: true,
	}).Create(&req)
	if resp.Error != nil {
		return resp.Error, nil
	}
	if resp.RowsAffected == 0 {
		return gorm.ErrRecordNotFound, nil
	}
	return nil, &req.FarmAreaId
}

//-------------------------------------------------------------------------------//
//							Delete Data
//-------------------------------------------------------------------------------//
func (r *Repository) DeleteOneSocket (socketId string) error {
	resp := r.db.Debug().Where("socket_id = ?", socketId).Delete(&model_db.Socket{})
	//resp := r.db.Debug().Table(config.DB_SOCKET).Where("socket_id = ?", socketId).Update("status_id",
	//	config.GetStatus().Inactive)
	if resp.Error != nil {
		return resp.Error
	}
	if resp.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *Repository) DeactivateOneMainbox (mainboxId string) error {
	// Assign status pending
	statusId := config.GetStatus().Inactive
	resp := r.db.Debug().Model(&model_db.Mainbox{}).Where("mainbox_id = ?", mainboxId).Update("status_id", statusId)
	if resp.Error != nil {
		return resp.Error
	}
	if resp.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *Repository) DeleteOneFarm (farmId string) error {
	resp := r.db.Debug().Where("farm_id = ?", farmId).Delete(&model_db.Farm{})
	//resp := r.db.Debug().Table(config.DB_FARM).Where("farm_id = ?", farmId).Update("status_id",
	//	config.GetStatus().Inactive)
	if resp.Error != nil {
		return resp.Error
	}
	if resp.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *Repository) DeleteOneFarmArea (farmAreaId string) error {
	resp := r.db.Debug().Where("farm_area_id = ?", farmAreaId).Delete(&model_db.FarmArea{})
	//resp := r.db.Debug().Table(config.DB_FARM_AREA).Where("farm_area_id = ?", farmAreaId).Update("status_id",
	//	config.GetStatus().Inactive)
	if resp.Error != nil {
		return resp.Error
	}
	if resp.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
