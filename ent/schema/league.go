package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type League struct {
	ent.Schema
}

func (League) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.String("name"),
	}
}

func (League) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("league_certificate_type", LeagueCertificateType.Type),
	}
}
