package models

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type Media struct {
	ent.Schema
}

// this is VERY opinionated
func (Media) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "media"},
	}
}

func (Media) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").StorageKey("idmedia"),
		field.Time("creationdatemedia"),
		field.Bytes("datamedia"),
		field.String("nomemedia"),
		field.String("mimemedia"),
	}
}
