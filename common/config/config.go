package config

/*-------------------------------------------------------------------------------------------*/
//                                 CONSTART
/*-------------------------------------------------------------------------------------------*/
const (
	DSN = "host=103.212.181.187 user=ln02t password=ln-0110-2 dbname=smartlife port=5432 sslmode=disable TimeZone=Asia/Bangkok"
	SERVER_HOST = "0.0.0.0:5200"
)
//const SERVER_HOST = "103.212.181.187:9000"
const PLANT_TYPE_ALL = "ab137f17-16fd-4214-b08d-3c43b238b542"
/*-------------------------------------------------------------------------------------------*/
//                                 Response Type
/*-------------------------------------------------------------------------------------------*/
const (
	STRUCT	string = "struct"
	MAP		string = "map"
	KEY		string = "key"
	ALL		string = "all"
	LIST	string = "list"
)
type ResType struct {
	Struct	string
	Map		string
	Key		string
	All		string
	List	string
}
func GetResType() ResType {
	return ResType{
		Struct: STRUCT,
		Map:    MAP,
		Key:    KEY,
		All:    ALL,
		List:   LIST,
	}
}

/*-------------------------------------------------------------------------------------------*/
//                                 Status Type
/*-------------------------------------------------------------------------------------------*/
const (
	ACTIVE		string = "07de88ec-4f91-4235-9975-8b2ef1e6149e"
	INACTIVE 	string = "b918a30d-fcf8-4567-8170-9732744f5707"
	CANCELLED 	string = "c7dde653-fabf-45fe-b589-73297ab5401f"
	PENDING		string = "fe13e5d7-f467-48e8-9ce1-a997ae2c0d9f"
)
type StatusType struct {
	Active		string
	Inactive 	string
	Cancelled 	string
	Pending		string
}
func GetStatus() StatusType {
	return StatusType{
		Active:    ACTIVE,
		Inactive:  INACTIVE,
		Cancelled: CANCELLED,
		Pending:   PENDING,
	}
}


