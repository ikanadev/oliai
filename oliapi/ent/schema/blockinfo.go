package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// BlockInfo holds the schema definition for the BlockInfo entity.
type BlockInfo struct {
	ent.Schema
}

// Fields of the BlockInfo.
func (BlockInfo) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.Text("content"),
		field.Time("created_at").Immutable().Default(time.Now),
		field.Time("updated_at").Default(time.Now),
		field.Time("archived_at").Nillable().Default(nil),
	}
}

// Edges of the BlockInfo.
func (BlockInfo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			From("category", BlockCategory.Type).
			Ref("blocks").
			Field("category_id").
			Required().Immutable().Unique(),
	}
}
