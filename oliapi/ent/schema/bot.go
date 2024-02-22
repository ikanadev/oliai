package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Bot holds the schema definition for the Bot entity.
type Bot struct {
	ent.Schema
}

// Fields of the Bot.
func (Bot) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("name"),
		field.String("greeting_message"),
		field.Text("custom_propmt"),
		field.Time("created_at").Immutable().Default(time.Now),
		field.Time("updated_at").Default(time.Now),
		field.Time("archived_at").Nillable().Default(nil),
	}
}

// Edges of the Bot.
func (Bot) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("blocks", BlockCategory.Type),
		edge.
			From("company", Company.Type).
			Ref("bots").
			Required().Immutable().Unique(),
	}
}
