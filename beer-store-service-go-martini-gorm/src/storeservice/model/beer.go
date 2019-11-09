package model

import (
	"time"
)

// Beer is our struct to represent database content
type Beer struct {
	Idbeer           int64      `json:"idbeer,omitempty" gorm:"column:idbeer;primary_key"`
	Titlebeer        string     `json:"titlebeer,omitempty" gorm:"column:titlebeer"`
	Descriptionbeer  string     `json:"descriptionbeer,omitempty" gorm:"column:descriptionbeer"`
	Creationdatebeer *time.Time `json:"creationdatebeer,omitempty" gorm:"column:creationdatebeer"`
	Idmedia          *int       `json:"idmedia,omitempty" gorm:"column:idmedia"`
}

func (Beer) TableName() string {
	return "beer"
}
