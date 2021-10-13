package sf_my_farm

import (
	"errors"
	"github.com/jjoykkm/ln-backend/common/models/model_db"
	"gorm.io/gorm"
)

type Repositorier interface {
	FindAllPlantType(status string) ([]model_db.PlantType, error)
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

func (r *Repository) FindAllPlantType(status string) ([]model_db.PlantType, error) {
	var result []model_db.PlantType

	err := r.db.Where("status_id = ? ", status).Find(&result).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return result, nil
}

func (r *Repository) GetCountMainbox(status, mainboxId string) int64 {
	var forPlant []model_db.FormulaPlant
	var count int64

	resp := r.db.Debug().Model(&forPlant).Where("status_id = ? AND mainbox_id = ?", status, mainboxId).Count(&count)
	if resp.Error != nil && !errors.Is(resp.Error, gorm.ErrRecordNotFound) {
		return 0
	}
	return count
}


func (r *Repository) GetCountFarmArea(status, farmId string) int64 {
	var forPlant []model_db.FarmArea
	var count int64

	resp := r.db.Debug().Model(&forPlant).Where("status_id = ? AND farm_id = ?", status, farmId).Count(&count)
	if resp.Error != nil && !errors.Is(resp.Error, gorm.ErrRecordNotFound) {
		return 0
	}
	return count
}
