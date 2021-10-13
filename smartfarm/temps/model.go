package sf_dashboard

import (
	"github.com/jjoykkm/ln-backend/common/config"
	"github.com/jjoykkm/ln-backend/common/models/model_db"
)
//-------------------------------------------------------------------------------//
//				 	    	Socket Detail
//-------------------------------------------------------------------------------//
//Model
type SocketDetail struct {
	Socket        	model_db.Socket	 	 	`json:"socket" gorm:"embedded"`
	StatusSensor    model_db.StatusSensor	`json:"status_sensor" gorm:"foreignkey:StatusSensorId; references:StatusSensorId"`
	Sensor         	[]SensorDetail			`json:"sensor" gorm:"foreignkey:SensorId; references:SensorId"`
	//SocketAction    model_db.SocketAction	`json:"socket_action" gorm:"foreignkey:StatusSensorId; references:StatusSensorId"`
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
	Mainbox        	[]model_db.Mainbox			`json:"mainbox_detail" gorm:"foreignkey:MainboxId; references:MainboxId"`
}
func (SenSocMainList) TableName() string {
	return config.DB_TRANS_SOCKET_AREA
}

//-------------------------------------------------------------------------------//
//				 	    Plant And PlantType
//-------------------------------------------------------------------------------//
//Model
type FarmSensorDetail struct {
	FarmArea    	model_db.FarmArea		`json:"farm_area" gorm:"embedded"`
	FormulaPlant	model_db.FormulaPlant	`json:"formula_plant" gorm:"foreignkey:FormulaPlantId; references:FormulaPlantId"`
	SensorDetail	[]SenSocMainList		`json:"sensor_detail" gorm:"foreignkey:FarmAreaId; references:FarmAreaId"`
}
func (FarmSensorDetail) TableName() string {
	return config.DB_FARM_AREA
}

