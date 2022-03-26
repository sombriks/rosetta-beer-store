package models

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type Media struct {
	ent.Schema
}

func (Media) Fields() []ent.Field {
	return []ent.Field{
		field.Int("idmedia"),
		field.Time("creationdatemedia"),
		field.Bytes("datamedia"),
		field.String("nomemedia"),
		field.String("mimemedia"),
	}
}
