package model_services

import (
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	//"time"
)

//-------------------------------------------------------------------------------//
//				 	    	FarmList
//-------------------------------------------------------------------------------//
//Model
type DashboardFarmList struct {
	Uid      	uuid.UUID	 `mapstructure:"uid" json:"uid"`
	FarmId      uuid.UUID	 `mapstructure:"farm_id" json:"farm_id"`
	FarmName    string		 `mapstructure:"farm_name" json:"farm_name"`
	FarmDesc    string		 `mapstructure:"farm_desc" json:"farm_desc"`
}
// New instance
func (u *DashboardFarmList) New() *DashboardFarmList {
	return &DashboardFarmList{
		Uid:		u.Uid ,
		FarmId:		u.FarmId ,
		FarmName:	u.FarmName ,
		FarmDesc:	u.FarmDesc ,
	}
}

//-------------------------------------------------------------------------------//
//				 	    	FarmList
//-------------------------------------------------------------------------------//
//Model
type DashboardFarmAreaList struct {
	FarmId				uuid.UUID	 `mapstructure:"farm_id" json:"farm_id"`
	FarmAreaId			uuid.UUID	 `mapstructure:"farm_area_id" json:"farm_area_id"`
	FarmAreaName		string	 	 `mapstructure:"farm_area_name" json:"farm_area_name"`
	FormulaPlantId		string	 	 `mapstructure:"formula_plant_id" json:"formula_plant_id"`
	FormulaName			string	 	 `mapstructure:"formula_name" json:"formula_name"`
	FormulaDesc			string	 	 `mapstructure:"formula_desc" json:"formula_desc"`

}
// New instance
func (u *DashboardFarmAreaList) New() *DashboardFarmAreaList {
	return &DashboardFarmAreaList{
		FarmId:				u.FarmId ,
		FarmAreaId:			u.FarmAreaId ,
		FarmAreaName:		u.FarmAreaName ,
		FormulaPlantId:		u.FormulaPlantId ,
		FormulaName:		u.FormulaName	 ,
		FormulaDesc:		u.FormulaDesc	 ,
	}
}