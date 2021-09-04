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
	Uid      	uuid.UUID	 `mapstructure:"uid" json:"uid"`
	FarmId      uuid.UUID	 `mapstructure:"farm_id" json:"farm_id"`
	FarmName    string		 `mapstructure:"farm_name" json:"farm_name"`
	FarmDesc    string		 `mapstructure:"farm_desc" json:"farm_desc"`
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
	FarmId				uuid.UUID	 		`mapstructure:"farm_id" json:"farm_id"`
	FarmAreaId			uuid.UUID	 		`mapstructure:"farm_area_id" json:"farm_area_id"`
	FarmAreaName		string	 	 		`mapstructure:"farm_area_name" json:"farm_area_name"`
	FormulaPlantId		string	 	 		`mapstructure:"formula_plant_id" json:"formula_plant_id"`
	FormulaName			string	 	 		`mapstructure:"formula_name" json:"formula_name"`
	FormulaDesc			string	 	 		`mapstructure:"formula_desc" json:"formula_desc"`
	SensorDetail		[]SenSocMainList	`mapstructure:"sensor_detail" json:"sensor_detail"`
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
	SocketId      	uuid.UUID	 `mapstructure:"socket_id" json:"socket_id"`
	MainboxId      	uuid.UUID	 `mapstructure:"mainbox_id" json:"mainbox_id"`
	SensorId      	uuid.UUID	 `mapstructure:"sensor_id" json:"sensor_id"`
	StatusId		uuid.UUID	 `mapstructure:"status_id" json:"status_id"`
	SocketName      string	 	 `mapstructure:"socket_name" json:"socket_name"`
	CreateDate		time.Time	 `mapstructure:"create_date" json:"create_date"`
	ChangeDate	    time.Time	 `mapstructure:"change_date" json:"change_date"`
	SocketNumber	int64		 `mapstructure:"socket_number" json:"socket_number"`
	StatusSensorId	uuid.UUID	 `mapstructure:"status_sensor_id" json:"status_sensor_id"`
	FarmAreaId      uuid.UUID	 `mapstructure:"farm_area_id" json:"farm_area_id"`
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
	SensorId      		uuid.UUID	 `mapstructure:"sensor_id" json:"sensor_id"`
	SensorModel     	string	 	 `mapstructure:"sensor_model" json:"sensor_model"`
	SensorLots      	string	 	 `mapstructure:"sensor_lots" json:"sensor_lots"`
	BitTransfer	    	pgtype.Bit	 `mapstructure:"bit_transfer" json:"bit_transfer"`
	SensorTypeId		uuid.UUID	 `mapstructure:"sensor_type_id" json:"sensor_type_id"`
	SensorTypeName  	string	 	 `mapstructure:"sensor_type_name" json:"sensor_type_name"`
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
	MainboxId      		uuid.UUID	 `mapstructure:"mainbox_id" json:"mainbox_id"`
	MainboxName     	string		 `mapstructure:"mainbox_name" json:"mainbox_name"`
	MainboxModel    	string		 `mapstructure:"mainbox_model" json:"mainbox_model"`
	MainboxLots     	string		 `mapstructure:"mainbox_lots" json:"mainbox_lots"`
	StartWarranty		time.Time	 `mapstructure:"start_warranty" json:"start_warranty"`
	EndWarranty			time.Time	 `mapstructure:"end_warranty" json:"end_warranty"`
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
	SocketId       		uuid.UUID 		`mapstructure:"socket_id" json:"socket_id"`
	SocketName     		string    		`mapstructure:"socket_name" json:"socket_name"`
	SocketNumber		int64		 	`mapstructure:"socket_number" json:"socket_number"`
	StatusSensorId 		uuid.UUID 		`mapstructure:"status_sensor_id" json:"status_sensor_id"`
	StatusSensorName 	string			`mapstructure:"status_sensor_name" json:"status_sensor_name"`
	Sensor         		SensorDetail	`mapstructure:"sensor_detail" json:"sensor_detail"`
	Mainbox        		MainboxDetail	`mapstructure:"mainbox_detail" json:"mainbox_detail"`
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
//	SocketId      		uuid.UUID	 `mapstructure:"socket_id" json:"socket_id"`
//	SocketName      	string	 	 `mapstructure:"socket_name" json:"socket_name"`
//	StatusSensorId		uuid.UUID	 `mapstructure:"status_sensor_id" json:"status_sensor_id"`
//	SensorId      		uuid.UUID	 `mapstructure:"sensor_id" json:"sensor_id"`
//	SensorModel     	string	 	 `mapstructure:"sensor_model" json:"sensor_model"`
//	SensorLots      	string	 	 `mapstructure:"sensor_lots" json:"sensor_lots"`
//	BitTransfer	    	pgtype.Bit	 `mapstructure:"bit_transfer" json:"bit_transfer"`
//	SensorTypeId		uuid.UUID	 `mapstructure:"sensor_type_id" json:"sensor_type_id"`
//	SensorTypeName  	string	 	 `mapstructure:"sensor_type_name" json:"sensor_type_name"`
//	MainboxId      		uuid.UUID	 `mapstructure:"mainbox_id" json:"mainbox_id"`
//	MainboxName     	string		 `mapstructure:"mainbox_name" json:"mainbox_name"`
//	MainboxModel    	string		 `mapstructure:"mainbox_model" json:"mainbox_model"`
//	MainboxLots     	string		 `mapstructure:"mainbox_lots" json:"mainbox_lots"`
//	StartWarranty		time.Time	 `mapstructure:"start_warranty" json:"start_warranty"`
//	EndWarranty			time.Time	 `mapstructure:"end_warranty" json:"end_warranty"`
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