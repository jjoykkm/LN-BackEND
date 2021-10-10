package model_db

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/common/config"
	"image"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table users
//-------------------------------------------------------------------------------//
//model users
type Users struct {
	Uid      		uuid.UUID	 	`json:"uid,omitempty"`
	Username     	string	 	 	`json:"username,omitempty"`
	Password      	string	 	 	`json:"password,omitempty"`
	FullName      	string	 	 	`json:"full_name,omitempty"`
	SurName      	string	 	 	`json:"sur_name,omitempty"`
	NickName      	string	 	 	`json:"nick_name,omitempty"`
	Gender			uuid.UUID	 	`json:"gender,omitempty"`
	BirthDate		time.Time	 	`json:"birth_date,omitempty"`
	MobilePhone     string	 	 	`json:"mobile_phone,omitempty"`
	Telephone      	string	 	 	`json:"telephone,omitempty"`
	Mail      		string	 	 	`json:"mail,omitempty"`
	Image	      	image.Image	 	`json:"image,omitempty"`
	CreateDate		time.Time	 	`json:"create_date,omitempty"`
	ChangeDate	    time.Time	 	`json:"change_date,omitempty"`
	StatusId		uuid.UUID	 	`json:"status_id,omitempty"`
	UserNo			string		 	`json:"user_no,omitempty"`
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

// Custom table name for GORM
func (Users) TableName() string {
	return config.DB_USERS
}

//-------------------------------------------------------------------------------//
//							Table users only data
//-------------------------------------------------------------------------------//
//model users
type UsersShort struct {
	Uid      		uuid.UUID	 	`json:"-"`
	Username     	string	 	 	`json:"username,omitempty"`
	NickName      	string	 	 	`json:"nick_name,omitempty" gorm:"column:nickname"`
	Image	      	string		 	`json:"image,omitempty"`
	CreateDate		time.Time	 	`json:"create_date,omitempty" gorm:"column:createdate"`
	ChangeDate	    time.Time	 	`json:"change_date,omitempty" gorm:"column:changedate"`
	StatusId		uuid.UUID	 	`json:"status_id,omitempty" gorm:"column:statusid"`
	UserNo			string		 	`json:"user_no,omitempty" gorm:"column:userno"`
}
// New instance users
func (u *UsersShort) New() *UsersShort {
	return &UsersShort{
		Uid:       		u.Uid ,
		Username:       u.Username ,
		NickName:       u.NickName ,
		Image:       	u.Image ,
		CreateDate:     u.CreateDate ,
		ChangeDate:     u.ChangeDate ,
		StatusId:       u.StatusId ,
		UserNo:       	u.UserNo ,
	}
}

// Custom table name for GORM
func (UsersShort) TableName() string {
	return config.DB_USERS
}

