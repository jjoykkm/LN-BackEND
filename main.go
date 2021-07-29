package main

import (
	"LN-BackEND/config"
	"LN-BackEND/controllers"
	"database/sql"
	"fmt"
	//"github.com/gin-gonic/gin"
	//"github.com/gin-gonic/gin"
	"log"
	"reflect"

	//"time"
	//"github.com/gin-gonic/gin/binding"
	_ "github.com/lib/pq"
	//"gorm.io/driver/postgres"
	//"gorm.io/gorm"
	//"time"
)

var dbs *sql.DB

func main() {
	db, err := sql.Open(config.DRIVER_NAME, config.CONNECTION)
	if err != nil {
		log.Fatal(err)
	}
	dbs = db
	fmt.Println(reflect.TypeOf(db))
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Successfully connected to DB!\n")

	////GetCountryName(db *sql.DB, status string, countryId string, language string) string
	//countryName := controllers.GetCountryName(db, config.STATUS_ACTIVE,"067ea4ff-25ef-47b4-b566-fc2ee28aa07e", config.LANGUAGE_EN)
	//fmt.Println(countryName)


	//router := gin.Default()
	//router.POST("/find", jjoy)
	//router.POST("/test", test)
	//router.Run(config.SERVER_HOST)

	////GetPlantCategoryList (db *sql.DB, status string, language string) []model_services.ForPlantCatList
	//plantCategoryList := controllers.GetPlantCategoryList(db, config.STATUS_ACTIVE, config.LANGUAGE_EN)
	//fmt.Println(plantCategoryList)

	////GetPlantCategoryItem (db *sql.DB, status string, plantTypeId string, language string, offset string) ([]model_services.ForPlantCat, int)
	//plantCategoryItem, nextOffset := controllers.GetPlantCategoryItem(db, config.STATUS_ACTIVE, "",  config.LANGUAGE_EN, 0)
	//fmt.Println(plantCategoryItem)
	//fmt.Println(nextOffset)

	////GetFavoriteFormulaPlant (db *sql.DB, status string, uid string, language string) ([]model_databases.FavoritePlant, []uuid.UUID)
	//favoritePlant := controllers.GetFavoriteFormulaPlant(db, config.STATUS_ACTIVE, "6f08ea87-47dd-4511-be6c-3f2f6603de6c", config.LANGUAGE_EN)
	//fmt.Println(favoritePlant)

	////GetPlantOverviewFavorite(db *sql.DB, status string, uid string, language string, offset int) ([]model_services.ForPlantItem, int)
	jj,_ := controllers.GetPlantOverviewFavorite(db, config.STATUS_ACTIVE, "6f08ea87-47dd-4511-be6c-3f2f6603de6c", config.LANGUAGE_EN, 0)
	fmt.Println(jj)
}

//func jjoy(c *gin.Context)  {
//	plantCategoryList := controllers.GetPlantCategoryList(dbs, config.STATUS_ACTIVE, config.LANGUAGE_EN)
//	fmt.Println(plantCategoryList)
//}
//func test(c *gin.Context) {
//	c.JSON(http.StatusOK, gin.H{
//		"message": "00001",
//	})
//}