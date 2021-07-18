package models

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table mainbox
//-------------------------------------------------------------------------------//
//model mainbox
type Mainbox struct {
	MainboxId      	uuid.UUID	 `gorm:"mainbox_id" json:"mainbox_id,omitempty"`
	MainboxName     string		 `gorm:"mainbox_name" json:"mainbox_name,omitempty"`
	MainboxModel    string		 `gorm:"mainbox_model" json:"mainbox_model,omitempty"`
	MainboxLots     string		 `gorm:"mainbox_lots" json:"mainbox_lots,omitempty"`
	StartWarranty	time.Time	 `gorm:"start_warranty" json:"start_warranty,omitempty"`
	EndWarranty		time.Time	 `gorm:"end_warranty" json:"end_warranty,omitempty"`
	CreateDate		time.Time	 `gorm:"create_date" json:"create_date,omitempty"`
	ChangeDate	    time.Time	 `gorm:"change_date" json:"change_date,omitempty"`
	StatusId		uuid.UUID	 `gorm:"status_id" json:"status_id,omitempty"`
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
	}
}
