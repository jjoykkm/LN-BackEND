package helper_gorm

import (
	"fmt"
	"github.com/jjoykkm/ln-backend/common/models/model_db"
)

type Servicer interface {
	//GetFarmAreaByFarm(status, farmId, resultType string) ([]model_databases.FarmArea, []string)
}

type Service struct {
	repo Repositorier
}

func NewService(repo Repositorier) Servicer {
	return &Service{
		repo:  repo,
	}
}
//-------------------------------------------------------------------------------//
//						Function Common Database (Create)
//-------------------------------------------------------------------------------//
func BeforeCreate(u *model_db.DBCommonCreateUpdate, userNo string) {
	fmt.Println("BeforeCreate")
	u.CreateBy = nil
	u.ChangeBy = userNo
	fmt.Printf("%+v\n", u)
	//u.UUID = uuid.New()
	//
}

//-------------------------------------------------------------------------------//
//						Function Common Database (Update)
//-------------------------------------------------------------------------------//
func BeforeUpdate(u *model_db.DBCommonCreateUpdate, userNo string) {
	fmt.Println("BeforeUpdate")
	u.CreateBy = &userNo
	u.ChangeBy = userNo
	fmt.Printf("%+v\n", u)
	//u.UUID = uuid.New()
	//
	//if !u.IsValid() {
	//	err := errors.New("can't save invalid data")
	//	return err
	//}
	//return
}