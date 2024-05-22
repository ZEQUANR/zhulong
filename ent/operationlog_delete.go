// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ZEQUANR/zhulong/ent/operationlog"
	"github.com/ZEQUANR/zhulong/ent/predicate"
)

// OperationLogDelete is the builder for deleting a OperationLog entity.
type OperationLogDelete struct {
	config
	hooks    []Hook
	mutation *OperationLogMutation
}

// Where appends a list predicates to the OperationLogDelete builder.
func (old *OperationLogDelete) Where(ps ...predicate.OperationLog) *OperationLogDelete {
	old.mutation.Where(ps...)
	return old
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (old *OperationLogDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, old.sqlExec, old.mutation, old.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (old *OperationLogDelete) ExecX(ctx context.Context) int {
	n, err := old.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (old *OperationLogDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(operationlog.Table, sqlgraph.NewFieldSpec(operationlog.FieldID, field.TypeInt))
	if ps := old.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, old.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	old.mutation.done = true
	return affected, err
}

// OperationLogDeleteOne is the builder for deleting a single OperationLog entity.
type OperationLogDeleteOne struct {
	old *OperationLogDelete
}

// Where appends a list predicates to the OperationLogDelete builder.
func (oldo *OperationLogDeleteOne) Where(ps ...predicate.OperationLog) *OperationLogDeleteOne {
	oldo.old.mutation.Where(ps...)
	return oldo
}

// Exec executes the deletion query.
func (oldo *OperationLogDeleteOne) Exec(ctx context.Context) error {
	n, err := oldo.old.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{operationlog.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (oldo *OperationLogDeleteOne) ExecX(ctx context.Context) {
	if err := oldo.Exec(ctx); err != nil {
		panic(err)
	}
}
