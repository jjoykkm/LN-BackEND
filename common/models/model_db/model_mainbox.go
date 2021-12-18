package model_db

import (
	"fmt"
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/common/config"
	"gorm.io/gorm"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table mainbox
//-------------------------------------------------------------------------------//
//model mainbox
type Mainbox struct {
	DBCommonGet
	MainboxId      		uuid.UUID	 `json:"mainbox_id"`
	MainboxName     	string		 `json:"mainbox_name"`
	MainboxModel    	string		 `json:"mainbox_model"`
	MainboxLots     	string		 `json:"mainbox_lots"`
	StartWarranty		time.Time	 `json:"start_warranty"`
	EndWarranty			time.Time	 `json:"end_warranty"`
	FarmId				uuid.UUID	 `json:"farm_id"`
	MainboxSerialNo		string		 `json:"serial_no"`
}
// New instance mainbox
func (u *Mainbox) New() *Mainbox {
	return &Mainbox{
		DBCommonGet:      		u.DBCommonGet ,
		MainboxId:			u.MainboxId ,
		MainboxName:		u.MainboxName ,
		MainboxModel:		u.MainboxModel ,
		MainboxLots:		u.MainboxLots ,
		StartWarranty:		u.StartWarranty ,
		EndWarranty:		u.EndWarranty ,
		FarmId:				u.FarmId ,
		MainboxSerialNo:	u.MainboxSerialNo ,
	}
}

// Custom table name for GORM
func (Mainbox) TableName() string {
	return config.DB_MAINBOX
}

//-------------------------------------------------------------------------------//
//							Upsert Mainbox
//-------------------------------------------------------------------------------//
type MainboxSerialUS struct {
	MainboxName     	string		 `json:"mainbox_name"`
	MainboxSerialNo		string		 `json:"serial_no"`
	DBCommonCreateUpdate
}
func (MainboxSerialUS) TableName() string {
	return config.DB_MAINBOX
}

type MainboxUS struct {
	MainboxId		string	 `json:"mainbox_id"`
	MainboxName     string	 `json:"mainbox_name"`
	StatusId		string	 `json:"status_id"`
	DBCommonCreateUpdate
}
func (MainboxUS) TableName() string {
	return config.DB_MAINBOX
}
func (u *MainboxSerialUS) BeforeUpdate(tx *gorm.DB) (err error) {
	fmt.Println("BeforeUpdate")
	u.CreateBy = nil
	//u.CreateBy = sf_my_farm.Handler{}.Username
	//u.ChangeBy = *sf_my_farm.Handler{}.Username
	jj := "test"
	//u.CreateBy = &jj
	//u.ChangeDate = dataValue.Interface().(time.Time)
	u.ChangeBy = jj
	fmt.Printf("%+v\n", u)
	//u.UUID = uuid.New()
	//
	//if !u.IsValid() {
	//	err = errors.New("can't save invalid data")
	//}
	return
}

//func (u *MainboxUS) BeforeUpdate(tx *gorm.DB) (err error) {
//	fmt.Println("BeforeUpdate")
//	// if MainboxName changed
//	if tx.Statement.Changed("MainboxName") {
//		fmt.Println("Changed")
//		return errors.New("MainboxName not allowed to change")
//	}else {
//		fmt.Println("Not Changed")
//		tx.Statement.SetColumn("StatusId", "fe13e5d7-f467-48e8-9ce1-a997ae2c0d9f")
//	}
//
//	return nil
//}