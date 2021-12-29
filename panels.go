package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	//uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/common/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Panels struct {
	PanelId 			string `gorm:"primaryKey;foreignKey:panelListId;references:panelListId"`
	PanelOwnId			string
	PanelFilename 		string
	PanelName 			string
	PanelNameEn 		string
	PanelType 			string
	PanelSide 			string
	PanelAccess 		string
	PanelDisplay 		string
	PanelStatus 		string
	PanelRestriction 	string
	PanelLanguages 		string
}

type PanelsLists struct {
	PanelListId 		string	`json:"panel_list_id"`	//`gorm:"primaryKey;foreignKey:panelId;references:panelId"`
	PanelListOwnId 		string	`json:"panel_list_own_id"`
	PanelListName 		string	`json:"panel_list_name"`
	PanelListType 		string	`json:"panel_list_type"`
	PanelListStatus 	string	`json:"panel_list_status"`
	PanelListUser 		string	`json:"panel_list_user"`
	PanelListSide 		string	`json:"panel_list_side"`
	PanelListAccess 	string	`json:"panel_list_access"`
	PanelListDisplay 	string	`json:"panel_list_display"`
}
func (PanelsLists) TableName() string {
	return "panels_lists"
}

type JoinPanels struct {
	PanelId 			string 		 `json:"panel_id"`//`gorm:"primaryKey"`
	PanelOwnId			string	 `json:"panel_own_id"`
	PanelFilename 		string	 `json:"panel_filename"`
	PanelName 			string	 `json:"panel_name"`
	PanelNameEn 		string	 `json:"panel_name_en"`
	PanelType 			string	 `json:"panel_type"`
	PanelSide 			string	 `json:"panel_side"`
	PanelAccess 		string	 `json:"panel_access"`
	PanelDisplay 		string	 `json:"panel_display"`
	PanelStatus 		string	 `json:"panel_status"`
	PanelRestriction 	string	 `json:"panel_restriction"`
	PanelLanguages 		string	 `json:"panel_languages"`
	PanelList			PanelsLists `gorm:"embedded"`//`gorm:"embedded;foreignKey:panelId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func test()  {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  config.DSN,
		PreferSimpleProtocol: false, // disables implicit prepared statement usage
	}), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	http := gin.Default()
	http.POST("/panel", func(c *gin.Context) {
		result := panell(db)
		c.JSON(200, gin.H{
			"result": result,
			"total": len(result),
		})
	})

	http.Run(config.SERVER_HOST)
}

func panell(db *gorm.DB) []JoinPanels {
	var joinPanel []JoinPanels
//func panell(db *gorm.DB) []JoinPanels {
	//var joinPanel []JoinPanels
	//errs := db.Joins("panels_lists").Find(&joinPanel).Error
	//errs := db.Find(&joinPanel).Error
	////errs := db.Table("panels").Joins("LEFT JOIN panels_lists ON panels_lists.panelId = panels.panelListId").Find(&joinPanel).Error
	////errs := db.Table("panels_lists").Joins("LEFT JOIN panels_lists ON panels.panel_id = panels_lists.panel_list_id").Find(&joinPanel).Error
	sql := "SELECT * FROM panels LEFT JOIN panels_lists ON panels.panel_id = panels_lists.panel_list_id"
	fmt.Println(sql)
	errs := db.Raw(sql).Scan(&joinPanel).Error
	//fmt.Printf("%+v\n", joinPanel)
	if errs != nil && !errors.Is(errs, gorm.ErrRecordNotFound) {
		panic(errs)
	}
	return joinPanel
}