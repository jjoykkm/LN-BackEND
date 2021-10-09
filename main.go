package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jjoykkm/ln-backend/smartfarm/sf_formula_plant"
	"github.com/jjoykkm/ln-backend/config"
	"github.com/jjoykkm/ln-backend/controllers"
	"github.com/jjoykkm/ln-backend/models/model_other"
	"github.com/jjoykkm/ln-backend/models/model_services"
	"github.com/jjoykkm/ln-backend/utility"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"reflect"
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
	//db, err := sqlx.Connect("postgres", "host=103.212.181.187 user=ln02t password=ln-0110-2 dbname=smartlife port=5432 sslmode=disable")
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//fmt.Println(db)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  config.DSN,
		PreferSimpleProtocol: false, // disables implicit prepared statement usage
	}), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	//var oo []Jj
	////db.Raw("SELECT jjoy, eiei FROM tests").Scan(&oo)
	//db.Exec("SELECT jjoy, eiei FROM tests").Scan(&oo)
	////db.Raw("SELECT * FROM tests WHERE jjoy = ?", "111").Scan(&oo)
	//fmt.Printf("%+v/n",oo)
	//db, err := sql.Open(config.DRIVER_NAME, config.CONNECTION)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//defer db.Close()
	//
	//err = db.Ping()
	//if err != nil {
	//	panic(err)
	//}
	fmt.Printf("Successfully connected to DB!\n")

	controller := controllers.Ln{db}
	me = Ln{db, controller}

	http := gin.Default()
	http.Use(cors.Default())

	repoFormulaPlant := sf_formula_plant.NewRepository(db)
	//cache := cache.New(1*time.Hour, 1*time.Hour)
	serviceFormulaPlant := sf_formula_plant.NewService(repoFormulaPlant)
	handlerFormulaPlant := sf_formula_plant.NewHandler(serviceFormulaPlant)


	http.GET("/formulaPlant/plantCategoryList/api/v1/run", func(c *gin.Context) {
		handlerFormulaPlant.GetPlantCategoryList(c)
	})
	http.POST("/formulaPlant/plantCategoryItem/api/v1/run", func(c *gin.Context) {
		handlerFormulaPlant.GetPlantCategoryItem(c)
	})
	http.POST("/formulaPlant/plantOverviewByPlant/api/v1/run", func(c *gin.Context) {
		handlerFormulaPlant.GetPlantOverviewByPlant(c)
	})
	http.POST("/formulaPlant/plantOverviewFavorite/api/v1/run", func(c *gin.Context) {
		handlerFormulaPlant.GetPlantOverviewFavorite(c)
	})
	http.POST("/formulaPlant/myPlantOverview/api/v1/run", func(c *gin.Context) {
		handlerFormulaPlant.GetMyPlantOverview(c)
	})
	http.POST("/formulaPlant/formulaPlantDetail/api/v1/run", func(c *gin.Context) {
		handlerFormulaPlant.GetFormulaPlantDetail(c)
	})


	http.POST("/jjoy", func(c *gin.Context) {
		handlerFormulaPlant.GetPlantCategoryList(c)
	})


	// Formula Plant
	http.POST("/test", Test)
	//http.POST("/plantCategoryList", GetPlantCategoryList)
	//http.POST("/plantCategoryItem", GetPlantCategoryItem)
	//http.POST("/plantOverviewFavorite", GetPlantOverviewFavorite)
	//http.POST("/myPlantOverview", GetMyPlantOverview)
	//http.POST("/plantOverviewByPlant", GetPlantOverviewByPlant)
	//http.POST("/formulaPlantDetail", GetFormulaPlantDetail)

	// Dashboard
	http.POST("/farmList", GetFarmList)
	http.POST("/farmAreaDashboardList", GetFarmAreaDashboardList)
	http.POST("/farmAreaDetailSensor", GetFarmAreaDetailSensor)

	// My Farm
	http.POST("/overviewFarm", GetOverviewFarm)
	http.POST("/manageMainbox", GetManageMainbox)
	http.POST("/manageFarmArea", GetManageFarmArea)
	http.POST("/manageRole", GetManageRole)

	// Schedule + Reminder
	http.POST("/farmAreaList", GetFarmAreaList)
	http.POST("/scheRemind", GetScheRemind)

	http.Run(config.SERVER_HOST)

}

