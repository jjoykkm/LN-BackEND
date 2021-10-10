package model_databases

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/config"
	"image"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table users
//-------------------------------------------------------------------------------//
//model users
type Users struct {
	Uid      		uuid.UUID	 	`mapstructure:"uid" json:"uid,omitempty"`
	Username     	string	 	 	`mapstructure:"username" json:"username,omitempty"`
	Password      	string	 	 	`mapstructure:"password" json:"password,omitempty"`
	FullName      	string	 	 	`mapstructure:"full_name" json:"full_name,omitempty"`
	SurName      	string	 	 	`mapstructure:"sur_name" json:"sur_name,omitempty"`
	NickName      	string	 	 	`mapstructure:"nick_name" json:"nick_name,omitempty"`
	Gender			uuid.UUID	 	`mapstructure:"gender" json:"gender,omitempty"`
	BirthDate		time.Time	 	`mapstructure:"birth_date" json:"birth_date,omitempty"`
	MobilePhone     string	 	 	`mapstructure:"mobile_phone" json:"mobile_phone,omitempty"`
	Telephone      	string	 	 	`mapstructure:"telephone" json:"telephone,omitempty"`
	Mail      		string	 	 	`mapstructure:"mail" json:"mail,omitempty"`
	Image	      	image.Image	 	`mapstructure:"image" json:"image,omitempty"`
	CreateDate		time.Time	 	`mapstructure:"create_date" json:"create_date,omitempty"`
	ChangeDate	    time.Time	 	`mapstructure:"change_date" json:"change_date,omitempty"`
	StatusId		uuid.UUID	 	`mapstructure:"status_id" json:"status_id,omitempty"`
	UserNo			string		 	`mapstructure:"user_no" json:"user_no,omitempty"`
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
	Uid      		uuid.UUID	 	`json:"uid,omitempty"`
	Username     	string	 	 	`json:"username,omitempty"`
	NickName      	string	 	 	`json:"nick_name,omitempty"`
	Image	      	string	 		`json:"image,omitempty"`
	CreateDate		time.Time	 	`json:"create_date,omitempty"`
	ChangeDate	    time.Time	 	`json:"change_date,omitempty"`
	StatusId		uuid.UUID	 	`json:"status_id,omitempty"`
	UserNo			string		 	`json:"user_no,omitempty"`
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

