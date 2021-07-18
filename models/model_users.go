package models

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"image"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table users
//-------------------------------------------------------------------------------//
//model users
type Users struct {
	Uid      		uuid.UUID	 	`gorm:"uid" json:"uid,omitempty"`
	Username     	string	 	 	`gorm:"username" json:"username,omitempty"`
	Password      	string	 	 	`gorm:"password" json:"password,omitempty"`
	FullName      	string	 	 	`gorm:"full_name" json:"full_name,omitempty"`
	SurName      	string	 	 	`gorm:"sur_name" json:"sur_name,omitempty"`
	NickName      	string	 	 	`gorm:"nick_name" json:"nick_name,omitempty"`
	Gender			uuid.UUID	 	`gorm:"gender" json:"gender,omitempty"`
	BirthDate		time.Time	 	`gorm:"birth_date" json:"birth_date,omitempty"`
	MobilePhone     string	 	 	`gorm:"mobile_phone" json:"mobile_phone,omitempty"`
	Telephone      	string	 	 	`gorm:"telephone" json:"telephone,omitempty"`
	Mail      		string	 	 	`gorm:"mail" json:"mail,omitempty"`
	Image	      	image.Image	 	`gorm:"image" json:"image,omitempty"`
	CreateDate		time.Time	 	`gorm:"create_date" json:"create_date,omitempty"`
	ChangeDate	    time.Time	 	`gorm:"change_date" json:"change_date,omitempty"`
	StatusId		uuid.UUID	 	`gorm:"status_id" json:"status_id,omitempty"`
	UserNo			string		 	`gorm:"user_no" json:"user_no,omitempty"`
}
// New instance users
func (u *Users) New() *Users {
	return &Users{
		Uid:       		u.Uid ,
		Username:       u.Username ,
		Password:       u.Password ,
		FullName:       u.FullName ,
		SurName:       	u.SurName ,
		NickName:       u.NickName ,
		Gender:       	u.Gender ,
		BirthDate:      u.BirthDate ,
		MobilePhone:    u.MobilePhone ,
		Telephone:      u.Telephone ,
		Mail:       	u.Mail ,
		Image:       	u.Image ,
		CreateDate:     u.CreateDate ,
		ChangeDate:     u.ChangeDate ,
		StatusId:       u.StatusId ,
		UserNo:       	u.UserNo ,

	}
}
