package sf_formula_plant

import (
	"errors"
	"fmt"
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/config"
	"github.com/jjoykkm/ln-backend/helper"
	"github.com/jjoykkm/ln-backend/models/model_databases"
	"gorm.io/gorm"
)

type Repositorier interface {
	FindAllPlantType(status string) ([]model_databases.PlantType, error)
	FindAllPlantWithPlantType(status, plantTypeId string, offset int) ([]JoinPlantAndPlantType, error)
	GetCountFormulaPlant(status, plantId string) int64
	//FindAllFavoriteFormulaPlant(status, uid string) ([]model_databases.FavoritePlant, error)
	FindAllFormulaPlantByPlant(status, plantId string, offset int) ([]FormulaPlantItem, error)
	FindAllFavForPlantId(status, resultType, uid string) ([]uuid.UUID, map[string]bool, error)
	FindAllPlantedForPlantId(status, resultType, uid string) ([]uuid.UUID, map[string]bool, error)
	FindAllFormulaPlantFavorite(status, uid string, offset int) ([]FormulaPlantItem, error)
	FindAllMyFormulaPlant(status, uid string, offset int) ([]FormulaPlantItem, error)
	FindAllFormulaPlantDetail(status, forPlantId string) ([]ForPlantFormula, error)
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repositorier {
	return &Repository{db: db}
}

func (r *Repository) FindAllPlantType(status string) ([]model_databases.PlantType, error) {
	var result []model_databases.PlantType

	resp := r.db.Debug().Where("status_id = ? ", status).Find(&result)
	if resp.Error != nil && !errors.Is(resp.Error, gorm.ErrRecordNotFound) {
		return nil, resp.Error
	}
	return result, nil
}

func (r *Repository) FindAllPlantWithPlantType(status, plantTypeId string, offset int) ([]JoinPlantAndPlantType, error) {
	var result []JoinPlantAndPlantType
	var sqlWhere string
	// Generate condition when get plant
	sqlWhere = fmt.Sprintf("%s.status_id = ?",config.DB_PLANT)
	if plantTypeId != "" {
		sqlWhere = sqlWhere + fmt.Sprintf(" AND %s.plant_type_id = ?", config.DB_PLANT)
	}
	resp := r.db.Debug().Where(sqlWhere, status, plantTypeId).Preload("PlantType", "status_id = ?", config.GetStatus().Active).Limit(LIMIT_GET_DATA).Offset(offset).Find(&result)
	if resp.Error != nil && !errors.Is(resp.Error, gorm.ErrRecordNotFound) {
		return nil, resp.Error
	}
	return result, nil
}

func (r *Repository) GetCountFormulaPlant(status, plantId string) int64 {
	var forPlant []model_databases.FormulaPlant
	var count int64

	resp := r.db.Debug().Model(&forPlant).Where("status_id = ? AND plant_id = ?", status, plantId).Count(&count)
	if resp.Error != nil && !errors.Is(resp.Error, gorm.ErrRecordNotFound) {
		return 0
	}
	return count
}

//func (r *Repository) FindAllFavoriteFormulaPlant(status, uid string) ([]model_databases.FavoritePlant, error) {
//	var result []model_databases.FavoritePlant
//
//	resp := r.db.Debug().Where("status_id = ? AND uid = ?", status, uid).Order("change_date desc").Find(&result)
//	if resp.Error != nil && !errors.Is(resp.Error, gorm.ErrRecordNotFound) {
//		return nil, resp.Error
//	}
//	return result, nil
//}

func (r *Repository) FindAllFormulaPlantByPlant(status, plantId string, offset int) ([]FormulaPlantItem, error) {
	var result []FormulaPlantItem
	var sqlWhere string

	sqlWhere = fmt.Sprintf("status_id = '%s'",status)
	// Generate condition when get plant
	if plantId != "" {
		sqlWhere = sqlWhere + fmt.Sprintf(" AND plant_id = '%s'", plantId)
	}
	resp := r.db.Debug().Order("create_date desc, formula_name").Limit(LIMIT_GET_DATA).Offset(offset).Where(sqlWhere).Preload("Plant", func(db *gorm.DB) *gorm.DB {
		return db.Where("status_id = ?", config.GetStatus().Active).Preload("PlantType", "status_id = ?", config.GetStatus().Active)
	}).Preload("Province", "status_id = ?", config.GetStatus().Active).Preload("Country", "status_id = ?", config.GetStatus().Active).Find(&result)

	if resp.Error != nil && !errors.Is(resp.Error, gorm.ErrRecordNotFound) {
		return nil, resp.Error
	}
	return result, nil
}

func (r *Repository) FindAllFavForPlantId(status, resultType, uid string) ([]uuid.UUID, map[string]bool, error) {
	var resultMap map[string]bool
	result:= []uuid.UUID{}

	resp := r.db.Debug().Table(config.DB_FAVORITE_PLANT).Select("formula_plant_id").Find(&result, "status_id = ? AND uid = ?", status, uid)
	if resp.Error != nil && !errors.Is(resp.Error, gorm.ErrRecordNotFound) {
		return nil, nil, resp.Error
	}
	if resultType == config.GetResType().Map {
		resultMap = helper.ConvertUUIDtoStringMap(result)
	}
	return result, resultMap, nil
}

func (r *Repository) FindAllPlantedForPlantId(status, resultType, uid string) ([]uuid.UUID, map[string]bool, error) {
	var resultMap map[string]bool
	result:= []uuid.UUID{}

	resp := r.db.Debug().Table(config.DB_FARM_AREA).Where(
		"status_id = ? AND farm_id IN (?)", status, r.db.Table(config.DB_TRANS_MANAGEMENT).Select("farm_id").Where(
			"status_id = ? AND uid = ?",config.GetStatus().Active, uid)).Select("formula_plant_id").Find(&result)
	if resp.Error != nil && !errors.Is(resp.Error, gorm.ErrRecordNotFound) {
		return nil, nil, resp.Error
	}
	if resultType == config.GetResType().Map {
		resultMap = helper.ConvertUUIDtoStringMap(result)
	}
	return result, resultMap, nil
}

func (r *Repository) FindAllFormulaPlantFavorite(status, uid string, offset int) ([]FormulaPlantItem, error) {
	var result []FormulaPlantItem

	resp := r.db.Debug().Order("create_date desc, formula_name").Limit(LIMIT_GET_DATA).Offset(offset).Where(
		"status_id = ? AND formula_plant_id IN (?)", config.GetStatus().Active, r.db.Debug().Table(config.DB_FAVORITE_PLANT).Select("formula_plant_id").Find(&result,
			"status_id = ? AND uid = ?", status, uid)).Preload("Plant", func(db *gorm.DB) *gorm.DB {
		return db.Where("status_id = ?", config.GetStatus().Active).Preload("PlantType", "status_id = ?", config.GetStatus().Active)
	}).Preload("Province", "status_id = ?", config.GetStatus().Active).Preload("Country", "status_id = ?", config.GetStatus().Active).Find(&result)

	if resp.Error != nil && !errors.Is(resp.Error, gorm.ErrRecordNotFound) {
		return nil, resp.Error
	}
	return result, nil
}

func (r *Repository) FindAllMyFormulaPlant(status, uid string, offset int) ([]FormulaPlantItem, error) {
	var result []FormulaPlantItem

	resp := r.db.Debug().Order("create_date desc, formula_name").Limit(LIMIT_GET_DATA).Offset(offset).Where(
		"status_id = ? AND uid = ?", status, uid).Preload("Plant", func(db *gorm.DB) *gorm.DB {
		return db.Where("status_id = ?", config.GetStatus().Active).Preload("PlantType", "status_id = ?", config.GetStatus().Active)
	}).Preload("Province", "status_id = ?", config.GetStatus().Active).Preload("Country", "status_id = ?", config.GetStatus().Active).Find(&result)

	if resp.Error != nil && !errors.Is(resp.Error, gorm.ErrRecordNotFound) {
		return nil, resp.Error
	}
	return result, nil
}

func (r *Repository) FindAllFormulaPlantDetail(status, forPlantId string) ([]ForPlantFormula, error) {
	var result []ForPlantFormula

	resp := r.db.Debug().Where("status_id = ? AND formula_plant_id = ?", status, forPlantId).Preload("ForPlantFert",func(db *gorm.DB) *gorm.DB {
		return db.Debug().Where("status_id = ?", config.GetStatus().Active).Preload("Fertilizer",func(db *gorm.DB) *gorm.DB {
			return db.Debug().Where("status_id = ?", config.GetStatus().Active).Preload("FertilizerCat","status_id = ?", config.GetStatus().Active)
		}) } ).Preload("ForPlantSensor",func(db *gorm.DB) *gorm.DB {
		return db.Debug().Where("status_id = ?", config.GetStatus().Active).Preload(
			"SensorType","status_id = ?", config.GetStatus().Active)
	}).Find(&result)
	if resp.Error != nil && !errors.Is(resp.Error, gorm.ErrRecordNotFound) {
		return nil, resp.Error
	}
	return result, nil
}