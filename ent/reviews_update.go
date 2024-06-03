// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ZEQUANR/zhulong/ent/predicate"
	"github.com/ZEQUANR/zhulong/ent/reviews"
	"github.com/ZEQUANR/zhulong/ent/thesis"
	"github.com/ZEQUANR/zhulong/ent/user"
)

// ReviewsUpdate is the builder for updating Reviews entities.
type ReviewsUpdate struct {
	config
	hooks    []Hook
	mutation *ReviewsMutation
}

// Where appends a list predicates to the ReviewsUpdate builder.
func (ru *ReviewsUpdate) Where(ps ...predicate.Reviews) *ReviewsUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetFileName sets the "file_name" field.
func (ru *ReviewsUpdate) SetFileName(s string) *ReviewsUpdate {
	ru.mutation.SetFileName(s)
	return ru
}

// SetNillableFileName sets the "file_name" field if the given value is not nil.
func (ru *ReviewsUpdate) SetNillableFileName(s *string) *ReviewsUpdate {
	if s != nil {
		ru.SetFileName(*s)
	}
	return ru
}

// ClearFileName clears the value of the "file_name" field.
func (ru *ReviewsUpdate) ClearFileName() *ReviewsUpdate {
	ru.mutation.ClearFileName()
	return ru
}

// SetFileURL sets the "file_url" field.
func (ru *ReviewsUpdate) SetFileURL(s string) *ReviewsUpdate {
	ru.mutation.SetFileURL(s)
	return ru
}

// SetNillableFileURL sets the "file_url" field if the given value is not nil.
func (ru *ReviewsUpdate) SetNillableFileURL(s *string) *ReviewsUpdate {
	if s != nil {
		ru.SetFileURL(*s)
	}
	return ru
}

// ClearFileURL clears the value of the "file_url" field.
func (ru *ReviewsUpdate) ClearFileURL() *ReviewsUpdate {
	ru.mutation.ClearFileURL()
	return ru
}

