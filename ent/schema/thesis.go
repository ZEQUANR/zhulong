package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Thesis holds the schema definition for the Thesis entity.
type Thesis struct {
	ent.Schema
}

// Fields of the Thesis.
func (Thesis) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Time("time").
			Default(time.Now),
		field.String("url").
			Unique(),
		field.Int("type"),
		field.Int("status"),
	}
}

// Edges of the Thesis.
func (Thesis) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("uploaders", User.Type).
			Ref("files").
			Unique(),
	}
}
