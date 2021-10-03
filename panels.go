package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Panels struct {
	panelId 			string `gorm:"primaryKey;foreignKey:panelListId;references:panelListId"`
	panelOwnId			uuid.UUID
	panelFilename 		string
	panelName 			string
	panelNameEn 		string
	panelType 			string
	panelSide 			string
	panelAccess 		string
	panelDisplay 		string
	panelStatus 		string
	panelRestriction 	string
	panelLanguages 		string
}
//// New instance
//func (u *Panels) New() *Panels {
//	return &Panels{
//		panelId 		:      			u.panelId 		 ,
//		panelOwnId		:      			u.panelOwnId		 ,
//		panelFilename 	:      			u.panelFilename 	 ,
//		panelName 		:      			u.panelName 		 ,
//		panelNameEn 	:      			u.panelNameEn 	 ,
//		panelType 		:      			u.panelType 		 ,
//		panelSide 		:      			u.panelSide 		 ,
//		panelAccess 	:      			u.panelAccess 	 ,
//		panelDisplay 	:      			u.panelDisplay 	 ,
//		panelStatus 	:      			u.panelStatus 	 ,
//		panelRestriction:      			u.panelRestriction ,
//		panelLanguages 	:      			u.panelLanguages 	 ,
//	}
//}
func (Panels) TableName() string {
	return "panels"
}

type PanelsLists struct {
	panelListId 		string	`"json:panel_list_id"`	//`gorm:"primaryKey;foreignKey:panelId;references:panelId"`
	panelListOwnId 		string	`"json:panel_list_own_id"`
	panelListName 		string	`"json:panel_list_name"`
	panelListType 		string	`"json:panel_list_type"`
	panelListStatus 	string	`"json:panel_list_status"`
	panelListUser 		string	`"json:panel_list_user"`
	panelListSide 		string	`"json:panel_list_side"`
	panelListAccess 	string	`"json:panel_list_access"`
	panelListDisplay 	string	`"json:panel_list_display"`
}
func (PanelsLists) TableName() string {
	return "panels_lists"
}

type JoinPanels struct {
	panelId 			string 		 `"json:panel_id"`//`gorm:"primaryKey"`
	panelOwnId			uuid.UUID	 `"json:panel_own_id"`
	panelFilename 		string		 `"json:panel_filename"`
	panelName 			string		 `"json:panel_name"`
	panelNameEn 		string		 `"json:panel_name_en"`
	panelType 			string		 `"json:panel_type"`
	panelSide 			string		 `"json:panel_side"`
	panelAccess 		string		 `"json:panel_access"`
	panelDisplay 		string		 `"json:panel_display"`
	panelStatus 		string		 `"json:panel_status"`
	panelRestriction 	string		 `"json:panel_restriction"`
	panelLanguages 		string		 `"json:panel_languages"`
	PanelList			PanelsLists `gorm:"embedded"`//`gorm:"embedded;foreignKey:panelId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func main()  {
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
	sql := "SELECT * FROM panels INNER JOIN panels_lists ON panels.panel_id = panels_lists.panel_list_id"
	fmt.Println(sql)
	errs := db.Raw(sql).Scan(&joinPanel).Error
	fmt.Printf("%+v\n", joinPanel)
	if errs != nil && !errors.Is(errs, gorm.ErrRecordNotFound) {
		panic(errs)
	}
	return joinPanel
}