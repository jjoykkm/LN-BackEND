package controllers

import (
	"LN-BackEND/config"
	"LN-BackEND/models/model_services"
	"fmt"
	"log"
)

/*-------------------------------------------------------------------------------------------*/
//                                 INTERFACE
/*-------------------------------------------------------------------------------------------*/
type IntDashboard interface {
	GetFarmLister(status, uid string) ([]model_services.DashboardFarmList, int)
	GetFarmAreaLister(status, farmId string) ([]model_services.DashboardFarmList, int)
}

/*-------------------------------------------------------------------------------------------*/
//                                   METHOD
/*-------------------------------------------------------------------------------------------*/
func (ln Ln) GetFarmLister(status, uid string) ([]model_services.DashboardFarmList, int) {
	var farmList []model_services.DashboardFarmList
	var total int

	sql := fmt.Sprintf("SELECT * FROM %s INNER JOIN %s ON %s.farm_id = %s.farm_id WHERE %s.status_id = '%s' AND %s.uid = '%s'",
		config.DB_FARM, config.DB_TRANS_MANAGEMENT, config.DB_FARM, config.DB_TRANS_MANAGEMENT, config.DB_FARM, status, config.DB_TRANS_MANAGEMENT, uid)
	fmt.Println(sql)
	err := ln.Db.Raw(sql).Scan(&farmList).Error
	if err != nil {
		log.Print(err)
	}

	total = len(farmList)
	return farmList, total
}

func (ln Ln) GetFarmAreaLister(status, farmId string) ([]model_services.DashboardFarmAreaList, int) {
	var farmAreaList []model_services.DashboardFarmAreaList
	var total int

	sql := fmt.Sprintf("SELECT * FROM %s INNER JOIN %s ON %s.formula_plant_id = %s.formula_plant_id WHERE %s.status_id = '%s' AND %s.farm_id = '%s'",
		config.DB_FARM_AREA, config.DB_FORMULA_PLANT, config.DB_FARM_AREA, config.DB_FORMULA_PLANT, config.DB_FARM_AREA, status, config.DB_FARM_AREA, farmId)
	fmt.Println(sql)
	err := ln.Db.Raw(sql).Scan(&farmAreaList).Error
	if err != nil {
		log.Print(err)
	}

	total = len(farmAreaList)
	return farmAreaList, total
}