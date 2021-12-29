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
	OPENED		string = "081e9cf1-20c5-428b-9828-3fe8f4e40cd8"
	CLOSED		string = "caee866d-1fa4-4eac-a15d-89cca9e73baa"
	UNKNOWN		string = "f2531663-7cd5-4644-9ef0-44713a26a9bb"
)
type SensorStatusConst struct {
	Opened		string
	Closed		string
	Unknown		string
}
func GetSensorStatus() SensorStatusConst {
	return SensorStatusConst{
		Opened: 	OPENED,
		Closed:     CLOSED,
		Unknown:    UNKNOWN,
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
type StatusTypeConst struct {
	Active		string
	Inactive 	string
	Cancelled 	string
	Pending		string
}
func GetStatus() StatusTypeConst {
	return StatusTypeConst{
		Active:    ACTIVE,
		Inactive:  INACTIVE,
		Cancelled: CANCELLED,
		Pending:   PENDING,
	}
}

/*-------------------------------------------------------------------------------------------*/
//                                 Role Type
/*-------------------------------------------------------------------------------------------*/
const (
	OWNER	string = "cd30c998-7390-4108-b6bc-ab0fb5dbc698"
	VIEW 	string = "a6a5e973-ce67-4d13-9387-06b5457ea602"
	CHANGE 	string = "8496b7e8-e23b-45bc-a6c0-d604a7119cc2"
)
type RoleTypeConst struct {
	Owner	string
	View 	string
	Change 	string
}
func GetRole() RoleTypeConst {
	return RoleTypeConst{
		Owner:  	OWNER,
		View:  		VIEW,
		Change: 	CHANGE,
	}
}



