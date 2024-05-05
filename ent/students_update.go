// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ZEQUANR/zhulong/ent/predicate"
	"github.com/ZEQUANR/zhulong/ent/students"
	"github.com/ZEQUANR/zhulong/ent/user"
)

// StudentsUpdate is the builder for updating Students entities.
type StudentsUpdate struct {
	config
	hooks    []Hook
	mutation *StudentsMutation
}

// Where appends a list predicates to the StudentsUpdate builder.
func (su *StudentsUpdate) Where(ps ...predicate.Students) *StudentsUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetName sets the "name" field.
func (su *StudentsUpdate) SetName(s string) *StudentsUpdate {
	su.mutation.SetName(s)
	return su
}

// SetNillableName sets the "name" field if the given value is not nil.
func (su *StudentsUpdate) SetNillableName(s *string) *StudentsUpdate {
	if s != nil {
		su.SetName(*s)
	}
	return su
}

// SetCollege sets the "college" field.
func (su *StudentsUpdate) SetCollege(s string) *StudentsUpdate {
	su.mutation.SetCollege(s)
	return su
}

// SetNillableCollege sets the "college" field if the given value is not nil.
func (su *StudentsUpdate) SetNillableCollege(s *string) *StudentsUpdate {
	if s != nil {
		su.SetCollege(*s)
	}
	return su
}

