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
type IntMyFarm interface {
	GetFarmListWithRoleer(status, uid, roleId string) ([]model_services.DashboardFarmList, int)
}

/*-------------------------------------------------------------------------------------------*/
//                                   METHOD
/*-------------------------------------------------------------------------------------------*/
func (ln Ln) GetFarmListWithRoleer(status, uid, roleId string) ([]model_services.DashboardFarmList, int) {
	var farmList []model_services.DashboardFarmList
	var total int

	sql := fmt.Sprintf("SELECT * FROM %s INNER JOIN %s ON %s.farm_id = %s.farm_id WHERE %s.status_id = '%s' AND %s.uid = '%s' AND %s.role_id = '%s'",
		config.DB_FARM, config.DB_TRANS_MANAGEMENT, config.DB_FARM, config.DB_TRANS_MANAGEMENT, config.DB_FARM, status, config.DB_TRANS_MANAGEMENT, uid, config.DB_TRANS_MANAGEMENT, roleId)
	fmt.Println(sql)
	err := ln.Db.Raw(sql).Scan(&farmList).Error
	if err != nil {
		log.Print(err)
	}

	total = len(farmList)
	return farmList, total
}
