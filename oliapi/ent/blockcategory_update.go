// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"oliapi/ent/blockcategory"
	"oliapi/ent/blockinfo"
	"oliapi/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// BlockCategoryUpdate is the builder for updating BlockCategory entities.
type BlockCategoryUpdate struct {
	config
	hooks    []Hook
	mutation *BlockCategoryMutation
}

// Where appends a list predicates to the BlockCategoryUpdate builder.
func (bcu *BlockCategoryUpdate) Where(ps ...predicate.BlockCategory) *BlockCategoryUpdate {
	bcu.mutation.Where(ps...)
	return bcu
}

// SetName sets the "name" field.
func (bcu *BlockCategoryUpdate) SetName(s string) *BlockCategoryUpdate {
	bcu.mutation.SetName(s)
	return bcu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (bcu *BlockCategoryUpdate) SetNillableName(s *string) *BlockCategoryUpdate {
	if s != nil {
		bcu.SetName(*s)
	}
	return bcu
}

// SetUpdatedAt sets the "updated_at" field.
func (bcu *BlockCategoryUpdate) SetUpdatedAt(t time.Time) *BlockCategoryUpdate {
	bcu.mutation.SetUpdatedAt(t)
	return bcu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (bcu *BlockCategoryUpdate) SetNillableUpdatedAt(t *time.Time) *BlockCategoryUpdate {
	if t != nil {
		bcu.SetUpdatedAt(*t)
	}
	return bcu
}

// SetArchivedAt sets the "archived_at" field.
func (bcu *BlockCategoryUpdate) SetArchivedAt(t time.Time) *BlockCategoryUpdate {
	bcu.mutation.SetArchivedAt(t)
	return bcu
}

// SetNillableArchivedAt sets the "archived_at" field if the given value is not nil.
func (bcu *BlockCategoryUpdate) SetNillableArchivedAt(t *time.Time) *BlockCategoryUpdate {
	if t != nil {
		bcu.SetArchivedAt(*t)
	}
	return bcu
}

// AddBlockIDs adds the "blocks" edge to the BlockInfo entity by IDs.
func (bcu *BlockCategoryUpdate) AddBlockIDs(ids ...uuid.UUID) *BlockCategoryUpdate {
	bcu.mutation.AddBlockIDs(ids...)
	return bcu
}

// AddBlocks adds the "blocks" edges to the BlockInfo entity.
func (bcu *BlockCategoryUpdate) AddBlocks(b ...*BlockInfo) *BlockCategoryUpdate {
	ids := make([]uuid.UUID, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bcu.AddBlockIDs(ids...)
}

// Mutation returns the BlockCategoryMutation object of the builder.
func (bcu *BlockCategoryUpdate) Mutation() *BlockCategoryMutation {
	return bcu.mutation
}

// ClearBlocks clears all "blocks" edges to the BlockInfo entity.
func (bcu *BlockCategoryUpdate) ClearBlocks() *BlockCategoryUpdate {
	bcu.mutation.ClearBlocks()
	return bcu
}

// RemoveBlockIDs removes the "blocks" edge to BlockInfo entities by IDs.
func (bcu *BlockCategoryUpdate) RemoveBlockIDs(ids ...uuid.UUID) *BlockCategoryUpdate {
	bcu.mutation.RemoveBlockIDs(ids...)
	return bcu
}

// RemoveBlocks removes "blocks" edges to BlockInfo entities.
func (bcu *BlockCategoryUpdate) RemoveBlocks(b ...*BlockInfo) *BlockCategoryUpdate {
	ids := make([]uuid.UUID, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bcu.RemoveBlockIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bcu *BlockCategoryUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, bcu.sqlSave, bcu.mutation, bcu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bcu *BlockCategoryUpdate) SaveX(ctx context.Context) int {
	affected, err := bcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bcu *BlockCategoryUpdate) Exec(ctx context.Context) error {
	_, err := bcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcu *BlockCategoryUpdate) ExecX(ctx context.Context) {
	if err := bcu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bcu *BlockCategoryUpdate) check() error {
	if _, ok := bcu.mutation.BotID(); bcu.mutation.BotCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "BlockCategory.bot"`)
	}
	return nil
}

func (bcu *BlockCategoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := bcu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(blockcategory.Table, blockcategory.Columns, sqlgraph.NewFieldSpec(blockcategory.FieldID, field.TypeUUID))
	if ps := bcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bcu.mutation.Name(); ok {
		_spec.SetField(blockcategory.FieldName, field.TypeString, value)
	}
	if value, ok := bcu.mutation.UpdatedAt(); ok {
		_spec.SetField(blockcategory.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := bcu.mutation.ArchivedAt(); ok {
		_spec.SetField(blockcategory.FieldArchivedAt, field.TypeTime, value)
	}
	if bcu.mutation.BlocksCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bcu.mutation.RemovedBlocksIDs(); len(nodes) > 0 && !bcu.mutation.BlocksCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bcu.mutation.BlocksIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{blockcategory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	bcu.mutation.done = true
	return n, nil
}

// BlockCategoryUpdateOne is the builder for updating a single BlockCategory entity.
type BlockCategoryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BlockCategoryMutation
}

// SetName sets the "name" field.
func (bcuo *BlockCategoryUpdateOne) SetName(s string) *BlockCategoryUpdateOne {
	bcuo.mutation.SetName(s)
	return bcuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (bcuo *BlockCategoryUpdateOne) SetNillableName(s *string) *BlockCategoryUpdateOne {
	if s != nil {
		bcuo.SetName(*s)
	}
	return bcuo
}

// SetUpdatedAt sets the "updated_at" field.
func (bcuo *BlockCategoryUpdateOne) SetUpdatedAt(t time.Time) *BlockCategoryUpdateOne {
	bcuo.mutation.SetUpdatedAt(t)
	return bcuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (bcuo *BlockCategoryUpdateOne) SetNillableUpdatedAt(t *time.Time) *BlockCategoryUpdateOne {
	if t != nil {
		bcuo.SetUpdatedAt(*t)
	}
	return bcuo
}

// SetArchivedAt sets the "archived_at" field.
func (bcuo *BlockCategoryUpdateOne) SetArchivedAt(t time.Time) *BlockCategoryUpdateOne {
	bcuo.mutation.SetArchivedAt(t)
	return bcuo
}

// SetNillableArchivedAt sets the "archived_at" field if the given value is not nil.
func (bcuo *BlockCategoryUpdateOne) SetNillableArchivedAt(t *time.Time) *BlockCategoryUpdateOne {
	if t != nil {
		bcuo.SetArchivedAt(*t)
	}
	return bcuo
}

// AddBlockIDs adds the "blocks" edge to the BlockInfo entity by IDs.
func (bcuo *BlockCategoryUpdateOne) AddBlockIDs(ids ...uuid.UUID) *BlockCategoryUpdateOne {
	bcuo.mutation.AddBlockIDs(ids...)
	return bcuo
}

// AddBlocks adds the "blocks" edges to the BlockInfo entity.
func (bcuo *BlockCategoryUpdateOne) AddBlocks(b ...*BlockInfo) *BlockCategoryUpdateOne {
	ids := make([]uuid.UUID, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bcuo.AddBlockIDs(ids...)
}

// Mutation returns the BlockCategoryMutation object of the builder.
func (bcuo *BlockCategoryUpdateOne) Mutation() *BlockCategoryMutation {
	return bcuo.mutation
}

// ClearBlocks clears all "blocks" edges to the BlockInfo entity.
func (bcuo *BlockCategoryUpdateOne) ClearBlocks() *BlockCategoryUpdateOne {
	bcuo.mutation.ClearBlocks()
	return bcuo
}

// RemoveBlockIDs removes the "blocks" edge to BlockInfo entities by IDs.
func (bcuo *BlockCategoryUpdateOne) RemoveBlockIDs(ids ...uuid.UUID) *BlockCategoryUpdateOne {
	bcuo.mutation.RemoveBlockIDs(ids...)
	return bcuo
}

// RemoveBlocks removes "blocks" edges to BlockInfo entities.
func (bcuo *BlockCategoryUpdateOne) RemoveBlocks(b ...*BlockInfo) *BlockCategoryUpdateOne {
	ids := make([]uuid.UUID, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bcuo.RemoveBlockIDs(ids...)
}

// Where appends a list predicates to the BlockCategoryUpdate builder.
func (bcuo *BlockCategoryUpdateOne) Where(ps ...predicate.BlockCategory) *BlockCategoryUpdateOne {
	bcuo.mutation.Where(ps...)
	return bcuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (bcuo *BlockCategoryUpdateOne) Select(field string, fields ...string) *BlockCategoryUpdateOne {
	bcuo.fields = append([]string{field}, fields...)
	return bcuo
}

// Save executes the query and returns the updated BlockCategory entity.
func (bcuo *BlockCategoryUpdateOne) Save(ctx context.Context) (*BlockCategory, error) {
	return withHooks(ctx, bcuo.sqlSave, bcuo.mutation, bcuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bcuo *BlockCategoryUpdateOne) SaveX(ctx context.Context) *BlockCategory {
	node, err := bcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (bcuo *BlockCategoryUpdateOne) Exec(ctx context.Context) error {
	_, err := bcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcuo *BlockCategoryUpdateOne) ExecX(ctx context.Context) {
	if err := bcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bcuo *BlockCategoryUpdateOne) check() error {
	if _, ok := bcuo.mutation.BotID(); bcuo.mutation.BotCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "BlockCategory.bot"`)
	}
	return nil
}

func (bcuo *BlockCategoryUpdateOne) sqlSave(ctx context.Context) (_node *BlockCategory, err error) {
	if err := bcuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(blockcategory.Table, blockcategory.Columns, sqlgraph.NewFieldSpec(blockcategory.FieldID, field.TypeUUID))
	id, ok := bcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "BlockCategory.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := bcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, blockcategory.FieldID)
		for _, f := range fields {
			if !blockcategory.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != blockcategory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := bcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bcuo.mutation.Name(); ok {
		_spec.SetField(blockcategory.FieldName, field.TypeString, value)
	}
	if value, ok := bcuo.mutation.UpdatedAt(); ok {
		_spec.SetField(blockcategory.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := bcuo.mutation.ArchivedAt(); ok {
		_spec.SetField(blockcategory.FieldArchivedAt, field.TypeTime, value)
	}
	if bcuo.mutation.BlocksCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bcuo.mutation.RemovedBlocksIDs(); len(nodes) > 0 && !bcuo.mutation.BlocksCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bcuo.mutation.BlocksIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &BlockCategory{config: bcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, bcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{blockcategory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	bcuo.mutation.done = true
	return _node, nil
}
