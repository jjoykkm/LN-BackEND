package main

import (
	"LN-BackEND/models/model_databases"
	"fmt"
	//"time"
)
type test struct {
	jjoy         string
	eiei         string
}

func main() {



	//db, err := gorm.Open(postgres.New(postgres.Config{
	//	DSN: "host=103.212.181.187 user=ln02t password=ln-0110-2 dbname=smartlife port=5432 sslmode=disable TimeZone=Asia/Bangkok",
	//	PreferSimpleProtocol: false, // disables implicit prepared statement usage
	//}), &gorm.Config{})
	//if err != nil {
	//	panic(err.Error())
	//}
	//fmt.Println(db)
	//user := test{jjoy: "Jinzhu", eiei: "18"}
	//db.Select("jjoy", "eiei").Create(&user)

	var jj model_databases.Country
	jj.CountryEN = "jjjj"
	fmt.Println(jj.CountryEN)



}