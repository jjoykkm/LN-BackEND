package model_services

import (
	"github.com/jackc/pgtype"
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"time"

	//"time"
)

//-------------------------------------------------------------------------------//
//				 	    	FarmList
//-------------------------------------------------------------------------------//
//Model
type DashboardFarmList struct {
	Uid      	uuid.UUID	 `json:"uid"`
	FarmId      uuid.UUID	 `json:"farm_id"`
	FarmName    string		 `json:"farm_name"`
	FarmDesc    string		 `json:"farm_desc"`
}
// New instance
func (u *DashboardFarmList) New() *DashboardFarmList {
	return &DashboardFarmList{
		Uid:		u.Uid ,
		FarmId:		u.FarmId ,
		FarmName:	u.FarmName ,
		FarmDesc:	u.FarmDesc ,
	}
}

//-------------------------------------------------------------------------------//
//				 	    	FarmList
//-------------------------------------------------------------------------------//
//Model
type DashboardFarmAreaList struct {
	FarmId				uuid.UUID	 		`json:"farm_id"`
	FarmAreaId			uuid.UUID	 		`json:"farm_area_id"`
	FarmAreaName		string	 	 		`json:"farm_area_name"`
	FormulaPlantId		string	 	 		`json:"formula_plant_id"`
	FormulaName			string	 	 		`json:"formula_name"`
	FormulaDesc			string	 	 		`json:"formula_desc"`
	SensorDetail		[]SenSocMainList	`json:"sensor_detail"`
}
// New instance
func (u *DashboardFarmAreaList) New() *DashboardFarmAreaList {
	return &DashboardFarmAreaList{
		FarmId:				u.FarmId ,
		FarmAreaId:			u.FarmAreaId ,
		FarmAreaName:		u.FarmAreaName ,
		FormulaPlantId:		u.FormulaPlantId ,
		FormulaName:		u.FormulaName ,
		FormulaDesc:		u.FormulaDesc ,
		SensorDetail:		u.SensorDetail ,
	}
}

//-------------------------------------------------------------------------------//
//			Join Socket And TransSocketArea (View/Add/Edit)
//-------------------------------------------------------------------------------//
//Model
type JoinSocketAndTrans struct {
	SocketId      	uuid.UUID	 `json:"socket_id"`
	MainboxId      	uuid.UUID	 `json:"mainbox_id"`
	SensorId      	uuid.UUID	 `json:"sensor_id"`
	StatusId		uuid.UUID	 `json:"status_id"`
	SocketName      string	 	 `json:"socket_name"`
	CreateDate		time.Time	 `json:"create_date"`
	ChangeDate	    time.Time	 `json:"change_date"`
	SocketNumber	int64		 `json:"socket_number"`
	StatusSensorId	uuid.UUID	 `json:"status_sensor_id"`
	FarmAreaId      uuid.UUID	 `json:"farm_area_id"`
}

// New instance
func (u *JoinSocketAndTrans) New() *JoinSocketAndTrans {
	return &JoinSocketAndTrans{
		SocketId:			u.SocketId ,
		MainboxId:			u.MainboxId ,
		SensorId:			u.SensorId ,
		StatusId:			u.StatusId ,
		SocketName:			u.SocketName ,
		CreateDate:			u.CreateDate ,
		ChangeDate:			u.ChangeDate ,
		SocketNumber:		u.SocketNumber ,
		StatusSensorId:		u.StatusSensorId ,
		FarmAreaId:			u.FarmAreaId ,
	}
}

//-------------------------------------------------------------------------------//
//				 	    	Sensor Detail
//-------------------------------------------------------------------------------//
//Model
type SensorDetail struct {
	SensorId      		uuid.UUID	 `json:"sensor_id"`
	SensorModel     	string	 	 `json:"sensor_model"`
	SensorLots      	string	 	 `json:"sensor_lots"`
	BitTransfer	    	pgtype.Bit	 `json:"bit_transfer"`
	SensorTypeId		uuid.UUID	 `json:"sensor_type_id"`
	SensorTypeName  	string	 	 `json:"sensor_type_name"`
}
// New instance
func (u *SensorDetail) New() *SensorDetail {
	return &SensorDetail{
		SensorId:			u.SensorId ,
		SensorModel:		u.SensorModel ,
		SensorLots:			u.SensorLots ,
		BitTransfer:		u.BitTransfer ,
		SensorTypeId:		u.SensorTypeId ,
		SensorTypeName:		u.SensorTypeName ,
	}
}

