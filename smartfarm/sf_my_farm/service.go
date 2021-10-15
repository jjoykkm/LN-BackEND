package sf_my_farm

import (
	"github.com/jjoykkm/ln-backend/common/config"
	"github.com/jjoykkm/ln-backend/common/models/model_other"
)

type Servicer interface {
	//GetFarmListWithRoleer(status, uid, roleId string) ([]model_services.DashboardFarmList, int)
	GetAuthorizeCheckForManageFarm(uid, farmId string) (bool, error)
	// FarmId
	GetOverviewFarm(status string, ReqModel *model_other.ReqModel) (*model_other.RespModel, error)
	//GetTransSocketAreaer(status, farmId string) ([]model_databases.TransSocketArea, []string, []string, []string, int)
	//GetSocketByIder(status string, socketIdList []string) ([]model_databases.Socket, map[string]model_databases.Socket)
	//GetSocketWithSensorer(status, language string, socketIdList []string) ([]model_services.MyFarmSenSocDetail, map[string]model_services.MyFarmSenSocDetail, int)
	//GetSocSenByKeyer(mainboxId, farmAreaId string, tranSoc []model_databases.TransSocketArea, socSenMap map[string]model_services.MyFarmSenSocDetail) ([]model_services.MyFarmSenSocDetail, int)
	//GetManageMainboxer(status, language, farmId string) ([]model_services.MyFarmManageMainbox, int)
	//GetFarmAreaByIder(status string, farmAreaIdList []string) ([]model_databases.FarmArea, map[string]model_databases.FarmArea)
	//GetManageFarmAreaer(status, language, farmId string) ([]model_services.MyFarmManageFarmArea, int)
	GetManageRole(status string, ReqModel *model_other.ReqModel) (*model_other.RespModel, error)
	GetManageFarmArea(status string, ReqModel *model_other.ReqModel) (*model_other.RespModel, error)
	GetManageMainbox(status string, ReqModel *model_other.ReqModel) (*model_other.RespModel, error)
}

type Service struct {
	repo  Repositorier
}

func NewService(repo Repositorier) Servicer {
	return &Service{
		repo:  repo,
	}
}

func (s *Service) GetAuthorizeCheckForManageFarm(uid, farmId string) (bool, error) {
	authManage := false
	trans, err := s.repo.FindOneTransManagement(uid, farmId)
	if err != nil{
		return false, err
	}
	if trans.RoleId.UUID.String() != config.GetRole().View {
		authManage = true
	}
	return authManage, err
}

func (s *Service) GetOverviewFarm(status string, ReqModel *model_other.ReqModel) (*model_other.RespModel, error) {
	farm, err := s.repo.FindOneFarm(status, ReqModel.FarmId)
	if err != nil{
		return nil, err
	}
	// Get Mainbox count
	farm.MainboxCount, err = s.repo.GetCountMainbox(status, ReqModel.FarmId)
	if err != nil{
		return nil, err
	}
	// Get Farm area count
	farm.FarmAreaCount, err = s.repo.GetCountFarmArea(status, ReqModel.FarmId)
	if err != nil{
		return nil, err
	}
	return &model_other.RespModel{
		Item: farm,
		Total: 1,
	}, nil
}

