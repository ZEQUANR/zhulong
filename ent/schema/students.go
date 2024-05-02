package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Students holds the schema definition for the Students entity.
type Students struct {
	ent.Schema
}

// Fields of the Students.
func (Students) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("college"),
		field.String("phone"),
		field.String("subject"),
		field.String("class"),
		field.String("identity"),
	}
}

// Edges of the Students.
func (Students) Edges() []ent.Edge {
	return nil
}
