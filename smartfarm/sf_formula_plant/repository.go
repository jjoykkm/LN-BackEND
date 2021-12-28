package sf_formula_plant

import (
	"errors"
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/common/config"
	"github.com/jjoykkm/ln-backend/common/models/model_db"
	"github.com/jjoykkm/ln-backend/smartfarm/sf_common/cm_other"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repositorier interface {
	Begin() *Repository
	Commit()
	Rollback()
	FindAllPlantType(status string) ([]model_db.PlantType, error)
	FindAllPlantWithPlantType(status, plantTypeId string, offset int) ([]PlantTypeAndPlantList, error)
	GetCountFormulaPlant(status, plantId string) (int64, error)
	FindAllFormulaPlantByPlant(status, plantId string, offset int) (*ForPlantlist, error)
	FindAllFavForPlantId(status, uid string) (map[string]bool, error)
	FindAllPlantedForPlantId(status, uid string) (map[string]bool, error)
	FindAllFormulaPlantFavorite(status, uid string, offset int) ([]FormulaPlantItem, error)
	FindAllMyFormulaPlant(status, uid string, offset int) ([]FormulaPlantItem, error)
	FindAllFormulaPlantDetail(status, forPlantId string) (*ForPlantFormula, error)
	UpsertFormulaPlant (req *model_db.FormulaPlantUS) (error, *string)
	UpsertForPlantSensor (req []model_db.TransSenValueRecUS) error
	UpsertForPlantFert (req []model_db.TransFertRatioUS) error
	UpsertFertilizer (req *model_db.FertilizerUS) error
	CreateFavFormulaPlant (req *model_db.FavForPlantUS) error
	DeleteFavFormulaPlant (req *model_db.FavForPlantUS) error
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

func (r *Repository) FindAllPlantType(status string) ([]model_db.PlantType, error) {
	var result []model_db.PlantType

	resp := r.db.Debug().Where("status_id = ? ", status).Order("change_date desc").Find(&result)
	if resp.Error != nil && !errors.Is(resp.Error, gorm.ErrRecordNotFound) {
		return nil, resp.Error
	}
	return result, nil
}

func (r *Repository) FindAllPlantWithPlantType(status, plantTypeId string, offset int) ([]PlantTypeAndPlantList, error) {
	var (
		result []PlantTypeAndPlantList
		resp *gorm.DB
	)

	if plantTypeId == config.PLANT_TYPE_ALL || plantTypeId == "" {
		// Get data exclude type all
		resp = r.db.Debug().Where("status_id = ? AND plant_type_id != ?",
			config.GetStatus().Active,config.PLANT_TYPE_ALL).Preload("Plant",
			"status_id = ?", status).Limit(LIMIT_GET_DATA).Offset(offset).Order(
			"change_date desc").Find(&result)
	}else {
		// Get data specific type
		resp = r.db.Debug().Where("status_id = ? AND plant_type_id = ?",
			config.GetStatus().Active, plantTypeId).Preload("Plant", "status_id = ?",
			status).Limit(LIMIT_GET_DATA).Offset(offset).Order("change_date desc").Find(&result)
	}

	if resp.Error != nil && !errors.Is(resp.Error, gorm.ErrRecordNotFound) {
		return nil, resp.Error
	}
	return result, nil
}

func (r *Repository) GetCountFormulaPlant(status, plantId string) (int64, error) {
	var forPlant []model_db.FormulaPlant
	var count int64

	resp := r.db.Debug().Model(&forPlant).Where("status_id = ? AND plant_id = ?", status, plantId).Count(&count)
	if resp.Error != nil && !errors.Is(resp.Error, gorm.ErrRecordNotFound) {
		return 0, resp.Error
	}
	return count, nil
}

func (r *Repository) FindAllFormulaPlantByPlant(status, plantId string, offset int) (*ForPlantlist, error) {
	var result *ForPlantlist
	// Get Province
	province := r.db.Debug().Where("status_id = ?", config.GetStatus().Active)
	// Get Country
	country := r.db.Debug().Where("status_id = ?", config.GetStatus().Active)
	// Get Owner
	//owner := r.db.Debug().Where("status_id = ?", config.GetStatus().Active)
	// Get Formula Plant
	forPlant := r.db.Debug().Where("status_id = ?",
		status).Order("change_date desc").Limit(LIMIT_GET_DATA).Offset(offset).Preload("Province",
		func(db *gorm.DB) *gorm.DB {
			return province
		}).Preload("Country",
		func(db *gorm.DB) *gorm.DB {
			return country
		}).Preload("Owner")
		//func(db *gorm.DB) *gorm.DB {
		//	return owner
		//})
	// Get PlantType
	plantType := r.db.Debug().Where("status_id = ?", config.GetStatus().Active)
	resp := r.db.Debug().Where("status_id = ? AND plant_id = ?",
		config.GetStatus().Active, plantId).Preload("PlantType",
			func(db *gorm.DB) *gorm.DB {
				return plantType
			}).Preload("ForPlantDetail",
			func(db *gorm.DB) *gorm.DB {
				return forPlant
			}).Find(&result)

	if resp.Error != nil && !errors.Is(resp.Error, gorm.ErrRecordNotFound) {
		return nil, resp.Error
	}
	return result, nil
}

func (r *Repository) FindAllFavForPlantId(status, uid string) (map[string]bool, error) {
	resultMap := make(map[string]bool)
	result:= []uuid.UUID{}

	resp := r.db.Debug().Table(config.DB_FAVORITE_PLANT).Select("formula_plant_id").Find(&result,
		"status_id = ? AND uid = ?", status, uid)

	if resp.Error != nil && !errors.Is(resp.Error, gorm.ErrRecordNotFound) {
		return nil, resp.Error
	}
	//Convert UUID to String Map
	resultMap = cm_other.ConvertUUIDtoStringMap(result)

	return resultMap, nil
}

func (r *Repository) FindAllPlantedForPlantId(status, uid string) (map[string]bool, error) {
	resultMap := make(map[string]bool)

	result:= []uuid.UUID{}
	// Get farm_id filter by uid
	subQuery := r.db.Debug().Table(config.DB_TRANS_MANAGEMENT).Select("farm_id").Where(
		"status_id = ? AND uid = ?", config.GetStatus().Active, uid)

	// Get data by farm list
	resp := r.db.Debug().Table(config.DB_FARM_AREA).Where(
		"status_id = ? AND farm_id IN (?)", status, subQuery).Select("formula_plant_id").Find(&result)
	if resp.Error != nil && !errors.Is(resp.Error, gorm.ErrRecordNotFound) {
		return nil, resp.Error
	}
	//Convert UUID to String Map
	resultMap = cm_other.ConvertUUIDtoStringMap(result)

	return resultMap, nil
}

func (r *Repository) FindAllFormulaPlantFavorite(status, uid string, offset int) ([]FormulaPlantItem, error) {
	var result []FormulaPlantItem
	// Get Province
	province := r.db.Debug().Where("status_id = ?", config.GetStatus().Active)
	// Get Country
	country := r.db.Debug().Where("status_id = ?", config.GetStatus().Active)
	// Get Owner
	//owner := r.db.Debug().Where("status_id = ?", config.GetStatus().Active)
	// Get formula_plant_id filter by uid
	forPlantId := r.db.Debug().Table(config.DB_FAVORITE_PLANT).Select("formula_plant_id").Find(&result,
		"status_id = ? AND uid = ?", status, uid)
	// Get PlantType
	plantType := r.db.Debug().Where("status_id = ?", config.GetStatus().Active).Preload("PlantType",
		"status_id = ?", config.GetStatus().Active)

	resp := r.db.Debug().Order("change_date desc").Limit(LIMIT_GET_DATA).Offset(offset).Where(
		"status_id = ? AND formula_plant_id IN (?)",
		config.GetStatus().Active, forPlantId).Preload("Plant",
			func(db *gorm.DB) *gorm.DB {
				return plantType
			}).Preload("Province",
			func(db *gorm.DB) *gorm.DB {
				return province
			}).Preload("Country",
			func(db *gorm.DB) *gorm.DB {
				return country
			}).Preload("Owner").Find(&result)
		//func(db *gorm.DB) *gorm.DB {
		//	return owner
		//})
	if resp.Error != nil && !errors.Is(resp.Error, gorm.ErrRecordNotFound) {
		return nil, resp.Error
	}
	return result, nil
}

func (r *Repository) FindAllMyFormulaPlant(status, uid string, offset int) ([]FormulaPlantItem, error) {
	var result []FormulaPlantItem
	// Get Province
	province := r.db.Debug().Where("status_id = ?", config.GetStatus().Active)
	// Get Country
	country := r.db.Debug().Where("status_id = ?", config.GetStatus().Active)
	// Get Owner
	//owner := r.db.Debug().Where("status_id = ?", config.GetStatus().Active)
	// Get PlantType
	plantCond := r.db.Debug().Where("status_id = ?", config.GetStatus().Active).Preload("PlantType",
		"status_id = ?", config.GetStatus().Active)

	resp := r.db.Debug().Order("change_date desc").Limit(LIMIT_GET_DATA).Offset(offset).Where(
		"status_id = ? AND uid = ?", status, uid).Preload("Plant",
			func(db *gorm.DB) *gorm.DB {
				return plantCond
			}).Preload("Province",
			func(db *gorm.DB) *gorm.DB {
				return province
			}).Preload("Country",
			func(db *gorm.DB) *gorm.DB {
				return country
			}).Preload("Owner").Find(&result)
		//func(db *gorm.DB) *gorm.DB {
		//	return owner
		//})

	if resp.Error != nil && !errors.Is(resp.Error, gorm.ErrRecordNotFound) {
		return nil, resp.Error
	}
	return result, nil
}

func (r *Repository) FindAllFormulaPlantDetail(status, forPlantId string) (*ForPlantFormula, error) {
	var result *ForPlantFormula
	// Get FertilizerCat
	FertCatId := r.db.Debug().Where("status_id = ?", config.GetStatus().Active).Preload("FertilizerCat",
		"status_id = ?", config.GetStatus().Active)
	// Get Fertilizer
	FertId := r.db.Debug().Where("status_id = ?", config.GetStatus().Active).Preload("Fertilizer",
		func(db *gorm.DB) *gorm.DB {
			return FertCatId
		})
	// Get SensorType
	SensorTypeId := r.db.Debug().Where("status_id = ?", config.GetStatus().Active).Preload(
		"SensorType","status_id = ?", config.GetStatus().Active)
	// Get Province
	province := r.db.Debug().Where("status_id = ?", config.GetStatus().Active)
	// Get Country
	country := r.db.Debug().Where("status_id = ?", config.GetStatus().Active)
	// Get Owner
	//owner := r.db.Debug().Where("status_id = ?", config.GetStatus().Active)
	// Get PlantType
	plantType := r.db.Debug().Where("status_id = ?", config.GetStatus().Active).Preload("PlantType",
		"status_id = ?", config.GetStatus().Active)
	resp := r.db.Debug().Where("status_id = ? AND formula_plant_id = ?",
		status, forPlantId).Preload("Plant",
		func(db *gorm.DB) *gorm.DB {
			return plantType
		}).Preload("Province",
		func(db *gorm.DB) *gorm.DB {
			return province
		}).Preload("Country",
		func(db *gorm.DB) *gorm.DB {
			return country
		}).Preload("Owner").Preload("ForPlantFert",
		func(db *gorm.DB) *gorm.DB {
			return FertId
		}).Preload("ForPlantSensor",
		func(db *gorm.DB) *gorm.DB {
			return SensorTypeId
		}).Find(&result)
	if resp.Error != nil && !errors.Is(resp.Error, gorm.ErrRecordNotFound) {
		return nil, resp.Error
	}
	return result, nil
}

//-------------------------------------------------------------------------------//
//									Upsert
//-------------------------------------------------------------------------------//
func (r *Repository) UpsertFormulaPlant (req *model_db.FormulaPlantUS) (error, *string) {
		resp := r.db.Debug().Model(model_db.FormulaPlantUS{}).Clauses(clause.OnConflict{
			Columns: []clause.Column{
				{Name: "formula_plant_id"},
			},
			UpdateAll: true,
		}).Create(&req)
		if resp.Error != nil {
			return resp.Error, nil
		}
	return nil, &req.FormulaPlantId
}

func (r *Repository) UpsertForPlantSensor (req []model_db.TransSenValueRecUS) error {
	resp := r.db.Debug().Model(model_db.TransSenValueRecUS{}).Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: "formula_plant_id"},
			{Name: "sensor_type_id"},
		},
		UpdateAll: true,
	}).Create(&req)
	if resp.Error != nil {
		return resp.Error
	}
	return nil
}

