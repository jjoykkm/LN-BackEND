package model_controllers

import (
	"fmt"
	"github.com/jjoykkm/ln-backend/common/models/model_db"
)

//-------------------------------------------------------------------------------//
//						Function Common Database (Create)
//-------------------------------------------------------------------------------//
func BeforeCreate(u model_db.DBCommonCreateUpdate, user string) {
	fmt.Println("BeforeCreate")
	u.CreateBy = nil
	u.ChangeBy = user
	fmt.Printf("%+v\n", u)
	//u.UUID = uuid.New()
	//
}

//-------------------------------------------------------------------------------//
//						Function Common Database (Update)
//-------------------------------------------------------------------------------//
func BeforeUpdate(u model_db.DBCommonCreateUpdate, user string) {
	fmt.Println("BeforeUpdate")
	u.CreateBy = &user
	u.ChangeBy = user
	fmt.Printf("%+v\n", u)
	//u.UUID = uuid.New()
	//
	//if !u.IsValid() {
	//	err := errors.New("can't save invalid data")
	//	return err
	//}
	//return
}

func Greeting(name string)  {
	fmt.Println(name)
}