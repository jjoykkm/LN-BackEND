package model_db

//-------------------------------------------------------------------------------//
//							Table Sensor Value
//-------------------------------------------------------------------------------//
//model country
type SensorValue struct {
	Humidity       	float64	 `json:"humidity"`
	Temperature     float64	 `json:"temperature"`
}

// Custom table name for GORM
func (SensorValue) TableName() string {
	return "tbl_sensor_value"
}
