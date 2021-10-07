package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jjoykkm/ln-backend/config"
	"github.com/jjoykkm/ln-backend/models/model_databases"
	"github.com/jjoykkm/ln-backend/smartfarm/sf_formula_plant"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//type Test struct {
//	Plant  model_databases.Plant `gorm:"embedded"`
//	PlantType  model_databases.PlantType `gorm:"foreignkey:PlantTypeId; references:PlantTypeId"`
//}

type JJ struct {
	//PlantId         uuid.UUID	 `gorm:"primaryKey" mapstructure:"plant_id" json:"plant_id,omitempty"`
	//PlantTypeId     uuid.UUID	 //`gorm:"column:plant_type_id" mapstructure:"plant_type_id" json:"plant_type_id"`
	////CreateDate		time.Time	 `mapstructure:"create_date" json:"create_date,omitempty"`
	////ChangeDate	    time.Time	 `mapstructure:"change_date" json:"change_date,omitempty"`
	////StatusId		uuid.UUID	 `mapstructure:"status_id" json:"status_id,omitempty"`
	////Plant1     	Plant	 	`gorm:"embedded"`//`gorm:"embeddedPrefix:p_"` // //
	Plant  model_databases.Plant `gorm:"embedded"`
	PlantType  model_databases.PlantType `gorm:"foreignkey:PlantTypeId; references:PlantTypeId"` //`gorm:"embeddedPrefix:pt_"` //
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
	http := gin.Default()
	http.Use(cors.Default())

	var result []sf_formula_plant.JoinPlantAndFormulaPlant
	//var result []JJ
	//db.Table("plant").Joins("inner join plant_type on plant.plant_type_id = plant_type.plant_type_id").Scan(&results)
	//plantTypeId := "caea388c-aa95-4612-9399-6e07cf42709a"
	//sqlWhere := fmt.Sprintf("%s.status_id = ? AND %s.plant_type_id = ?", config.DB_PLANT, config.DB_PLANT)
	//db.Joins("Company").Joins("Manager").Joins("Account").Find(&users, "users.id IN ?",
	//db.Debug().Joins("FormulaPlant").Find(&result, "plant.plant_id = ?","bde0c9e9-8dea-485f-8ec1-cb94dfd4b14e")
	//db.Debug().Where(sqlWhere, config.STATUS_ACTIVE, plantTypeId).Preload("PlantType", "status_id = ?", config.STATUS_ACTIVE).Find(&results)
	//db.Debug().Preload("Plant").Preload("Province").Preload("Country").Find(&result)
	//db.Debug().Preload(clause.Associations).Find(&result)
	//db.Debug().Joins("inner join plant_type on plant_type.plant_type_id = plant.plant_type_id").Find(&result)
	//db.Debug().Joins("PlantType").Find(&result)
	//db.Debug().Preload("PlantType").Find(&result)

	//db.Debug().Preload("Plant", func(db *gorm.DB) *gorm.DB {
	//		return db.Preload("PlantType")
	//	}).Preload("Province").Preload("Country").Limit(5).Offset(20).Find(&result)
	plantId:="bde0c9e9-8dea-485f-8ec1-cb94dfd4b14e"
	var sqlWhere string
	//// Generate condition when get plant
	//status:=config.STATUS_ACTIVE
	sqlWhere = fmt.Sprintf("status_id = '%s'",config.STATUS_ACTIVE)
	if plantId != "" {
		sqlWhere = sqlWhere + fmt.Sprintf(" AND plant_id = '%s'", plantId)
	}
	//db.Debug().Preload("Plant", func(db *gorm.DB) *gorm.DB {
	//	return db.Where("status_id = ?", config.STATUS_ACTIVE).Preload("PlantType", "status_id = ?", config.STATUS_ACTIVE)
	//}).Preload("Province", "status_id = ?", config.STATUS_ACTIVE).Preload("Country", "status_id = ?", config.STATUS_ACTIVE).Limit(5).Offset(20).Where(sqlWhere, status, plantId).Find(&result)
	db.Debug().Limit(5).Offset(0).Where(sqlWhere).Preload("Plant", func(db *gorm.DB) *gorm.DB {
		return db.Where("status_id = ?", config.STATUS_ACTIVE).Preload("PlantType", "status_id = ?", config.STATUS_ACTIVE)
	}).Preload("Province", "status_id = ?", config.STATUS_ACTIVE).Preload("Country", "status_id = ?", config.STATUS_ACTIVE).Find(&result)


	//.Joins("inner join plant_type on plant.plant_type_id = plant_type.plant_type_id")
	//.Preload("Plant", func(db *gorm.DB) *gorm.DB {
	//	return db.Limit(5).Offset(20)
	//})
	//.Where("status_id = ? AND plant_id = ? ", config.STATUS_ACTIVE, "bde0c9e9-8dea-485f-8ec1-cb94dfd4b14e")
	//.Preload("PlantType")
	//.Joins("Province")
	//.Preload("JoinPlantAndFormulaPlant.FormulaPlant.Country.CountryId")
	//db.Debug().Set("gorm:auto_preload", true).Find(&result)

//db.Debug().Where("status_id = ? AND plant_id = ? ", config.STATUS_ACTIVE, "bde0c9e9-8dea-485f-8ec1-cb94dfd4b14e").Preload("Plant").Limit(20).Preload("PlantType").Find(&result)
	//db.Debug().Where("status_id = ? AND plant_id = ? ", config.STATUS_ACTIVE, "bde0c9e9-8dea-485f-8ec1-cb94dfd4b14e").Preload("FormulaPlant").Find(&result)

	fmt.Println("-----------------------------")
	//fmt.Printf("%+v\n", result)
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
	http.POST("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"01": result,
		})
	})
	http.Run(config.SERVER_HOST)
}