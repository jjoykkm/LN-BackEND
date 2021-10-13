package model_db

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/common/config"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table indicate_type
//-------------------------------------------------------------------------------//
//model indicate_type
type ColorRGB struct {
	IndColorCodeR      	string	 	 `json:"code_r"`
	IndColorCodeG      	string	 	 `json:"code_g"`
	IndColorCodeB      	string	 	 `json:"code_b"`
}
// New instance indicate_type
func (u *ColorRGB) New() *ColorRGB {
	return &ColorRGB{
		IndColorCodeR:		u.IndColorCodeR ,
		IndColorCodeG:		u.IndColorCodeG ,
		IndColorCodeB:		u.IndColorCodeB ,
	}
}

//-------------------------------------------------------------------------------//
//							Table indicate_type
//-------------------------------------------------------------------------------//
//model indicate_type
type IndicateType struct {
	IndicateTypeId      uuid.UUID	 `json:"indicate_type_id"`
	IndicateName      	string	 	 `json:"indicate_name"`
	IndicateDesc      	string	 	 `json:"indicate_desc"`
	Important	      	int			 `json:"important"`
	IndColorName      	string	 	 `json:"ind_color_name"`
	IndColorCode      	string	 	 `json:"ind_color_code"`
	ColorRGB			ColorRGB 	 `json:"color_rgb" gorm:"embedded"`
	CreateDate			time.Time	 `json:"create_date"`
	ChangeDate	    	time.Time	 `json:"change_date"`
	StatusId			uuid.UUID	 `json:"status_id"`
}
// New instance indicate_type
func (u *IndicateType) New() *IndicateType {
	return &IndicateType{
		IndicateTypeId:		u.IndicateTypeId ,
		IndicateName:		u.IndicateName ,
		IndicateDesc:		u.IndicateDesc ,
		Important:			u.Important ,
		IndColorName:		u.IndColorName ,
		IndColorCode:		u.IndColorCode ,
		ColorRGB:			u.ColorRGB ,
		CreateDate:			u.CreateDate ,
		ChangeDate:			u.ChangeDate ,
		StatusId:			u.StatusId ,
	}
}

// Custom table name for GORM
func (IndicateType) TableName() string {
	return config.DB_INDICATE_TYPE
}

