package main

import (
	"errors"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jjoykkm/ln-backend/common/config"
	"gorm.io/gorm/clause"

	//"github.com/jjoykkm/ln-backend/modelsOld/model_databases"
	//"github.com/jjoykkm/ln-backend/smartfarm/sf_my_farm"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
)
type Test101 struct {
	Id 		int
	Name 	string
	Fname 	string
	Lname 	string
	Jjoy 	string
}
func (Test101) TableName() string {
	return "test101"
}
type JJ struct {
	Jjoy string
	Eiei string
	Kiki string
	Test101 []Test101	`gorm:"foreignkey:Jjoy; references:Jjoy"`
}
func (JJ) TableName() string {
	return "tests"
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
	v1 := http.Group("/v1")
	{
		v1.GET("/login", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"01": "result",
				//"01": result,
			})
		})
		t2 := v1.Group("run")
		t2.GET("/test", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"01": "222222",
				//"01": result,
			})
		})
	}
	//activeId := uuid.UUID{}.Get()

	fmt.Println("***********************************")

	var (
		data JJ
		wa_test Test101
		test1 []Test101
	)
	data.Jjoy = "777"
	data.Eiei = "222"
	data.Kiki = "333"
	//db.Debug().Save(&data)
	wa_test.Id = 111
	wa_test.Name = "1111"
	test1 = append(test1, wa_test)
	wa_test.Id = 222
	wa_test.Name = "test2"
	test1 = append(test1, wa_test)
	data.Test101 = test1
	db.Debug().Omit("Eiei").Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "jjoy"}},
		UpdateAll: true,
	}).Create(&data)//.Model(Test101{})
	//db.Debug().Clauses(clause.OnConflict{
	//	Columns:   []clause.Column{{Name: "jjoy"}},
	//	DoUpdates: clause.AssignmentColumns([]string{"eiei"}),
	//}).Create(&data)
	//data.Test101 = test1
	fmt.Printf("%+v,\n", data)
	//db.Debug().Where("jjoy = ?", "333").Updates(&data)
	//db.Debug().Create(&data)
	//mb := sf_my_farm.model_db.MainboxDatailUS{""}
	//db.Debug().Create(&sf_my_farm.ReqConfMainbox{
	//	Mainbox: "jinzhu",
	//	SocketSensor: CreditCard{Number: "411111111111"}
	//})
	//res1B, _ := json.Marshal(sf_my_farm.ReqConfMainbox{})
	//fmt.Println(res1B)
	//fmt.Printf("%+v\n",res1B)
	//fmt.Fprint(w, string(res1B))
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
	//result := []sf_dashboard.FarmSensorDetail{}
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

	//db.Debug().Preload("FormulaPlant").Find(&result)
	//.Preload("SensorDetail",
	//	func(db *gorm.DB) *gorm.DB {
	//		return senSocMain
	//	})
	//db.Debug().Preload("Mainbox").Find(&result)
	//result := []sf_my_farm.SensorDetail{}
	////Get Sensor Detail
	//sensorDet := db.Debug().Where("status_id = ?", config.GetStatus().Active).Preload("SensorType",
	//	"status_id = ?", config.GetStatus().Active)
	//// Get Socket Detail
	//socketDet := db.Debug().Where("status_id = ?", config.GetStatus().Active).Preload("StatusSensor",
	//	"status_id = ?", config.GetStatus().Active).Preload("Sensor",
	//	func(db *gorm.DB) *gorm.DB {
	//		return sensorDet
	//	})
	// Get Sensor, Socket
	//db.Debug().Where("status_id = ?", config.GetStatus().Active).Preload("Socket").Preload("StatusSensor").Preload("Sensor").Find(&result)
	//,
	//	func(db *gorm.DB) *gorm.DB {
	//		return socketDet
	//	}).Find(&result)
	//db.Debug().Select("farm_area_id").Where("status_id = ? AND farm_id = ?",
	//	config.GetStatus().Active, "41470e4b-005d-4df9-aa4d-c59f37f6390b").Find(&result)//.Table(config.DB_FARM_AREA)
	//fmt.Println("-----------------------------")
	//fmt.Printf("%+v\n", result)
	//var trans []model_db.TransSocketArea
	//var count int64
	//
	//// Get farm_id
	//farmAreaId := db.Debug().Distinct("farm_area_id").Where("status_id = ? AND farm_id = ?",
	//	config.GetStatus().Active, "41470e4b-005d-4df9-aa4d-c59f37f6390b").Table(config.DB_FARM_AREA)
	//
	//db.Debug().Model(&trans).Distinct("mainbox_id").Where("farm_area_id IN (?)", farmAreaId).Count(&count)
	//fmt.Println(count)



	http.POST("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			//"01": result,
			//"01": result,
		})
	})
	http.Run(config.SERVER_HOST)
}