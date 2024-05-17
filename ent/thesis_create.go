// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ZEQUANR/zhulong/ent/thesis"
	"github.com/ZEQUANR/zhulong/ent/user"
)

// ThesisCreate is the builder for creating a Thesis entity.
type ThesisCreate struct {
	config
	mutation *ThesisMutation
	hooks    []Hook
}

// SetFileName sets the "file_name" field.
func (tc *ThesisCreate) SetFileName(s string) *ThesisCreate {
	tc.mutation.SetFileName(s)
	return tc
}

// SetNillableFileName sets the "file_name" field if the given value is not nil.
func (tc *ThesisCreate) SetNillableFileName(s *string) *ThesisCreate {
	if s != nil {
		tc.SetFileName(*s)
	}
	return tc
}

// SetFileURL sets the "file_url" field.
func (tc *ThesisCreate) SetFileURL(s string) *ThesisCreate {
	tc.mutation.SetFileURL(s)
	return tc
}

// SetNillableFileURL sets the "file_url" field if the given value is not nil.
func (tc *ThesisCreate) SetNillableFileURL(s *string) *ThesisCreate {
	if s != nil {
		tc.SetFileURL(*s)
	}
	return tc
}

// SetFileState sets the "file_state" field.
func (tc *ThesisCreate) SetFileState(i int) *ThesisCreate {
	tc.mutation.SetFileState(i)
	return tc
}

// SetNillableFileState sets the "file_state" field if the given value is not nil.
func (tc *ThesisCreate) SetNillableFileState(i *int) *ThesisCreate {
	if i != nil {
		tc.SetFileState(*i)
	}
	return tc
}

// SetUploadTime sets the "upload_time" field.
func (tc *ThesisCreate) SetUploadTime(t time.Time) *ThesisCreate {
	tc.mutation.SetUploadTime(t)
	return tc
}

// SetNillableUploadTime sets the "upload_time" field if the given value is not nil.
func (tc *ThesisCreate) SetNillableUploadTime(t *time.Time) *ThesisCreate {
	if t != nil {
		tc.SetUploadTime(*t)
	}
	return tc
}

// SetChineseTitle sets the "chinese_title" field.
func (tc *ThesisCreate) SetChineseTitle(s string) *ThesisCreate {
	tc.mutation.SetChineseTitle(s)
	return tc
}

// SetEnglishTitle sets the "english_title" field.
func (tc *ThesisCreate) SetEnglishTitle(s string) *ThesisCreate {
	tc.mutation.SetEnglishTitle(s)
	return tc
}

// SetAuthors sets the "authors" field.
func (tc *ThesisCreate) SetAuthors(s string) *ThesisCreate {
	tc.mutation.SetAuthors(s)
	return tc
}

// SetTeachers sets the "teachers" field.
func (tc *ThesisCreate) SetTeachers(s string) *ThesisCreate {
	tc.mutation.SetTeachers(s)
	return tc
}

// SetFirstAdvance sets the "first_advance" field.
func (tc *ThesisCreate) SetFirstAdvance(s string) *ThesisCreate {
	tc.mutation.SetFirstAdvance(s)
	return tc
}

// SetSecondAdvance sets the "second_advance" field.
func (tc *ThesisCreate) SetSecondAdvance(s string) *ThesisCreate {
	tc.mutation.SetSecondAdvance(s)
	return tc
}

// SetThirdAdvance sets the "third_advance" field.
func (tc *ThesisCreate) SetThirdAdvance(s string) *ThesisCreate {
	tc.mutation.SetThirdAdvance(s)
	return tc
}

// SetDrawback sets the "drawback" field.
func (tc *ThesisCreate) SetDrawback(s string) *ThesisCreate {
	tc.mutation.SetDrawback(s)
	return tc
}

