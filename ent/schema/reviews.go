package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Reviews holds the schema definition for the Reviews entity.
type Reviews struct {
	ent.Schema
}

// Fields of the Reviews.
func (Reviews) Fields() []ent.Field {
	return []ent.Field{
		field.String("file_name").
			Optional(),
		field.String("file_url").
			Optional().
			Unique(),
		field.Time("upload_time").
			Optional(),
		field.Time("create_time").
			Default(time.Now),
		field.String("reviews_title"),
	}
}

// Edges of the Reviews.
func (Reviews) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("uploaders", User.Type).
			Ref("reviews").
			Unique(),
	}
}
