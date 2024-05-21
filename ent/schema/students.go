package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Students holds the schema definition for the Students entity.
type Students struct {
	ent.Schema
}

// Fields of the Students.
func (Students) Fields() []ent.Field {
	return []ent.Field{
		field.String("avatar"),
		field.String("name"),
		field.String("college"),
		field.String("phone"),
		field.String("major"),
		field.String("class"),
		field.String("number"),
	}
}

// Edges of the Students.
func (Students) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users", User.Type).
			Ref("students").
			Unique().
			Required(),
	}
}