//-------------------------------------------------------------------------------//
//				 	    	Mainbox Detail
//-------------------------------------------------------------------------------//
//Model
type MainboxDetail struct {
	MainboxId      		uuid.UUID	 `json:"mainbox_id"`
	MainboxName     	string		 `json:"mainbox_name"`
	MainboxModel    	string		 `json:"mainbox_model"`
	MainboxLots     	string		 `json:"mainbox_lots"`
	StartWarranty		time.Time	 `json:"start_warranty"`
	EndWarranty			time.Time	 `json:"end_warranty"`
}
// New instance
func (u *MainboxDetail) New() *MainboxDetail {
	return &MainboxDetail{
		MainboxId:			u.MainboxId ,
		MainboxName:		u.MainboxName ,
		MainboxModel:		u.MainboxModel ,
		MainboxLots:		u.MainboxLots ,
		StartWarranty:		u.StartWarranty ,
		EndWarranty:		u.EndWarranty ,
	}
}

//-------------------------------------------------------------------------------//
//				 	    	Sensor List
//-------------------------------------------------------------------------------//
//Model
type SenSocMainList struct {
	SocketId       		uuid.UUID 		`json:"socket_id"`
	SocketName     		string    		`json:"socket_name"`
	SocketNumber		int64		 	`json:"socket_number"`
	StatusSensorId 		uuid.UUID 		`json:"status_sensor_id"`
	StatusSensorName 	string			`json:"status_sensor_name"`
	Sensor         		SensorDetail	`json:"sensor_detail"`
	Mainbox        		MainboxDetail	`json:"mainbox_detail"`
}
// New instance
func (u *SenSocMainList) New() *SenSocMainList {
	return &SenSocMainList{
		SocketId:			u.SocketId ,
		SocketName:			u.SocketName ,
		SocketNumber:		u.SocketNumber ,
		StatusSensorId:		u.StatusSensorId ,
		StatusSensorName:	u.StatusSensorName ,
		Sensor:				u.Sensor ,
		Mainbox:			u.Mainbox ,
	}
}
//
////-------------------------------------------------------------------------------//
////				 	    	Sensor List
////-------------------------------------------------------------------------------//
////Model
//type SenSocMainList struct {
//	SocketId      		uuid.UUID	 `json:"socket_id"`
//	SocketName      	string	 	 `json:"socket_name"`
//	StatusSensorId		uuid.UUID	 `json:"status_sensor_id"`
//	SensorId      		uuid.UUID	 `json:"sensor_id"`
//	SensorModel     	string	 	 `json:"sensor_model"`
//	SensorLots      	string	 	 `json:"sensor_lots"`
//	BitTransfer	    	pgtype.Bit	 `json:"bit_transfer"`
//	SensorTypeId		uuid.UUID	 `json:"sensor_type_id"`
//	SensorTypeName  	string	 	 `json:"sensor_type_name"`
//	MainboxId      		uuid.UUID	 `json:"mainbox_id"`
//	MainboxName     	string		 `json:"mainbox_name"`
//	MainboxModel    	string		 `json:"mainbox_model"`
//	MainboxLots     	string		 `json:"mainbox_lots"`
//	StartWarranty		time.Time	 `json:"start_warranty"`
//	EndWarranty			time.Time	 `json:"end_warranty"`
//}
//// New instance
//func (u *SenSocMainList) New() *SenSocMainList {
//	return &SenSocMainList{
//		SocketId:			u.SocketId ,
//		SocketName:			u.SocketName ,
//		StatusSensorId:		u.StatusSensorId ,
//		SensorId:			u.SensorId ,
//		SensorModel:		u.SensorModel ,
//		SensorLots:			u.SensorLots ,
//		BitTransfer:		u.BitTransfer ,
//		SensorTypeId:		u.SensorTypeId ,
//		SensorTypeName:		u.SensorTypeName ,
//		MainboxId:			u.MainboxId ,
//		MainboxName:		u.MainboxName ,
//		MainboxModel:		u.MainboxModel ,
//		MainboxLots:		u.MainboxLots ,
//		StartWarranty:		u.StartWarranty ,
//		EndWarranty:		u.EndWarranty ,
//	}
//}