func (r *Repository) UpsertForPlantFert (req []model_db.TransFertRatioUS) error {
	resp := r.db.Debug().Model(model_db.TransFertRatioUS{}).Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: "formula_plant_id"},
			{Name: "fertilizer_id"},
		},
		UpdateAll: true,
	}).Create(&req)
	if resp.Error != nil {
		return resp.Error
	}
	return nil
}

func (r *Repository) UpsertFertilizer (req *model_db.FertilizerUS) error {
	resp := r.db.Debug().Model(model_db.FertilizerUS{}).Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: "fertilizer_id"},
		},
		UpdateAll: true,
	}).Create(&req)
	if resp.Error != nil {
		return resp.Error
	}
	return nil
}

//-------------------------------------------------------------------------------//
//									Create
//-------------------------------------------------------------------------------//
func (r *Repository) CreateFavFormulaPlant (req *model_db.FavForPlantUS) error {
	req.StatusId = config.GetStatus().Active
	resp := r.db.Debug().Create(&req)
	if resp.Error != nil {
		return resp.Error
	}
	return nil
}

//-------------------------------------------------------------------------------//
//									Delete
//-------------------------------------------------------------------------------//
func (r *Repository) DeleteFavFormulaPlant (req *model_db.FavForPlantUS) error {
	resp := r.db.Debug().Where("uid = ? AND formula_plant_id = ?", req.Uid, req.FormulaPlantId).Delete(&req)
	if resp.Error != nil {
		return resp.Error
	}
	return nil
}