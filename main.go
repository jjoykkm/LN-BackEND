package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
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

	////GetFertilizerRatioRelate(status, formulaPlantId, language string) ([]model_services.ForPlantFert, int)
	//fertilizerRatio, _ := controller.GetFertilizerRatioRelate(config.STATUS_ACTIVE, "0c5c896d-dede-41de-ab78-1f1d41dd95cf", config.LANGUAGE_EN)
	//fmt.Println(fertilizerRatio)

	////GetSensorValueRecRelate(status, formulaPlantId, language string) ([]model_services.ForPlantSensor, int)
	//sensorValueRec, _ := controller.GetSensorValueRecRelate(config.STATUS_ACTIVE, "0c5c896d-dede-41de-ab78-1f1d41dd95cf", config.LANGUAGE_EN)
	//fmt.Println(sensorValueRec)

	//_, _, formulaPlantMap := controller.GetFavoriteFormulaPlanter(config.STATUS_ACTIVE, "6f08ea87-47dd-4511-be6c-3f2f6603de6c")
	//fmt.Println(formulaPlantMap)

	//plantCategoryItem, _,_ := me.Ctrl.GetPlantCategoryItemer(config.STATUS_ACTIVE, "",  "EN", 0)
	//fmt.Println(plantCategoryItem)
	//
	//
	////http.GET("/test", func(c *gin.Context) {
	////	c.JSON(http.StatusOK, gin.H{
	////		"message": "test",
	////	})
	////})
	//GetFarmLister(status, uid string) ([]model_services.DashboardFarmList, int)
	//farmList, _ := me.Ctrl.GetFarmLister(config.STATUS_ACTIVE, "17ac6921-ece0-43bc-9d88-7b9bfc59ffd3")
	//fmt.Println(farmList)

	//GetSensorLister(status, language string, socketIdList []string) []model_services.SenSocList
	//ss:=[]string{"ec2dcb62-e741-47da-9288-4a9ca4f797ed","2b71a9c7-1d59-44d5-80f0-e805ddd82d55"}
	//senSocList,_ := me.Ctrl.GetSensorBySocketLister(config.STATUS_ACTIVE, config.LANGUAGE_EN, ss)
	//fmt.Printf("%+v\n",senSocList)
	//GetAllDetailSensor(status, farmId, language string) ([]model_services.SenSocMainList, int)

	http := gin.Default()

	// Formula Plant
	http.POST("/test", Test)
	http.POST("/plantCategoryList", GetPlantCategoryList)
	http.POST("/plantCategoryItem", GetPlantCategoryItem)
	http.POST("/plantOverviewFavorite", GetPlantOverviewFavorite)
	http.POST("/myPlantOverview", GetMyPlantOverview)
	http.POST("/plantOverviewByPlant", GetPlantOverviewByPlant)
	http.POST("/formulaPlantDetail", GetFormulaPlantDetail)

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

	////GetFertilizerer(status, language string, fertilizerIdList []string) ([]model_services.ForPlantFert, int)
	//var ff []string
	//ff = append(ff, "37be0915-ff7a-4539-a6a2-a04a6cb643d1")
	//ff = append(ff, "7d74391f-64ca-422e-83aa-78a507338230")
	//fmt.Println(reflect.TypeOf(ff))
	//fertilizer,_ := controller.GetFertilizerer(config.STATUS_ACTIVE, config.LANGUAGE_EN, ff)
	//fmt.Println(fertilizer)

	////GetCountryName(db *sql.DB, countryId string, language string) string
	//countryName := controllers.GetCountryName(db, "067ea4ff-25ef-47b4-b566-fc2ee28aa07e", config.LANGUAGE_EN)
	//fmt.Println(countryName)

	////GetPlantCategoryList (db *sql.DB, status string, language string) []model_services.ForPlantCatList
	//plantCategoryList := controllers.GetPlantCategoryList(config.STATUS_ACTIVE, config.LANGUAGE_EN)
	//fmt.Println(plantCategoryList)

	////GetPlantCategoryItem (db *sql.DB, status string, plantTypeId string, language string, offset string) ([]model_services.ForPlantCat, int)
	//plantCategoryItem, nextOffset := controllers.GetPlantCategoryItem(config.STATUS_ACTIVE, "",  config.LANGUAGE_EN, 0)
	//fmt.Println(plantCategoryItem)
	//fmt.Println(nextOffset)

	////GetFavoriteFormulaPlant (db *sql.DB, status string, uid string, language string) ([]model_databases.FavoritePlant, []uuid.UUID)
	//favoritePlant := controllers.GetFavoriteFormulaPlant(config.STATUS_ACTIVE, "6f08ea87-47dd-4511-be6c-3f2f6603de6c", config.LANGUAGE_EN)
	//fmt.Println(favoritePlant)

	////GetPlantOverviewFavorite(db *sql.DB, status string, uid string, language string, offset int) ([]model_services.ForPlantItem, int)
	//plantOverviewFavorite, _ := controllers.GetPlantOverviewFavorite(config.STATUS_ACTIVE, "6f08ea87-47dd-4511-be6c-3f2f6603de6c", config.LANGUAGE_EN, 0)
	//fmt.Println(plantOverviewFavorite)

	////GetMyPlantOverview(status, uid, language string, offset int) ([]model_services.ForPlantItem, int)
	//myPlantOverview, _ := controller.GetMyPlantOverview(config.STATUS_ACTIVE, "9c21bd63-a1f0-490a-bb5b-1de7ab502a1d", config.LANGUAGE_EN, 0)
	//fmt.Println(myPlantOverview)
}

