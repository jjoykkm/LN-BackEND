package main

import (
	"errors"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jjoykkm/ln-backend/common/config"
	"github.com/jjoykkm/ln-backend/modelsOld/model_databases"
	"github.com/jjoykkm/ln-backend/smartfarm/sf_dashboard"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
)

//type Test struct {
//	Plant  model_db.Plant `gorm:"embedded"`
//	PlantType  model_db.PlantType `gorm:"foreignkey:PlantTypeId; references:PlantTypeId"`
//}

type JJ struct {
	//PlantId         uuid.UUID	 `gorm:"primaryKey" json:"plant_id"`
	//PlantTypeId     uuid.UUID	 //`gorm:"column:plant_type_id" json:"plant_type_id"`
	////CreateDate		time.Time	 `json:"create_date"`
	////ChangeDate	    time.Time	 `json:"change_date"`
	////StatusId		uuid.UUID	 `json:"status_id"`
	////Plant1     	Plant	 	`gorm:"embedded"`//`gorm:"embeddedPrefix:p_"` // //
	Plant  model_databases.Plant `gorm:"embedded"`
	PlantType  model_databases.PlantType `gorm:"foreignkey:PlantTypeId; references:PlantTypeId"` //`gorm:"embeddedPrefix:pt_"` //
}
func (JJ) TableName() string {
	return "plant"
}
type RequestError struct {
	StatusCode int
	Err error
}

func (r *RequestError) Error() string {
	return r.Err.Error()
}

func (r *RequestError) Temporary() bool {
	return r.StatusCode == http.StatusNoContent //http.StatusServiceUnavailable // 503
}

func doRequest() error {
	return &RequestError{
		StatusCode: 503,
		Err:        errors.New("unavailable"),
	}
}

func main()  {
	//var fm Farm
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  config.DSN,
		PreferSimpleProtocol: false, // disables implicit prepared statement usage
	}), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	http := gin.Default()
	http.Use(cors.Default())
	//fmt.Println("***********************************")
	//ss := doRequest()
	//if ss != nil {
	//	fmt.Println(ss)
	//	re, ok := ss.(*RequestError)
	//	fmt.Println(re)
	//	if ok {
	//		if re.Temporary() {
	//			fmt.Println("This request can be tried again")
	//		} else {
	//			fmt.Println("This request cannot be tried again")
	//		}
	//	}
	//}
	result := []sf_dashboard.FarmSensorDetail{}
	//result := []model_dbmodel_db.SensorDetail{}

	//sensorDet := db.Debug().Where("status_id = ?", config.GetStatus().Active).Preload("SensorType",
	//	"status_id = ?", config.GetStatus().Active)
	//
	//socketDet := db.Debug().Where("status_id = ?", config.GetStatus().Active).Preload("StatusSensor",
	//	"status_id = ?", config.GetStatus().Active).Preload("Sensor",
	//	func(db *gorm.DB) *gorm.DB {
	//		return sensorDet
	//	})
	//
	//
	//mainbox := db.Debug().Where("status_id = ?", config.GetStatus().Active)
	//
	//senSocMain := db.Debug().Preload("Socket",
	//	func(db *gorm.DB) *gorm.DB {
	//		return socketDet
	//	}).Preload("Mainbox",
	//	func(db *gorm.DB) *gorm.DB {
	//		return mainbox
	//	})

	db.Debug().Preload("FormulaPlant").Find(&result)
	//.Preload("SensorDetail",
	//	func(db *gorm.DB) *gorm.DB {
	//		return senSocMain
	//	})
	//db.Debug().Preload("Mainbox").Find(&result)
	fmt.Println("-----------------------------")
	//fmt.Printf("%+v\n", result)



	http.POST("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"01": result,
		})
	})
	http.Run(config.SERVER_HOST)
}