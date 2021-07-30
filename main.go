package main

import (
	"LN-BackEND/config"
	"LN-BackEND/controllers"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	_ "github.com/lib/pq"

)

func main() {
	db, err := sql.Open(config.DRIVER_NAME, config.CONNECTION)
	if err != nil {
		log.Fatal(err)
	}

	formulaPlant := controllers.Ln{db}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Successfully connected to DB!\n")



	http := gin.Default()
	http.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})

	http.POST("/plantCategoryList", func(c *gin.Context) {
		plantCategoryList := formulaPlant.GetPlantCategoryLister(config.STATUS_ACTIVE, config.LANGUAGE_EN)
		c.JSON(200, gin.H{
			"item": plantCategoryList,
		})
	})

	http.POST("/plantCategoryItem", func(c *gin.Context) {
		plantCategoryItem, nextOffset := formulaPlant.GetPlantCategoryItemer(config.STATUS_ACTIVE, "",  config.LANGUAGE_EN, 0)
		c.JSON(200, gin.H{
			"item": plantCategoryItem,
			"offset": nextOffset,
		})
	})


	http.POST("/plantOverviewFavorite", func(c *gin.Context) {
		plantOverviewFavorite, nextOffset := formulaPlant.GetPlantOverviewFavoriteer(config.STATUS_ACTIVE, "6f08ea87-47dd-4511-be6c-3f2f6603de6c", config.LANGUAGE_EN, 0)
		c.JSON(200, gin.H{
			"item": plantOverviewFavorite,
			"offset": nextOffset,
		})
	})

	http.Run(config.SERVER_HOST)



	////GetCountryName(db *sql.DB, countryId string, language string) string
	//countryName := controllers.GetCountryName(db, "067ea4ff-25ef-47b4-b566-fc2ee28aa07e", config.LANGUAGE_EN)
	//fmt.Println(countryName)


	//router := gin.Default()
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
	//plantOverviewFavorite, _ := controllers.GetPlantOverviewFavorite(db, config.STATUS_ACTIVE, "6f08ea87-47dd-4511-be6c-3f2f6603de6c", config.LANGUAGE_EN, 0)
	//fmt.Println(plantOverviewFavorite)
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