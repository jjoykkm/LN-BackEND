package sf_dashboard

import (
	"github.com/jjoykkm/ln-backend/common/models/model_other"
)

type Servicer interface {
	// status, uid string
	GetFarmList(status string, ReqModel *model_other.ReqModel) (*model_other.RespModel, error)
	//GetFarmAreaDashboardLister(status, language, farmId string) ([]model_services.DashboardFarmAreaList, int)
	//GetSocketLister(status, farmId string) ([]model_services.JoinSocketAndTrans, []string, []string)
	//GetSensorByIder(status string, socketIdList []string) ([]model_db.Sensor, map[string]model_db.Sensor)
	//GetMainboxByIder(status string, mainboxIdList []string) ([]model_db.Mainbox, map[string]model_db.Mainbox)
	//GetFarmAreaDetailSensorer(status, farmId, language string) ([]model_services.SenSocMainList, int)
	//GetStatusSensorer(sensorStatusId string) (model_db.StatusSensor, string)
}

type Service struct {
	repo  Repositorier
}

func NewService(repo Repositorier) Servicer {
	return &Service{
		repo:  repo,
	}
}

func (s *Service) GetFarmList(status string, ReqModel *model_other.ReqModel) (*model_other.RespModel, error) {
	myFarm, err := s.repo.FindAllMyFarm(status, ReqModel.Uid)
	if err != nil{
		return nil, err
	}
	return &model_other.RespModel{
		Item: myFarm,
		Total: len(myFarm),
	}, nil
}

