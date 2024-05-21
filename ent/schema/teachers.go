package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Teachers holds the schema definition for the Teachers entity.
type Teachers struct {
	ent.Schema
}

// Fields of the Teachers.
func (Teachers) Fields() []ent.Field {
	return []ent.Field{
		field.String("avatar"),
		field.String("name"),
		field.String("college"),
		field.String("phone"),
		field.String("number"),
	}
}

// Edges of the Teachers.
func (Teachers) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users", User.Type).
			Ref("teachers").
			Unique().
			Required(),
	}
}
