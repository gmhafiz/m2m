package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type CertificateType struct {
	ent.Schema
}

func (CertificateType) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.String("name"),
	}
}

func (CertificateType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("league_certificate_type_league_type_id", LeagueCertificateType.Type).
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
	}
}

func (CertificateType) Annotations() []schema.Annotation {
	return nil
}
