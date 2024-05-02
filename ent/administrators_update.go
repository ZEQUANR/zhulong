// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ZEQUANR/zhulong/ent/administrators"
	"github.com/ZEQUANR/zhulong/ent/predicate"
)

// AdministratorsUpdate is the builder for updating Administrators entities.
type AdministratorsUpdate struct {
	config
	hooks    []Hook
	mutation *AdministratorsMutation
}

// Where appends a list predicates to the AdministratorsUpdate builder.
func (au *AdministratorsUpdate) Where(ps ...predicate.Administrators) *AdministratorsUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetName sets the "name" field.
func (au *AdministratorsUpdate) SetName(s string) *AdministratorsUpdate {
	au.mutation.SetName(s)
	return au
}

// SetNillableName sets the "name" field if the given value is not nil.
func (au *AdministratorsUpdate) SetNillableName(s *string) *AdministratorsUpdate {
	if s != nil {
		au.SetName(*s)
	}
	return au
}

// SetCollege sets the "college" field.
func (au *AdministratorsUpdate) SetCollege(s string) *AdministratorsUpdate {
	au.mutation.SetCollege(s)
	return au
}

// SetNillableCollege sets the "college" field if the given value is not nil.
func (au *AdministratorsUpdate) SetNillableCollege(s *string) *AdministratorsUpdate {
	if s != nil {
		au.SetCollege(*s)
	}
	return au
}

// SetPhone sets the "phone" field.
func (au *AdministratorsUpdate) SetPhone(s string) *AdministratorsUpdate {
	au.mutation.SetPhone(s)
	return au
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (au *AdministratorsUpdate) SetNillablePhone(s *string) *AdministratorsUpdate {
	if s != nil {
		au.SetPhone(*s)
	}
	return au
}

// SetIdentity sets the "identity" field.
func (au *AdministratorsUpdate) SetIdentity(s string) *AdministratorsUpdate {
	au.mutation.SetIdentity(s)
	return au
}

// SetNillableIdentity sets the "identity" field if the given value is not nil.
func (au *AdministratorsUpdate) SetNillableIdentity(s *string) *AdministratorsUpdate {
	if s != nil {
		au.SetIdentity(*s)
	}
	return au
}

// Mutation returns the AdministratorsMutation object of the builder.
func (au *AdministratorsUpdate) Mutation() *AdministratorsMutation {
	return au.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AdministratorsUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *AdministratorsUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AdministratorsUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AdministratorsUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

func (au *AdministratorsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(administrators.Table, administrators.Columns, sqlgraph.NewFieldSpec(administrators.FieldID, field.TypeInt))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.Name(); ok {
		_spec.SetField(administrators.FieldName, field.TypeString, value)
	}
	if value, ok := au.mutation.College(); ok {
		_spec.SetField(administrators.FieldCollege, field.TypeString, value)
	}
	if value, ok := au.mutation.Phone(); ok {
		_spec.SetField(administrators.FieldPhone, field.TypeString, value)
	}
	if value, ok := au.mutation.Identity(); ok {
		_spec.SetField(administrators.FieldIdentity, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{administrators.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// AdministratorsUpdateOne is the builder for updating a single Administrators entity.
type AdministratorsUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AdministratorsMutation
}

// SetName sets the "name" field.
func (auo *AdministratorsUpdateOne) SetName(s string) *AdministratorsUpdateOne {
	auo.mutation.SetName(s)
	return auo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (auo *AdministratorsUpdateOne) SetNillableName(s *string) *AdministratorsUpdateOne {
	if s != nil {
		auo.SetName(*s)
	}
	return auo
}

// SetCollege sets the "college" field.
func (auo *AdministratorsUpdateOne) SetCollege(s string) *AdministratorsUpdateOne {
	auo.mutation.SetCollege(s)
	return auo
}

// SetNillableCollege sets the "college" field if the given value is not nil.
func (auo *AdministratorsUpdateOne) SetNillableCollege(s *string) *AdministratorsUpdateOne {
	if s != nil {
		auo.SetCollege(*s)
	}
	return auo
}

// SetPhone sets the "phone" field.
func (auo *AdministratorsUpdateOne) SetPhone(s string) *AdministratorsUpdateOne {
	auo.mutation.SetPhone(s)
	return auo
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (auo *AdministratorsUpdateOne) SetNillablePhone(s *string) *AdministratorsUpdateOne {
	if s != nil {
		auo.SetPhone(*s)
	}
	return auo
}

// SetIdentity sets the "identity" field.
func (auo *AdministratorsUpdateOne) SetIdentity(s string) *AdministratorsUpdateOne {
	auo.mutation.SetIdentity(s)
	return auo
}

// SetNillableIdentity sets the "identity" field if the given value is not nil.
func (auo *AdministratorsUpdateOne) SetNillableIdentity(s *string) *AdministratorsUpdateOne {
	if s != nil {
		auo.SetIdentity(*s)
	}
	return auo
}

// Mutation returns the AdministratorsMutation object of the builder.
func (auo *AdministratorsUpdateOne) Mutation() *AdministratorsMutation {
	return auo.mutation
}

// Where appends a list predicates to the AdministratorsUpdate builder.
func (auo *AdministratorsUpdateOne) Where(ps ...predicate.Administrators) *AdministratorsUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AdministratorsUpdateOne) Select(field string, fields ...string) *AdministratorsUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Administrators entity.
func (auo *AdministratorsUpdateOne) Save(ctx context.Context) (*Administrators, error) {
	return withHooks(ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AdministratorsUpdateOne) SaveX(ctx context.Context) *Administrators {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AdministratorsUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AdministratorsUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (auo *AdministratorsUpdateOne) sqlSave(ctx context.Context) (_node *Administrators, err error) {
	_spec := sqlgraph.NewUpdateSpec(administrators.Table, administrators.Columns, sqlgraph.NewFieldSpec(administrators.FieldID, field.TypeInt))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Administrators.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, administrators.FieldID)
		for _, f := range fields {
			if !administrators.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != administrators.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.Name(); ok {
		_spec.SetField(administrators.FieldName, field.TypeString, value)
	}
	if value, ok := auo.mutation.College(); ok {
		_spec.SetField(administrators.FieldCollege, field.TypeString, value)
	}
	if value, ok := auo.mutation.Phone(); ok {
		_spec.SetField(administrators.FieldPhone, field.TypeString, value)
	}
	if value, ok := auo.mutation.Identity(); ok {
		_spec.SetField(administrators.FieldIdentity, field.TypeString, value)
	}
	_node = &Administrators{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{administrators.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}
