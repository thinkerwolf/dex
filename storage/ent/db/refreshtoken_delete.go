// Code generated by entc, DO NOT EDIT.

package db

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/koderover/dex/storage/ent/db/predicate"
	"github.com/koderover/dex/storage/ent/db/refreshtoken"
)

// RefreshTokenDelete is the builder for deleting a RefreshToken entity.
type RefreshTokenDelete struct {
	config
	hooks    []Hook
	mutation *RefreshTokenMutation
}

// Where adds a new predicate to the RefreshTokenDelete builder.
func (rtd *RefreshTokenDelete) Where(ps ...predicate.RefreshToken) *RefreshTokenDelete {
	rtd.mutation.predicates = append(rtd.mutation.predicates, ps...)
	return rtd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (rtd *RefreshTokenDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(rtd.hooks) == 0 {
		affected, err = rtd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RefreshTokenMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			rtd.mutation = mutation
			affected, err = rtd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(rtd.hooks) - 1; i >= 0; i-- {
			mut = rtd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rtd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (rtd *RefreshTokenDelete) ExecX(ctx context.Context) int {
	n, err := rtd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (rtd *RefreshTokenDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: refreshtoken.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: refreshtoken.FieldID,
			},
		},
	}
	if ps := rtd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, rtd.driver, _spec)
}

// RefreshTokenDeleteOne is the builder for deleting a single RefreshToken entity.
type RefreshTokenDeleteOne struct {
	rtd *RefreshTokenDelete
}

// Exec executes the deletion query.
func (rtdo *RefreshTokenDeleteOne) Exec(ctx context.Context) error {
	n, err := rtdo.rtd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{refreshtoken.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (rtdo *RefreshTokenDeleteOne) ExecX(ctx context.Context) {
	rtdo.rtd.ExecX(ctx)
}
