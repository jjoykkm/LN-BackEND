package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jjoykkm/ln-backend/common/config"
	"github.com/jjoykkm/ln-backend/controllers"
	"github.com/jjoykkm/ln-backend/modelsOld/model_other"
	"github.com/jjoykkm/ln-backend/smartfarm/sf_common"
	"github.com/jjoykkm/ln-backend/smartfarm/sf_dashboard"
	"github.com/jjoykkm/ln-backend/smartfarm/sf_formula_plant"
	"github.com/jjoykkm/ln-backend/smartfarm/sf_my_farm"
	"github.com/jjoykkm/ln-backend/utility"
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

type Jj struct {
	Jjoy string
	Eiei string
}

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

	// Common
	repoCommon := sf_common.NewRepository(db)
	serviceCommon := sf_common.NewService(repoCommon)
	handlerCommon := sf_common.NewHandler(serviceCommon)
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
				address := common.Group("/address")
				{
					address.POST("/farmList", handlerCommon.GetFarmList)
					address.POST("/farmAndFarmAreaList", handlerCommon.GetFarmAndFarmAreaList)
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
					address.GET("/provinceList", handlerCommon.GetProvinceList)
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
	bodyModel = utility.GetModelFromBody(c)
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
	bodyModel = utility.GetModelFromBody(c)
	// GetScheReminder(status string, farmAreaId []string) model_services.ScheduleScheRemind
	scheRemind := me.Ctrl.GetScheReminder(config.GetStatus().Active, bodyModel.FarmAreaIdList)

	//fmt.Printf("%+v\n", scheRemind)
	c.JSON(http.StatusOK, gin.H{
		"item": scheRemind,
	})
}
