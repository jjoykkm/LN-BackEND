package sf_remote_switch

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
	FindOneSocket(socketId string) (*model_db.Socket, error)
	FindOneSocketDetail(socketId string) (*SocketDetail, error)
	FindAllRemoteSwitch(status, uid string) ([]RemoteSwitch, error)
	UpsertRemoteSwitch (req *model_db.RemoteSwitchUS) error
	UpdateSocketFieldRemote (req *RemoteDetailUS) error
	UpdateNullSocketFieldRemote (req []string) error
	DeleteRemoteSwitch (req *string) error
	UpdateStatusSensor (req *ControlSwitch, statusSensor string) error
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

//-------------------------------------------------------------------------------//
//									Get
//-------------------------------------------------------------------------------//
func (r *Repository) FindAllRemoteSwitch(status, uid string) ([]RemoteSwitch, error) {
	var result []RemoteSwitch
	//Get farm id
	farmId := r.db.Debug().Table(config.DB_TRANS_MANAGEMENT).Select("farm_id").Where("status_id = ? AND uid = ?",
		config.GetStatus().Active, uid)
	//Get farm area
	farmAreaId := r.db.Debug().Table(config.DB_FARM_AREA).Select("farm_area_id").Where("status_id = ? AND farm_id IN (?)",
		config.GetStatus().Active, farmId)
	//Get RemoteIdList
	remoteIdList := r.db.Debug().Table(config.DB_SOCKET).Select("remote_id").Where("status_id = ? AND farm_area_id IN (?)",
		config.GetStatus().Active, farmAreaId)
	//Get Sensor Detail
	sensorDet := r.db.Debug().Where("status_id = ?", config.GetStatus().Active).Preload("SensorType",
		"status_id = ?", config.GetStatus().Active)
	// Get Socket Detail
	socketDet := r.db.Debug().Where("status_id = ?", config.GetStatus().Active).Preload("StatusSensor",
		"status_id = ?", config.GetStatus().Active).Preload("Sensor",
		func(db *gorm.DB) *gorm.DB {
			return sensorDet
		})
	resp := r.db.Debug().Where("status_id = ? AND remote_id IN (?)", status, remoteIdList).Preload("SocSenDetail",
		func(db *gorm.DB) *gorm.DB {
			return socketDet
		}).Find(&result)

	if resp.Error != nil && !errors.Is(resp.Error, gorm.ErrRecordNotFound) {
		return nil, resp.Error
	}
	return result, nil
}

func (r *Repository) FindOneSocket(socketId string) (*model_db.Socket, error) {
	var result *model_db.Socket
	resp := r.db.Debug().Where("status_id = ? AND socket_id = ?",
		config.GetStatus().Active, socketId).First(&result)
	if resp.Error != nil {
		return nil, resp.Error
	}
	return result, nil
}

func (r *Repository) FindOneSocketDetail(socketId string) (*SocketDetail, error) {
	var result *SocketDetail
	//Get StatusSensor
	statusSensor := r.db.Debug().Where("status_id = ?", config.GetStatus().Active)
	resp := r.db.Debug().Where("status_id = ? AND socket_id = ?",
		config.GetStatus().Active, socketId).Preload("StatusSensor",
		func(db *gorm.DB) *gorm.DB {
			return statusSensor
		}).First(&result)
	if resp.Error != nil {
		return nil, resp.Error
	}
	return result, nil
}
//-------------------------------------------------------------------------------//
//									Upsert
//-------------------------------------------------------------------------------//
func (r *Repository) UpsertRemoteSwitch (req *model_db.RemoteSwitchUS) error {
	resp := r.db.Debug().Model(model_db.RemoteSwitchUS{}).Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: "remote_id"},
		},
		UpdateAll: true,
	}).Create(&req)
	if resp.Error != nil {
		return resp.Error
	}
	return nil
}

//-------------------------------------------------------------------------------//
//									Update
//-------------------------------------------------------------------------------//
func (r *Repository) UpdateSocketFieldRemote (req *RemoteDetailUS) error {
	resp := r.db.Debug().Table(config.DB_SOCKET).Where("status_id = ? AND socket_id IN (?)",
		config.GetStatus().Active, req.SocketList).Update("remote_id", req.RemoteSwitch.RemoteId)
	if resp.Error != nil {
		return resp.Error
	}
	return nil
}

func (r *Repository) UpdateNullSocketFieldRemote (socketList []string) error {
	resp := r.db.Debug().Table(config.DB_SOCKET).Where("status_id = ? AND socket_id IN (?)",
		config.GetStatus().Active, socketList).Update("remote_id", sql.NullString{})
	if resp.Error != nil {
		return resp.Error
	}
	return nil
}

func (r *Repository) UpdateStatusSensor (req *ControlSwitch, statusSensor string) error {
	resp := r.db.Debug().Table(config.DB_SOCKET).Where("status_id = ? AND socket_id = ?",
		config.GetStatus().Active, req.SocketId).Update("status_sensor_id", statusSensor)
	if resp.Error != nil {
		return resp.Error
	}
	return nil
}
//-------------------------------------------------------------------------------//
//									Delete
//-------------------------------------------------------------------------------//
func (r *Repository) DeleteRemoteSwitch (req *string) error {
	model := RemoteDetailDel{}
	resp := r.db.Debug().Table(config.DB_REMOTE_SWITCH).Delete(&model, "status_id = ? AND remote_id = ?",
		config.GetStatus().Active, req)
	if resp.Error != nil {
		return resp.Error
	}
	return nil
}



