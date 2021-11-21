package model_db

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/common/config"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table mainbox
//-------------------------------------------------------------------------------//
//model mainbox
type Mainbox struct {
	MainboxId      		uuid.UUID	 `json:"mainbox_id"`
	MainboxName     	string		 `json:"mainbox_name"`
	MainboxModel    	string		 `json:"mainbox_model"`
	MainboxLots     	string		 `json:"mainbox_lots"`
	StartWarranty		time.Time	 `json:"start_warranty"`
	EndWarranty			time.Time	 `json:"end_warranty"`
	CreateDate			time.Time	 `json:"create_date"`
	ChangeDate	    	time.Time	 `json:"change_date"`
	StatusId			uuid.UUID	 `json:"status_id"`
	FarmId				uuid.UUID	 `json:"farm_id"`
	MainboxSerialNo		string		 `json:"serial_no"`
}
// New instance mainbox
func (u *Mainbox) New() *Mainbox {
	return &Mainbox{
		MainboxId:			u.MainboxId ,
		MainboxName:		u.MainboxName ,
		MainboxModel:		u.MainboxModel ,
		MainboxLots:		u.MainboxLots ,
		StartWarranty:		u.StartWarranty ,
		EndWarranty:		u.EndWarranty ,
		CreateDate:			u.CreateDate ,
		ChangeDate:			u.ChangeDate ,
		StatusId:			u.StatusId ,
		FarmId:				u.FarmId ,
		MainboxSerialNo:	u.MainboxSerialNo ,
	}
}

// Custom table name for GORM
func (Mainbox) TableName() string {
	return config.DB_MAINBOX
}

//-------------------------------------------------------------------------------//
//							Request Mainbox
//-------------------------------------------------------------------------------//
type MainboxSerialUS struct {
	MainboxName     	string		 `json:"mainbox_name"`
	MainboxSerialNo		string		 `json:"serial_no"`
}
func (MainboxSerialUS) TableName() string {
	return config.DB_MAINBOX
}

//type MainboxUS struct {
//	MainboxId			uuid.UUID	 `json:"mainbox_id"`
//	MainboxName     	string		 `json:"mainbox_name"`
//	StatusId			uuid.UUID	 `json:"status_id"`
//}
type MainboxUS struct {
	MainboxId		string	 `json:"mainbox_id"`
	MainboxName     string	 `json:"mainbox_name"`
	StatusId		string	 `json:"status_id"`
}
func (MainboxUS) TableName() string {
	return config.DB_MAINBOX
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