// SetPhone sets the "phone" field.
func (su *StudentsUpdate) SetPhone(s string) *StudentsUpdate {
	su.mutation.SetPhone(s)
	return su
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (su *StudentsUpdate) SetNillablePhone(s *string) *StudentsUpdate {
	if s != nil {
		su.SetPhone(*s)
	}
	return su
}

// SetSubject sets the "subject" field.
func (su *StudentsUpdate) SetSubject(s string) *StudentsUpdate {
	su.mutation.SetSubject(s)
	return su
}

// SetNillableSubject sets the "subject" field if the given value is not nil.
func (su *StudentsUpdate) SetNillableSubject(s *string) *StudentsUpdate {
	if s != nil {
		su.SetSubject(*s)
	}
	return su
}

// SetClass sets the "class" field.
func (su *StudentsUpdate) SetClass(s string) *StudentsUpdate {
	su.mutation.SetClass(s)
	return su
}

// SetNillableClass sets the "class" field if the given value is not nil.
func (su *StudentsUpdate) SetNillableClass(s *string) *StudentsUpdate {
	if s != nil {
		su.SetClass(*s)
	}
	return su
}

// SetIdentity sets the "identity" field.
func (su *StudentsUpdate) SetIdentity(s string) *StudentsUpdate {
	su.mutation.SetIdentity(s)
	return su
}

// SetNillableIdentity sets the "identity" field if the given value is not nil.
func (su *StudentsUpdate) SetNillableIdentity(s *string) *StudentsUpdate {
	if s != nil {
		su.SetIdentity(*s)
	}
	return su
}

// SetUsersID sets the "users" edge to the User entity by ID.
func (su *StudentsUpdate) SetUsersID(id int) *StudentsUpdate {
	su.mutation.SetUsersID(id)
	return su
}

// SetUsers sets the "users" edge to the User entity.
func (su *StudentsUpdate) SetUsers(u *User) *StudentsUpdate {
	return su.SetUsersID(u.ID)
}

// Mutation returns the StudentsMutation object of the builder.
func (su *StudentsUpdate) Mutation() *StudentsMutation {
	return su.mutation
}

// ClearUsers clears the "users" edge to the User entity.
func (su *StudentsUpdate) ClearUsers() *StudentsUpdate {
	su.mutation.ClearUsers()
	return su
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *StudentsUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *StudentsUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *StudentsUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *StudentsUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (su *StudentsUpdate) check() error {
	if _, ok := su.mutation.UsersID(); su.mutation.UsersCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Students.users"`)
	}
	return nil
}

func (su *StudentsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := su.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(students.Table, students.Columns, sqlgraph.NewFieldSpec(students.FieldID, field.TypeInt))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.Name(); ok {
		_spec.SetField(students.FieldName, field.TypeString, value)
	}
	if value, ok := su.mutation.College(); ok {
		_spec.SetField(students.FieldCollege, field.TypeString, value)
	}
	if value, ok := su.mutation.Phone(); ok {
		_spec.SetField(students.FieldPhone, field.TypeString, value)
	}
	if value, ok := su.mutation.Subject(); ok {
		_spec.SetField(students.FieldSubject, field.TypeString, value)
	}
	if value, ok := su.mutation.Class(); ok {
		_spec.SetField(students.FieldClass, field.TypeString, value)
	}
	if value, ok := su.mutation.Identity(); ok {
		_spec.SetField(students.FieldIdentity, field.TypeString, value)
	}
	if su.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   students.UsersTable,
			Columns: []string{students.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   students.UsersTable,
			Columns: []string{students.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{students.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// StudentsUpdateOne is the builder for updating a single Students entity.
type StudentsUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *StudentsMutation
}

// SetName sets the "name" field.
func (suo *StudentsUpdateOne) SetName(s string) *StudentsUpdateOne {
	suo.mutation.SetName(s)
	return suo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (suo *StudentsUpdateOne) SetNillableName(s *string) *StudentsUpdateOne {
	if s != nil {
		suo.SetName(*s)
	}
	return suo
}

// SetCollege sets the "college" field.
func (suo *StudentsUpdateOne) SetCollege(s string) *StudentsUpdateOne {
	suo.mutation.SetCollege(s)
	return suo
}

// SetNillableCollege sets the "college" field if the given value is not nil.
func (suo *StudentsUpdateOne) SetNillableCollege(s *string) *StudentsUpdateOne {
	if s != nil {
		suo.SetCollege(*s)
	}
	return suo
}

// SetPhone sets the "phone" field.
func (suo *StudentsUpdateOne) SetPhone(s string) *StudentsUpdateOne {
	suo.mutation.SetPhone(s)
	return suo
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (suo *StudentsUpdateOne) SetNillablePhone(s *string) *StudentsUpdateOne {
	if s != nil {
		suo.SetPhone(*s)
	}
	return suo
}

// SetSubject sets the "subject" field.
func (suo *StudentsUpdateOne) SetSubject(s string) *StudentsUpdateOne {
	suo.mutation.SetSubject(s)
	return suo
}

// SetNillableSubject sets the "subject" field if the given value is not nil.
func (suo *StudentsUpdateOne) SetNillableSubject(s *string) *StudentsUpdateOne {
	if s != nil {
		suo.SetSubject(*s)
	}
	return suo
}

// SetClass sets the "class" field.
func (suo *StudentsUpdateOne) SetClass(s string) *StudentsUpdateOne {
	suo.mutation.SetClass(s)
	return suo
}

// SetNillableClass sets the "class" field if the given value is not nil.
func (suo *StudentsUpdateOne) SetNillableClass(s *string) *StudentsUpdateOne {
	if s != nil {
		suo.SetClass(*s)
	}
	return suo
}

// SetIdentity sets the "identity" field.
func (suo *StudentsUpdateOne) SetIdentity(s string) *StudentsUpdateOne {
	suo.mutation.SetIdentity(s)
	return suo
}

// SetNillableIdentity sets the "identity" field if the given value is not nil.
func (suo *StudentsUpdateOne) SetNillableIdentity(s *string) *StudentsUpdateOne {
	if s != nil {
		suo.SetIdentity(*s)
	}
	return suo
}

// SetUsersID sets the "users" edge to the User entity by ID.
func (suo *StudentsUpdateOne) SetUsersID(id int) *StudentsUpdateOne {
	suo.mutation.SetUsersID(id)
	return suo
}

// SetUsers sets the "users" edge to the User entity.
func (suo *StudentsUpdateOne) SetUsers(u *User) *StudentsUpdateOne {
	return suo.SetUsersID(u.ID)
}

// Mutation returns the StudentsMutation object of the builder.
func (suo *StudentsUpdateOne) Mutation() *StudentsMutation {
	return suo.mutation
}

// ClearUsers clears the "users" edge to the User entity.
func (suo *StudentsUpdateOne) ClearUsers() *StudentsUpdateOne {
	suo.mutation.ClearUsers()
	return suo
}

// Where appends a list predicates to the StudentsUpdate builder.
func (suo *StudentsUpdateOne) Where(ps ...predicate.Students) *StudentsUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *StudentsUpdateOne) Select(field string, fields ...string) *StudentsUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Students entity.
func (suo *StudentsUpdateOne) Save(ctx context.Context) (*Students, error) {
	return withHooks(ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *StudentsUpdateOne) SaveX(ctx context.Context) *Students {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *StudentsUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *StudentsUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (suo *StudentsUpdateOne) check() error {
	if _, ok := suo.mutation.UsersID(); suo.mutation.UsersCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Students.users"`)
	}
	return nil
}

func (suo *StudentsUpdateOne) sqlSave(ctx context.Context) (_node *Students, err error) {
	if err := suo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(students.Table, students.Columns, sqlgraph.NewFieldSpec(students.FieldID, field.TypeInt))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Students.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, students.FieldID)
		for _, f := range fields {
			if !students.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != students.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.Name(); ok {
		_spec.SetField(students.FieldName, field.TypeString, value)
	}
	if value, ok := suo.mutation.College(); ok {
		_spec.SetField(students.FieldCollege, field.TypeString, value)
	}
	if value, ok := suo.mutation.Phone(); ok {
		_spec.SetField(students.FieldPhone, field.TypeString, value)
	}
	if value, ok := suo.mutation.Subject(); ok {
		_spec.SetField(students.FieldSubject, field.TypeString, value)
	}
	if value, ok := suo.mutation.Class(); ok {
		_spec.SetField(students.FieldClass, field.TypeString, value)
	}
	if value, ok := suo.mutation.Identity(); ok {
		_spec.SetField(students.FieldIdentity, field.TypeString, value)
	}
	if suo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   students.UsersTable,
			Columns: []string{students.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   students.UsersTable,
			Columns: []string{students.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Students{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{students.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}
