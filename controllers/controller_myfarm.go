package controllers

import (
	"LN-BackEND/config"
	"LN-BackEND/models/model_databases"
	"LN-BackEND/models/model_services"
	"fmt"
	"github.com/mitchellh/mapstructure"
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
	GetSocketByIder(status string, socketIdList []string) ([]model_databases.Socket, map[string]model_databases.Socket)
	GetSocketWithSensorer(status, language string, socketIdList []string) ([]model_services.MyFarmSenSocDetail, map[string]model_services.MyFarmSenSocDetail, int)
	GetSocSenByKeyer(mainboxId, farmAreaId string, tranSoc []model_databases.TransSocketArea, socSenMap map[string]model_services.MyFarmSenSocDetail) ([]model_services.MyFarmSenSocDetail, int)
	GetManageMainboxer(status, language, farmId string) ([]model_services.MyFarmManageMainbox, int)
	GetFarmAreaByIder(status string, farmAreaIdList []string) ([]model_databases.FarmArea, map[string]model_databases.FarmArea)
	GetManageFarmAreaer(status, language, farmId string) ([]model_services.MyFarmManageFarmArea, int)
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

func (ln Ln) GetSocketByIder(status string, socketIdList []string) ([]model_databases.Socket, map[string]model_databases.Socket) {
	var socketAr []model_databases.Socket
	var socketMap map[string]model_databases.Socket

	socketMap = make(map[string]model_databases.Socket)

	sqlIn := "('" + strings.Join(socketIdList, "','") + "')"
	sql := fmt.Sprintf("SELECT * FROM %s WHERE status_id = '%s' AND socket_id IN %s",
		config.DB_SOCKET, status, sqlIn)
	fmt.Println(sql)
	err := ln.Db.Raw(sql).Scan(&socketAr).Error
	if err != nil {
		log.Print(err)
	}

	for _, wa := range socketAr {
		socketMap[wa.SocketId.UUID.String()] = wa
	}
	return socketAr, socketMap
}

func (ln Ln) GetSocketWithSensorer(status, language string, socketIdList []string) ([]model_services.MyFarmSenSocDetail, map[string]model_services.MyFarmSenSocDetail, int) {
	var joinArray []model_services.MyFarmSenSocDetail
	var found bool
	var total int

	sensorTypeMap := make(map[string]string)
	statusSensorMap := make(map[string]string)
	socSenMap := make(map[string]model_services.MyFarmSenSocDetail)

	sqlIn := "('" + strings.Join(socketIdList, "','") + "')"
	sql := fmt.Sprintf("SELECT * FROM %s INNER JOIN %s ON %s.sensor_id = %s.sensor_id WHERE %s.status_id = '%s' AND %s.socket_id IN %s",
		config.DB_SOCKET, config.DB_SENSOR, config.DB_SOCKET, config.DB_SENSOR, config.DB_SOCKET, status, config.DB_SOCKET, sqlIn)
	fmt.Println(sql)
	err := ln.Db.Raw(sql).Scan(&joinArray).Error
	if err != nil {
		log.Print(err)
	}

	for idx, wa := range joinArray {
		//Get Sensor Type name
		wa.SensorTypeName, found = sensorTypeMap[wa.SensorTypeId.UUID.String()]
		if !found {
			_, wa.SensorTypeName = IntCommon.GetSensorTypeNameer(ln, wa.SensorTypeId.UUID.String(), language)
			sensorTypeMap[wa.SensorTypeId.UUID.String()] = wa.SensorTypeName
		}

		//Get Status Sensor name
		wa.StatusSensorName, found = statusSensorMap[wa.StatusSensorId.UUID.String()]
		if !found {
			_, wa.StatusSensorName= IntDashboard.GetStatusSensorer(ln, wa.StatusSensorId.UUID.String())
			statusSensorMap[wa.StatusSensorId.UUID.String()] = wa.StatusSensorName
		}
		joinArray[idx] = wa
		socSenMap[wa.SocketId.UUID.String()] = wa
	}

	total = len(joinArray)

	return joinArray, socSenMap, total
}

func (ln Ln) GetSocSenByKeyer(mainboxId, farmAreaId string, tranSoc []model_databases.TransSocketArea, socSenMap map[string]model_services.MyFarmSenSocDetail) ([]model_services.MyFarmSenSocDetail, int){
	var list []model_services.MyFarmSenSocDetail
	var total int

	for _, wa := range tranSoc {
		if mainboxId != "" {
			if wa.MainboxId.UUID.String() == mainboxId {
				socSen, _ := socSenMap[wa.SocketId.UUID.String()]
				list = append(list, socSen)
			}
		}else if farmAreaId != "" {
			if wa.FarmAreaId.UUID.String() == farmAreaId {
				socSen, _ := socSenMap[wa.SocketId.UUID.String()]
				list = append(list, socSen)
			}
		}
	}
	total = len(list)

	return list, total
}

func (ln Ln) GetManageMainboxer(status, language, farmId string) ([]model_services.MyFarmManageMainbox, int) {
	var manageMB model_services.MyFarmManageMainbox
	var manageMBList []model_services.MyFarmManageMainbox
	var total int

	tranSoc, _, socketIdList, mainboxIdList, total := IntMyFarm.GetTransSocketAreaer(ln, config.STATUS_ACTIVE, farmId)
	if total == 0 {
		return manageMBList, total
	}

	// Get Socket and Sensor
	_, socSenMap, _ := IntMyFarm.GetSocketWithSensorer(ln, status, language, socketIdList)

	// Get Mainbox
	mainbox, _ := IntDashboard.GetMainboxByIder(ln, status, mainboxIdList)
	for _, m := range mainbox {
		mapstructure.Decode(m, &manageMB)
		// Get Socket and Sensor by mainbox
		manageMB.SenSocDetail, _ = IntMyFarm.GetSocSenByKeyer(ln, m.MainboxId.UUID.String(), "", tranSoc, socSenMap)
		manageMBList = append(manageMBList, manageMB)
	}

	total = len(manageMBList)

	return manageMBList, total
}

func (ln Ln) GetFarmAreaByIder(status string, farmAreaIdList []string) ([]model_databases.FarmArea, map[string]model_databases.FarmArea) {
	var farmAreaAr []model_databases.FarmArea
	var farmAreaMap map[string]model_databases.FarmArea

	farmAreaMap = make(map[string]model_databases.FarmArea)

	sqlIn := "('" + strings.Join(farmAreaIdList, "','") + "')"
	sql := fmt.Sprintf("SELECT * FROM %s WHERE status_id = '%s' AND farm_area_id IN %s",
		config.DB_FARM_AREA, status, sqlIn)
	fmt.Println(sql)
	err := ln.Db.Raw(sql).Scan(&farmAreaAr).Error
	if err != nil {
		log.Print(err)
	}

	for _, wa := range farmAreaAr {
		farmAreaMap[wa.FarmAreaId.UUID.String()] = wa
	}
	return farmAreaAr, farmAreaMap
}

func (ln Ln) GetManageFarmAreaer(status, language, farmId string) ([]model_services.MyFarmManageFarmArea, int) {
	var manageFA model_services.MyFarmManageFarmArea
	var manageFAList []model_services.MyFarmManageFarmArea
	var total int

	tranSoc, farmAreaIdList, socketIdList, _, total := IntMyFarm.GetTransSocketAreaer(ln, config.STATUS_ACTIVE, farmId)
	if total == 0 {
		return manageFAList, total
	}

	// Get Socket and Sensor
	_, socSenMap, _ := IntMyFarm.GetSocketWithSensorer(ln, status, language, socketIdList)

	// Get FarmArea
	farmArea, _ := IntMyFarm.GetFarmAreaByIder(ln, status, farmAreaIdList)
	for _, fa := range farmArea {
		mapstructure.Decode(fa, &manageFA)
		// Get Socket and Sensor by farm area
		manageFA.SenSocDetail, _ = IntMyFarm.GetSocSenByKeyer(ln, "", fa.FarmAreaId.UUID.String(), tranSoc, socSenMap)
		manageFAList = append(manageFAList, manageFA)
	}

	total = len(manageFAList)

	return manageFAList, total
}

