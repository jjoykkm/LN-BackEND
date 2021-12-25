package sf_formula_plant

import (
	"errors"
	"fmt"
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/common/config"
	"github.com/jjoykkm/ln-backend/common/models/model_db"
	"github.com/jjoykkm/ln-backend/helper"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repositorier interface {
	Begin() *Repository
	Commit()
	Rollback()
	FindAllPlantType(status string) ([]model_db.PlantType, error)
	FindAllPlantWithPlantType(status, plantTypeId string, offset int) ([]PlantAndPlantType, error)
	GetCountFormulaPlant(status, plantId string) (int64, error)
	FindAllFormulaPlantByPlant(status, plantId string, offset int) ([]FormulaPlantItem, error)
	FindAllFavForPlantId(status, resultType, uid string) ([]uuid.UUID, map[string]bool, error)
	FindAllPlantedForPlantId(status, resultType, uid string) ([]uuid.UUID, map[string]bool, error)
	FindAllFormulaPlantFavorite(status, uid string, offset int) ([]FormulaPlantItem, error)
	FindAllMyFormulaPlant(status, uid string, offset int) ([]FormulaPlantItem, error)
	FindAllFormulaPlantDetail(status, forPlantId string) ([]ForPlantFormula, error)
	UpsertFormulaPlant (req *model_db.FormulaPlantUS) (error, *string)
	UpsertForPlantSensor (req []model_db.TransSenValueRecUS) error
	UpsertForPlantFert (req []model_db.TransFertRatioUS) error
	UpsertFertilizer (req *model_db.FertilizerUS) error
	CreateFavFormulaPlant (req *model_db.FavForPlantUS) error
	DeleteFavFormulaPlant (req *model_db.FavForPlantUS) error
	Test (req *ForPlantUS) error
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

func (r *Repository) FindAllPlantWithPlantType(status, plantTypeId string, offset int) ([]PlantAndPlantType, error) {
	var result []PlantAndPlantType
	var sqlWhere string
	// Generate condition when get plant
	sqlWhere = fmt.Sprintf("%s.status_id = ?", config.DB_PLANT)
	if plantTypeId == config.PLANT_TYPE_ALL || plantTypeId == "" {
		sqlWhere = sqlWhere + fmt.Sprintf(" AND %s.plant_type_id = ?", config.DB_PLANT)
	}
	resp := r.db.Debug().Where(sqlWhere, status, plantTypeId).Preload("PlantType",
		"status_id = ?", config.GetStatus().Active).Limit(LIMIT_GET_DATA).Offset(offset).Order(
			"change_date desc").Find(&result)
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

func (r *Repository) FindAllFormulaPlantByPlant(status, plantId string, offset int) ([]FormulaPlantItem, error) {
	var result []FormulaPlantItem
	var sqlWhere string

	sqlWhere = fmt.Sprintf("status_id = '%s'",status)
	// Generate condition when get plant
	if plantId != "" {
		sqlWhere = sqlWhere + fmt.Sprintf(" AND plant_id = '%s'", plantId)
	}
	// Get PlantType
	plantTypeId := r.db.Debug().Where("status_id = ?", config.GetStatus().Active).Preload("PlantType",
		"status_id = ?", config.GetStatus().Active)

	resp := r.db.Debug().Order("change_date desc, formula_name").Limit(LIMIT_GET_DATA).Offset(offset).Where(sqlWhere).Preload("Plant",
		func(db *gorm.DB) *gorm.DB {
			return plantTypeId
		}).Preload("Province", "status_id = ?", config.GetStatus().Active).Preload("Country",
			"status_id = ?", config.GetStatus().Active).Preload("Owner").Find(&result)
	//.Preload("Owner", "status_id = ?", config.GetStatus().Active)

	if resp.Error != nil && !errors.Is(resp.Error, gorm.ErrRecordNotFound) {
		return nil, resp.Error
	}
	return result, nil
}

func (r *Repository) FindAllFavForPlantId(status, resultType, uid string) ([]uuid.UUID, map[string]bool, error) {
	resultMap := make(map[string]bool)
	result:= []uuid.UUID{}

	resp := r.db.Debug().Table(config.DB_FAVORITE_PLANT).Select("formula_plant_id").Find(&result,
		"status_id = ? AND uid = ?", status, uid)
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
	// Get farm_id filter by uid
	subQuery := r.db.Debug().Table(config.DB_TRANS_MANAGEMENT).Select("farm_id").Where(
		"status_id = ? AND uid = ?", config.GetStatus().Active, uid)

	// Get data by farm list
	resp := r.db.Debug().Table(config.DB_FARM_AREA).Where(
		"status_id = ? AND farm_id IN (?)", status, subQuery).Select("formula_plant_id").Find(&result)
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
	// Get formula_plant_id filter by uid
	forPlantId := r.db.Debug().Table(config.DB_FAVORITE_PLANT).Select("formula_plant_id").Find(&result,
		"status_id = ? AND uid = ?", status, uid)

	// Get PlantType
	plantCond := r.db.Debug().Where("status_id = ?", config.GetStatus().Active).Preload("PlantType",
		"status_id = ?", config.GetStatus().Active)

	resp := r.db.Debug().Order("change_date desc, formula_name").Limit(LIMIT_GET_DATA).Offset(offset).Where(
		"status_id = ? AND formula_plant_id IN (?)", config.GetStatus().Active, forPlantId).Preload("Plant",
			func(db *gorm.DB) *gorm.DB {
				return plantCond
			}).Preload("Province", "status_id = ?", config.GetStatus().Active).Preload("Country",
				"status_id = ?", config.GetStatus().Active).Preload("Owner").Find(&result)
	//.Preload("Owner", "status_id = ?", config.GetStatus().Active)

	if resp.Error != nil && !errors.Is(resp.Error, gorm.ErrRecordNotFound) {
		return nil, resp.Error
	}
	return result, nil
}

func (r *Repository) FindAllMyFormulaPlant(status, uid string, offset int) ([]FormulaPlantItem, error) {
	var result []FormulaPlantItem
	// Get PlantType
	plantCond := r.db.Debug().Where("status_id = ?", config.GetStatus().Active).Preload("PlantType",
		"status_id = ?", config.GetStatus().Active)

	resp := r.db.Debug().Order("change_date desc, formula_name").Limit(LIMIT_GET_DATA).Offset(offset).Where(
		"status_id = ? AND uid = ?", status, uid).Preload("Plant",
			func(db *gorm.DB) *gorm.DB {
				return plantCond
			}).Preload("Province", "status_id = ?", config.GetStatus().Active).Preload("Country",
				"status_id = ?", config.GetStatus().Active).Preload("Owner").Find(&result)
	//.Preload("Owner", "status_id = ?", config.GetStatus().Active)

	if resp.Error != nil && !errors.Is(resp.Error, gorm.ErrRecordNotFound) {
		return nil, resp.Error
	}
	return result, nil
}

func (r *Repository) FindAllFormulaPlantDetail(status, forPlantId string) ([]ForPlantFormula, error) {
	var result []ForPlantFormula
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

	resp := r.db.Debug().Where("status_id = ? AND formula_plant_id = ?", status, forPlantId).Preload("ForPlantFert",
			func(db *gorm.DB) *gorm.DB {
				return FertId
			}).Preload("ForPlantSensor",
			func(db *gorm.DB) *gorm.DB {
				return SensorTypeId
			}).Preload("Owner").Find(&result)
	//.Preload("Owner", "status_id = ?", config.GetStatus().Active)
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

func (r *Repository) Test (req *ForPlantUS) error {
	req.FormulaPlant.StatusId = config.GetStatus().Active
	for idx, wa := range req.SensorValue {
		wa.FormulaPlantId = req.FormulaPlant.FormulaPlantId
		wa.StatusId = config.GetStatus().Active
		req.SensorValue[idx] = wa
	}
	fmt.Printf("%+v\n", req)
	//resp := r.db.Debug().Model(ForPlantUS{}).Clauses(clause.OnConflict{
	//	Columns: []clause.Column{
	//		{Name: "formula_plant_id"},
	//	},
	//	UpdateAll: true,
	//}).Create(&req)
	//if resp.Error != nil {
	//	return resp.Error
	//}


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