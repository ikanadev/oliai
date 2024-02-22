// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"oliapi/ent/blockcategory"
	"oliapi/ent/blockinfo"
	"oliapi/ent/bot"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// BlockCategoryCreate is the builder for creating a BlockCategory entity.
type BlockCategoryCreate struct {
	config
	mutation *BlockCategoryMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (bcc *BlockCategoryCreate) SetName(s string) *BlockCategoryCreate {
	bcc.mutation.SetName(s)
	return bcc
}

// SetCreatedAt sets the "created_at" field.
func (bcc *BlockCategoryCreate) SetCreatedAt(t time.Time) *BlockCategoryCreate {
	bcc.mutation.SetCreatedAt(t)
	return bcc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (bcc *BlockCategoryCreate) SetNillableCreatedAt(t *time.Time) *BlockCategoryCreate {
	if t != nil {
		bcc.SetCreatedAt(*t)
	}
	return bcc
}

// SetUpdatedAt sets the "updated_at" field.
func (bcc *BlockCategoryCreate) SetUpdatedAt(t time.Time) *BlockCategoryCreate {
	bcc.mutation.SetUpdatedAt(t)
	return bcc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (bcc *BlockCategoryCreate) SetNillableUpdatedAt(t *time.Time) *BlockCategoryCreate {
	if t != nil {
		bcc.SetUpdatedAt(*t)
	}
	return bcc
}

// SetArchivedAt sets the "archived_at" field.
func (bcc *BlockCategoryCreate) SetArchivedAt(t time.Time) *BlockCategoryCreate {
	bcc.mutation.SetArchivedAt(t)
	return bcc
}

// SetID sets the "id" field.
func (bcc *BlockCategoryCreate) SetID(u uuid.UUID) *BlockCategoryCreate {
	bcc.mutation.SetID(u)
	return bcc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (bcc *BlockCategoryCreate) SetNillableID(u *uuid.UUID) *BlockCategoryCreate {
	if u != nil {
		bcc.SetID(*u)
	}
	return bcc
}

// AddBlockIDs adds the "blocks" edge to the BlockInfo entity by IDs.
func (bcc *BlockCategoryCreate) AddBlockIDs(ids ...uuid.UUID) *BlockCategoryCreate {
	bcc.mutation.AddBlockIDs(ids...)
	return bcc
}

// AddBlocks adds the "blocks" edges to the BlockInfo entity.
func (bcc *BlockCategoryCreate) AddBlocks(b ...*BlockInfo) *BlockCategoryCreate {
	ids := make([]uuid.UUID, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bcc.AddBlockIDs(ids...)
}

// SetBotID sets the "bot" edge to the Bot entity by ID.
func (bcc *BlockCategoryCreate) SetBotID(id uuid.UUID) *BlockCategoryCreate {
	bcc.mutation.SetBotID(id)
	return bcc
}

// SetBot sets the "bot" edge to the Bot entity.
func (bcc *BlockCategoryCreate) SetBot(b *Bot) *BlockCategoryCreate {
	return bcc.SetBotID(b.ID)
}

// Mutation returns the BlockCategoryMutation object of the builder.
func (bcc *BlockCategoryCreate) Mutation() *BlockCategoryMutation {
	return bcc.mutation
}

// Save creates the BlockCategory in the database.
func (bcc *BlockCategoryCreate) Save(ctx context.Context) (*BlockCategory, error) {
	bcc.defaults()
	return withHooks(ctx, bcc.sqlSave, bcc.mutation, bcc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (bcc *BlockCategoryCreate) SaveX(ctx context.Context) *BlockCategory {
	v, err := bcc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bcc *BlockCategoryCreate) Exec(ctx context.Context) error {
	_, err := bcc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcc *BlockCategoryCreate) ExecX(ctx context.Context) {
	if err := bcc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bcc *BlockCategoryCreate) defaults() {
	if _, ok := bcc.mutation.CreatedAt(); !ok {
		v := blockcategory.DefaultCreatedAt()
		bcc.mutation.SetCreatedAt(v)
	}
	if _, ok := bcc.mutation.UpdatedAt(); !ok {
		v := blockcategory.DefaultUpdatedAt()
		bcc.mutation.SetUpdatedAt(v)
	}
	if _, ok := bcc.mutation.ID(); !ok {
		v := blockcategory.DefaultID()
		bcc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bcc *BlockCategoryCreate) check() error {
	if _, ok := bcc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "BlockCategory.name"`)}
	}
	if _, ok := bcc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "BlockCategory.created_at"`)}
	}
	if _, ok := bcc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "BlockCategory.updated_at"`)}
	}
	if _, ok := bcc.mutation.ArchivedAt(); !ok {
		return &ValidationError{Name: "archived_at", err: errors.New(`ent: missing required field "BlockCategory.archived_at"`)}
	}
	if _, ok := bcc.mutation.BotID(); !ok {
		return &ValidationError{Name: "bot", err: errors.New(`ent: missing required edge "BlockCategory.bot"`)}
	}
	return nil
}

func (bcc *BlockCategoryCreate) sqlSave(ctx context.Context) (*BlockCategory, error) {
	if err := bcc.check(); err != nil {
		return nil, err
	}
	_node, _spec := bcc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bcc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	bcc.mutation.id = &_node.ID
	bcc.mutation.done = true
	return _node, nil
}

func (bcc *BlockCategoryCreate) createSpec() (*BlockCategory, *sqlgraph.CreateSpec) {
	var (
		_node = &BlockCategory{config: bcc.config}
		_spec = sqlgraph.NewCreateSpec(blockcategory.Table, sqlgraph.NewFieldSpec(blockcategory.FieldID, field.TypeUUID))
	)
	if id, ok := bcc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := bcc.mutation.Name(); ok {
		_spec.SetField(blockcategory.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := bcc.mutation.CreatedAt(); ok {
		_spec.SetField(blockcategory.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := bcc.mutation.UpdatedAt(); ok {
		_spec.SetField(blockcategory.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := bcc.mutation.ArchivedAt(); ok {
		_spec.SetField(blockcategory.FieldArchivedAt, field.TypeTime, value)
		_node.ArchivedAt = &value
	}
	if nodes := bcc.mutation.BlocksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   blockcategory.BlocksTable,
			Columns: []string{blockcategory.BlocksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(blockinfo.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bcc.mutation.BotIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   blockcategory.BotTable,
			Columns: []string{blockcategory.BotColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(bot.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.bot_blocks = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// BlockCategoryCreateBulk is the builder for creating many BlockCategory entities in bulk.
type BlockCategoryCreateBulk struct {
	config
	err      error
	builders []*BlockCategoryCreate
}

// Save creates the BlockCategory entities in the database.
func (bccb *BlockCategoryCreateBulk) Save(ctx context.Context) ([]*BlockCategory, error) {
	if bccb.err != nil {
		return nil, bccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(bccb.builders))
	nodes := make([]*BlockCategory, len(bccb.builders))
	mutators := make([]Mutator, len(bccb.builders))
	for i := range bccb.builders {
		func(i int, root context.Context) {
			builder := bccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BlockCategoryMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, bccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, bccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bccb *BlockCategoryCreateBulk) SaveX(ctx context.Context) []*BlockCategory {
	v, err := bccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bccb *BlockCategoryCreateBulk) Exec(ctx context.Context) error {
	_, err := bccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bccb *BlockCategoryCreateBulk) ExecX(ctx context.Context) {
	if err := bccb.Exec(ctx); err != nil {
		panic(err)
	}
}
