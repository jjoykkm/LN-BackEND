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

type Plant struct {
	PlantId          uuid.UUID	 `gorm:"primaryKey" mapstructure:"plant_id" json:"plant_id,omitempty"`
	PlantTypeId     uuid.UUID	 //`gorm:"column:plant_type_id" mapstructure:"plant_type_id" json:"plant_type_id"`
	CreateDate		time.Time	 `mapstructure:"create_date" json:"create_date,omitempty"`
	ChangeDate	    time.Time	 `mapstructure:"change_date" json:"change_date,omitempty"`
	StatusId		uuid.UUID	 `mapstructure:"status_id" json:"status_id,omitempty"`
}
func (Plant) TableName() string {
	return "plant"
}
type PlantType struct {
	PlantTypeId     uuid.UUID	 //`gorm:"primaryKey" mapstructure:"plant_type_id" json:"plant_type_id"`
	PlantTypeEN      string		 //`mapstructure:"plant_type_en" json:"plant_type_en,omitempty"`
	//CreateDate		time.Time	 `mapstructure:"create_date" json:"create_date,omitempty"`
	//ChangeDate	    time.Time	 `mapstructure:"change_date" json:"change_date,omitempty"`
	//StatusId		uuid.UUID	 `mapstructure:"status_id" json:"status_id,omitempty"`
}
func (PlantType) TableName() string {
	return "plant_type"
}
type JJ struct {
	//PlantId         uuid.UUID	 `gorm:"primaryKey" mapstructure:"plant_id" json:"plant_id,omitempty"`
	//PlantTypeId     uuid.UUID	 //`gorm:"column:plant_type_id" mapstructure:"plant_type_id" json:"plant_type_id"`
	////CreateDate		time.Time	 `mapstructure:"create_date" json:"create_date,omitempty"`
	////ChangeDate	    time.Time	 `mapstructure:"change_date" json:"change_date,omitempty"`
	////StatusId		uuid.UUID	 `mapstructure:"status_id" json:"status_id,omitempty"`
	////Plant1     	Plant	 	`gorm:"embedded"`//`gorm:"embeddedPrefix:p_"` // //
	Plant  model_databases.Plant `gorm:"embedded"`
	PlantType  model_databases.PlantType 	`gorm:"foreignkey:PlantTypeId; references:PlantTypeId"` //`gorm:"embeddedPrefix:pt_"` //
}
func (JJ) TableName() string {
	return "plant"
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
	var results []JJ
	//db.Table("plant").Joins("inner join plant_type on plant.plant_type_id = plant_type.plant_type_id").Scan(&results)
	plantTypeId := "caea388c-aa95-4612-9399-6e07cf42709a"
	sqlWhere := fmt.Sprintf("%s.status_id = ? AND %s.plant_type_id = ?", config.DB_PLANT, config.DB_PLANT)
	db.Debug().Where(sqlWhere, config.STATUS_ACTIVE, plantTypeId).Preload("PlantType", "status_id = ?", config.STATUS_ACTIVE).Find(&results)
	fmt.Printf("%+v\n", results)
	jj := uuid.New()
	fmt.Println()
	//helper.ConvertToJson(results)

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