// SetCreateTime sets the "create_time" field.
func (ru *ReviewsUpdate) SetCreateTime(t time.Time) *ReviewsUpdate {
	ru.mutation.SetCreateTime(t)
	return ru
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (ru *ReviewsUpdate) SetNillableCreateTime(t *time.Time) *ReviewsUpdate {
	if t != nil {
		ru.SetCreateTime(*t)
	}
	return ru
}

// SetUploadersID sets the "uploaders" edge to the User entity by ID.
func (ru *ReviewsUpdate) SetUploadersID(id int) *ReviewsUpdate {
	ru.mutation.SetUploadersID(id)
	return ru
}

// SetNillableUploadersID sets the "uploaders" edge to the User entity by ID if the given value is not nil.
func (ru *ReviewsUpdate) SetNillableUploadersID(id *int) *ReviewsUpdate {
	if id != nil {
		ru = ru.SetUploadersID(*id)
	}
	return ru
}

// SetUploaders sets the "uploaders" edge to the User entity.
func (ru *ReviewsUpdate) SetUploaders(u *User) *ReviewsUpdate {
	return ru.SetUploadersID(u.ID)
}

// SetThesisResultID sets the "thesisResult" edge to the Thesis entity by ID.
func (ru *ReviewsUpdate) SetThesisResultID(id int) *ReviewsUpdate {
	ru.mutation.SetThesisResultID(id)
	return ru
}

// SetNillableThesisResultID sets the "thesisResult" edge to the Thesis entity by ID if the given value is not nil.
func (ru *ReviewsUpdate) SetNillableThesisResultID(id *int) *ReviewsUpdate {
	if id != nil {
		ru = ru.SetThesisResultID(*id)
	}
	return ru
}

// SetThesisResult sets the "thesisResult" edge to the Thesis entity.
func (ru *ReviewsUpdate) SetThesisResult(t *Thesis) *ReviewsUpdate {
	return ru.SetThesisResultID(t.ID)
}

// Mutation returns the ReviewsMutation object of the builder.
func (ru *ReviewsUpdate) Mutation() *ReviewsMutation {
	return ru.mutation
}

// ClearUploaders clears the "uploaders" edge to the User entity.
func (ru *ReviewsUpdate) ClearUploaders() *ReviewsUpdate {
	ru.mutation.ClearUploaders()
	return ru
}

// ClearThesisResult clears the "thesisResult" edge to the Thesis entity.
func (ru *ReviewsUpdate) ClearThesisResult() *ReviewsUpdate {
	ru.mutation.ClearThesisResult()
	return ru
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *ReviewsUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ru.sqlSave, ru.mutation, ru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ru *ReviewsUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *ReviewsUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *ReviewsUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ru *ReviewsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(reviews.Table, reviews.Columns, sqlgraph.NewFieldSpec(reviews.FieldID, field.TypeInt))
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.FileName(); ok {
		_spec.SetField(reviews.FieldFileName, field.TypeString, value)
	}
	if ru.mutation.FileNameCleared() {
		_spec.ClearField(reviews.FieldFileName, field.TypeString)
	}
	if value, ok := ru.mutation.FileURL(); ok {
		_spec.SetField(reviews.FieldFileURL, field.TypeString, value)
	}
	if ru.mutation.FileURLCleared() {
		_spec.ClearField(reviews.FieldFileURL, field.TypeString)
	}
	if value, ok := ru.mutation.CreateTime(); ok {
		_spec.SetField(reviews.FieldCreateTime, field.TypeTime, value)
	}
	if ru.mutation.UploadersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   reviews.UploadersTable,
			Columns: []string{reviews.UploadersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.UploadersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   reviews.UploadersTable,
			Columns: []string{reviews.UploadersColumn},
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
	if ru.mutation.ThesisResultCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   reviews.ThesisResultTable,
			Columns: []string{reviews.ThesisResultColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(thesis.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.ThesisResultIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   reviews.ThesisResultTable,
			Columns: []string{reviews.ThesisResultColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(thesis.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{reviews.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ru.mutation.done = true
	return n, nil
}

// ReviewsUpdateOne is the builder for updating a single Reviews entity.
type ReviewsUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ReviewsMutation
}

// SetFileName sets the "file_name" field.
func (ruo *ReviewsUpdateOne) SetFileName(s string) *ReviewsUpdateOne {
	ruo.mutation.SetFileName(s)
	return ruo
}

// SetNillableFileName sets the "file_name" field if the given value is not nil.
func (ruo *ReviewsUpdateOne) SetNillableFileName(s *string) *ReviewsUpdateOne {
	if s != nil {
		ruo.SetFileName(*s)
	}
	return ruo
}

// ClearFileName clears the value of the "file_name" field.
func (ruo *ReviewsUpdateOne) ClearFileName() *ReviewsUpdateOne {
	ruo.mutation.ClearFileName()
	return ruo
}

// SetFileURL sets the "file_url" field.
func (ruo *ReviewsUpdateOne) SetFileURL(s string) *ReviewsUpdateOne {
	ruo.mutation.SetFileURL(s)
	return ruo
}

// SetNillableFileURL sets the "file_url" field if the given value is not nil.
func (ruo *ReviewsUpdateOne) SetNillableFileURL(s *string) *ReviewsUpdateOne {
	if s != nil {
		ruo.SetFileURL(*s)
	}
	return ruo
}

// ClearFileURL clears the value of the "file_url" field.
func (ruo *ReviewsUpdateOne) ClearFileURL() *ReviewsUpdateOne {
	ruo.mutation.ClearFileURL()
	return ruo
}

// SetCreateTime sets the "create_time" field.
func (ruo *ReviewsUpdateOne) SetCreateTime(t time.Time) *ReviewsUpdateOne {
	ruo.mutation.SetCreateTime(t)
	return ruo
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (ruo *ReviewsUpdateOne) SetNillableCreateTime(t *time.Time) *ReviewsUpdateOne {
	if t != nil {
		ruo.SetCreateTime(*t)
	}
	return ruo
}

// SetUploadersID sets the "uploaders" edge to the User entity by ID.
func (ruo *ReviewsUpdateOne) SetUploadersID(id int) *ReviewsUpdateOne {
	ruo.mutation.SetUploadersID(id)
	return ruo
}

// SetNillableUploadersID sets the "uploaders" edge to the User entity by ID if the given value is not nil.
func (ruo *ReviewsUpdateOne) SetNillableUploadersID(id *int) *ReviewsUpdateOne {
	if id != nil {
		ruo = ruo.SetUploadersID(*id)
	}
	return ruo
}

// SetUploaders sets the "uploaders" edge to the User entity.
func (ruo *ReviewsUpdateOne) SetUploaders(u *User) *ReviewsUpdateOne {
	return ruo.SetUploadersID(u.ID)
}

// SetThesisResultID sets the "thesisResult" edge to the Thesis entity by ID.
func (ruo *ReviewsUpdateOne) SetThesisResultID(id int) *ReviewsUpdateOne {
	ruo.mutation.SetThesisResultID(id)
	return ruo
}

// SetNillableThesisResultID sets the "thesisResult" edge to the Thesis entity by ID if the given value is not nil.
func (ruo *ReviewsUpdateOne) SetNillableThesisResultID(id *int) *ReviewsUpdateOne {
	if id != nil {
		ruo = ruo.SetThesisResultID(*id)
	}
	return ruo
}

// SetThesisResult sets the "thesisResult" edge to the Thesis entity.
func (ruo *ReviewsUpdateOne) SetThesisResult(t *Thesis) *ReviewsUpdateOne {
	return ruo.SetThesisResultID(t.ID)
}

// Mutation returns the ReviewsMutation object of the builder.
func (ruo *ReviewsUpdateOne) Mutation() *ReviewsMutation {
	return ruo.mutation
}

// ClearUploaders clears the "uploaders" edge to the User entity.
func (ruo *ReviewsUpdateOne) ClearUploaders() *ReviewsUpdateOne {
	ruo.mutation.ClearUploaders()
	return ruo
}

// ClearThesisResult clears the "thesisResult" edge to the Thesis entity.
func (ruo *ReviewsUpdateOne) ClearThesisResult() *ReviewsUpdateOne {
	ruo.mutation.ClearThesisResult()
	return ruo
}

// Where appends a list predicates to the ReviewsUpdate builder.
func (ruo *ReviewsUpdateOne) Where(ps ...predicate.Reviews) *ReviewsUpdateOne {
	ruo.mutation.Where(ps...)
	return ruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *ReviewsUpdateOne) Select(field string, fields ...string) *ReviewsUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Reviews entity.
func (ruo *ReviewsUpdateOne) Save(ctx context.Context) (*Reviews, error) {
	return withHooks(ctx, ruo.sqlSave, ruo.mutation, ruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *ReviewsUpdateOne) SaveX(ctx context.Context) *Reviews {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *ReviewsUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *ReviewsUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ruo *ReviewsUpdateOne) sqlSave(ctx context.Context) (_node *Reviews, err error) {
	_spec := sqlgraph.NewUpdateSpec(reviews.Table, reviews.Columns, sqlgraph.NewFieldSpec(reviews.FieldID, field.TypeInt))
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Reviews.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, reviews.FieldID)
		for _, f := range fields {
			if !reviews.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != reviews.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.FileName(); ok {
		_spec.SetField(reviews.FieldFileName, field.TypeString, value)
	}
	if ruo.mutation.FileNameCleared() {
		_spec.ClearField(reviews.FieldFileName, field.TypeString)
	}
	if value, ok := ruo.mutation.FileURL(); ok {
		_spec.SetField(reviews.FieldFileURL, field.TypeString, value)
	}
	if ruo.mutation.FileURLCleared() {
		_spec.ClearField(reviews.FieldFileURL, field.TypeString)
	}
	if value, ok := ruo.mutation.CreateTime(); ok {
		_spec.SetField(reviews.FieldCreateTime, field.TypeTime, value)
	}
	if ruo.mutation.UploadersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   reviews.UploadersTable,
			Columns: []string{reviews.UploadersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.UploadersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   reviews.UploadersTable,
			Columns: []string{reviews.UploadersColumn},
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
	if ruo.mutation.ThesisResultCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   reviews.ThesisResultTable,
			Columns: []string{reviews.ThesisResultColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(thesis.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.ThesisResultIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   reviews.ThesisResultTable,
			Columns: []string{reviews.ThesisResultColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(thesis.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Reviews{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{reviews.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ruo.mutation.done = true
	return _node, nil
}
