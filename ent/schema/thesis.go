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
		field.String("file_name").
			Optional(),
		field.String("file_url").
			Optional().
			Unique(),
		field.Int("file_state").
			Optional(),
		field.Time("upload_time").
			Optional(),
		field.Time("create_time").
			Default(time.Now),
		field.String("thesis_title"),
	}
}

// Edges of the Thesis.
func (Thesis) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("uploaders", User.Type).
			Ref("thesis").
			Unique(),
	}
}