//func (s *Service) GetFarmAreaDashboardLister(status, language, farmId string) ([]model_services.DashboardFarmAreaList, int) {
//	var farmAreaList []model_services.DashboardFarmAreaList
//	var total int
//
//	sql := fmt.Sprintf("SELECT * FROM %s INNER JOIN %s ON %s.formula_plant_id = %s.formula_plant_id WHERE %s.status_id = '%s' AND %s.farm_id = '%s'",
//		config.DB_FARM_AREA, config.DB_FORMULA_PLANT, config.DB_FARM_AREA, config.DB_FORMULA_PLANT, config.DB_FARM_AREA, status, config.DB_FARM_AREA, farmId)
//	fmt.Println(sql)
//	err := ln.Db.Raw(sql).Scan(&farmAreaList).Error
//	if err != nil {
//		log.Print(err)
//	}
//
//	for idx, wa := range farmAreaList {
//		wa.SensorDetail, _ = IntDashboard.GetFarmAreaDetailSensorer(ln, config.GetStatus().Active, wa.FarmAreaId.UUID.String(), language)
//		farmAreaList[idx] = wa
//	}
//
//	total = len(farmAreaList)
//	return farmAreaList, total
//}
//
//func (s *Service) GetSocketLister(status, farmAreaId string) ([]model_services.JoinSocketAndTrans, []string, []string) {
//	var joinArray []model_services.JoinSocketAndTrans
//	var sensorStr string
//	var sensorIdList []string
//	var mainboxStr string
//	var mainboxIdList []string
//
//	sql := fmt.Sprintf("SELECT * FROM %s INNER JOIN %s ON %s.socket_id = %s.socket_id WHERE %s.status_id = '%s' AND %s.farm_area_id = '%s'",
//		config.DB_TRANS_SOCKET_AREA, config.DB_SOCKET, config.DB_TRANS_SOCKET_AREA, config.DB_SOCKET, config.DB_TRANS_SOCKET_AREA, status, config.DB_TRANS_SOCKET_AREA, farmAreaId)
//	fmt.Println(sql)
//	err := ln.Db.Raw(sql).Scan(&joinArray).Error
//	if err != nil {
//		log.Print(err)
//	}
//
//	for _, join := range joinArray {
//		sensorStr = join.SensorId.UUID.String()
//		mainboxStr = join.MainboxId.UUID.String()
//		sensorIdList = append(sensorIdList, sensorStr)
//		mainboxIdList = append(mainboxIdList, mainboxStr)
//	}
//
//	return joinArray, sensorIdList, mainboxIdList
//}
//
//func (s *Service) GetSensorByIder(status string, sensorIdList []string) ([]model_db.Sensor, map[string]model_db.Sensor) {
//	var sensorAr []model_db.Sensor
//	var sensorMap map[string]model_db.Sensor
//
//	sensorMap = make(map[string]model_db.Sensor)
//
//	sqlIn := utility.ConvertListToStringIn(sensorIdList)
//	sql := fmt.Sprintf("SELECT * FROM %s WHERE status_id = '%s' AND sensor_id IN %s",
//		config.DB_SENSOR, status, sqlIn)
//	fmt.Println(sql)
//	err := ln.Db.Raw(sql).Scan(&sensorAr).Error
//	if err != nil {
//		log.Print(err)
//	}
//
//	for _, wa := range sensorAr {
//		sensorMap[wa.SensorId.UUID.String()] = wa
//	}
//	return sensorAr, sensorMap
//}
//
//func (s *Service) GetMainboxByIder(status string, mainboxIdList []string) ([]model_db.Mainbox, map[string]model_db.Mainbox) {
//	var mainboxAr []model_db.Mainbox
//	var mainboxMap map[string]model_db.Mainbox
//
//	mainboxMap = make(map[string]model_db.Mainbox)
//
//	sqlIn := utility.ConvertListToStringIn(mainboxIdList)
//	sql := fmt.Sprintf("SELECT * FROM %s WHERE status_id = '%s' AND mainbox_id IN %s",
//		config.DB_MAINBOX, status, sqlIn)
//	fmt.Println(sql)
//	err := ln.Db.Raw(sql).Scan(&mainboxAr).Error
//	if err != nil {
//		log.Print(err)
//	}
//
//	for _, wa := range mainboxAr {
//		mainboxMap[wa.MainboxId.UUID.String()] = wa
//	}
//	return mainboxAr, mainboxMap
//}
//
//func (s *Service) GetStatusSensorer(sensorStatusId string) (model_db.StatusSensor, string) {
//	var model model_db.StatusSensor
//	var status string
//
//	sql := fmt.Sprintf("SELECT * FROM %s WHERE status_id = '%s' AND status_sensor_id = '%s'",
//		config.DB_STATUS_SENSOR, config.GetStatus().Active, sensorStatusId)
//	fmt.Println(sql)
//	err := ln.Db.Raw(sql).Scan(&model).Error
//	if err != nil {
//		log.Print(err)
//	}
//	status = model.StatusName
//	return model, status
//}
//
//func (s *Service) GetFarmAreaDetailSensorer(status, farmAreaId, language string) ([]model_services.SenSocMainList, int) {
//	var senSocMain model_services.SenSocMainList
//	var senSocMainList []model_services.SenSocMainList
//	var found bool
//	var total int
//
//	sensorTypeMap := make(map[string]string)
//	statusSensorMap := make(map[string]string)
//
//	socAreaAr, sensorIdList, mainboxIdList := IntDashboard.GetSocketLister(ln, status, farmAreaId)
//
//	_, sensorMap := IntDashboard.GetSensorByIder(ln, config.GetStatus().Active, sensorIdList)
//	_, mainboxMap := IntDashboard.GetMainboxByIder(ln, config.GetStatus().Active, mainboxIdList)
//
//	for _, wa := range socAreaAr {
//		mapstructure.Decode(wa, &senSocMain)
//		senSocMain.Sensor.SensorId = wa.SensorId
//		senSocMain.Mainbox.MainboxId = wa.MainboxId
//
//		//Get Status Sensor name
//		senSocMain.StatusSensorName, found = statusSensorMap[senSocMain.StatusSensorId.UUID.String()]
//		if !found {
//			_, senSocMain.StatusSensorName = IntDashboard.GetStatusSensorer(ln, senSocMain.StatusSensorId.UUID.String())
//			statusSensorMap[senSocMain.StatusSensorId.UUID.String()] = senSocMain.StatusSensorName
//		}
//
//		//Get Mainbox
//		mb, fmb := mainboxMap[senSocMain.Mainbox.MainboxId.UUID.String()]
//		if fmb {
//			mapstructure.Decode(mb, &senSocMain.Mainbox)
//		}
//		//Get Sensor
//		ss, fss := sensorMap[senSocMain.Sensor.SensorId.UUID.String()]
//		if fss {
//			mapstructure.Decode(ss, &senSocMain.Sensor)
//		}
//		//Get Sensor Type name
//		senSocMain.Sensor.SensorTypeName, found = sensorTypeMap[senSocMain.Sensor.SensorTypeId.UUID.String()]
//		if !found {
//			_, senSocMain.Sensor.SensorTypeName = IntCommon.GetSensorTypeNameer(ln, senSocMain.Sensor.SensorTypeId.UUID.String(), language)
//			sensorTypeMap[senSocMain.Sensor.SensorTypeId.UUID.String()] = senSocMain.Sensor.SensorTypeName
//		}
//
//		senSocMainList = append(senSocMainList, senSocMain)
//	}
//	total = len(senSocMainList)
//
//	return senSocMainList, total
//}
