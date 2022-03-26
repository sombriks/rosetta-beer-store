package models

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type Beer struct {
	ent.Schema
}

func (Beer) Fields() []ent.Field {
	return []ent.Field{
		field.Int("idbeer"),
		field.Time("creationdatebeer"),
		field.String("titlebeer"),
		field.String("descriptionbeer"),
		field.Int("idmedia"),
	}
}
