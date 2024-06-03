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
	"github.com/ZEQUANR/zhulong/ent/reviews"
	"github.com/ZEQUANR/zhulong/ent/thesis"
	"github.com/ZEQUANR/zhulong/ent/user"
)

// ReviewsQuery is the builder for querying Reviews entities.
type ReviewsQuery struct {
	config
	ctx              *QueryContext
	order            []reviews.OrderOption
	inters           []Interceptor
	predicates       []predicate.Reviews
	withUploaders    *UserQuery
	withThesisResult *ThesisQuery
	withFKs          bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ReviewsQuery builder.
func (rq *ReviewsQuery) Where(ps ...predicate.Reviews) *ReviewsQuery {
	rq.predicates = append(rq.predicates, ps...)
	return rq
}

// Limit the number of records to be returned by this query.
func (rq *ReviewsQuery) Limit(limit int) *ReviewsQuery {
	rq.ctx.Limit = &limit
	return rq
}

// Offset to start from.
func (rq *ReviewsQuery) Offset(offset int) *ReviewsQuery {
	rq.ctx.Offset = &offset
	return rq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (rq *ReviewsQuery) Unique(unique bool) *ReviewsQuery {
	rq.ctx.Unique = &unique
	return rq
}

// Order specifies how the records should be ordered.
func (rq *ReviewsQuery) Order(o ...reviews.OrderOption) *ReviewsQuery {
	rq.order = append(rq.order, o...)
	return rq
}

// QueryUploaders chains the current query on the "uploaders" edge.
func (rq *ReviewsQuery) QueryUploaders() *UserQuery {
	query := (&UserClient{config: rq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(reviews.Table, reviews.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, reviews.UploadersTable, reviews.UploadersColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryThesisResult chains the current query on the "thesisResult" edge.
func (rq *ReviewsQuery) QueryThesisResult() *ThesisQuery {
	query := (&ThesisClient{config: rq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(reviews.Table, reviews.FieldID, selector),
			sqlgraph.To(thesis.Table, thesis.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, reviews.ThesisResultTable, reviews.ThesisResultColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Reviews entity from the query.
// Returns a *NotFoundError when no Reviews was found.
func (rq *ReviewsQuery) First(ctx context.Context) (*Reviews, error) {
	nodes, err := rq.Limit(1).All(setContextOp(ctx, rq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{reviews.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rq *ReviewsQuery) FirstX(ctx context.Context) *Reviews {
	node, err := rq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Reviews ID from the query.
// Returns a *NotFoundError when no Reviews ID was found.
func (rq *ReviewsQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = rq.Limit(1).IDs(setContextOp(ctx, rq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{reviews.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (rq *ReviewsQuery) FirstIDX(ctx context.Context) int {
	id, err := rq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Reviews entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Reviews entity is found.
// Returns a *NotFoundError when no Reviews entities are found.
func (rq *ReviewsQuery) Only(ctx context.Context) (*Reviews, error) {
	nodes, err := rq.Limit(2).All(setContextOp(ctx, rq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{reviews.Label}
	default:
		return nil, &NotSingularError{reviews.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rq *ReviewsQuery) OnlyX(ctx context.Context) *Reviews {
	node, err := rq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Reviews ID in the query.
// Returns a *NotSingularError when more than one Reviews ID is found.
// Returns a *NotFoundError when no entities are found.
func (rq *ReviewsQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = rq.Limit(2).IDs(setContextOp(ctx, rq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{reviews.Label}
	default:
		err = &NotSingularError{reviews.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (rq *ReviewsQuery) OnlyIDX(ctx context.Context) int {
	id, err := rq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ReviewsSlice.
func (rq *ReviewsQuery) All(ctx context.Context) ([]*Reviews, error) {
	ctx = setContextOp(ctx, rq.ctx, "All")
	if err := rq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Reviews, *ReviewsQuery]()
	return withInterceptors[[]*Reviews](ctx, rq, qr, rq.inters)
}

// AllX is like All, but panics if an error occurs.
func (rq *ReviewsQuery) AllX(ctx context.Context) []*Reviews {
	nodes, err := rq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Reviews IDs.
func (rq *ReviewsQuery) IDs(ctx context.Context) (ids []int, err error) {
	if rq.ctx.Unique == nil && rq.path != nil {
		rq.Unique(true)
	}
	ctx = setContextOp(ctx, rq.ctx, "IDs")
	if err = rq.Select(reviews.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (rq *ReviewsQuery) IDsX(ctx context.Context) []int {
	ids, err := rq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (rq *ReviewsQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, rq.ctx, "Count")
	if err := rq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, rq, querierCount[*ReviewsQuery](), rq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (rq *ReviewsQuery) CountX(ctx context.Context) int {
	count, err := rq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rq *ReviewsQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, rq.ctx, "Exist")
	switch _, err := rq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (rq *ReviewsQuery) ExistX(ctx context.Context) bool {
	exist, err := rq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ReviewsQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rq *ReviewsQuery) Clone() *ReviewsQuery {
	if rq == nil {
		return nil
	}
	return &ReviewsQuery{
		config:           rq.config,
		ctx:              rq.ctx.Clone(),
		order:            append([]reviews.OrderOption{}, rq.order...),
		inters:           append([]Interceptor{}, rq.inters...),
		predicates:       append([]predicate.Reviews{}, rq.predicates...),
		withUploaders:    rq.withUploaders.Clone(),
		withThesisResult: rq.withThesisResult.Clone(),
		// clone intermediate query.
		sql:  rq.sql.Clone(),
		path: rq.path,
	}
}

// WithUploaders tells the query-builder to eager-load the nodes that are connected to
// the "uploaders" edge. The optional arguments are used to configure the query builder of the edge.
func (rq *ReviewsQuery) WithUploaders(opts ...func(*UserQuery)) *ReviewsQuery {
	query := (&UserClient{config: rq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rq.withUploaders = query
	return rq
}

// WithThesisResult tells the query-builder to eager-load the nodes that are connected to
// the "thesisResult" edge. The optional arguments are used to configure the query builder of the edge.
func (rq *ReviewsQuery) WithThesisResult(opts ...func(*ThesisQuery)) *ReviewsQuery {
	query := (&ThesisClient{config: rq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rq.withThesisResult = query
	return rq
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
//	client.Reviews.Query().
//		GroupBy(reviews.FieldFileName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (rq *ReviewsQuery) GroupBy(field string, fields ...string) *ReviewsGroupBy {
	rq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ReviewsGroupBy{build: rq}
	grbuild.flds = &rq.ctx.Fields
	grbuild.label = reviews.Label
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
//	client.Reviews.Query().
//		Select(reviews.FieldFileName).
//		Scan(ctx, &v)
func (rq *ReviewsQuery) Select(fields ...string) *ReviewsSelect {
	rq.ctx.Fields = append(rq.ctx.Fields, fields...)
	sbuild := &ReviewsSelect{ReviewsQuery: rq}
	sbuild.label = reviews.Label
	sbuild.flds, sbuild.scan = &rq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ReviewsSelect configured with the given aggregations.
func (rq *ReviewsQuery) Aggregate(fns ...AggregateFunc) *ReviewsSelect {
	return rq.Select().Aggregate(fns...)
}

func (rq *ReviewsQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range rq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, rq); err != nil {
				return err
			}
		}
	}
	for _, f := range rq.ctx.Fields {
		if !reviews.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if rq.path != nil {
		prev, err := rq.path(ctx)
		if err != nil {
			return err
		}
		rq.sql = prev
	}
	return nil
}

func (rq *ReviewsQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Reviews, error) {
	var (
		nodes       = []*Reviews{}
		withFKs     = rq.withFKs
		_spec       = rq.querySpec()
		loadedTypes = [2]bool{
			rq.withUploaders != nil,
			rq.withThesisResult != nil,
		}
	)
	if rq.withUploaders != nil || rq.withThesisResult != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, reviews.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Reviews).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Reviews{config: rq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, rq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := rq.withUploaders; query != nil {
		if err := rq.loadUploaders(ctx, query, nodes, nil,
			func(n *Reviews, e *User) { n.Edges.Uploaders = e }); err != nil {
			return nil, err
		}
	}
	if query := rq.withThesisResult; query != nil {
		if err := rq.loadThesisResult(ctx, query, nodes, nil,
			func(n *Reviews, e *Thesis) { n.Edges.ThesisResult = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (rq *ReviewsQuery) loadUploaders(ctx context.Context, query *UserQuery, nodes []*Reviews, init func(*Reviews), assign func(*Reviews, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Reviews)
	for i := range nodes {
		if nodes[i].user_reviews == nil {
			continue
		}
		fk := *nodes[i].user_reviews
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
			return fmt.Errorf(`unexpected foreign-key "user_reviews" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (rq *ReviewsQuery) loadThesisResult(ctx context.Context, query *ThesisQuery, nodes []*Reviews, init func(*Reviews), assign func(*Reviews, *Thesis)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Reviews)
	for i := range nodes {
		if nodes[i].thesis_reviews == nil {
			continue
		}
		fk := *nodes[i].thesis_reviews
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(thesis.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "thesis_reviews" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (rq *ReviewsQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rq.querySpec()
	_spec.Node.Columns = rq.ctx.Fields
	if len(rq.ctx.Fields) > 0 {
		_spec.Unique = rq.ctx.Unique != nil && *rq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, rq.driver, _spec)
}

func (rq *ReviewsQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(reviews.Table, reviews.Columns, sqlgraph.NewFieldSpec(reviews.FieldID, field.TypeInt))
	_spec.From = rq.sql
	if unique := rq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if rq.path != nil {
		_spec.Unique = true
	}
	if fields := rq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, reviews.FieldID)
		for i := range fields {
			if fields[i] != reviews.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := rq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := rq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := rq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := rq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (rq *ReviewsQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(rq.driver.Dialect())
	t1 := builder.Table(reviews.Table)
	columns := rq.ctx.Fields
	if len(columns) == 0 {
		columns = reviews.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if rq.sql != nil {
		selector = rq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if rq.ctx.Unique != nil && *rq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range rq.predicates {
		p(selector)
	}
	for _, p := range rq.order {
		p(selector)
	}
	if offset := rq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := rq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ReviewsGroupBy is the group-by builder for Reviews entities.
type ReviewsGroupBy struct {
	selector
	build *ReviewsQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rgb *ReviewsGroupBy) Aggregate(fns ...AggregateFunc) *ReviewsGroupBy {
	rgb.fns = append(rgb.fns, fns...)
	return rgb
}

// Scan applies the selector query and scans the result into the given value.
func (rgb *ReviewsGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rgb.build.ctx, "GroupBy")
	if err := rgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ReviewsQuery, *ReviewsGroupBy](ctx, rgb.build, rgb, rgb.build.inters, v)
}

func (rgb *ReviewsGroupBy) sqlScan(ctx context.Context, root *ReviewsQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(rgb.fns))
	for _, fn := range rgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*rgb.flds)+len(rgb.fns))
		for _, f := range *rgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*rgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ReviewsSelect is the builder for selecting fields of Reviews entities.
type ReviewsSelect struct {
	*ReviewsQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (rs *ReviewsSelect) Aggregate(fns ...AggregateFunc) *ReviewsSelect {
	rs.fns = append(rs.fns, fns...)
	return rs
}

// Scan applies the selector query and scans the result into the given value.
func (rs *ReviewsSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rs.ctx, "Select")
	if err := rs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ReviewsQuery, *ReviewsSelect](ctx, rs.ReviewsQuery, rs, rs.inters, v)
}

func (rs *ReviewsSelect) sqlScan(ctx context.Context, root *ReviewsQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(rs.fns))
	for _, fn := range rs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*rs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
