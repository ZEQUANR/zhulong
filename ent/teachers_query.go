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
	"github.com/ZEQUANR/zhulong/ent/teachers"
	"github.com/ZEQUANR/zhulong/ent/user"
)

// TeachersQuery is the builder for querying Teachers entities.
type TeachersQuery struct {
	config
	ctx        *QueryContext
	order      []teachers.OrderOption
	inters     []Interceptor
	predicates []predicate.Teachers
	withUsers  *UserQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TeachersQuery builder.
func (tq *TeachersQuery) Where(ps ...predicate.Teachers) *TeachersQuery {
	tq.predicates = append(tq.predicates, ps...)
	return tq
}

// Limit the number of records to be returned by this query.
func (tq *TeachersQuery) Limit(limit int) *TeachersQuery {
	tq.ctx.Limit = &limit
	return tq
}

// Offset to start from.
func (tq *TeachersQuery) Offset(offset int) *TeachersQuery {
	tq.ctx.Offset = &offset
	return tq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tq *TeachersQuery) Unique(unique bool) *TeachersQuery {
	tq.ctx.Unique = &unique
	return tq
}

// Order specifies how the records should be ordered.
func (tq *TeachersQuery) Order(o ...teachers.OrderOption) *TeachersQuery {
	tq.order = append(tq.order, o...)
	return tq
}

// QueryUsers chains the current query on the "users" edge.
func (tq *TeachersQuery) QueryUsers() *UserQuery {
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
			sqlgraph.From(teachers.Table, teachers.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, teachers.UsersTable, teachers.UsersColumn),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Teachers entity from the query.
// Returns a *NotFoundError when no Teachers was found.
func (tq *TeachersQuery) First(ctx context.Context) (*Teachers, error) {
	nodes, err := tq.Limit(1).All(setContextOp(ctx, tq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{teachers.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tq *TeachersQuery) FirstX(ctx context.Context) *Teachers {
	node, err := tq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Teachers ID from the query.
// Returns a *NotFoundError when no Teachers ID was found.
func (tq *TeachersQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tq.Limit(1).IDs(setContextOp(ctx, tq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{teachers.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tq *TeachersQuery) FirstIDX(ctx context.Context) int {
	id, err := tq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Teachers entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Teachers entity is found.
// Returns a *NotFoundError when no Teachers entities are found.
func (tq *TeachersQuery) Only(ctx context.Context) (*Teachers, error) {
	nodes, err := tq.Limit(2).All(setContextOp(ctx, tq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{teachers.Label}
	default:
		return nil, &NotSingularError{teachers.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tq *TeachersQuery) OnlyX(ctx context.Context) *Teachers {
	node, err := tq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Teachers ID in the query.
// Returns a *NotSingularError when more than one Teachers ID is found.
// Returns a *NotFoundError when no entities are found.
func (tq *TeachersQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tq.Limit(2).IDs(setContextOp(ctx, tq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{teachers.Label}
	default:
		err = &NotSingularError{teachers.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tq *TeachersQuery) OnlyIDX(ctx context.Context) int {
	id, err := tq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TeachersSlice.
func (tq *TeachersQuery) All(ctx context.Context) ([]*Teachers, error) {
	ctx = setContextOp(ctx, tq.ctx, "All")
	if err := tq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Teachers, *TeachersQuery]()
	return withInterceptors[[]*Teachers](ctx, tq, qr, tq.inters)
}

// AllX is like All, but panics if an error occurs.
func (tq *TeachersQuery) AllX(ctx context.Context) []*Teachers {
	nodes, err := tq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Teachers IDs.
func (tq *TeachersQuery) IDs(ctx context.Context) (ids []int, err error) {
	if tq.ctx.Unique == nil && tq.path != nil {
		tq.Unique(true)
	}
	ctx = setContextOp(ctx, tq.ctx, "IDs")
	if err = tq.Select(teachers.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tq *TeachersQuery) IDsX(ctx context.Context) []int {
	ids, err := tq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tq *TeachersQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, tq.ctx, "Count")
	if err := tq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, tq, querierCount[*TeachersQuery](), tq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (tq *TeachersQuery) CountX(ctx context.Context) int {
	count, err := tq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tq *TeachersQuery) Exist(ctx context.Context) (bool, error) {
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
func (tq *TeachersQuery) ExistX(ctx context.Context) bool {
	exist, err := tq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TeachersQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tq *TeachersQuery) Clone() *TeachersQuery {
	if tq == nil {
		return nil
	}
	return &TeachersQuery{
		config:     tq.config,
		ctx:        tq.ctx.Clone(),
		order:      append([]teachers.OrderOption{}, tq.order...),
		inters:     append([]Interceptor{}, tq.inters...),
		predicates: append([]predicate.Teachers{}, tq.predicates...),
		withUsers:  tq.withUsers.Clone(),
		// clone intermediate query.
		sql:  tq.sql.Clone(),
		path: tq.path,
	}
}

// WithUsers tells the query-builder to eager-load the nodes that are connected to
// the "users" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *TeachersQuery) WithUsers(opts ...func(*UserQuery)) *TeachersQuery {
	query := (&UserClient{config: tq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tq.withUsers = query
	return tq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Avatar string `json:"avatar,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Teachers.Query().
//		GroupBy(teachers.FieldAvatar).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (tq *TeachersQuery) GroupBy(field string, fields ...string) *TeachersGroupBy {
	tq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &TeachersGroupBy{build: tq}
	grbuild.flds = &tq.ctx.Fields
	grbuild.label = teachers.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Avatar string `json:"avatar,omitempty"`
//	}
//
//	client.Teachers.Query().
//		Select(teachers.FieldAvatar).
//		Scan(ctx, &v)
func (tq *TeachersQuery) Select(fields ...string) *TeachersSelect {
	tq.ctx.Fields = append(tq.ctx.Fields, fields...)
	sbuild := &TeachersSelect{TeachersQuery: tq}
	sbuild.label = teachers.Label
	sbuild.flds, sbuild.scan = &tq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a TeachersSelect configured with the given aggregations.
func (tq *TeachersQuery) Aggregate(fns ...AggregateFunc) *TeachersSelect {
	return tq.Select().Aggregate(fns...)
}

func (tq *TeachersQuery) prepareQuery(ctx context.Context) error {
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
		if !teachers.ValidColumn(f) {
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

func (tq *TeachersQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Teachers, error) {
	var (
		nodes       = []*Teachers{}
		withFKs     = tq.withFKs
		_spec       = tq.querySpec()
		loadedTypes = [1]bool{
			tq.withUsers != nil,
		}
	)
	if tq.withUsers != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, teachers.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Teachers).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Teachers{config: tq.config}
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
	if query := tq.withUsers; query != nil {
		if err := tq.loadUsers(ctx, query, nodes, nil,
			func(n *Teachers, e *User) { n.Edges.Users = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (tq *TeachersQuery) loadUsers(ctx context.Context, query *UserQuery, nodes []*Teachers, init func(*Teachers), assign func(*Teachers, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Teachers)
	for i := range nodes {
		if nodes[i].user_teachers == nil {
			continue
		}
		fk := *nodes[i].user_teachers
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
			return fmt.Errorf(`unexpected foreign-key "user_teachers" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (tq *TeachersQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tq.querySpec()
	_spec.Node.Columns = tq.ctx.Fields
	if len(tq.ctx.Fields) > 0 {
		_spec.Unique = tq.ctx.Unique != nil && *tq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, tq.driver, _spec)
}

func (tq *TeachersQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(teachers.Table, teachers.Columns, sqlgraph.NewFieldSpec(teachers.FieldID, field.TypeInt))
	_spec.From = tq.sql
	if unique := tq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if tq.path != nil {
		_spec.Unique = true
	}
	if fields := tq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, teachers.FieldID)
		for i := range fields {
			if fields[i] != teachers.FieldID {
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

func (tq *TeachersQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tq.driver.Dialect())
	t1 := builder.Table(teachers.Table)
	columns := tq.ctx.Fields
	if len(columns) == 0 {
		columns = teachers.Columns
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

// TeachersGroupBy is the group-by builder for Teachers entities.
type TeachersGroupBy struct {
	selector
	build *TeachersQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tgb *TeachersGroupBy) Aggregate(fns ...AggregateFunc) *TeachersGroupBy {
	tgb.fns = append(tgb.fns, fns...)
	return tgb
}

// Scan applies the selector query and scans the result into the given value.
func (tgb *TeachersGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tgb.build.ctx, "GroupBy")
	if err := tgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TeachersQuery, *TeachersGroupBy](ctx, tgb.build, tgb, tgb.build.inters, v)
}

func (tgb *TeachersGroupBy) sqlScan(ctx context.Context, root *TeachersQuery, v any) error {
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

// TeachersSelect is the builder for selecting fields of Teachers entities.
type TeachersSelect struct {
	*TeachersQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ts *TeachersSelect) Aggregate(fns ...AggregateFunc) *TeachersSelect {
	ts.fns = append(ts.fns, fns...)
	return ts
}

// Scan applies the selector query and scans the result into the given value.
func (ts *TeachersSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ts.ctx, "Select")
	if err := ts.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TeachersQuery, *TeachersSelect](ctx, ts.TeachersQuery, ts, ts.inters, v)
}

func (ts *TeachersSelect) sqlScan(ctx context.Context, root *TeachersQuery, v any) error {
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
