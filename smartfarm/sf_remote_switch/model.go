package sf_remote_switch

import (
	"github.com/jjoykkm/ln-backend/common/config"
	"github.com/jjoykkm/ln-backend/common/models/model_db"
)

//-------------------------------------------------------------------------------//
//									Get
//-------------------------------------------------------------------------------//
type SensorDetail struct {
	Sensor        model_db.Sensor	 	 `json:"sensor" gorm:"embedded"`
	SensorType    model_db.SensorType	 `json:"sensor_type" gorm:"foreignkey:SensorTypeId; references:SensorTypeId"`
}
func (SensorDetail) TableName() string {
	return config.DB_SENSOR
}

//-------------------------------------------------------------------------------//
//				 	    	Socket Sensor Detail
//-------------------------------------------------------------------------------//
//Model
type SocSenDetail struct {
	Socket        	model_db.Socket	 	 	`json:"socket" gorm:"embedded"`
	StatusSensor    model_db.StatusSensor	`json:"status_sensor" gorm:"foreignkey:StatusSensorId; references:StatusSensorId"`
	Sensor         	SensorDetail			`json:"sensor" gorm:"foreignkey:SensorId; references:SensorId"`
}
func (SocSenDetail) TableName() string {
	return config.DB_SOCKET
}

type RemoteSwitch struct {
	RemoteSwitch	model_db.RemoteSwitch 	`json:"remote_switch" gorm:"embedded"`
	SocSenDetail	[]SocSenDetail			`json:"socket_sensor_detail" gorm:"foreignkey:RemoteId; references:RemoteId"`
}
func (RemoteSwitch) TableName() string {
	return config.DB_REMOTE_SWITCH
}

//-------------------------------------------------------------------------------//
//									Upsert
//-------------------------------------------------------------------------------//
type RemoteDetailUS struct {
	RemoteSwitch	model_db.RemoteSwitchUS 	`json:"remote_switch" gorm:"embedded"`
	SocketList		[]string					`json:"socket_list"`
}
func (RemoteDetailUS) TableName() string {
	return config.DB_REMOTE_SWITCH
}

//-------------------------------------------------------------------------------//
//									Delete
//-------------------------------------------------------------------------------//
type RemoteDetailDel struct {
	RemoteId       string	 	 `json:"remote_id"`
	SocketList	   []string		 `json:"socket_list"`
}
func (RemoteDetailDel) TableName() string {
	return config.DB_REMOTE_SWITCH
}



