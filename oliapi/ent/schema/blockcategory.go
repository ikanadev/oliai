package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// BlockCategory holds the schema definition for the BlockCategory entity.
type BlockCategory struct {
	ent.Schema
}

// Fields of the BlockCategory.
func (BlockCategory) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("name"),
		field.Time("created_at").Immutable().Default(time.Now),
		field.Time("updated_at").Default(time.Now),
		field.Time("archived_at").Nillable().Default(nil),
	}
}

// Edges of the BlockCategory.
func (BlockCategory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("blocks", BlockInfo.Type),
		edge.
			From("bot", Bot.Type).
			Ref("blocks").
			Field("bot_id").
			Required().Immutable().Unique(),
	}
}
