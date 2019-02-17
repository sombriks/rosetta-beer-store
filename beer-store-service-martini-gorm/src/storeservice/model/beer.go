package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Beer is our struct to represent database content
type Beer struct {
	gorm.Model
	Idbeer           int64      `json:"idbeer,omitempty" gorm:"column:idbeer"`
	Titlebeer        string     `json:"titlebeer,omitempty" gorm:"column:titlebeer"`
	Descriptionbeer  string     `json:"descriptionbeer,omitempty" gorm:"column:descriptionbeer"`
	Creationdatebeer *time.Time `json:"creationdatebeer,omitempty" gorm:"column:creationdatebeer"`
	Idmedia          *int       `json:"idmedia,omitempty" gorm:"column:idmedia"`
}
