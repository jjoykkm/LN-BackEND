package model_services

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"time"
)

//Model
type Country struct {
	CountryId       uuid.UUID	 `gorm:"country_id" json:"country_id,omitempty"`
	CountryEN       string		 `gorm:"country_en" json:"country_en,omitempty"`
	CountryTH       string		 `gorm:"country_th" json:"country_th,omitempty"`
	CreateDate		time.Time	 `gorm:"create_date" json:"create_date,omitempty"`
	ChangeDate	   	time.Time	 `gorm:"change_date" json:"change_date,omitempty"`
	StatusId		uuid.UUID	 `gorm:"status_id" json:"status_id,omitempty"`
}
// New instance
func (u *Country) New() *Country {
	return &Country{
		CountryId:		u.CountryId ,
		CountryEN:		u.CountryEN ,
		CountryTH:		u.CountryTH ,
		CreateDate:		u.CreateDate ,
		ChangeDate:		u.ChangeDate ,
		StatusId:		u.StatusId ,
	}
}
