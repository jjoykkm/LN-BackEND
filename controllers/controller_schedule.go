package controllers

import (
	"fmt"
	"github.com/jjoykkm/ln-backend/config"
	"github.com/jjoykkm/ln-backend/models/model_databases"
	"github.com/jjoykkm/ln-backend/models/model_services"
	"github.com/jjoykkm/ln-backend/utility"
	"github.com/mitchellh/mapstructure"
	"log"
)

/*-------------------------------------------------------------------------------------------*/
//                                 INTERFACE
/*-------------------------------------------------------------------------------------------*/
type IntSchedule interface {
	GetFarmAreaLister(status, farmId string) ([]model_services.ScheduleFarmArea, int)
	GetScheduleIder(status string, farmAreaIdList []string) ([]model_databases.TransScheduleFarm, []string, int)
	GetScheduleer(status string, scheIdList []string) ([]model_databases.Schedule, []string, []string, int)
	GetFrequencyTypeer(status string, freqTypeIdList []string) ([]model_databases.FrequencyType, map[string]model_databases.FrequencyType, int)
	GetIndicateTypeer(status string, inTypeIdList []string) ([]model_databases.IndicateType, map[string]model_databases.IndicateType, int)
	GetScheReminder(status string, farmAreaId []string) model_services.ScheduleScheRemind
}

func (ln Ln) GetFarmAreaLister(status, farmId string) ([]model_services.ScheduleFarmArea, int) {
	var farmAreaList []model_services.ScheduleFarmArea
	var farmArea model_services.ScheduleFarmArea

	list, _, total := IntCommon.GetFarmAreaByFarmId(ln, status, farmId)

	for _, wa := range list {
		mapstructure.Decode(wa, &farmArea)
		farmAreaList = append(farmAreaList, farmArea)
	}
	return farmAreaList, total
}

func (ln Ln) GetScheduleIder(status string, farmAreaIdList []string) ([]model_databases.TransScheduleFarm, []string, int) {
	var scheduleFarm []model_databases.TransScheduleFarm
	var scheIdList []string

	sqlIn := utility.ConvertListToStringIn(farmAreaIdList)
	sql := fmt.Sprintf("SELECT * FROM %s WHERE status_id = '%s' AND farm_area_id IN %s",
		config.DB_TRANS_SCHEDULE_FARM, status, sqlIn)
	fmt.Println(sql)
	err := ln.Db.Raw(sql).Scan(&scheduleFarm).Error
	if err != nil {
		log.Print(err)
	}
	for _, wa := range scheduleFarm {
		scheIdList = append(scheIdList, wa.ScheduleId.UUID.String())
	}
	total := len(scheduleFarm)

	return scheduleFarm, scheIdList, total
}

func (ln Ln) GetScheduleer(status string, scheIdList []string) ([]model_databases.Schedule, []string, []string, int) {
	var schedule []model_databases.Schedule
	var structSR []model_services.ScheduleStruct
	var freqTypeIdList []string
	var inTypeIdList []string

	sqlIn := utility.ConvertListToStringIn(scheIdList)
	sql := fmt.Sprintf("SELECT * FROM %s INNER JOIN %s ON %s.schedule_id = %s.schedule_id WHERE %s.status_id = '%s' AND %s.schedule_id IN %s",
		config.DB_SCHEDULE, config.DB_TRANS_SCHEDULE_FARM, config.DB_SCHEDULE, config.DB_TRANS_SCHEDULE_FARM, config.DB_SCHEDULE, status, config.DB_TRANS_SCHEDULE_FARM, sqlIn)
	fmt.Println(sql)
	//sql := fmt.Sprintf("SELECT * FROM %s WHERE status_id = '%s' AND schedule_id IN %s",
	//	config.DB_SCHEDULE, status, sqlIn)
	fmt.Println(sql)
	err := ln.Db.Raw(sql).Scan(&structSR).Error
	if err != nil {
		log.Print(err)
	}
	fmt.Printf("%+v\n", structSR)
	for _, wa := range structSR {
		freqTypeIdList = append(freqTypeIdList, wa.FreqType.FrequencyTypeId.UUID.String())
		inTypeIdList = append(inTypeIdList, wa.IndicateType.IndicateTypeId.UUID.String())
	}
	total := len(schedule)

	return schedule, freqTypeIdList, inTypeIdList, total
}

