package main

import (
	"errors"
	"fmt"
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Farm struct {
	FarmId      	uuid.UUID	 `mapstructure:"farm_id" json:"farm_id,omitempty"`
	FarmName    	string		 `mapstructure:"farm_name" json:"farm_name,omitempty"`
	FarmDesc    	string		 `mapstructure:"farm_desc" json:"farm_desc,omitempty"`
	//CreateDate		time.Time	 `mapstructure:"create_date" json:"create_date,omitempty"`
	//ChangeDate	    time.Time	 `mapstructure:"change_date" json:"change_date,omitempty"`
	//StatusId		uuid.UUID	 `mapstructure:"status_id" json:"status_id,omitempty"`
}
func (Farm) TableName() string {
	return "farm"
}

func main()  {
	var fm Farm
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  config.DSN,
		PreferSimpleProtocol: false, // disables implicit prepared statement usage
	}), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	//add data to map
	//jj := map[string]interface{}{}
	//result := db.Model(&Farm{}).First(&jj)
	result := db.First(&fm)
	//fmt.Printf("%+v\n", jj)
	fmt.Printf("%+v\n", fm)
	fmt.Println(result.RowsAffected) // returns count of records found
	fmt.Println(result.Error) // returns error or nil
	fmt.Println(gorm.ErrRecordNotFound) // returns error or nil

	errors.Is(result.Error, gorm.ErrRecordNotFound)
}