package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jjoykkm/ln-backend/common/config"
	"github.com/jjoykkm/ln-backend/common/models/model_db"
	"github.com/jjoykkm/ln-backend/errs"
	"github.com/jjoykkm/ln-backend/smartfarm/sf_common/cm_address"
	"github.com/jjoykkm/ln-backend/smartfarm/sf_common/cm_farm"
	"github.com/jjoykkm/ln-backend/smartfarm/sf_common/cm_plant"
	"github.com/jjoykkm/ln-backend/smartfarm/sf_dashboard"
	"github.com/jjoykkm/ln-backend/smartfarm/sf_formula_plant"
	"github.com/jjoykkm/ln-backend/smartfarm/sf_my_farm"
	"github.com/jjoykkm/ln-backend/smartfarm/sf_remote_switch"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type Ln struct {
	Db   *gorm.DB
}

var me Ln

func main() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,           // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,          // Disable color
		},
	)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  config.DSN,
		PreferSimpleProtocol: false, // disables implicit prepared statement usage
	//}), &gorm.Config{})
	}),  &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err.Error())
	}

	me = Ln{db}

	http := gin.Default()
	http.Use(cors.Default())

	http.POST("/sensor_value_insert", func(c *gin.Context) {
		var model model_db.SensorValue
		if err := c.Bind(&model); err != nil {
			(&errs.Service{}).ErrMsgBindData(c, err)
		}
		err := db.Create(&model).Error
		if err != nil {
			c.String(400, err.Error())
		}
		c.JSON(200, "Insert Success")
	})
	http.GET("/sensor_value_get", func(c *gin.Context) {
		var model []model_db.SensorValue

		err := db.Find(&model).Error
		if err != nil {
			c.String(400, err.Error())
		}
		c.JSON(200, model)
	})
	http.GET("/sensor_value_delete", func(c *gin.Context) {
		var model []model_db.SensorValue

		err := db.Where("1 = 1").Delete(&model).Error
		if err != nil {
			c.String(400, err.Error())
		}
		c.JSON(200, "Delete Success")
	})


	//------------------------------  Common  ------------------------------------//
	// Farm
	repoCommonFarm := cm_farm.NewRepository(db)
	serviceCommonFarm := cm_farm.NewService(repoCommonFarm)
	handlerCommonFarm := cm_farm.NewHandler(serviceCommonFarm)
	// Address
	repoCommonAD := cm_address.NewRepository(db)
	serviceCommonAD := cm_address.NewService(repoCommonAD)
	handlerCommonAD := cm_address.NewHandler(serviceCommonAD)
	// Plant
	repoCommonPlant := cm_plant.NewRepository(db)
	serviceCommonPlant := cm_plant.NewService(repoCommonPlant)
	handlerCommonPlant := cm_plant.NewHandler(serviceCommonPlant)

	//------------------------------  App  ------------------------------------//
	// Formula Plant
	repoFormulaPlant := sf_formula_plant.NewRepository(db)
	serviceFormulaPlant := sf_formula_plant.NewService(repoFormulaPlant)
	handlerFormulaPlant := sf_formula_plant.NewHandler(serviceFormulaPlant)
	// Dashboard
	repoDashboard := sf_dashboard.NewRepository(db)
	serviceDashboard := sf_dashboard.NewService(repoDashboard)
	handlerDashboard := sf_dashboard.NewHandler(serviceDashboard)
	// My Farm
	repoMyFarm := sf_my_farm.NewRepository(db)
	serviceMyFarm := sf_my_farm.NewService(repoMyFarm)
	handlerMyFarm := sf_my_farm.NewHandler(serviceMyFarm)
	// Schedule + Reminder
	//repoScheRem := sf_my_farm.NewRepository(db)
	//serviceScheRem := sf_my_farm.NewService(repoScheRem)
	//handlerScheRem := sf_my_farm.NewHandler(serviceScheRem)
	//RemoteSwitch
	repoRemoteSwitch := sf_remote_switch.NewRepository(db)
	serviceRemoteSwitch := sf_remote_switch.NewService(repoRemoteSwitch)
	handlerRemoteSwitch := sf_remote_switch.NewHandler(serviceRemoteSwitch)

	v1 := http.Group("/v1")
	{
		api := v1.Group("/api")
		{
			common := api.Group("/common")
			{
				farm := common.Group("/farm")
				{
					farm.POST("/farmList", handlerCommonFarm.GetFarmList)
					farm.POST("/farmAndFarmAreaList", handlerCommonFarm.GetFarmAndFarmAreaList)
				}
				plant := common.Group("/plant")
				{
					plant.POST("/addChangeFertCat", handlerCommonPlant.AddChangeFertCat)
				}
			}
			forPlant := api.Group("/formulaPlant")
			{
				forPlant.POST("/forPlantCategoryItemList", handlerFormulaPlant.GetPlantCategoryItem)
				forPlant.POST("/forPlantListByPlant", handlerFormulaPlant.GetPlantOverviewByPlant)
				forPlant.POST("/forPlantListFavorite", handlerFormulaPlant.GetPlantOverviewFavorite)
				forPlant.POST("/forPlantListOfMine", handlerFormulaPlant.GetPlantOfMine)
				forPlant.POST("/forPlantDetail", handlerFormulaPlant.GetFormulaPlantDetail)
				forPlant.POST("/addChangeFormulaPlant", handlerFormulaPlant.AddChangeFormulaPlant)
				forPlant.POST("/addFavoriteForPlant", handlerFormulaPlant.AddFavoriteForPlant)
				forPlant.POST("/removeFavoriteForPlant", handlerFormulaPlant.RemoveFavoriteForPlant)
				forPlant.POST("/addChangeFertilizer", handlerFormulaPlant.AddChangeFertilizer)
			}
			dashboard := api.Group("/dashboard")
			{
				dashboard.POST("/farmAreaDashboardList", handlerDashboard.GetFarmAreaDashboardList)
			}
			myFarm := api.Group("/myFarm")
			{
				myFarm.POST("/overviewFarm", handlerMyFarm.GetOverviewFarm)
				myFarm.POST("/manageRole", handlerMyFarm.GetManageRole)
				myFarm.POST("/manageFarmArea", handlerMyFarm.GetManageFarmArea)
				myFarm.POST("/manageMainbox", handlerMyFarm.GetManageMainbox)
				myFarm.POST("/activateMainbox", handlerMyFarm.ActivateMainbox)
				myFarm.POST("/configMainbox", handlerMyFarm.ConfigMainbox)
				myFarm.POST("/configFarm", handlerMyFarm.ConfigFarm)
				myFarm.POST("/configCustomSensor", handlerMyFarm.ConfigCustomSensor)
				myFarm.POST("/configDeleteSocket", handlerMyFarm.ConfigDeleteSocket)
				myFarm.POST("/configDeleteMainbox", handlerMyFarm.ConfigDeleteMainbox)
				myFarm.POST("/configDeleteFarm", handlerMyFarm.ConfigDeleteFarm)
				myFarm.POST("/configDeleteFarmArea", handlerMyFarm.ConfigDeleteFarmArea)
				myFarm.POST("/configFarmArea", handlerMyFarm.ConfigFarmArea)
				myFarm.POST("/removeSocketLinkedFarm", handlerMyFarm.RemoveSocketLinkedFarm)
			}
			//scheRem := app.Group("/scheRem")
			//{
			//	//scheRem.GET("/farmAreaDashboardList", handlerScheRem.GetFarmAreaDashboardList)
			//}

			remoteSwitch := api.Group("/remoteSwitch")
			{
				remoteSwitch.POST("/getRemoteSwitch", handlerRemoteSwitch.GetRemoteSwitch)
				remoteSwitch.POST("/configRemoteSwitch", handlerRemoteSwitch.ConfigRemoteSwitch)
				remoteSwitch.POST("/unlinkSocketRemote", handlerRemoteSwitch.UnlinkSocketRemote)
				remoteSwitch.POST("/removeRemoteSwitch", handlerRemoteSwitch.RemoveRemoteSwitch)
				remoteSwitch.POST("/changeStatusSensor", handlerRemoteSwitch.ChangeStatusSensor)
			}

		}

		get := v1.Group("/get")
		{
			forPlant := get.Group("/formulaPlant")
			{
				forPlant.GET("/forPlantCategoryList", handlerFormulaPlant.GetPlantCategoryList)
			}
			common := get.Group("/common")
			{
				address := common.Group("/address")
				{
					address.GET("/provinceList", handlerCommonAD.GetProvinceList)
				}
				plant := common.Group("/plant")
				{
					plant.GET("/fertAndCatList", handlerCommonPlant.GetFertAndCatList)
					plant.GET("/sensorTypeList", handlerCommonPlant.GetSensorTypeList)
					plant.GET("/fertCatList", handlerCommonPlant.GetFertCatList)
				}
			}
		}
	}

	http.Run(config.SERVER_HOST)

}
