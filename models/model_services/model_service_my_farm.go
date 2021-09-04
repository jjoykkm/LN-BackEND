package model_services

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
)
//-------------------------------------------------------------------------------//
//							Table Farm Overview
//-------------------------------------------------------------------------------//
//Model
type MyFarmOverviewFarm struct {
	FarmId      	uuid.UUID	 `mapstructure:"farm_id" json:"farm_id"`
	FarmName    	string		 `mapstructure:"farm_name" json:"farm_name"`
	FarmDesc    	string		 `mapstructure:"farm_desc" json:"farm_desc"`
	MainboxCount	int			 `mapstructure:"mainbox_count" json:"mainbox_count"`
	FarmAreaCount	int			 `mapstructure:"farm_area_count" json:"farm_area_count"`
}

// New instance
func (u *MyFarmOverviewFarm) New() *MyFarmOverviewFarm {
	return &MyFarmOverviewFarm{
		FarmId:			u.FarmId ,
		FarmName:		u.FarmName ,
		FarmDesc:		u.FarmDesc ,
		MainboxCount:	u.MainboxCount ,
		FarmAreaCount:	u.FarmAreaCount ,
	}
}