package sf_my_farm

import (
	"github.com/jjoykkm/ln-backend/common/config"
	"github.com/jjoykkm/ln-backend/common/models/model_db"
)

//-------------------------------------------------------------------------------//
//							Table Farm Overview
//-------------------------------------------------------------------------------//
//Model
type FarmOverview struct {
	Farm			model_db.Farm	`json:"farm" gorm:"embedded"`
	MainboxCount	int64			`json:"mainbox_count"`
	FarmAreaCount	int64		 	`json:"farm_area_count"`
}
// New instance
func (u *FarmOverview) New() *FarmOverview {
	return &FarmOverview{
		Farm:			u.Farm ,
		MainboxCount:	u.MainboxCount ,
		FarmAreaCount:	u.FarmAreaCount ,
	}
}
func (FarmOverview) TableName() string {
	return config.DB_FARM
}
//-------------------------------------------------------------------------------//
//				 	    	Socket Detail
//-------------------------------------------------------------------------------//
//Model
type SocketDetail struct {
	Socket        	model_db.Socket	 	 	`json:"socket" gorm:"embedded"`
	StatusSensor    model_db.StatusSensor	`json:"status_sensor" gorm:"foreignkey:StatusSensorId; references:StatusSensorId"`
	Sensor         	[]SensorDetail			`json:"sensor" gorm:"foreignkey:SensorId; references:SensorId"`
}
func (SocketDetail) TableName() string {
	return config.DB_SOCKET
}

//-------------------------------------------------------------------------------//
//				 	    	Sensor Detail
//-------------------------------------------------------------------------------//
//Model
type SensorDetail struct {
	Sensor        model_db.Sensor	 	 `json:"sensor" gorm:"embedded"`
	SensorType    model_db.SensorType	 `json:"sensor_type" gorm:"foreignkey:SensorTypeId; references:SensorTypeId"`
}
func (SensorDetail) TableName() string {
	return config.DB_SENSOR
}

//-------------------------------------------------------------------------------//
//				 	    	Sensor List
//-------------------------------------------------------------------------------//
//Model
type SenSocMainList struct {
	SocketArea		model_db.TransSocketArea	`json:"socket_area" gorm:"embedded"`
	Socket         	[]SocketDetail				`json:"socket_detail" gorm:"foreignkey:SocketId; references:SocketId"`
}
func (SenSocMainList) TableName() string {
	return config.DB_TRANS_SOCKET_AREA
}

//-------------------------------------------------------------------------------//
//				 	    Manage FarmArea
//-------------------------------------------------------------------------------//
//Model
type ManageFarmArea struct {
	FarmArea    	model_db.FarmArea		`json:"farm_area" gorm:"embedded"`
	SensorDetail	[]SenSocMainList		`json:"sensor_detail" gorm:"foreignkey:FarmAreaId; references:FarmAreaId"`
}
func (ManageFarmArea) TableName() string {
	return config.DB_FARM_AREA
}

//-------------------------------------------------------------------------------//
//							Table Manage Mainbox
//-------------------------------------------------------------------------------//
//Model
type ManageMainbox struct {
	Mainbox			model_db.Mainbox	`json:"mainbox" gorm:"embedded"`
	SenSocDetail	[]SenSocMainList	`json:"socket_detail" gorm:"foreignkey:MainboxId; references:MainboxId"`
}
// New instance
func (u *ManageMainbox) New() *ManageMainbox {
	return &ManageMainbox{
		Mainbox:		u.Mainbox ,
		SenSocDetail:	u.SenSocDetail ,
	}
}
func (ManageMainbox) TableName() string {
	return config.DB_MAINBOX
}

//-------------------------------------------------------------------------------//
//							Table Manage Role
//-------------------------------------------------------------------------------//
//Model
type ManageRole struct {
	Management	model_db.TransManagement	`json:"farm_area_id" gorm:"embedded"`
	User		model_db.UsersShort			`json:"user_detail" gorm:"foreignkey:Uid; references:Uid"`
	Role		model_db.Role				`json:"role_detail" gorm:"foreignkey:RoleId; references:RoleId"`
}

// New instance
func (u *ManageRole) New() *ManageRole {
	return &ManageRole{
		Management:		u.Management ,
		User:       	u.User ,
		Role:       	u.Role ,
	}
}
func (ManageRole) TableName() string {
	return config.DB_TRANS_MANAGEMENT
}

//-------------------------------------------------------------------------------//
//							Request Config Mainbox
//-------------------------------------------------------------------------------//
//Model
type ReqConfMainbox struct {
	Mainbox    	*model_db.MainboxUS	 `json:"mainbox"`
	Socket		[]model_db.SocketUS	 `json:"socket"`
	Sensor     	*model_db.SensorUS	 `json:"sensor"`
}

type ReqDeleteConfig struct {
	MainboxId      	string	 `json:"mainbox_id"`
	SensorId      	string	 `json:"sensor_id"`
	SocketId      	string	 `json:"socket_id"`
	FarmId     		string	 `json:"farm_id"`
	FarmAreaId     	string	 `json:"farm_area_id"`
}

type ReqConfFarmArea struct {
	FarmArea    		*model_db.FarmAreaUS			`json:"farm_area"`
	TransSocketArea		[]model_db.TransSocketAreaUS	`json:"trans_socket_area"`
}


type ReqConfFarm struct {
	Farm   	*model_db.FarmUS	`json:"farm"`
}