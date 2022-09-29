package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type LeagueCertificateType struct {
	ent.Schema
}

func (LeagueCertificateType) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("league_id"),
		field.Uint64("certificate_type_id"),
	}
}

func (LeagueCertificateType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("league", League.Type).
			Ref("league_certificate_type").
			Unique().
			Required().
			Field("league_id"),
		edge.From("certificates", CertificateType.Type).
			Ref("league_certificate_type_league_type_id").
			Unique().
			Required().
			Field("certificate_type_id"),
	}
}

func (LeagueCertificateType) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: "league_certificate_type",
		},
	}
}

func (LeagueCertificateType) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("league_id", "certificate_type_id").
			Unique(),
	}
}
