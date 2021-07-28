package model_other

import uuid "github.com/jackc/pgtype/ext/gofrs-uuid"

//-------------------------------------------------------------------------------//
//				 			 	Uid
//-------------------------------------------------------------------------------//
//Model
type ModelUuid struct {
	Uuid     	uuid.UUID	`mapstructure:"uid" json:"uid,omitempty"`
}

// New instance
func (u *ModelUuid) New() *ModelUuid {
	return &ModelUuid{
		Uuid:      	u.Uuid ,
	}
}

