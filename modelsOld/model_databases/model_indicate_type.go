package model_databases

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/common/config"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table indicate_type
//-------------------------------------------------------------------------------//
//model indicate_type
type IndicateType struct {
	IndicateTypeId      uuid.UUID	 `json:"indicate_type_id,omitempty"`
	IndicateName      	string	 	 `json:"indicate_name,omitempty"`
	IndicateDesc      	string	 	 `json:"indicate_desc,omitempty"`
	Important	      	int			 `json:"important,omitempty"`
	IndColorName      	string	 	 `json:"ind_color_name,omitempty"`
	IndColorCode      	string	 	 `json:"ind_color_code,omitempty"`
	IndColorCodeR      	string	 	 `json:"ind_color_code_r,omitempty"`
	IndColorCodeG      	string	 	 `json:"ind_color_code_g,omitempty"`
	IndColorCodeB      	string	 	 `json:"ind_color_code_b,omitempty"`
	CreateDate			time.Time	 `json:"create_date,omitempty"`
	ChangeDate	    	time.Time	 `json:"change_date,omitempty"`
	StatusId			uuid.UUID	 `json:"status_id,omitempty"`
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
		IndColorCodeR:		u.IndColorCodeR ,
		IndColorCodeG:		u.IndColorCodeG ,
		IndColorCodeB:		u.IndColorCodeB ,
		CreateDate:			u.CreateDate ,
		ChangeDate:			u.ChangeDate ,
		StatusId:			u.StatusId ,
	}
}

// Custom table name for GORM
func (IndicateType) TableName() string {
	return config.DB_INDICATE_TYPE
}

