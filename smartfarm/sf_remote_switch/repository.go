package sf_remote_switch

import (
	"database/sql"
	"github.com/jjoykkm/ln-backend/common/config"
	"github.com/jjoykkm/ln-backend/common/models/model_db"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repositorier interface {
	Begin() *Repository
	Commit()
	Rollback()
	UpsertRemoteSwitch (req *model_db.RemoteSwitchUS) error
	UpdateSocketFieldRemote (req *RemoteDetailUS) error
	UpdateNullSocketFieldRemote (req []string) error
	DeleteRemoteSwitch (req *string) error
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



