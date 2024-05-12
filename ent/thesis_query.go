// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ZEQUANR/zhulong/ent/predicate"
	"github.com/ZEQUANR/zhulong/ent/thesis"
	"github.com/ZEQUANR/zhulong/ent/user"
)

// ThesisQuery is the builder for querying Thesis entities.
type ThesisQuery struct {
	config
	ctx           *QueryContext
	order         []thesis.OrderOption
	inters        []Interceptor
	predicates    []predicate.Thesis
	withUploaders *UserQuery
	withFKs       bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ThesisQuery builder.
func (tq *ThesisQuery) Where(ps ...predicate.Thesis) *ThesisQuery {
	tq.predicates = append(tq.predicates, ps...)
	return tq
}

// Limit the number of records to be returned by this query.
func (tq *ThesisQuery) Limit(limit int) *ThesisQuery {
	tq.ctx.Limit = &limit
	return tq
}

// Offset to start from.
func (tq *ThesisQuery) Offset(offset int) *ThesisQuery {
	tq.ctx.Offset = &offset
	return tq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tq *ThesisQuery) Unique(unique bool) *ThesisQuery {
	tq.ctx.Unique = &unique
	return tq
}

// Order specifies how the records should be ordered.
func (tq *ThesisQuery) Order(o ...thesis.OrderOption) *ThesisQuery {
	tq.order = append(tq.order, o...)
	return tq
}

// QueryUploaders chains the current query on the "uploaders" edge.
func (tq *ThesisQuery) QueryUploaders() *UserQuery {
	query := (&UserClient{config: tq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(thesis.Table, thesis.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, thesis.UploadersTable, thesis.UploadersColumn),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Thesis entity from the query.
// Returns a *NotFoundError when no Thesis was found.
func (tq *ThesisQuery) First(ctx context.Context) (*Thesis, error) {
	nodes, err := tq.Limit(1).All(setContextOp(ctx, tq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{thesis.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tq *ThesisQuery) FirstX(ctx context.Context) *Thesis {
	node, err := tq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Thesis ID from the query.
// Returns a *NotFoundError when no Thesis ID was found.
func (tq *ThesisQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tq.Limit(1).IDs(setContextOp(ctx, tq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{thesis.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tq *ThesisQuery) FirstIDX(ctx context.Context) int {
	id, err := tq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Thesis entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Thesis entity is found.
// Returns a *NotFoundError when no Thesis entities are found.
func (tq *ThesisQuery) Only(ctx context.Context) (*Thesis, error) {
	nodes, err := tq.Limit(2).All(setContextOp(ctx, tq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{thesis.Label}
	default:
		return nil, &NotSingularError{thesis.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tq *ThesisQuery) OnlyX(ctx context.Context) *Thesis {
	node, err := tq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Thesis ID in the query.
// Returns a *NotSingularError when more than one Thesis ID is found.
// Returns a *NotFoundError when no entities are found.
func (tq *ThesisQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tq.Limit(2).IDs(setContextOp(ctx, tq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{thesis.Label}
	default:
		err = &NotSingularError{thesis.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tq *ThesisQuery) OnlyIDX(ctx context.Context) int {
	id, err := tq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Theses.
func (tq *ThesisQuery) All(ctx context.Context) ([]*Thesis, error) {
	ctx = setContextOp(ctx, tq.ctx, "All")
	if err := tq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Thesis, *ThesisQuery]()
	return withInterceptors[[]*Thesis](ctx, tq, qr, tq.inters)
}

// AllX is like All, but panics if an error occurs.
func (tq *ThesisQuery) AllX(ctx context.Context) []*Thesis {
	nodes, err := tq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Thesis IDs.
func (tq *ThesisQuery) IDs(ctx context.Context) (ids []int, err error) {
	if tq.ctx.Unique == nil && tq.path != nil {
		tq.Unique(true)
	}
	ctx = setContextOp(ctx, tq.ctx, "IDs")
	if err = tq.Select(thesis.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tq *ThesisQuery) IDsX(ctx context.Context) []int {
	ids, err := tq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tq *ThesisQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, tq.ctx, "Count")
	if err := tq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, tq, querierCount[*ThesisQuery](), tq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (tq *ThesisQuery) CountX(ctx context.Context) int {
	count, err := tq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tq *ThesisQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, tq.ctx, "Exist")
	switch _, err := tq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (tq *ThesisQuery) ExistX(ctx context.Context) bool {
	exist, err := tq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ThesisQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tq *ThesisQuery) Clone() *ThesisQuery {
	if tq == nil {
		return nil
	}
	return &ThesisQuery{
		config:        tq.config,
		ctx:           tq.ctx.Clone(),
		order:         append([]thesis.OrderOption{}, tq.order...),
		inters:        append([]Interceptor{}, tq.inters...),
		predicates:    append([]predicate.Thesis{}, tq.predicates...),
		withUploaders: tq.withUploaders.Clone(),
		// clone intermediate query.
		sql:  tq.sql.Clone(),
		path: tq.path,
	}
}

// WithUploaders tells the query-builder to eager-load the nodes that are connected to
// the "uploaders" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *ThesisQuery) WithUploaders(opts ...func(*UserQuery)) *ThesisQuery {
	query := (&UserClient{config: tq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tq.withUploaders = query
	return tq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		FileName string `json:"file_name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Thesis.Query().
//		GroupBy(thesis.FieldFileName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (tq *ThesisQuery) GroupBy(field string, fields ...string) *ThesisGroupBy {
	tq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ThesisGroupBy{build: tq}
	grbuild.flds = &tq.ctx.Fields
	grbuild.label = thesis.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		FileName string `json:"file_name,omitempty"`
//	}
//
//	client.Thesis.Query().
//		Select(thesis.FieldFileName).
//		Scan(ctx, &v)
func (tq *ThesisQuery) Select(fields ...string) *ThesisSelect {
	tq.ctx.Fields = append(tq.ctx.Fields, fields...)
	sbuild := &ThesisSelect{ThesisQuery: tq}
	sbuild.label = thesis.Label
	sbuild.flds, sbuild.scan = &tq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ThesisSelect configured with the given aggregations.
func (tq *ThesisQuery) Aggregate(fns ...AggregateFunc) *ThesisSelect {
	return tq.Select().Aggregate(fns...)
}

func (tq *ThesisQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range tq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, tq); err != nil {
				return err
			}
		}
	}
	for _, f := range tq.ctx.Fields {
		if !thesis.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if tq.path != nil {
		prev, err := tq.path(ctx)
		if err != nil {
			return err
		}
		tq.sql = prev
	}
	return nil
}

func (tq *ThesisQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Thesis, error) {
	var (
		nodes       = []*Thesis{}
		withFKs     = tq.withFKs
		_spec       = tq.querySpec()
		loadedTypes = [1]bool{
			tq.withUploaders != nil,
		}
	)
	if tq.withUploaders != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, thesis.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Thesis).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Thesis{config: tq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, tq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := tq.withUploaders; query != nil {
		if err := tq.loadUploaders(ctx, query, nodes, nil,
			func(n *Thesis, e *User) { n.Edges.Uploaders = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (tq *ThesisQuery) loadUploaders(ctx context.Context, query *UserQuery, nodes []*Thesis, init func(*Thesis), assign func(*Thesis, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Thesis)
	for i := range nodes {
		if nodes[i].user_thesis == nil {
			continue
		}
		fk := *nodes[i].user_thesis
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_thesis" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (tq *ThesisQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tq.querySpec()
	_spec.Node.Columns = tq.ctx.Fields
	if len(tq.ctx.Fields) > 0 {
		_spec.Unique = tq.ctx.Unique != nil && *tq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, tq.driver, _spec)
}

func (tq *ThesisQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(thesis.Table, thesis.Columns, sqlgraph.NewFieldSpec(thesis.FieldID, field.TypeInt))
	_spec.From = tq.sql
	if unique := tq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if tq.path != nil {
		_spec.Unique = true
	}
	if fields := tq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, thesis.FieldID)
		for i := range fields {
			if fields[i] != thesis.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tq *ThesisQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tq.driver.Dialect())
	t1 := builder.Table(thesis.Table)
	columns := tq.ctx.Fields
	if len(columns) == 0 {
		columns = thesis.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tq.sql != nil {
		selector = tq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tq.ctx.Unique != nil && *tq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range tq.predicates {
		p(selector)
	}
	for _, p := range tq.order {
		p(selector)
	}
	if offset := tq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ThesisGroupBy is the group-by builder for Thesis entities.
type ThesisGroupBy struct {
	selector
	build *ThesisQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tgb *ThesisGroupBy) Aggregate(fns ...AggregateFunc) *ThesisGroupBy {
	tgb.fns = append(tgb.fns, fns...)
	return tgb
}

// Scan applies the selector query and scans the result into the given value.
func (tgb *ThesisGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tgb.build.ctx, "GroupBy")
	if err := tgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ThesisQuery, *ThesisGroupBy](ctx, tgb.build, tgb, tgb.build.inters, v)
}

func (tgb *ThesisGroupBy) sqlScan(ctx context.Context, root *ThesisQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(tgb.fns))
	for _, fn := range tgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*tgb.flds)+len(tgb.fns))
		for _, f := range *tgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*tgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ThesisSelect is the builder for selecting fields of Thesis entities.
type ThesisSelect struct {
	*ThesisQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ts *ThesisSelect) Aggregate(fns ...AggregateFunc) *ThesisSelect {
	ts.fns = append(ts.fns, fns...)
	return ts
}

// Scan applies the selector query and scans the result into the given value.
func (ts *ThesisSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ts.ctx, "Select")
	if err := ts.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ThesisQuery, *ThesisSelect](ctx, ts.ThesisQuery, ts, ts.inters, v)
}

func (ts *ThesisSelect) sqlScan(ctx context.Context, root *ThesisQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ts.fns))
	for _, fn := range ts.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ts.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