func Test(c *gin.Context) {
	aa, bb := me.Ctrl.GetManageMainboxer(config.STATUS_ACTIVE, config.LANGUAGE_EN, "41470e4b-005d-4df9-aa4d-c59f37f6390b")

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
	plantCategoryList, total := me.Ctrl.GetPlantCategoryLister(config.STATUS_ACTIVE, bodyModel.Language)
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
	plantCategoryItem, nextOffset, total := me.Ctrl.GetPlantCategoryItemer(config.STATUS_ACTIVE, bodyModel.PlantTypeId, bodyModel.Language, bodyModel.Offset)
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
	plantOverviewFavorite, nextOffset, total := me.Ctrl.GetPlantOverviewFavoriteer(config.STATUS_ACTIVE, bodyModel.Uid, bodyModel.Language, bodyModel.Offset)
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
	myPlantOverview, nextOffset, total := me.Ctrl.GetMyPlantOverviewer(config.STATUS_ACTIVE, bodyModel.Uid, bodyModel.Language, bodyModel.Offset)
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
	plantOverviewByPlant, nextOffset, total := me.Ctrl.GetPlantOverviewByPlanter(config.STATUS_ACTIVE, bodyModel.Uid, bodyModel.PlantId, bodyModel.Language, bodyModel.Offset)
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
	//_, _, formulaPlant := me.Ctrl.GetSensorValueRecRelate(config.STATUS_ACTIVE, "243367fe-fc14-4074-8ff5-374220dadf8f", bodyModel.Language)
	formulaPlant := me.Ctrl.GetFormulaPlantDetailer(config.STATUS_ACTIVE, bodyModel.FormulaPlantId, bodyModel.Language)
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
	farmList, total := me.Ctrl.GetFarmLister(config.STATUS_ACTIVE, bodyModel.Uid)
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
	farmAreaList, total := me.Ctrl.GetFarmAreaDashboardLister(config.STATUS_ACTIVE, bodyModel.Language, bodyModel.FarmId)
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
	senSocMainList, total := me.Ctrl.GetFarmAreaDetailSensorer(config.STATUS_ACTIVE, bodyModel.FarmAreaId, bodyModel.Language)
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
	overviewFarm := me.Ctrl.GetOverviewFarmer(config.STATUS_ACTIVE, bodyModel.FarmId)
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
	manageMainbox, total := me.Ctrl.GetManageMainboxer(config.STATUS_ACTIVE, bodyModel.Language, bodyModel.FarmId)
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
	manageFarmArea, total := me.Ctrl.GetManageFarmAreaer(config.STATUS_ACTIVE, bodyModel.Language, bodyModel.FarmId)
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
	manageRole, total := me.Ctrl.GetManageRoleer(config.STATUS_ACTIVE, bodyModel.Language, bodyModel.FarmId)
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
	farmArea, total := me.Ctrl.GetFarmAreaLister(config.STATUS_ACTIVE, bodyModel.FarmId)
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
	scheRemind := me.Ctrl.GetScheReminder(config.STATUS_ACTIVE, bodyModel.FarmAreaIdList)

	fmt.Printf("%+v\n", scheRemind)
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
