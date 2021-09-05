package model_services

import (
	"github.com/jackc/pgtype"
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table Farm Overview
//-------------------------------------------------------------------------------//
//Model
type MyFarmOverviewFarm struct {
	FarmId      	uuid.UUID	 `mapstructure:"farm_id" json:"farm_id"`
	FarmName    	string		 `mapstructure:"farm_name" json:"farm_name"`
	FarmDesc    	string		 `mapstructure:"farm_desc" json:"farm_desc"`
	MainboxCount	int			 `mapstructure:"mainbox_count" json:"mainbox_count"`
	FarmAreaCount	int			 `mapstructure:"farm_area_count" json:"farm_area_count"`
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
	SocketId       		uuid.UUID 		`mapstructure:"socket_id" json:"socket_id"`
	SocketName     		string    		`mapstructure:"socket_name" json:"socket_name"`
	SocketNumber		int64		 	`mapstructure:"socket_number" json:"socket_number"`
	StatusSensorId 		uuid.UUID 		`mapstructure:"status_sensor_id" json:"status_sensor_id"`
	StatusSensorName 	string			`mapstructure:"status_sensor_name" json:"status_sensor_name"`
	SensorId      		uuid.UUID	 	`mapstructure:"sensor_id" json:"sensor_id"`
	SensorModel     	string	 	 	`mapstructure:"sensor_model" json:"sensor_model"`
	SensorLots      	string	 	 	`mapstructure:"sensor_lots" json:"sensor_lots"`
	BitTransfer	    	pgtype.Bit	 	`mapstructure:"bit_transfer" json:"bit_transfer"`
	SensorTypeId		uuid.UUID	 	`mapstructure:"sensor_type_id" json:"sensor_type_id"`
	SensorTypeName  	string	 	 	`mapstructure:"sensor_type_name" json:"sensor_type_name"`
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
	MainboxId      		uuid.UUID	 			`mapstructure:"mainbox_id" json:"mainbox_id"`
	MainboxName     	string		 			`mapstructure:"mainbox_name" json:"mainbox_name"`
	MainboxModel    	string		 			`mapstructure:"mainbox_model" json:"mainbox_model"`
	MainboxLots     	string		 			`mapstructure:"mainbox_lots" json:"mainbox_lots"`
	StartWarranty		time.Time	 			`mapstructure:"start_warranty" json:"start_warranty"`
	EndWarranty			time.Time	 			`mapstructure:"end_warranty" json:"end_warranty"`
	SenSocDetail		[]MyFarmSenSocDetail	`mapstructure:"sen_soc_detail" json:"sen_soc_detail"`
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
	FarmAreaId      	uuid.UUID	 			`mapstructure:"farm_area_id" json:"farm_area_id"`
	FarmAreaName    	string		 			`mapstructure:"farm_area_name" json:"farm_area_name"`
	SenSocDetail		[]MyFarmSenSocDetail	`mapstructure:"sen_soc_detail" json:"sen_soc_detail"`
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
	Uid      		uuid.UUID	 `mapstructure:"uid" json:"uid"`
	Username     	string	 	 `mapstructure:"username" json:"username"`
	FullName      	string	 	 `mapstructure:"full_name" json:"full_name"`
	SurName      	string	 	 `mapstructure:"sur_name" json:"sur_name"`
	NickName      	string	 	 `mapstructure:"nick_name" json:"nick_name"`
	FarmId     		uuid.UUID	 `mapstructure:"farm_id" json:"farm_id"`
	RoleId      	uuid.UUID	 `mapstructure:"role_id" json:"role_id"`
	RoleName      	string		 `mapstructure:"role_name" json:"role_name"`
	RoleDesc      	string		 `mapstructure:"role_desc" json:"role_desc"`
	CreateDate		time.Time	 `mapstructure:"create_date" json:"create_date"`
	ChangeDate	    time.Time	 `mapstructure:"change_date" json:"change_date"`
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