func (ln Ln) GetFrequencyTypeer(status string, freqTypeIdList []string) ([]model_databases.FrequencyType, map[string]model_databases.FrequencyType, int) {
	var freqType []model_databases.FrequencyType

	mapFreqType := make(map[string]model_databases.FrequencyType)

	sqlIn := utility.ConvertListToStringIn(freqTypeIdList)
	sql := fmt.Sprintf("SELECT * FROM %s WHERE status_id = '%s' AND frequency_type_id IN %s",
		config.DB_FREQUENCY_TYPE, status, sqlIn)
	fmt.Println(sql)
	err := ln.Db.Raw(sql).Scan(&freqType).Error
	if err != nil {
		log.Print(err)
	}
	for _, wa := range freqType {
		mapFreqType[wa.FreqTypeId.UUID.String()] = wa
	}
	total := len(freqType)

	return freqType, mapFreqType, total
}

func (ln Ln) GetIndicateTypeer(status string, inTypeIdList []string) ([]model_databases.IndicateType, map[string]model_databases.IndicateType, int) {
	var inType []model_databases.IndicateType

	inTypeMap := make(map[string]model_databases.IndicateType)

	sqlIn := utility.ConvertListToStringIn(inTypeIdList)
	sql := fmt.Sprintf("SELECT * FROM %s WHERE status_id = '%s' AND indicate_type_id IN %s",
		config.DB_INDICATE_TYPE, status, sqlIn)
	fmt.Println(sql)
	err := ln.Db.Raw(sql).Scan(&inType).Error
	if err != nil {
		log.Print(err)
	}
	for _, wa := range inType {
		inTypeMap[wa.IndicateTypeId.UUID.String()] = wa
	}
	total := len(inType)

	return inType, inTypeMap, total
}

func (ln Ln) GetScheReminder(status string, farmAreaId []string) model_services.ScheduleScheRemind {
	var sRList []model_services.ScheduleStruct
	var structSR model_services.ScheduleStruct
	var scheRemind model_services.ScheduleScheRemind
	var freqTypeIdList []string
	var inTypeIdList []string
	var found bool

	farmAreaMap := make(map[string]string)

	//_, scheIdList, _ := IntSchedule.GetScheduleIder(ln, status, farmAreaId)

	sqlIn := utility.ConvertListToStringIn(farmAreaId)
	sql := fmt.Sprintf("SELECT * FROM %s INNER JOIN %s ON %s.schedule_id = %s.schedule_id WHERE %s.status_id = '%s' AND %s.farm_area_id IN %s",
		config.DB_SCHEDULE, config.DB_TRANS_SCHEDULE_FARM, config.DB_SCHEDULE, config.DB_TRANS_SCHEDULE_FARM, config.DB_SCHEDULE, status, config.DB_TRANS_SCHEDULE_FARM, sqlIn)
	fmt.Println(sql)
	err := ln.Db.Raw(sql).Scan(&sRList).Error
	if err != nil {
		log.Print(err)
	}

	//fmt.Printf("%+v\n",sRList)
	for _, wa := range sRList {
		freqTypeIdList = append(freqTypeIdList, wa.FreqType.FrequencyTypeId.UUID.String())
		inTypeIdList = append(inTypeIdList, wa.IndicateType.IndicateTypeId.UUID.String())
	}
	fmt.Println(inTypeIdList)
	//
	////scheList, freqTypeIdList, inTypeIdList, _ := IntSchedule.GetScheduleer(ln, config.STATUS_ACTIVE, scheIdList)
	//
	// Get Frequency Type
	_, freqTypeMap, _ := IntSchedule.GetFrequencyTypeer(ln, config.STATUS_ACTIVE, freqTypeIdList)

	// Get Indicate Type
	_, inTypeMap, _ := IntSchedule.GetIndicateTypeer(ln, config.STATUS_ACTIVE, inTypeIdList)

	for _, wa := range sRList {
		mapstructure.Decode(wa, &structSR)
		//Get Country name
		structSR.FarmAreaName, found = farmAreaMap[structSR.FarmAreaId.UUID.String()]
		if !found {
			_, structSR.FarmAreaName = IntCommon.GetFarmAreaNameer(ln, structSR.FarmAreaId.UUID.String())
			farmAreaMap[structSR.FarmAreaId.UUID.String()] = structSR.FarmAreaName
		}
		//Get Frequency Type
		freqType, f := freqTypeMap[wa.FreqType.FrequencyTypeId.UUID.String()]
		if f {
			mapstructure.Decode(freqType, &structSR.FreqType)
		}
		//Get Indicate Type
		inType, f := inTypeMap[wa.IndicateType.IndicateTypeId.UUID.String()]
		if f {
			mapstructure.Decode(inType, &structSR.IndicateType)
		}
		if wa.IsReminder {
			scheRemind.ScheduleList = append(scheRemind.ScheduleList, structSR)
		} else {
			scheRemind.ReminderList = append(scheRemind.ReminderList, structSR)
		}
	}
	//fmt.Printf("%+v\n",scheRemind)
	return scheRemind
}
