// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ZEQUANR/zhulong/ent/operationlog"
	"github.com/ZEQUANR/zhulong/ent/user"
)

// OperationLogCreate is the builder for creating a OperationLog entity.
type OperationLogCreate struct {
	config
	mutation *OperationLogMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (olc *OperationLogCreate) SetName(s string) *OperationLogCreate {
	olc.mutation.SetName(s)
	return olc
}

// SetContext sets the "context" field.
func (olc *OperationLogCreate) SetContext(i int) *OperationLogCreate {
	olc.mutation.SetContext(i)
	return olc
}

// SetStatus sets the "status" field.
func (olc *OperationLogCreate) SetStatus(i int) *OperationLogCreate {
	olc.mutation.SetStatus(i)
	return olc
}

// SetTime sets the "time" field.
func (olc *OperationLogCreate) SetTime(t time.Time) *OperationLogCreate {
	olc.mutation.SetTime(t)
	return olc
}

// SetNillableTime sets the "time" field if the given value is not nil.
func (olc *OperationLogCreate) SetNillableTime(t *time.Time) *OperationLogCreate {
	if t != nil {
		olc.SetTime(*t)
	}
	return olc
}

// SetOperatorID sets the "operator" edge to the User entity by ID.
func (olc *OperationLogCreate) SetOperatorID(id int) *OperationLogCreate {
	olc.mutation.SetOperatorID(id)
	return olc
}

// SetNillableOperatorID sets the "operator" edge to the User entity by ID if the given value is not nil.
func (olc *OperationLogCreate) SetNillableOperatorID(id *int) *OperationLogCreate {
	if id != nil {
		olc = olc.SetOperatorID(*id)
	}
	return olc
}

// SetOperator sets the "operator" edge to the User entity.
func (olc *OperationLogCreate) SetOperator(u *User) *OperationLogCreate {
	return olc.SetOperatorID(u.ID)
}

// Mutation returns the OperationLogMutation object of the builder.
func (olc *OperationLogCreate) Mutation() *OperationLogMutation {
	return olc.mutation
}

// Save creates the OperationLog in the database.
func (olc *OperationLogCreate) Save(ctx context.Context) (*OperationLog, error) {
	olc.defaults()
	return withHooks(ctx, olc.sqlSave, olc.mutation, olc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (olc *OperationLogCreate) SaveX(ctx context.Context) *OperationLog {
	v, err := olc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (olc *OperationLogCreate) Exec(ctx context.Context) error {
	_, err := olc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (olc *OperationLogCreate) ExecX(ctx context.Context) {
	if err := olc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (olc *OperationLogCreate) defaults() {
	if _, ok := olc.mutation.Time(); !ok {
		v := operationlog.DefaultTime()
		olc.mutation.SetTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (olc *OperationLogCreate) check() error {
	if _, ok := olc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "OperationLog.name"`)}
	}
	if _, ok := olc.mutation.Context(); !ok {
		return &ValidationError{Name: "context", err: errors.New(`ent: missing required field "OperationLog.context"`)}
	}
	if _, ok := olc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "OperationLog.status"`)}
	}
	if _, ok := olc.mutation.Time(); !ok {
		return &ValidationError{Name: "time", err: errors.New(`ent: missing required field "OperationLog.time"`)}
	}
	return nil
}

func (olc *OperationLogCreate) sqlSave(ctx context.Context) (*OperationLog, error) {
	if err := olc.check(); err != nil {
		return nil, err
	}
	_node, _spec := olc.createSpec()
	if err := sqlgraph.CreateNode(ctx, olc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	olc.mutation.id = &_node.ID
	olc.mutation.done = true
	return _node, nil
}

func (olc *OperationLogCreate) createSpec() (*OperationLog, *sqlgraph.CreateSpec) {
	var (
		_node = &OperationLog{config: olc.config}
		_spec = sqlgraph.NewCreateSpec(operationlog.Table, sqlgraph.NewFieldSpec(operationlog.FieldID, field.TypeInt))
	)
	if value, ok := olc.mutation.Name(); ok {
		_spec.SetField(operationlog.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := olc.mutation.Context(); ok {
		_spec.SetField(operationlog.FieldContext, field.TypeInt, value)
		_node.Context = value
	}
	if value, ok := olc.mutation.Status(); ok {
		_spec.SetField(operationlog.FieldStatus, field.TypeInt, value)
		_node.Status = value
	}
	if value, ok := olc.mutation.Time(); ok {
		_spec.SetField(operationlog.FieldTime, field.TypeTime, value)
		_node.Time = value
	}
	if nodes := olc.mutation.OperatorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   operationlog.OperatorTable,
			Columns: []string{operationlog.OperatorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_operating_record = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OperationLogCreateBulk is the builder for creating many OperationLog entities in bulk.
type OperationLogCreateBulk struct {
	config
	err      error
	builders []*OperationLogCreate
}

// Save creates the OperationLog entities in the database.
func (olcb *OperationLogCreateBulk) Save(ctx context.Context) ([]*OperationLog, error) {
	if olcb.err != nil {
		return nil, olcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(olcb.builders))
	nodes := make([]*OperationLog, len(olcb.builders))
	mutators := make([]Mutator, len(olcb.builders))
	for i := range olcb.builders {
		func(i int, root context.Context) {
			builder := olcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OperationLogMutation)
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
					_, err = mutators[i+1].Mutate(root, olcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, olcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, olcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (olcb *OperationLogCreateBulk) SaveX(ctx context.Context) []*OperationLog {
	v, err := olcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (olcb *OperationLogCreateBulk) Exec(ctx context.Context) error {
	_, err := olcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (olcb *OperationLogCreateBulk) ExecX(ctx context.Context) {
	if err := olcb.Exec(ctx); err != nil {
		panic(err)
	}
}