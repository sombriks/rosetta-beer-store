package models

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type Beer struct {
	ent.Schema
}

func (Beer) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "beer"},
	}
}

// https://github.com/ent/ent/issues/127#issuecomment-573030359
func (Beer) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").StorageKey("idbeer"),
		field.Time("creationdatebeer"),
		field.String("titlebeer"),
		field.String("descriptionbeer"),
		field.Int("idmedia"),
	}
}
