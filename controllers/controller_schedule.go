package controllers

import (
	"LN-BackEND/models/model_services"
	"github.com/mitchellh/mapstructure"
)

/*-------------------------------------------------------------------------------------------*/
//                                 INTERFACE
/*-------------------------------------------------------------------------------------------*/
type IntSchedule interface {
	GetFarmAreaLister(status, farmId string) ([]model_services.ScheduleFarmArea, int)
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