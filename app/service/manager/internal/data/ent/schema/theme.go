package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// Theme holds the schema definition for the Theme entity.
type Theme struct {
	ent.Schema
}

// Fields of the Theme.
func (Theme) Fields() []ent.Field {
	return []ent.Field{
		field.Time("createdAt").Default(time.Now),
		field.Time("updatedAt").Default(time.Now),
	}
}

// Edges of the Theme.
func (Theme) Edges() []ent.Edge {
	return nil
}
