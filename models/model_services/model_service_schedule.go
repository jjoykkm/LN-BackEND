package model_services

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
)
//-------------------------------------------------------------------------------//
//							Table Manage FarmArea
//-------------------------------------------------------------------------------//
//Model
type ScheduleFarmArea struct {
	FarmId      	uuid.UUID	 			`mapstructure:"farm_id" json:"farm_id"`
	FarmAreaId      uuid.UUID	 			`mapstructure:"farm_area_id" json:"farm_area_id"`
	FarmAreaName    string		 			`mapstructure:"farm_area_name" json:"farm_area_name"`
}

// New instance
func (u *ScheduleFarmArea) New() *ScheduleFarmArea {
	return &ScheduleFarmArea{
		FarmId:				u.FarmId ,
		FarmAreaId:			u.FarmAreaId ,
		FarmAreaName:		u.FarmAreaName ,
	}
}