func Test(c *gin.Context) {
	aa, bb := me.Ctrl.GetManageMainboxer(config.GetStatus().Active, config.GetLanguage().En, "41470e4b-005d-4df9-aa4d-c59f37f6390b")

	c.JSON(http.StatusOK, gin.H{
		"01": aa,
		"02": bb,
		//"03": cc,
		//"04": dd,
		//"05": ee,
	})

}

func GetPlantCategoryList(c *gin.Context) {
	var bodyModel model_other.PostBody
	bodyModel = utility.GetModelFromBody(c)
	//GetPlantCategoryLister(status, language string) ([]model_services.ForPlantCatList, int)
	plantCategoryList, total := me.Ctrl.GetPlantCategoryLister(config.GetStatus().Active, bodyModel.Language)
	if total == 0 {
		c.JSON(http.StatusNoContent, gin.H{})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"item":  plantCategoryList,
			"total": total,
		})
	}
}

func GetPlantCategoryItem(c *gin.Context) {
	var bodyModel model_other.PostBody
	bodyModel = utility.GetModelFromBody(c)
	//GetPlantCategoryItemer(status, plantTypeId, language string, offset int) ([]model_services.ForPlantCat, int, int)
	plantCategoryItem, nextOffset, total := me.Ctrl.GetPlantCategoryItemer(config.GetStatus().Active, bodyModel.PlantTypeId, bodyModel.Language, bodyModel.Offset)
	if total == 0 {
		c.JSON(http.StatusNoContent, gin.H{})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"item":   plantCategoryItem,
			"offset": nextOffset,
			"total":  total,
		})
	}
}

func GetPlantOverviewFavorite(c *gin.Context) {
	var bodyModel model_other.PostBody
	bodyModel = utility.GetModelFromBody(c)
	//GetPlantOverviewFavoriteer(status, uid, language string, offset int) ([]model_services.ForPlantItem, int, int)
	plantOverviewFavorite, nextOffset, total := me.Ctrl.GetPlantOverviewFavoriteer(config.GetStatus().Active, bodyModel.Uid, bodyModel.Language, bodyModel.Offset)
	if total == 0 {
		c.JSON(http.StatusNoContent, gin.H{})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"item":   plantOverviewFavorite,
			"offset": nextOffset,
			"total":  total,
		})
	}
}

func GetMyPlantOverview(c *gin.Context) {
	var bodyModel model_other.PostBody
	bodyModel = utility.GetModelFromBody(c)
	//GetMyPlantOverviewer(status, uid, language string, offset int) ([]model_services.ForPlantItem, int, int)
	myPlantOverview, nextOffset, total := me.Ctrl.GetMyPlantOverviewer(config.GetStatus().Active, bodyModel.Uid, bodyModel.Language, bodyModel.Offset)
	if total == 0 {
		c.JSON(http.StatusNoContent, gin.H{})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"item":   myPlantOverview,
			"offset": nextOffset,
			"total":  total,
		})
	}
}

func GetPlantOverviewByPlant(c *gin.Context) {
	var bodyModel model_other.PostBody
	bodyModel = utility.GetModelFromBody(c)
	//GetPlantOverviewByPlanter(status, uid, plantId, language string, offset int) ([]model_services.ForPlantItem, int, int)
	plantOverviewByPlant, nextOffset, total := me.Ctrl.GetPlantOverviewByPlanter(config.GetStatus().Active, bodyModel.Uid, bodyModel.PlantId, bodyModel.Language, bodyModel.Offset)
	if total == 0 {
		c.JSON(http.StatusNoContent, gin.H{})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"item":   plantOverviewByPlant,
			"offset": nextOffset,
			"total":  total,
		})
	}
}

func GetFormulaPlantDetail(c *gin.Context) {
	var bodyModel model_other.PostBody
	bodyModel = utility.GetModelFromBody(c)

	//GetFormulaPlantDetailer(status, formulaPlantId, language string) model_services.ForPlantFormula
	//_, _, formulaPlant := me.Ctrl.GetSensorValueRecRelate(config.GetStatus().Active, "243367fe-fc14-4074-8ff5-374220dadf8f", bodyModel.Language)
	formulaPlant := me.Ctrl.GetFormulaPlantDetailer(config.GetStatus().Active, bodyModel.FormulaPlantId, bodyModel.Language)
	if reflect.DeepEqual(formulaPlant, model_services.ForPlantFormula{}) {
		c.JSON(http.StatusNoContent, gin.H{})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"item": formulaPlant,
		})
	}
}

