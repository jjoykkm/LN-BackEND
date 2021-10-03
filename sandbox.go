package main

import (
	"fmt"
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/config"
	"github.com/jjoykkm/ln-backend/models/model_databases"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type Farm struct {
	FarmId      	uuid.UUID	 `mapstructure:"farm_id" json:"farm_id,omitempty"`
	FarmName    	string		 `mapstructure:"farm_name" json:"farm_name,omitempty"`
	FarmDesc    	string		 `mapstructure:"farm_desc" json:"farm_desc,omitempty"`
	CreateDate		time.Time	 `mapstructure:"create_date" json:"create_date,omitempty"`
	ChangeDate	    time.Time	 `mapstructure:"change_date" json:"change_date,omitempty"`
	StatusId		uuid.UUID	 `mapstructure:"status_id" json:"status_id,omitempty"`
}
func (Farm) TableName() string {
	return "farm"
}
type JJ struct {
	Plant     	model_databases.Plant	 		`mapstructure:"plant_type_id" json:"plant_type_id" gorm:"embedded"`
	PlantType   model_databases.PlantType		 `mapstructure:"plant_type_name" json:"plant_type_name" gorm:"embedded"`
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

	var results []Farm
	var count int64

	//var results []JJ
	db.Model(&results).Count(&count)
	fmt.Println(count)
	fmt.Println(results)
	//db.Model(&results).Association("StatusId").Count()
	//db.Table("plant").Joins("inner join plant_type on plant.plant_type_id = plant_type.plant_type_id").Scan(&results)
	//helper.ConvertToJson(results)
	//fmt.Printf("%+v\n", results)

	//add data to map
	//jj := map[string]interface{}{}
	//result := db.Model(&Farm{}).First(&jj)
	//result := db.First(&fm)
	//fmt.Printf("%+v\n", jj)
	//fmt.Printf("%+v\n", fm)
	//fmt.Println(result.RowsAffected) // returns count of records found
	//fmt.Println(result.Error) // returns error or nil
	//fmt.Println(gorm.ErrRecordNotFound) // returns error or nil
	//
	//errors.Is(result.Error, gorm.ErrRecordNotFound)
}