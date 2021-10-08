package config

/*-------------------------------------------------------------------------------------------*/
//                                 CONSTART
/*-------------------------------------------------------------------------------------------*/
const STATUS_ACTIVE = "07de88ec-4f91-4235-9975-8b2ef1e6149e"
const ROLE_VIEW_ONLY = "a6a5e973-ce67-4d13-9387-06b5457ea602"
const CONNECTION = "postgres://ln02t:ln-0110-2@103.212.181.187/smartlife?sslmode=disable"
const DRIVER_NAME = "postgres"
const DSN = "host=103.212.181.187 user=ln02t password=ln-0110-2 dbname=smartlife port=5432 sslmode=disable TimeZone=Asia/Bangkok"

//const SERVER_HOST = "103.212.181.187:9000"
const SERVER_HOST = "0.0.0.0:5200"

/*-------------------------------------------------------------------------------------------*/
//                                 Response Type
/*-------------------------------------------------------------------------------------------*/
type ResType struct {
	Struct	string
	Map		string
	Key		string
	All		string
	List	string
}
func GetResType() ResType {
	return ResType{
		Struct: "struct" ,
		Map: 	"map" ,
		Key: 	"key" ,
		All: 	"all" ,
		List: 	"list" ,
	}
}

/*-------------------------------------------------------------------------------------------*/
//                                 Join Type
/*-------------------------------------------------------------------------------------------*/
type JoinType struct {
	Inner	string
	Left 	string
	Right	string
}
func GetJoinType() JoinType {
	return JoinType{
		Inner: 	"inner" ,
		Left: 	"left" ,
		Right: 	"right" ,
	}
}

/*-------------------------------------------------------------------------------------------*/
//                                 Status Type
/*-------------------------------------------------------------------------------------------*/
type StatusType struct {
	Active		string
	Inactive 	string
	Cancelled 	string
	Pending		string
}
func GetStatus() StatusType {
	return StatusType{
		Active: 	"07de88ec-4f91-4235-9975-8b2ef1e6149e" ,
		Inactive: 	"b918a30d-fcf8-4567-8170-9732744f5707" ,
		Cancelled: 	"c7dde653-fabf-45fe-b589-73297ab5401f" ,
		Pending: 	"fe13e5d7-f467-48e8-9ce1-a997ae2c0d9f" ,
	}
}