// SetCreateTime sets the "create_time" field.
func (tc *ThesisCreate) SetCreateTime(t time.Time) *ThesisCreate {
	tc.mutation.SetCreateTime(t)
	return tc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (tc *ThesisCreate) SetNillableCreateTime(t *time.Time) *ThesisCreate {
	if t != nil {
		tc.SetCreateTime(*t)
	}
	return tc
}

// SetUploadersID sets the "uploaders" edge to the User entity by ID.
func (tc *ThesisCreate) SetUploadersID(id int) *ThesisCreate {
	tc.mutation.SetUploadersID(id)
	return tc
}

// SetNillableUploadersID sets the "uploaders" edge to the User entity by ID if the given value is not nil.
func (tc *ThesisCreate) SetNillableUploadersID(id *int) *ThesisCreate {
	if id != nil {
		tc = tc.SetUploadersID(*id)
	}
	return tc
}

// SetUploaders sets the "uploaders" edge to the User entity.
func (tc *ThesisCreate) SetUploaders(u *User) *ThesisCreate {
	return tc.SetUploadersID(u.ID)
}

// SetExamineID sets the "examine" edge to the User entity by ID.
func (tc *ThesisCreate) SetExamineID(id int) *ThesisCreate {
	tc.mutation.SetExamineID(id)
	return tc
}

// SetNillableExamineID sets the "examine" edge to the User entity by ID if the given value is not nil.
func (tc *ThesisCreate) SetNillableExamineID(id *int) *ThesisCreate {
	if id != nil {
		tc = tc.SetExamineID(*id)
	}
	return tc
}

// SetExamine sets the "examine" edge to the User entity.
func (tc *ThesisCreate) SetExamine(u *User) *ThesisCreate {
	return tc.SetExamineID(u.ID)
}

// Mutation returns the ThesisMutation object of the builder.
func (tc *ThesisCreate) Mutation() *ThesisMutation {
	return tc.mutation
}

// Save creates the Thesis in the database.
func (tc *ThesisCreate) Save(ctx context.Context) (*Thesis, error) {
	tc.defaults()
	return withHooks(ctx, tc.sqlSave, tc.mutation, tc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tc *ThesisCreate) SaveX(ctx context.Context) *Thesis {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *ThesisCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *ThesisCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *ThesisCreate) defaults() {
	if _, ok := tc.mutation.CreateTime(); !ok {
		v := thesis.DefaultCreateTime()
		tc.mutation.SetCreateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *ThesisCreate) check() error {
	if _, ok := tc.mutation.ChineseTitle(); !ok {
		return &ValidationError{Name: "chinese_title", err: errors.New(`ent: missing required field "Thesis.chinese_title"`)}
	}
	if _, ok := tc.mutation.EnglishTitle(); !ok {
		return &ValidationError{Name: "english_title", err: errors.New(`ent: missing required field "Thesis.english_title"`)}
	}
	if _, ok := tc.mutation.Authors(); !ok {
		return &ValidationError{Name: "authors", err: errors.New(`ent: missing required field "Thesis.authors"`)}
	}
	if _, ok := tc.mutation.Teachers(); !ok {
		return &ValidationError{Name: "teachers", err: errors.New(`ent: missing required field "Thesis.teachers"`)}
	}
	if _, ok := tc.mutation.FirstAdvance(); !ok {
		return &ValidationError{Name: "first_advance", err: errors.New(`ent: missing required field "Thesis.first_advance"`)}
	}
	if _, ok := tc.mutation.SecondAdvance(); !ok {
		return &ValidationError{Name: "second_advance", err: errors.New(`ent: missing required field "Thesis.second_advance"`)}
	}
	if _, ok := tc.mutation.ThirdAdvance(); !ok {
		return &ValidationError{Name: "third_advance", err: errors.New(`ent: missing required field "Thesis.third_advance"`)}
	}
	if _, ok := tc.mutation.Drawback(); !ok {
		return &ValidationError{Name: "drawback", err: errors.New(`ent: missing required field "Thesis.drawback"`)}
	}
	if _, ok := tc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "Thesis.create_time"`)}
	}
	return nil
}

func (tc *ThesisCreate) sqlSave(ctx context.Context) (*Thesis, error) {
	if err := tc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	tc.mutation.id = &_node.ID
	tc.mutation.done = true
	return _node, nil
}

func (tc *ThesisCreate) createSpec() (*Thesis, *sqlgraph.CreateSpec) {
	var (
		_node = &Thesis{config: tc.config}
		_spec = sqlgraph.NewCreateSpec(thesis.Table, sqlgraph.NewFieldSpec(thesis.FieldID, field.TypeInt))
	)
	if value, ok := tc.mutation.FileName(); ok {
		_spec.SetField(thesis.FieldFileName, field.TypeString, value)
		_node.FileName = value
	}
	if value, ok := tc.mutation.FileURL(); ok {
		_spec.SetField(thesis.FieldFileURL, field.TypeString, value)
		_node.FileURL = value
	}
	if value, ok := tc.mutation.FileState(); ok {
		_spec.SetField(thesis.FieldFileState, field.TypeInt, value)
		_node.FileState = value
	}
	if value, ok := tc.mutation.UploadTime(); ok {
		_spec.SetField(thesis.FieldUploadTime, field.TypeTime, value)
		_node.UploadTime = value
	}
	if value, ok := tc.mutation.ChineseTitle(); ok {
		_spec.SetField(thesis.FieldChineseTitle, field.TypeString, value)
		_node.ChineseTitle = value
	}
	if value, ok := tc.mutation.EnglishTitle(); ok {
		_spec.SetField(thesis.FieldEnglishTitle, field.TypeString, value)
		_node.EnglishTitle = value
	}
	if value, ok := tc.mutation.Authors(); ok {
		_spec.SetField(thesis.FieldAuthors, field.TypeString, value)
		_node.Authors = value
	}
	if value, ok := tc.mutation.Teachers(); ok {
		_spec.SetField(thesis.FieldTeachers, field.TypeString, value)
		_node.Teachers = value
	}
	if value, ok := tc.mutation.FirstAdvance(); ok {
		_spec.SetField(thesis.FieldFirstAdvance, field.TypeString, value)
		_node.FirstAdvance = value
	}
	if value, ok := tc.mutation.SecondAdvance(); ok {
		_spec.SetField(thesis.FieldSecondAdvance, field.TypeString, value)
		_node.SecondAdvance = value
	}
	if value, ok := tc.mutation.ThirdAdvance(); ok {
		_spec.SetField(thesis.FieldThirdAdvance, field.TypeString, value)
		_node.ThirdAdvance = value
	}
	if value, ok := tc.mutation.Drawback(); ok {
		_spec.SetField(thesis.FieldDrawback, field.TypeString, value)
		_node.Drawback = value
	}
	if value, ok := tc.mutation.CreateTime(); ok {
		_spec.SetField(thesis.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if nodes := tc.mutation.UploadersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   thesis.UploadersTable,
			Columns: []string{thesis.UploadersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_thesis = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.ExamineIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   thesis.ExamineTable,
			Columns: []string{thesis.ExamineColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.thesis_examine = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ThesisCreateBulk is the builder for creating many Thesis entities in bulk.
type ThesisCreateBulk struct {
	config
	err      error
	builders []*ThesisCreate
}

// Save creates the Thesis entities in the database.
func (tcb *ThesisCreateBulk) Save(ctx context.Context) ([]*Thesis, error) {
	if tcb.err != nil {
		return nil, tcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Thesis, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ThesisMutation)
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
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *ThesisCreateBulk) SaveX(ctx context.Context) []*Thesis {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *ThesisCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *ThesisCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}