//func (s *Service) GetTransSocketAreaer(status, farmId string) ([]model_databases.TransSocketArea, []string, []string, []string, int) {
//	var joinArray []model_databases.TransSocketArea
//	var socketIdList []string
//	var farmAreaIdList []string
//	var mainboxIdList []string
//	var total int
//
//	// Get farmArea
//	sql := fmt.Sprintf("SELECT * FROM %s INNER JOIN %s ON %s.farm_area_id = %s.farm_area_id WHERE %s.status_id = '%s' AND %s.farm_id = '%s'",
//		config.DB_FARM_AREA, config.DB_TRANS_SOCKET_AREA, config.DB_FARM_AREA, config.DB_TRANS_SOCKET_AREA, config.DB_FARM_AREA, status, config.DB_FARM_AREA, farmId)
//	fmt.Println(sql)
//	err := ln.Db.Raw(sql).Scan(&joinArray).Error
//	if err != nil {
//		log.Print(err)
//	}
//
//	for _, join := range joinArray {
//		socketIdList = append(socketIdList, join.SocketId.UUID.String())
//		farmAreaIdList = append(farmAreaIdList, join.FarmAreaId.UUID.String())
//		mainboxIdList = append(mainboxIdList, join.MainboxId.UUID.String())
//	}
//
//	total = len(joinArray)
//
//	return joinArray, farmAreaIdList, socketIdList, mainboxIdList, total
//}
//
//func (s *Service) GetSocketByIder(status string, socketIdList []string) ([]model_databases.Socket, map[string]model_databases.Socket) {
//	var socketAr []model_databases.Socket
//	var socketMap map[string]model_databases.Socket
//
//	socketMap = make(map[string]model_databases.Socket)
//
//	sqlIn := utility.ConvertListToStringIn(socketIdList)
//	sql := fmt.Sprintf("SELECT * FROM %s WHERE status_id = '%s' AND socket_id IN %s",
//		config.DB_SOCKET, status, sqlIn)
//	fmt.Println(sql)
//	err := ln.Db.Raw(sql).Scan(&socketAr).Error
//	if err != nil {
//		log.Print(err)
//	}
//
//	for _, wa := range socketAr {
//		socketMap[wa.SocketId.UUID.String()] = wa
//	}
//	return socketAr, socketMap
//}
//
//func (s *Service) GetSocketWithSensorer(status, language string, socketIdList []string) ([]model_services.MyFarmSenSocDetail, map[string]model_services.MyFarmSenSocDetail, int) {
//	var joinArray []model_services.MyFarmSenSocDetail
//	var found bool
//	var total int
//
//	sensorTypeMap := make(map[string]string)
//	statusSensorMap := make(map[string]string)
//	socSenMap := make(map[string]model_services.MyFarmSenSocDetail)
//
//	sqlIn := utility.ConvertListToStringIn(socketIdList)
//	sql := fmt.Sprintf("SELECT * FROM %s INNER JOIN %s ON %s.sensor_id = %s.sensor_id WHERE %s.status_id = '%s' AND %s.socket_id IN %s",
//		config.DB_SOCKET, config.DB_SENSOR, config.DB_SOCKET, config.DB_SENSOR, config.DB_SOCKET, status, config.DB_SOCKET, sqlIn)
//	fmt.Println(sql)
//	err := ln.Db.Raw(sql).Scan(&joinArray).Error
//	if err != nil {
//		log.Print(err)
//	}
//
//	for idx, wa := range joinArray {
//		//Get Sensor Type name
//		wa.SensorTypeName, found = sensorTypeMap[wa.SensorTypeId.UUID.String()]
//		if !found {
//			_, wa.SensorTypeName = IntCommon.GetSensorTypeNameer(ln, wa.SensorTypeId.UUID.String(), language)
//			sensorTypeMap[wa.SensorTypeId.UUID.String()] = wa.SensorTypeName
//		}
//
//		//Get Status Sensor name
//		wa.StatusSensorName, found = statusSensorMap[wa.StatusSensorId.UUID.String()]
//		if !found {
//			_, wa.StatusSensorName = IntDashboard.GetStatusSensorer(ln, wa.StatusSensorId.UUID.String())
//			statusSensorMap[wa.StatusSensorId.UUID.String()] = wa.StatusSensorName
//		}
//		joinArray[idx] = wa
//		socSenMap[wa.SocketId.UUID.String()] = wa
//	}
//
//	total = len(joinArray)
//
//	return joinArray, socSenMap, total
//}
//
//func (s *Service) GetSocSenByKeyer(mainboxId, farmAreaId string, tranSoc []model_databases.TransSocketArea, socSenMap map[string]model_services.MyFarmSenSocDetail) ([]model_services.MyFarmSenSocDetail, int) {
//	var list []model_services.MyFarmSenSocDetail
//	var total int
//
//	for _, wa := range tranSoc {
//		if mainboxId != "" {
//			if wa.MainboxId.UUID.String() == mainboxId {
//				socSen, _ := socSenMap[wa.SocketId.UUID.String()]
//				list = append(list, socSen)
//			}
//		} else if farmAreaId != "" {
//			if wa.FarmAreaId.UUID.String() == farmAreaId {
//				socSen, _ := socSenMap[wa.SocketId.UUID.String()]
//				list = append(list, socSen)
//			}
//		}
//	}
//	total = len(list)
//
//	return list, total
//}
//
//func (s *Service) GetManageMainboxer(status, language, farmId string) ([]model_services.MyFarmManageMainbox, int) {
//	var manageMB model_services.MyFarmManageMainbox
//	var manageMBList []model_services.MyFarmManageMainbox
//	var total int
//
//	tranSoc, _, socketIdList, mainboxIdList, total := IntMyFarm.GetTransSocketAreaer(ln, config.GetStatus().Active, farmId)
//	if total == 0 {
//		return manageMBList, total
//	}
//
//	// Get Socket and Sensor
//	_, socSenMap, _ := IntMyFarm.GetSocketWithSensorer(ln, status, language, socketIdList)
//
//	// Get Mainbox
//	mainbox, _ := IntDashboard.GetMainboxByIder(ln, status, mainboxIdList)
//	for _, m := range mainbox {
//		mapstructure.Decode(m, &manageMB)
//		// Get Socket and Sensor by mainbox
//		manageMB.SenSocDetail, _ = IntMyFarm.GetSocSenByKeyer(ln, m.MainboxId.UUID.String(), "", tranSoc, socSenMap)
//		manageMBList = append(manageMBList, manageMB)
//	}
//
//	total = len(manageMBList)
//
//	return manageMBList, total
//}
//
//func (s *Service) GetFarmAreaByIder(status string, farmAreaIdList []string) ([]model_databases.FarmArea, map[string]model_databases.FarmArea) {
//	var farmAreaAr []model_databases.FarmArea
//	var farmAreaMap map[string]model_databases.FarmArea
//
//	farmAreaMap = make(map[string]model_databases.FarmArea)
//
//	sqlIn := utility.ConvertListToStringIn(farmAreaIdList)
//	sql := fmt.Sprintf("SELECT * FROM %s WHERE status_id = '%s' AND farm_area_id IN %s",
//		config.DB_FARM_AREA, status, sqlIn)
//	fmt.Println(sql)
//	err := ln.Db.Raw(sql).Scan(&farmAreaAr).Error
//	if err != nil {
//		log.Print(err)
//	}
//
//	for _, wa := range farmAreaAr {
//		farmAreaMap[wa.FarmAreaId.UUID.String()] = wa
//	}
//	return farmAreaAr, farmAreaMap
//}
//
//func (s *Service) GetManageFarmAreaer(status, language, farmId string) ([]model_services.MyFarmManageFarmArea, int) {
//	var manageFA model_services.MyFarmManageFarmArea
//	var manageFAList []model_services.MyFarmManageFarmArea
//	var total int
//
//	tranSoc, farmAreaIdList, socketIdList, _, total := IntMyFarm.GetTransSocketAreaer(ln, config.GetStatus().Active, farmId)
//	if total == 0 {
//		return manageFAList, total
//	}
//
//	// Get Socket and Sensor
//	_, socSenMap, _ := IntMyFarm.GetSocketWithSensorer(ln, status, language, socketIdList)
//
//	// Get FarmArea
//	farmArea, _ := IntMyFarm.GetFarmAreaByIder(ln, status, farmAreaIdList)
//	for _, fa := range farmArea {
//		mapstructure.Decode(fa, &manageFA)
//		// Get Socket and Sensor by farm area
//		manageFA.SenSocDetail, _ = IntMyFarm.GetSocSenByKeyer(ln, "", fa.FarmAreaId.UUID.String(), tranSoc, socSenMap)
//		manageFAList = append(manageFAList, manageFA)
//	}
//
//	total = len(manageFAList)
//
//	return manageFAList, total
//}
//
func (s *Service) GetManageRole(status string, ReqModel *model_other.ReqModel) (*model_other.RespModel, error) {
	// Check auth for edit
	//isAuth, err := Servicer.GetAuthorizeCheckForManageFarm(s, ReqModel.Uid, ReqModel.FarmId)
	//if err != nil{
	//	return nil, err
	//}
	//// No Auth
	//if isAuth != true {
	//	return nil, &errs.ErrContext{
	//		Code: ERROR_4002005,
	//		Err:  err,
	//		Msg:  MSG_NO_AUTH,
	//	}
	//}

	manageRole, err := s.repo.FindAllManageRole(status, ReqModel.FarmId)
	if err != nil{
		return nil, err
	}
	return &model_other.RespModel{
		Item: manageRole,
		Total: len(manageRole),
	}, nil
}

func (s *Service) GetManageFarmArea(status string, ReqModel *model_other.ReqModel) (*model_other.RespModel, error) {
	manageFarmArea, err := s.repo.FindAllManageFarmArea(status, ReqModel.FarmId)
	if err != nil{
		return nil, err
	}
	return &model_other.RespModel{
		Item: manageFarmArea,
		Total: len(manageFarmArea),
	}, nil
}

func (s *Service) GetManageMainbox(status string, ReqModel *model_other.ReqModel) (*model_other.RespModel, error) {
	manageFarmArea, err := s.repo.FindAllManageMainbox(status, ReqModel.FarmId)
	if err != nil{
		return nil, err
	}
	return &model_other.RespModel{
		Item: manageFarmArea,
		Total: len(manageFarmArea),
	}, nil
}
