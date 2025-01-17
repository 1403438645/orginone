// Code generated by entc, DO NOT EDIT.

package schema

import (
	"context"
	"fmt"
	"orginone/common/schema/asmarketmenu"
	"orginone/common/schema/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AsMarketMenuDelete is the builder for deleting a AsMarketMenu entity.
type AsMarketMenuDelete struct {
	config
	hooks    []Hook
	mutation *AsMarketMenuMutation
}

// Where appends a list predicates to the AsMarketMenuDelete builder.
func (ammd *AsMarketMenuDelete) Where(ps ...predicate.AsMarketMenu) *AsMarketMenuDelete {
	ammd.mutation.Where(ps...)
	return ammd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ammd *AsMarketMenuDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ammd.hooks) == 0 {
		affected, err = ammd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AsMarketMenuMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ammd.mutation = mutation
			affected, err = ammd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ammd.hooks) - 1; i >= 0; i-- {
			if ammd.hooks[i] == nil {
				return 0, fmt.Errorf("schema: uninitialized hook (forgotten import schema/runtime?)")
			}
			mut = ammd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ammd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ammd *AsMarketMenuDelete) ExecX(ctx context.Context) int {
	n, err := ammd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ammd *AsMarketMenuDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: asmarketmenu.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: asmarketmenu.FieldID,
			},
		},
	}
	if ps := ammd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, ammd.driver, _spec)
}

// AsMarketMenuDeleteOne is the builder for deleting a single AsMarketMenu entity.
type AsMarketMenuDeleteOne struct {
	ammd *AsMarketMenuDelete
}

// Exec executes the deletion query.
func (ammdo *AsMarketMenuDeleteOne) Exec(ctx context.Context) error {
	n, err := ammdo.ammd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{asmarketmenu.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ammdo *AsMarketMenuDeleteOne) ExecX(ctx context.Context) {
	ammdo.ammd.ExecX(ctx)
}
