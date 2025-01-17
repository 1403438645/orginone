// Code generated by entc, DO NOT EDIT.

package schema

import (
	"context"
	"errors"
	"fmt"
	"math"
	"orginone/common/schema/asallgroup"
	"orginone/common/schema/aslayer"
	"orginone/common/schema/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AsLayerQuery is the builder for querying AsLayer entities.
type AsLayerQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.AsLayer
	// eager-loading edges.
	withGroup *AsAllGroupQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AsLayerQuery builder.
func (alq *AsLayerQuery) Where(ps ...predicate.AsLayer) *AsLayerQuery {
	alq.predicates = append(alq.predicates, ps...)
	return alq
}

// Limit adds a limit step to the query.
func (alq *AsLayerQuery) Limit(limit int) *AsLayerQuery {
	alq.limit = &limit
	return alq
}

// Offset adds an offset step to the query.
func (alq *AsLayerQuery) Offset(offset int) *AsLayerQuery {
	alq.offset = &offset
	return alq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (alq *AsLayerQuery) Unique(unique bool) *AsLayerQuery {
	alq.unique = &unique
	return alq
}

// Order adds an order step to the query.
func (alq *AsLayerQuery) Order(o ...OrderFunc) *AsLayerQuery {
	alq.order = append(alq.order, o...)
	return alq
}

// QueryGroup chains the current query on the "group" edge.
func (alq *AsLayerQuery) QueryGroup() *AsAllGroupQuery {
	query := &AsAllGroupQuery{config: alq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := alq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := alq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(aslayer.Table, aslayer.FieldID, selector),
			sqlgraph.To(asallgroup.Table, asallgroup.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, aslayer.GroupTable, aslayer.GroupColumn),
		)
		fromU = sqlgraph.SetNeighbors(alq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first AsLayer entity from the query.
// Returns a *NotFoundError when no AsLayer was found.
func (alq *AsLayerQuery) First(ctx context.Context) (*AsLayer, error) {
	nodes, err := alq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{aslayer.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (alq *AsLayerQuery) FirstX(ctx context.Context) *AsLayer {
	node, err := alq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AsLayer ID from the query.
// Returns a *NotFoundError when no AsLayer ID was found.
func (alq *AsLayerQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = alq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{aslayer.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (alq *AsLayerQuery) FirstIDX(ctx context.Context) int64 {
	id, err := alq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AsLayer entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AsLayer entity is found.
// Returns a *NotFoundError when no AsLayer entities are found.
func (alq *AsLayerQuery) Only(ctx context.Context) (*AsLayer, error) {
	nodes, err := alq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{aslayer.Label}
	default:
		return nil, &NotSingularError{aslayer.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (alq *AsLayerQuery) OnlyX(ctx context.Context) *AsLayer {
	node, err := alq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AsLayer ID in the query.
// Returns a *NotSingularError when more than one AsLayer ID is found.
// Returns a *NotFoundError when no entities are found.
func (alq *AsLayerQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = alq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{aslayer.Label}
	default:
		err = &NotSingularError{aslayer.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (alq *AsLayerQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := alq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AsLayers.
func (alq *AsLayerQuery) All(ctx context.Context) ([]*AsLayer, error) {
	if err := alq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return alq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (alq *AsLayerQuery) AllX(ctx context.Context) []*AsLayer {
	nodes, err := alq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AsLayer IDs.
func (alq *AsLayerQuery) IDs(ctx context.Context) ([]int64, error) {
	var ids []int64
	if err := alq.Select(aslayer.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (alq *AsLayerQuery) IDsX(ctx context.Context) []int64 {
	ids, err := alq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (alq *AsLayerQuery) Count(ctx context.Context) (int64, error) {
	if err := alq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return alq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (alq *AsLayerQuery) CountX(ctx context.Context) int64 {
	count, err := alq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (alq *AsLayerQuery) Exist(ctx context.Context) (bool, error) {
	if err := alq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return alq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (alq *AsLayerQuery) ExistX(ctx context.Context) bool {
	exist, err := alq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AsLayerQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (alq *AsLayerQuery) Clone() *AsLayerQuery {
	if alq == nil {
		return nil
	}
	return &AsLayerQuery{
		config:     alq.config,
		limit:      alq.limit,
		offset:     alq.offset,
		order:      append([]OrderFunc{}, alq.order...),
		predicates: append([]predicate.AsLayer{}, alq.predicates...),
		withGroup:  alq.withGroup.Clone(),
		// clone intermediate query.
		sql:    alq.sql.Clone(),
		path:   alq.path,
		unique: alq.unique,
	}
}

// WithGroup tells the query-builder to eager-load the nodes that are connected to
// the "group" edge. The optional arguments are used to configure the query builder of the edge.
func (alq *AsLayerQuery) WithGroup(opts ...func(*AsAllGroupQuery)) *AsLayerQuery {
	query := &AsAllGroupQuery{config: alq.config}
	for _, opt := range opts {
		opt(query)
	}
	alq.withGroup = query
	return alq
}

// ThenGroup tells the query-builder to eager-load the nodes that are connected to
// the "group" edge. The optional arguments are used to configure the query builder of the edge.
func (alq *AsLayerQuery) ThenGroup(opts ...func(*AsAllGroupQuery)) *AsLayerQuery {
	query := &AsAllGroupQuery{config: alq.config}
	for _, opt := range opts {
		opt(query.Where(asallgroup.IsDeleted(0)))
	}
	alq.withGroup = query
	return alq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Layer int64 `json:"layer"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.AsLayer.Query().
//		GroupBy(aslayer.FieldLayer).
//		Aggregate(schema.Count()).
//		Scan(ctx, &v)
//
func (alq *AsLayerQuery) GroupBy(field string, fields ...string) *AsLayerGroupBy {
	group := &AsLayerGroupBy{config: alq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := alq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return alq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Layer int64 `json:"layer"`
//	}
//
//	client.AsLayer.Query().
//		Select(aslayer.FieldLayer).
//		Scan(ctx, &v)
//
func (alq *AsLayerQuery) Select(fields ...string) *AsLayerSelect {
	alq.fields = append(alq.fields, fields...)
	return &AsLayerSelect{AsLayerQuery: alq}
}

func (alq *AsLayerQuery) prepareQuery(ctx context.Context) error {
	for _, f := range alq.fields {
		if !aslayer.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("schema: invalid field %q for query", f)}
		}
	}
	if alq.path != nil {
		prev, err := alq.path(ctx)
		if err != nil {
			return err
		}
		alq.sql = prev
	}
	return nil
}

func (alq *AsLayerQuery) sqlAll(ctx context.Context) ([]*AsLayer, error) {
	var (
		nodes       = []*AsLayer{}
		_spec       = alq.querySpec()
		loadedTypes = [1]bool{
			alq.withGroup != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &AsLayer{config: alq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("schema: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, alq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := alq.withGroup; query != nil {
		ids := make([]int64, 0, len(nodes))
		nodeids := make(map[int64][]*AsLayer)
		for i := range nodes {
			fk := nodes[i].GroupID
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(asallgroup.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "group_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Group = n
			}
		}
	}

	return nodes, nil
}

func (alq *AsLayerQuery) sqlCount(ctx context.Context) (int64, error) {
	_spec := alq.querySpec()
	_spec.Node.Columns = alq.fields
	if len(alq.fields) > 0 {
		_spec.Unique = alq.unique != nil && *alq.unique
	}
	c, err := sqlgraph.CountNodes(ctx, alq.driver, _spec)
	return int64(c), err
}

func (alq *AsLayerQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := alq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("schema: check existence: %w", err)
	}
	return n > 0, nil
}

func (alq *AsLayerQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   aslayer.Table,
			Columns: aslayer.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: aslayer.FieldID,
			},
		},
		From:   alq.sql,
		Unique: true,
	}
	if unique := alq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := alq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, aslayer.FieldID)
		for i := range fields {
			if fields[i] != aslayer.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := alq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := alq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := alq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := alq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (alq *AsLayerQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(alq.driver.Dialect())
	t1 := builder.Table(aslayer.Table)
	columns := alq.fields
	if len(columns) == 0 {
		columns = aslayer.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if alq.sql != nil {
		selector = alq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if alq.unique != nil && *alq.unique {
		selector.Distinct()
	}
	for _, p := range alq.predicates {
		p(selector)
	}
	for _, p := range alq.order {
		p(selector)
	}
	if offset := alq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := alq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AsLayerGroupBy is the group-by builder for AsLayer entities.
type AsLayerGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (algb *AsLayerGroupBy) Aggregate(fns ...AggregateFunc) *AsLayerGroupBy {
	algb.fns = append(algb.fns, fns...)
	return algb
}

// Scan applies the group-by query and scans the result into the given value.
func (algb *AsLayerGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := algb.path(ctx)
	if err != nil {
		return err
	}
	algb.sql = query
	return algb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (algb *AsLayerGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := algb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (algb *AsLayerGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(algb.fields) > 1 {
		return nil, errors.New("schema: AsLayerGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := algb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (algb *AsLayerGroupBy) StringsX(ctx context.Context) []string {
	v, err := algb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (algb *AsLayerGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = algb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{aslayer.Label}
	default:
		err = fmt.Errorf("schema: AsLayerGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (algb *AsLayerGroupBy) StringX(ctx context.Context) string {
	v, err := algb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (algb *AsLayerGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(algb.fields) > 1 {
		return nil, errors.New("schema: AsLayerGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := algb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (algb *AsLayerGroupBy) IntsX(ctx context.Context) []int {
	v, err := algb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (algb *AsLayerGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = algb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{aslayer.Label}
	default:
		err = fmt.Errorf("schema: AsLayerGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (algb *AsLayerGroupBy) IntX(ctx context.Context) int {
	v, err := algb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (algb *AsLayerGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(algb.fields) > 1 {
		return nil, errors.New("schema: AsLayerGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := algb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (algb *AsLayerGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := algb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (algb *AsLayerGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = algb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{aslayer.Label}
	default:
		err = fmt.Errorf("schema: AsLayerGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (algb *AsLayerGroupBy) Float64X(ctx context.Context) float64 {
	v, err := algb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (algb *AsLayerGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(algb.fields) > 1 {
		return nil, errors.New("schema: AsLayerGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := algb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (algb *AsLayerGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := algb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (algb *AsLayerGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = algb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{aslayer.Label}
	default:
		err = fmt.Errorf("schema: AsLayerGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (algb *AsLayerGroupBy) BoolX(ctx context.Context) bool {
	v, err := algb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int64s returns list of int64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (algb *AsLayerGroupBy) Int64s(ctx context.Context) ([]int64, error) {
	if len(algb.fields) > 1 {
		return nil, errors.New("schema: AsLayerGroupBy.Int64s is not achievable when grouping more than 1 field")
	}
	var v []int64
	if err := algb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Int64sX is like Int64s, but panics if an error occurs.
func (algb *AsLayerGroupBy) Int64sX(ctx context.Context) []int64 {
	v, err := algb.Int64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int64 returns a single int64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (algb *AsLayerGroupBy) Int64(ctx context.Context) (_ int64, err error) {
	var v []int64
	if v, err = algb.Int64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{aslayer.Label}
	default:
		err = fmt.Errorf("schema: AsLayerGroupBy.Int64s returned %d results when one was expected", len(v))
	}
	return
}

// Int64X is like Int64, but panics if an error occurs.
func (algb *AsLayerGroupBy) Int64X(ctx context.Context) int64 {
	v, err := algb.Int64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (algb *AsLayerGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range algb.fields {
		if !aslayer.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := algb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := algb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (algb *AsLayerGroupBy) sqlQuery() *sql.Selector {
	selector := algb.sql.Select()
	aggregation := make([]string, 0, len(algb.fns))
	for _, fn := range algb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(algb.fields)+len(algb.fns))
		for _, f := range algb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(algb.fields...)...)
}

// AsLayerSelect is the builder for selecting fields of AsLayer entities.
type AsLayerSelect struct {
	*AsLayerQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (als *AsLayerSelect) Scan(ctx context.Context, v interface{}) error {
	if err := als.prepareQuery(ctx); err != nil {
		return err
	}
	als.sql = als.AsLayerQuery.sqlQuery(ctx)
	return als.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (als *AsLayerSelect) ScanX(ctx context.Context, v interface{}) {
	if err := als.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (als *AsLayerSelect) Strings(ctx context.Context) ([]string, error) {
	if len(als.fields) > 1 {
		return nil, errors.New("schema: AsLayerSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := als.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (als *AsLayerSelect) StringsX(ctx context.Context) []string {
	v, err := als.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (als *AsLayerSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = als.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{aslayer.Label}
	default:
		err = fmt.Errorf("schema: AsLayerSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (als *AsLayerSelect) StringX(ctx context.Context) string {
	v, err := als.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (als *AsLayerSelect) Ints(ctx context.Context) ([]int, error) {
	if len(als.fields) > 1 {
		return nil, errors.New("schema: AsLayerSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := als.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (als *AsLayerSelect) IntsX(ctx context.Context) []int {
	v, err := als.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (als *AsLayerSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = als.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{aslayer.Label}
	default:
		err = fmt.Errorf("schema: AsLayerSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (als *AsLayerSelect) IntX(ctx context.Context) int {
	v, err := als.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (als *AsLayerSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(als.fields) > 1 {
		return nil, errors.New("schema: AsLayerSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := als.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (als *AsLayerSelect) Float64sX(ctx context.Context) []float64 {
	v, err := als.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (als *AsLayerSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = als.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{aslayer.Label}
	default:
		err = fmt.Errorf("schema: AsLayerSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (als *AsLayerSelect) Float64X(ctx context.Context) float64 {
	v, err := als.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (als *AsLayerSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(als.fields) > 1 {
		return nil, errors.New("schema: AsLayerSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := als.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (als *AsLayerSelect) BoolsX(ctx context.Context) []bool {
	v, err := als.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (als *AsLayerSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = als.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{aslayer.Label}
	default:
		err = fmt.Errorf("schema: AsLayerSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (als *AsLayerSelect) BoolX(ctx context.Context) bool {
	v, err := als.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int64s returns list of int64s from a selector. It is only allowed when selecting one field.
func (als *AsLayerSelect) Int64s(ctx context.Context) ([]int64, error) {
	if len(als.fields) > 1 {
		return nil, errors.New("schema: AsLayerSelect.Int64s is not achievable when selecting more than 1 field")
	}
	var v []int64
	if err := als.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Int64sX is like Int64s, but panics if an error occurs.
func (als *AsLayerSelect) Int64sX(ctx context.Context) []int64 {
	v, err := als.Int64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int64 returns a single int64 from a selector. It is only allowed when selecting one field.
func (als *AsLayerSelect) Int64(ctx context.Context) (_ int64, err error) {
	var v []int64
	if v, err = als.Int64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{aslayer.Label}
	default:
		err = fmt.Errorf("schema: AsLayerSelect.Int64s returned %d results when one was expected", len(v))
	}
	return
}

// Int64X is like Int64, but panics if an error occurs.
func (als *AsLayerSelect) Int64X(ctx context.Context) int64 {
	v, err := als.Int64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (als *AsLayerSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := als.sql.Query()
	if err := als.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
