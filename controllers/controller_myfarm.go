package controllers

import (
	"LN-BackEND/config"
	"LN-BackEND/models/model_databases"
	"LN-BackEND/models/model_services"
	"fmt"
	"log"
	"strings"
)

/*-------------------------------------------------------------------------------------------*/
//                                 INTERFACE
/*-------------------------------------------------------------------------------------------*/
type IntMyFarm interface {
	GetFarmListWithRoleer(status, uid, roleId string) ([]model_services.DashboardFarmList, int)
	GetOverviewFarmer(status, farmId string) (model_services.MyFarmOverviewFarm)
	GetTransSocketAreaer(status, farmId string) ([]model_databases.TransSocketArea, []string, []string, []string, int)
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

func (ln Ln) GetOverviewFarmer(status, farmId string) (model_services.MyFarmOverviewFarm) {
	var farmList []model_databases.TransSocketArea
	mbList := make(map[string]bool)
	faList := make(map[string]bool)
	var farmOver model_services.MyFarmOverviewFarm

	// Get Farm detail
	sqlFarm := fmt.Sprintf("SELECT * FROM %s WHERE status_id = '%s' AND farm_id = '%s' ",
		config.DB_FARM, config.STATUS_ACTIVE, farmId)
	fmt.Println(sqlFarm)
	errFarm := ln.Db.Raw(sqlFarm).Scan(&farmOver).Error
	if errFarm != nil {
		log.Print(errFarm)
	}

	//Get Farm Area detail
	_, farmAreaIdList, _ := IntCommon.GetFarmAreaByFarmId(ln, config.STATUS_ACTIVE, farmId)
	sqlIn := "('" + strings.Join(farmAreaIdList, "','") + "')"
	sql := fmt.Sprintf("SELECT * FROM %s WHERE status_id = '%s' AND farm_area_id IN %s ",
		config.DB_TRANS_SOCKET_AREA, status, sqlIn)
	fmt.Println(sql)
	err := ln.Db.Raw(sql).Scan(&farmList).Error
	if err != nil {
		log.Print(err)
	}

	for _, array := range farmList {
		//Mainbox unique
		if _, value := mbList[array.MainboxId.UUID.String()]; !value {
			mbList[array.MainboxId.UUID.String()] = true
		}
		//FarmArea unique
		if _, value := faList[array.FarmAreaId.UUID.String()]; !value {
			faList[array.FarmAreaId.UUID.String()] = true
		}
	}

	farmOver.MainboxCount = len(mbList)
	farmOver.FarmAreaCount = len(faList)

	return farmOver
}

func (ln Ln) GetTransSocketAreaer(status, farmId string) ([]model_databases.TransSocketArea, []string, []string, []string, int) {
	var joinArray []model_databases.TransSocketArea
	var socketIdList []string
	var farmAreaIdList []string
	var mainboxIdList []string
	var total int

	// Get farmArea
	sql := fmt.Sprintf("SELECT * FROM %s INNER JOIN %s ON %s.farm_area_id = %s.farm_area_id WHERE %s.status_id = '%s' AND %s.farm_id = '%s'",
		config.DB_FARM_AREA, config.DB_TRANS_SOCKET_AREA, config.DB_FARM_AREA, config.DB_TRANS_SOCKET_AREA, config.DB_FARM_AREA, status, config.DB_FARM_AREA, farmId)
	fmt.Println(sql)
	err := ln.Db.Raw(sql).Scan(&joinArray).Error
	if err != nil {
		log.Print(err)
	}

	for _, join := range joinArray {
		socketIdList = append(socketIdList, join.SocketId.UUID.String())
		farmAreaIdList = append(farmAreaIdList, join.FarmAreaId.UUID.String())
		mainboxIdList = append(mainboxIdList, join.MainboxId.UUID.String())
	}
	total = len(joinArray)

	return joinArray, farmAreaIdList, socketIdList, mainboxIdList, total
}

