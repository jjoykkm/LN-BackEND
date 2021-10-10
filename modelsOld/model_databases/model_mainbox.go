package model_databases

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
	MainboxId      	uuid.UUID	 `mapstructure:"mainbox_id" json:"mainbox_id,omitempty"`
	MainboxName     string		 `mapstructure:"mainbox_name" json:"mainbox_name,omitempty"`
	MainboxModel    string		 `mapstructure:"mainbox_model" json:"mainbox_model,omitempty"`
	MainboxLots     string		 `mapstructure:"mainbox_lots" json:"mainbox_lots,omitempty"`
	StartWarranty	time.Time	 `mapstructure:"start_warranty" json:"start_warranty,omitempty"`
	EndWarranty		time.Time	 `mapstructure:"end_warranty" json:"end_warranty,omitempty"`
	CreateDate		time.Time	 `mapstructure:"create_date" json:"create_date,omitempty"`
	ChangeDate	    time.Time	 `mapstructure:"change_date" json:"change_date,omitempty"`
	StatusId		uuid.UUID	 `mapstructure:"status_id" json:"status_id,omitempty"`
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
