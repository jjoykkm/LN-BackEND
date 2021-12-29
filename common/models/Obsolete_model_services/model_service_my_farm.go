package Obsolete_model_services

import (
	"github.com/jackc/pgtype"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table Farm Overview
//-------------------------------------------------------------------------------//
//Model
type MyFarmOverviewFarm struct {
	FarmId      	string	 `json:"farm_id"`
	FarmName    	string	 `json:"farm_name"`
	FarmDesc    	string	 `json:"farm_desc"`
	MainboxCount	int			 `json:"mainbox_count"`
	FarmAreaCount	int			 `json:"farm_area_count"`
}

// New instance
func (u *MyFarmOverviewFarm) New() *MyFarmOverviewFarm {
	return &MyFarmOverviewFarm{
		FarmId:			u.FarmId ,
		FarmName:		u.FarmName ,
		FarmDesc:		u.FarmDesc ,
		MainboxCount:	u.MainboxCount ,
		FarmAreaCount:	u.FarmAreaCount ,
	}
}
//-------------------------------------------------------------------------------//
//							Table Sensor Socket Detail
//-------------------------------------------------------------------------------//
//Model
type MyFarmSenSocDetail struct {
	SocketId       		string 		`json:"socket_id"`
	SocketName     		string    		`json:"socket_name"`
	SocketNumber		int64		 	`json:"socket_number"`
	StatusSensorId 		string 		`json:"status_sensor_id"`
	StatusSensorName 	string			`json:"status_sensor_name"`
	SensorId      		string	 	`json:"sensor_id"`
	SensorModel     	string	 	 	`json:"sensor_model"`
	SensorLots      	string	 	 	`json:"sensor_lots"`
	BitTransfer	    	pgtype.Bit	 	`json:"bit_transfer"`
	SensorTypeId		string	 	`json:"sensor_type_id"`
	SensorTypeName  	string	 	 	`json:"sensor_type_name"`
}

// New instance
func (u *MyFarmSenSocDetail) New() *MyFarmSenSocDetail {
	return &MyFarmSenSocDetail{
		SocketId:			u.SocketId ,
		SocketName:			u.SocketName ,
		SocketNumber:		u.SocketNumber ,
		StatusSensorId:		u.StatusSensorId ,
		StatusSensorName:	u.StatusSensorName ,
		SensorId:			u.SensorId ,
		SensorModel:		u.SensorModel ,
		SensorLots:			u.SensorLots ,
		BitTransfer:		u.BitTransfer ,
		SensorTypeId:		u.SensorTypeId ,
		SensorTypeName:		u.SensorTypeName ,
	}
}

//-------------------------------------------------------------------------------//
//							Table Manage Mainbox
//-------------------------------------------------------------------------------//
//Model
type MyFarmManageMainbox struct {
	MainboxId      		string	 			`json:"mainbox_id"`
	MainboxName     	string	 			`json:"mainbox_name"`
	MainboxModel    	string	 			`json:"mainbox_model"`
	MainboxLots     	string	 			`json:"mainbox_lots"`
	StartWarranty		time.Time	 			`json:"start_warranty"`
	EndWarranty			time.Time	 			`json:"end_warranty"`
	SenSocDetail		[]MyFarmSenSocDetail	`json:"sen_soc_detail"`
}

// New instance
func (u *MyFarmManageMainbox) New() *MyFarmManageMainbox {
	return &MyFarmManageMainbox{
		MainboxId:			u.MainboxId ,
		MainboxName:		u.MainboxName ,
		MainboxModel:		u.MainboxModel ,
		MainboxLots:		u.MainboxLots ,
		StartWarranty:		u.StartWarranty ,
		EndWarranty:		u.EndWarranty ,
		SenSocDetail:		u.SenSocDetail ,
	}
}

//-------------------------------------------------------------------------------//
//							Table Manage FarmArea
//-------------------------------------------------------------------------------//
//Model
type MyFarmManageFarmArea struct {
	FarmAreaId      	string	 			`json:"farm_area_id"`
	FarmAreaName    	string	 			`json:"farm_area_name"`
	SenSocDetail		[]MyFarmSenSocDetail	`json:"sen_soc_detail"`
}

// New instance
func (u *MyFarmManageFarmArea) New() *MyFarmManageFarmArea {
	return &MyFarmManageFarmArea{
		FarmAreaId:			u.FarmAreaId ,
		FarmAreaName:		u.FarmAreaName ,
		SenSocDetail:		u.SenSocDetail ,
	}
}


//-------------------------------------------------------------------------------//
//							Table Manage Role
//-------------------------------------------------------------------------------//
//Model
type MyFarmManageRole struct {
	Uid      		string	 `json:"uid"`
	Username     	string	 	 `json:"username"`
	FullName      	string	 	 `json:"full_name"`
	SurName      	string	 	 `json:"sur_name"`
	NickName      	string	 	 `json:"nick_name"`
	FarmId     		string	 `json:"farm_id"`
	RoleId      	string	 `json:"role_id"`
	RoleName      	string	 `json:"role_name"`
	RoleDesc      	string	 `json:"role_desc"`
	CreateDate		time.Time	 `json:"create_date"`
	ChangeDate	    time.Time	 `json:"change_date"`
}

// New instance
func (u *MyFarmManageRole) New() *MyFarmManageRole {
	return &MyFarmManageRole{
		Uid:			u.Uid ,
		Username:       u.Username ,
		FullName:       u.FullName ,
		SurName:       	u.SurName ,
		NickName:       u.NickName ,
		FarmId:			u.FarmId ,
		RoleId:			u.RoleId ,
		RoleName:		u.RoleName ,
		RoleDesc:		u.RoleDesc ,
		CreateDate:		u.CreateDate ,
		ChangeDate:		u.ChangeDate ,
	}
}