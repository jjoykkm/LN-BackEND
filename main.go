package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jjoykkm/ln-backend/common/config"
	"github.com/jjoykkm/ln-backend/controllers"
	"github.com/jjoykkm/ln-backend/modelsOld/model_other"
	"github.com/jjoykkm/ln-backend/smartfarm/sf_common/cm_address"
	"github.com/jjoykkm/ln-backend/smartfarm/sf_common/cm_farm"
	"github.com/jjoykkm/ln-backend/smartfarm/sf_common/cm_plant"
	"github.com/jjoykkm/ln-backend/smartfarm/sf_dashboard"
	"github.com/jjoykkm/ln-backend/smartfarm/sf_formula_plant"
	"github.com/jjoykkm/ln-backend/smartfarm/sf_my_farm"
	"github.com/jjoykkm/ln-backend/obsolete_utility"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
)

type Ln struct {
	Db   *gorm.DB
	Ctrl controllers.Ln
}

var me Ln

func main() {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  config.DSN,
		PreferSimpleProtocol: false, // disables implicit prepared statement usage
	}), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	controller := controllers.Ln{db}
	me = Ln{db, controller}

	http := gin.Default()
	http.Use(cors.Default())

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
					plant.POST("/testttt", handlerCommonPlant.GetFertAndCatList)
				}
			}
			forPlant := api.Group("/formulaPlant")
			{
				forPlant.POST("/forPlantCategoryItemList", handlerFormulaPlant.GetPlantCategoryItem)
				forPlant.POST("/forPlantListByPlant", handlerFormulaPlant.GetPlantOverviewByPlant)
				forPlant.POST("/forPlantListFavorite", handlerFormulaPlant.GetPlantOverviewFavorite)
				forPlant.POST("/forPlantListOfMine", handlerFormulaPlant.GetPlantOfMine)
				forPlant.POST("/forPlantDetail", handlerFormulaPlant.GetFormulaPlantDetail)
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
				}
			}
		}
	}

	// Dashboard
	http.POST("/farmAreaDetailSensor", GetFarmAreaDetailSensor)
	// Schedule + Reminder
	http.POST("/scheRemind", GetScheRemind)

	http.Run(config.SERVER_HOST)

}

func GetFarmAreaDetailSensor(c *gin.Context) {
	var bodyModel model_other.PostBody
	bodyModel = obsolete_utility.GetModelFromBody(c)
	//GetFarmAreaDetailSensorer(status, farmId, language string) ([]model_services.SenSocMainList, int)
	senSocMainList, total := me.Ctrl.GetFarmAreaDetailSensorer(config.GetStatus().Active, bodyModel.FarmAreaId, bodyModel.Language)
	if total == 0 {
		c.JSON(http.StatusNoContent, gin.H{})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"item":  senSocMainList,
			"total": total,
		})
	}
}

func GetScheRemind(c *gin.Context) {
	var bodyModel model_other.PostBody
	bodyModel = obsolete_utility.GetModelFromBody(c)
	// GetScheReminder(status string, farmAreaId []string) model_services.ScheduleScheRemind
	scheRemind := me.Ctrl.GetScheReminder(config.GetStatus().Active, bodyModel.FarmAreaIdList)

	//fmt.Printf("%+v\n", scheRemind)
	c.JSON(http.StatusOK, gin.H{
		"item": scheRemind,
	})
}
