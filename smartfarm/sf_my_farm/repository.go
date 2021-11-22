package sf_my_farm

import (
	"errors"
	"github.com/jjoykkm/ln-backend/common/config"
	"github.com/jjoykkm/ln-backend/common/models/model_db"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repositorier interface {
	FindOneFarm(status, farmId string) (*FarmOverview, error)
	FindOneMainboxBySerialNo (serialNo string) (*model_db.Mainbox, error)
	GetCountMainbox(status, farmId string) (int64, error)
	GetCountFarmArea(status, farmId string) (int64, error)
	FindOneTransManagement(uid, farmId string) (*model_db.TransManagement, error)
	FindAllManageRole(status, farmId string) ([]ManageRole, error)
	FindAllManageFarmArea(status, farmId string) ([]ManageFarmArea, error)
	FindAllManageMainbox(status, farmId string) ([]ManageMainbox, error)
	UpdateOneMainboxBySerialNo (req *model_db.MainboxSerialUS) error
	UpdateOneMainbox (req *model_db.MainboxUS) error
	UpdateOneSensor (req *model_db.Sensor) error
	UpsertSocket (req []model_db.SocketUS) error
	CreateOneSensor (req *model_db.SensorUS) error
	DeleteOneSocket (socketId string) error
	DeactivateOneMainbox (mainboxId string) error
	DeleteOneFarm (farmId string) error
	DeleteOneFarmArea (farmAreaId string) error
	UpsertFarmArea (req *model_db.FarmAreaUS) (error, *string)
	UpsertTransSocketArea (req []model_db.TransSocketAreaUS) error
	//FindAllPlantType(status string) ([]model_db.PlantType, error)
	//GetFarmListWithRoleer(status, uid, roleId string) ([]model_services.DashboardFarmList, int)
	//GetOverviewFarmer(status, farmId string) model_services.MyFarmOverviewFarm
	//GetTransSocketAreaer(status, farmId string) ([]model_databases.TransSocketArea, []string, []string, []string, int)
	//GetSocketByIder(status string, socketIdList []string) ([]model_databases.Socket, map[string]model_databases.Socket)
	//GetSocketWithSensorer(status, language string, socketIdList []string) ([]model_services.MyFarmSenSocDetail, map[string]model_services.MyFarmSenSocDetail, int)
	//GetSocSenByKeyer(mainboxId, farmAreaId string, tranSoc []model_databases.TransSocketArea, socSenMap map[string]model_services.MyFarmSenSocDetail) ([]model_services.MyFarmSenSocDetail, int)
	//GetManageMainboxer(status, language, farmId string) ([]model_services.MyFarmManageMainbox, int)
	//GetFarmAreaByIder(status string, farmAreaIdList []string) ([]model_databases.FarmArea, map[string]model_databases.FarmArea)
	//GetManageFarmAreaer(status, language, farmId string) ([]model_services.MyFarmManageFarmArea, int)
	//GetManageRoleer(status, language, farmId string) ([]model_services.MyFarmManageRole, int)
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repositorier {
	return &Repository{db: db}
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
	var trans []model_db.TransSocketArea
	var count int64

	// Get farm_area_id
	farmAreaId := r.db.Debug().Distinct("farm_area_id").Where("status_id = ? AND farm_id = ?",
		config.GetStatus().Active, farmId).Table(config.DB_FARM_AREA)

	resp := r.db.Debug().Model(&trans).Distinct("mainbox_id").Where("status_id = ? AND farm_area_id IN (?)",
		status, farmAreaId).Count(&count)
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

func (r *Repository) FindOneTransManagement(uid, farmId string) (*model_db.TransManagement, error) {
	var result model_db.TransManagement

	resp := r.db.Debug().Where("status_id = ? AND uid = ? AND farm_id = ?",
		config.GetStatus().Active, uid, farmId).First(&result)
	if resp.Error != nil && !errors.Is(resp.Error, gorm.ErrRecordNotFound) {
		return nil, resp.Error
	}
	return &result, nil
}

func (r *Repository) FindAllManageRole(status, farmId string) ([]ManageRole, error) {
	var result []ManageRole
	checkStatus := r.db.Debug().Where("status_id = ?", config.GetStatus().Active)
	resp := r.db.Debug().Where("status_id = ? AND farm_id = ?",
		status, farmId).Preload("User").Preload("Role",
			func(db *gorm.DB) *gorm.DB {
				return checkStatus
			}).Find(&result)
	//,
	//func(db *gorm.DB) *gorm.DB {
	//	return checkStatus
	//})
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
	// Get Sensor, Socket
	senSocMain := r.db.Debug().Where("status_id = ?", config.GetStatus().Active).Preload("Socket",
		func(db *gorm.DB) *gorm.DB {
			return socketDet
		})

	resp := r.db.Debug().Where("status_id = ? AND farm_id = ?", status, farmId).Preload("SensorDetail",
		func(db *gorm.DB) *gorm.DB {
			return senSocMain
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
	// Get Sensor, Socket
	senSoc := r.db.Debug().Where("status_id = ?", config.GetStatus().Active).Preload("Socket",
		func(db *gorm.DB) *gorm.DB {
			return socketDet
		})
	// Get Farm Area
	farmAreaId := r.db.Debug().Distinct("farm_area_id").Where("status_id = ? AND farm_id = ?",
		config.GetStatus().Active, farmId).Table(config.DB_FARM_AREA)
	// Get Mainbox
	mainboxId := r.db.Debug().Distinct("mainbox_id").Where("status_id = ? AND farm_area_id IN (?)",
		config.GetStatus().Active, farmAreaId).Table(config.DB_TRANS_SOCKET_AREA)

	resp := r.db.Debug().Where("status_id = ? AND mainbox_id IN (?)",
		config.GetStatus().Active, mainboxId).Preload("SenSocDetail",
		func(db *gorm.DB) *gorm.DB {
			return senSoc
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

//-------------------------------------------------------------------------------//
//							Upsert Data
//-------------------------------------------------------------------------------//
//func (r *Repository) UpsertOneMainbox (req *model_db.MainboxUS) error {
//	//resp := r.db.Debug().Model(&model_db.Mainbox{}).Where("mainbox_id = ?",
//	//	req.MainboxId).Updates(&req)
//	//resp := r.db.Debug().Clauses(clause.OnConflict{
//	//	Columns:   []clause.Column{{Name: "id"}},
//	//	DoUpdates: clause.AssignmentColumns([]{"name"}),
//	//}).Create(&req)
//	resp := r.db.Debug().Clauses(clause.OnConflict{
//		Columns:   []clause.Column{{Name: "id"}},
//		UpdateAll: true,
//	}).Create(&req)
//	if resp.Error != nil {
//		return resp.Error
//	}
//	if resp.RowsAffected == 0 {
//		return gorm.ErrRecordNotFound
//	}
//	return nil
//}
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

func (r *Repository) UpsertTransSocketArea (req []model_db.TransSocketAreaUS) error {
	resp := r.db.Debug().Model(model_db.TransSocketAreaUS{}).Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: "farm_area_id"},
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

//-------------------------------------------------------------------------------//
//							Delete Data
//-------------------------------------------------------------------------------//
func (r *Repository) DeleteOneSocket (socketId string) error {
	resp := r.db.Debug().Where("socket_id = ?", socketId).Delete(&model_db.Socket{})
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
	if resp.Error != nil {
		return resp.Error
	}
	if resp.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
