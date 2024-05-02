package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Administrators holds the schema definition for the Administrators entity.
type Administrators struct {
	ent.Schema
}

// Fields of the Administrators.
func (Administrators) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("college"),
		field.String("phone"),
		field.String("identity"),
	}
}

// Edges of the Administrators.
func (Administrators) Edges() []ent.Edge {
	return nil
}
