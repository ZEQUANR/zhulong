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
		field.Int("file_state"),
		field.Time("upload_time").
			Optional(),
		field.String("chinese_title"),
		field.String("english_title"),
		field.String("authors"),
		field.String("teachers"),
		field.String("first_advance"),
		field.String("second_advance"),
		field.String("third_advance"),
		field.String("drawback"),
		field.Time("create_time").
			Default(time.Now),
	}
}

// Edges of the Thesis.
func (Thesis) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("uploaders", User.Type).
			Ref("thesis").
			Unique(),

		edge.To("examine", User.Type).
			Unique(),

		edge.To("reviews", Reviews.Type).
			Unique(),
	}
}