// ------------------- dashboard ------------------------------ //
func GetFarmList(c *gin.Context) {
	var bodyModel model_other.PostBody
	bodyModel = utility.GetModelFromBody(c)
	//GetFarmLister(status, uid string) ([]model_services.DashboardFarmList, int)
	farmList, total := me.Ctrl.GetFarmLister(config.GetStatus().Active, bodyModel.Uid)
	if total == 0 {
		c.JSON(http.StatusNoContent, gin.H{})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"item":  farmList,
			"total": total,
		})
	}
}

func GetFarmAreaDashboardList(c *gin.Context) {
	var bodyModel model_other.PostBody
	bodyModel = utility.GetModelFromBody(c)
	//GetFarmAreaDashboardLister(status, language, farmId string) ([]model_services.DashboardFarmList, int)
	farmAreaList, total := me.Ctrl.GetFarmAreaDashboardLister(config.GetStatus().Active, bodyModel.Language, bodyModel.FarmId)
	if total == 0 {
		c.JSON(http.StatusNoContent, gin.H{})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"item":  farmAreaList,
			"total": total,
		})
	}
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

func GetOverviewFarm(c *gin.Context) {
	var bodyModel model_other.PostBody
	bodyModel = utility.GetModelFromBody(c)
	//GetOverviewFarmer(status, farmId string) (model_services.MyFarmOverviewFarm)
	overviewFarm := me.Ctrl.GetOverviewFarmer(config.GetStatus().Active, bodyModel.FarmId)
	if reflect.DeepEqual(overviewFarm, model_services.MyFarmOverviewFarm{}) {
		c.JSON(http.StatusNoContent, gin.H{})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"item": overviewFarm,
		})
	}
}

func GetManageMainbox(c *gin.Context) {
	var bodyModel model_other.PostBody
	bodyModel = utility.GetModelFromBody(c)
	// GetManageMainboxer(status, language, farmId string) ([]model_services.MyFarmManageMainbox, int)
	manageMainbox, total := me.Ctrl.GetManageMainboxer(config.GetStatus().Active, bodyModel.Language, bodyModel.FarmId)
	if total == 0 {
		c.JSON(http.StatusNoContent, gin.H{})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"item":  manageMainbox,
			"total": total,
		})
	}
}

func GetManageFarmArea(c *gin.Context) {
	var bodyModel model_other.PostBody
	bodyModel = utility.GetModelFromBody(c)
	// GetManageFarmAreaer(status, language, farmId string) ([]model_services.MyFarmManageFarmArea, int)
	manageFarmArea, total := me.Ctrl.GetManageFarmAreaer(config.GetStatus().Active, bodyModel.Language, bodyModel.FarmId)
	if total == 0 {
		c.JSON(http.StatusNoContent, gin.H{})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"item":  manageFarmArea,
			"total": total,
		})
	}
}

func GetManageRole(c *gin.Context) {
	var bodyModel model_other.PostBody
	bodyModel = utility.GetModelFromBody(c)
	// GetManageRoleer(status, language, farmId string) ([]model_services.MyFarmManageRole, int)
	manageRole, total := me.Ctrl.GetManageRoleer(config.GetStatus().Active, bodyModel.Language, bodyModel.FarmId)
	if total == 0 {
		c.JSON(http.StatusNoContent, gin.H{})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"item":  manageRole,
			"total": total,
		})
	}
}

func GetFarmAreaList(c *gin.Context) {
	var bodyModel model_other.PostBody
	bodyModel = utility.GetModelFromBody(c)
	// GetFarmAreaLister(status, farmId string) ([]model_services.ScheduleFarmArea, int)
	farmArea, total := me.Ctrl.GetFarmAreaLister(config.GetStatus().Active, bodyModel.FarmId)
	if total == 0 {
		c.JSON(http.StatusNoContent, gin.H{})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"item":  farmArea,
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
	//if reflect.ValueOf(scheRemind).IsZero() {
	//	c.JSON(http.StatusNoContent, gin.H{})
	//}else {
	//	c.JSON(http.StatusOK, gin.H{
	//		"item":  scheRemind,
	//	})
	//}
}
