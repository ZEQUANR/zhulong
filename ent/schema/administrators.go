package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Administrators holds the schema definition for the Administrators entity.
type Administrators struct {
	ent.Schema
}

// Fields of the Administrators.
func (Administrators) Fields() []ent.Field {
	return []ent.Field{
		field.String("avatar"),
		field.String("name"),
		field.String("college"),
		field.String("phone"),
		field.String("number"),
	}
}

// Edges of the Administrators.
func (Administrators) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users", User.Type).
			Ref("administrators").
			Unique().
			Required(),
	}
}
