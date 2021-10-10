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
	MainboxId      	uuid.UUID	 `json:"mainbox_id,omitempty"`
	MainboxName     string		 `json:"mainbox_name,omitempty"`
	MainboxModel    string		 `json:"mainbox_model,omitempty"`
	MainboxLots     string		 `json:"mainbox_lots,omitempty"`
	StartWarranty	time.Time	 `json:"start_warranty,omitempty"`
	EndWarranty		time.Time	 `json:"end_warranty,omitempty"`
	CreateDate		time.Time	 `json:"create_date,omitempty"`
	ChangeDate	    time.Time	 `json:"change_date,omitempty"`
	StatusId		uuid.UUID	 `json:"status_id,omitempty"`
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

// Custom table name for GORM
func (Mainbox) TableName() string {
	return config.DB_MAINBOX
}
