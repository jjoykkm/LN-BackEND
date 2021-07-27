package main

import (
	"LN-BackEND/config"
	"database/sql"
	"fmt"
	"log"
	//"time"

	_ "github.com/lib/pq"
	//"gorm.io/driver/postgres"
	//"gorm.io/gorm"
	//"time"
)
type test struct {
	jjoy         string	`mapstructure:"jjoy" json:"jjoy,omitempty"`
	eiei         string	`mapstructure:"eiei" json:"eiei,omitempty"`
}

func main() {
	db, err := sql.Open(config.DriverName, config.Connection)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Successfully connected to DB!\n")

	////GetPlantCategoryList (db *sql.DB, status string, language string) []model_services.ForPlantCatList
	//plantCategoryList := controllers.GetPlantCategoryList(db, config.StatusActive, config.LanguageEN)
	//fmt.Println(plantCategoryList)

	////GetPlantCategoryItem (db *sql.DB, status string, plantTypeId string, language string, offset string) ([]model_services.ForPlantCat, int)
	//plantCategoryItem, nextOffset := controllers.GetPlantCategoryItem(db, config.StatusActive, "",  config.LanguageEN, 0)
	//fmt.Println(plantCategoryItem)
	//fmt.Println(nextOffset)
}