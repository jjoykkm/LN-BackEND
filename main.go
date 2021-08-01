package main

import (
	"LN-BackEND/config"
	"LN-BackEND/controllers"
	"LN-BackEND/models/model_other"
	"LN-BackEND/utility"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

type Ln struct {
	Db			*sql.DB
	Ctrl 		controllers.Ln
}

var me Ln

func main() {
	db, err := sql.Open(config.DRIVER_NAME, config.CONNECTION)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Successfully connected to DB!\n")

	controller:= controllers.Ln{db}
	me = Ln{db, controller}

	////GetFertilizerRatioRelate(status, formulaPlantId, language string) ([]model_services.ForPlantFert, int)
	//fertilizerRatio, _ := controller.GetFertilizerRatioRelate(config.STATUS_ACTIVE, "0c5c896d-dede-41de-ab78-1f1d41dd95cf", config.LANGUAGE_EN)
	//fmt.Println(fertilizerRatio)

	////GetSensorValueRecRelate(status, formulaPlantId, language string) ([]model_services.ForPlantSensor, int)
	//sensorValueRec, _ := controller.GetSensorValueRecRelate(config.STATUS_ACTIVE, "0c5c896d-dede-41de-ab78-1f1d41dd95cf", config.LANGUAGE_EN)
	//fmt.Println(sensorValueRec)

	//_, _, formulaPlantMap := controller.GetFavoriteFormulaPlanter(config.STATUS_ACTIVE, "9c21bd63-a1f0-490a-bb5b-1de7ab502a1d")
	//fmt.Println(formulaPlantMap)
	//
	////http.GET("/test", func(c *gin.Context) {
	////	c.JSON(200, gin.H{
	////		"message": "test",
	////	})
	////})

	http := gin.Default()
	http.POST("/test", Test)
	http.POST("/plantCategoryList", GetPlantCategoryList)
	http.POST("/plantCategoryItem", GetPlantCategoryItem)
	http.POST("/plantOverviewFavorite", GetPlantOverviewFavorite)
	http.POST("/myPlantOverview", GetMyPlantOverview)
	http.POST("/plantOverviewByPlant", GetPlantOverviewByPlant)
	http.POST("/formulaPlantDetail", GetFormulaPlantDetail)
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
	c.JSON(http.StatusOK, gin.H{
		"message": "00001",
	})
}

func GetPlantCategoryList(c *gin.Context) {
	var bodyModel model_other.PostBody
	bodyModel = utility.GetModelFromBody(c)
	//GetPlantCategoryLister(status, language string) ([]model_services.ForPlantCatList, int)
	plantCategoryList, total := me.Ctrl.GetPlantCategoryLister(config.STATUS_ACTIVE, bodyModel.Language)
	c.JSON(200, gin.H{
		"item": plantCategoryList,
		"total": total,
	})
}

func GetPlantCategoryItem(c *gin.Context) {
	var bodyModel model_other.PostBody
	bodyModel = utility.GetModelFromBody(c)
	//GetPlantCategoryItemer(status, plantTypeId, language string, offset int) ([]model_services.ForPlantCat, int, int)
	plantCategoryItem, nextOffset, total := me.Ctrl.GetPlantCategoryItemer(config.STATUS_ACTIVE, bodyModel.PlantTypeId,  bodyModel.Language, bodyModel.Offset)
	c.JSON(200, gin.H{
		"item": plantCategoryItem,
		"offset": nextOffset,
		"total": total,
	})
}

func GetPlantOverviewFavorite(c *gin.Context) {
	var bodyModel model_other.PostBody
	bodyModel = utility.GetModelFromBody(c)
	//GetPlantOverviewFavoriteer(status, uid, language string, offset int) ([]model_services.ForPlantItem, int, int)
	plantOverviewFavorite, nextOffset, total := me.Ctrl.GetPlantOverviewFavoriteer(config.STATUS_ACTIVE, bodyModel.Uid, bodyModel.Language, bodyModel.Offset)
	c.JSON(200, gin.H{
		"item": plantOverviewFavorite,
		"offset": nextOffset,
		"total": total,
	})
}

func GetMyPlantOverview(c *gin.Context) {
	var bodyModel model_other.PostBody
	bodyModel = utility.GetModelFromBody(c)
	//GetMyPlantOverviewer(status, uid, language string, offset int) ([]model_services.ForPlantItem, int, int)
	myPlantOverview, nextOffset, total := me.Ctrl.GetMyPlantOverviewer(config.STATUS_ACTIVE, bodyModel.Uid, bodyModel.Language, bodyModel.Offset)
	c.JSON(200, gin.H{
		"item": myPlantOverview,
		"offset": nextOffset,
		"total": total,
	})
}

func GetPlantOverviewByPlant(c *gin.Context) {
	var bodyModel model_other.PostBody
	bodyModel = utility.GetModelFromBody(c)
	//GetPlantOverviewByPlanter(status, uid, plantId, language string, offset int) ([]model_services.ForPlantItem, int, int)
	plantOverviewByPlant, nextOffset, total := me.Ctrl.GetPlantOverviewByPlanter(config.STATUS_ACTIVE, bodyModel.Uid, bodyModel.PlantId, bodyModel.Language, bodyModel.Offset)
	c.JSON(200, gin.H{
		"item": plantOverviewByPlant,
		"offset": nextOffset,
		"total": total,
	})
}

func GetFormulaPlantDetail(c *gin.Context) {
	var bodyModel model_other.PostBody
	bodyModel = utility.GetModelFromBody(c)
	//GetFormulaPlantDetailer(status, formulaPlantId, language string) model_services.ForPlantFormula
	formulaPlant := me.Ctrl.GetFormulaPlantDetailer(config.STATUS_ACTIVE, bodyModel.FormulaPlantId, bodyModel.Language)
	c.JSON(200, gin.H{
		"item": formulaPlant,
	})
}
