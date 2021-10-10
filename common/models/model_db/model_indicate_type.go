package model_db

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jjoykkm/ln-backend/config"
	"time"
)

//-------------------------------------------------------------------------------//
//							Table indicate_type
//-------------------------------------------------------------------------------//
//model indicate_type
type IndicateType struct {
	IndicateTypeId      uuid.UUID	 `mapstructure:"indicate_type_id" json:"indicate_type_id,omitempty"`
	IndicateName      	string	 	 `mapstructure:"indicate_name" json:"indicate_name,omitempty"`
	IndicateDesc      	string	 	 `mapstructure:"indicate_desc" json:"indicate_desc,omitempty"`
	Important	      	int			 `mapstructure:"important" json:"important,omitempty"`
	IndColorName      	string	 	 `mapstructure:"ind_color_name" json:"ind_color_name,omitempty"`
	IndColorCode      	string	 	 `mapstructure:"ind_color_code" json:"ind_color_code,omitempty"`
	IndColorCodeR      	string	 	 `mapstructure:"ind_color_code_r" json:"ind_color_code_r,omitempty"`
	IndColorCodeG      	string	 	 `mapstructure:"ind_color_code_g" json:"ind_color_code_g,omitempty"`
	IndColorCodeB      	string	 	 `mapstructure:"ind_color_code_b" json:"ind_color_code_b,omitempty"`
	CreateDate			time.Time	 `mapstructure:"create_date" json:"create_date,omitempty"`
	ChangeDate	    	time.Time	 `mapstructure:"change_date" json:"change_date,omitempty"`
	StatusId			uuid.UUID	 `mapstructure:"status_id" json:"status_id,omitempty"`
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

