package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// React holds the schema definition for the React entity.
type React struct {
	ent.Schema
}

// Fields of the React.
func (React) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Immutable().
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.String("react_type").
			Default("").Optional(),
		field.String("post_type").
			Default("").Optional(),
	}
}

// Edges of the React.
func (React) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("posts", Post.Type).Ref("reacts"),
		edge.From("comments", Comment.Type).Ref("reacts"),
		edge.From("reacted_user", Comment.Type).Ref("reacts"),
	}
}
