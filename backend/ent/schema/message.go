package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Message holds the schema definition for the Message entity.
type Message struct {
	ent.Schema
}

// Fields of the Message.
func (Message) Fields() []ent.Field {
	return []ent.Field{
		field.String("content").NotEmpty(),
		field.String("author_name").NotEmpty(),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the Message.
func (Message) Edges() []ent.Edge {
	return nil
}
