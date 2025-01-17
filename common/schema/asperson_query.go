// Code generated by entc, DO NOT EDIT.

package schema

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"
	"orginone/common/schema/asinneragency"
	"orginone/common/schema/asjob"
	"orginone/common/schema/asperson"
	"orginone/common/schema/asuser"
	"orginone/common/schema/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AsPersonQuery is the builder for querying AsPerson entities.
type AsPersonQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.AsPerson
	// eager-loading edges.
	withUserx   *AsUserQuery
	withAgencys *AsInnerAgencyQuery
	withJobs    *AsJobQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AsPersonQuery builder.
func (apq *AsPersonQuery) Where(ps ...predicate.AsPerson) *AsPersonQuery {
	apq.predicates = append(apq.predicates, ps...)
	return apq
}

// Limit adds a limit step to the query.
func (apq *AsPersonQuery) Limit(limit int) *AsPersonQuery {
	apq.limit = &limit
	return apq
}

// Offset adds an offset step to the query.
func (apq *AsPersonQuery) Offset(offset int) *AsPersonQuery {
	apq.offset = &offset
	return apq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (apq *AsPersonQuery) Unique(unique bool) *AsPersonQuery {
	apq.unique = &unique
	return apq
}

// Order adds an order step to the query.
func (apq *AsPersonQuery) Order(o ...OrderFunc) *AsPersonQuery {
	apq.order = append(apq.order, o...)
	return apq
}

// QueryUserx chains the current query on the "userx" edge.
func (apq *AsPersonQuery) QueryUserx() *AsUserQuery {
	query := &AsUserQuery{config: apq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := apq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := apq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(asperson.Table, asperson.FieldID, selector),
			sqlgraph.To(asuser.Table, asuser.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, asperson.UserxTable, asperson.UserxColumn),
		)
		fromU = sqlgraph.SetNeighbors(apq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAgencys chains the current query on the "agencys" edge.
func (apq *AsPersonQuery) QueryAgencys() *AsInnerAgencyQuery {
	query := &AsInnerAgencyQuery{config: apq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := apq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := apq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(asperson.Table, asperson.FieldID, selector),
			sqlgraph.To(asinneragency.Table, asinneragency.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, asperson.AgencysTable, asperson.AgencysPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(apq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryJobs chains the current query on the "jobs" edge.
func (apq *AsPersonQuery) QueryJobs() *AsJobQuery {
	query := &AsJobQuery{config: apq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := apq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := apq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(asperson.Table, asperson.FieldID, selector),
			sqlgraph.To(asjob.Table, asjob.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, asperson.JobsTable, asperson.JobsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(apq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first AsPerson entity from the query.
// Returns a *NotFoundError when no AsPerson was found.
func (apq *AsPersonQuery) First(ctx context.Context) (*AsPerson, error) {
	nodes, err := apq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{asperson.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (apq *AsPersonQuery) FirstX(ctx context.Context) *AsPerson {
	node, err := apq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AsPerson ID from the query.
// Returns a *NotFoundError when no AsPerson ID was found.
func (apq *AsPersonQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = apq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{asperson.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (apq *AsPersonQuery) FirstIDX(ctx context.Context) int64 {
	id, err := apq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AsPerson entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AsPerson entity is found.
// Returns a *NotFoundError when no AsPerson entities are found.
func (apq *AsPersonQuery) Only(ctx context.Context) (*AsPerson, error) {
	nodes, err := apq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{asperson.Label}
	default:
		return nil, &NotSingularError{asperson.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (apq *AsPersonQuery) OnlyX(ctx context.Context) *AsPerson {
	node, err := apq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AsPerson ID in the query.
// Returns a *NotSingularError when more than one AsPerson ID is found.
// Returns a *NotFoundError when no entities are found.
func (apq *AsPersonQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = apq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{asperson.Label}
	default:
		err = &NotSingularError{asperson.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (apq *AsPersonQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := apq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AsPersons.
func (apq *AsPersonQuery) All(ctx context.Context) ([]*AsPerson, error) {
	if err := apq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return apq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (apq *AsPersonQuery) AllX(ctx context.Context) []*AsPerson {
	nodes, err := apq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AsPerson IDs.
func (apq *AsPersonQuery) IDs(ctx context.Context) ([]int64, error) {
	var ids []int64
	if err := apq.Select(asperson.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (apq *AsPersonQuery) IDsX(ctx context.Context) []int64 {
	ids, err := apq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (apq *AsPersonQuery) Count(ctx context.Context) (int64, error) {
	if err := apq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return apq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (apq *AsPersonQuery) CountX(ctx context.Context) int64 {
	count, err := apq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (apq *AsPersonQuery) Exist(ctx context.Context) (bool, error) {
	if err := apq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return apq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (apq *AsPersonQuery) ExistX(ctx context.Context) bool {
	exist, err := apq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AsPersonQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (apq *AsPersonQuery) Clone() *AsPersonQuery {
	if apq == nil {
		return nil
	}
	return &AsPersonQuery{
		config:      apq.config,
		limit:       apq.limit,
		offset:      apq.offset,
		order:       append([]OrderFunc{}, apq.order...),
		predicates:  append([]predicate.AsPerson{}, apq.predicates...),
		withUserx:   apq.withUserx.Clone(),
		withAgencys: apq.withAgencys.Clone(),
		withJobs:    apq.withJobs.Clone(),
		// clone intermediate query.
		sql:    apq.sql.Clone(),
		path:   apq.path,
		unique: apq.unique,
	}
}

// WithUserx tells the query-builder to eager-load the nodes that are connected to
// the "userx" edge. The optional arguments are used to configure the query builder of the edge.
func (apq *AsPersonQuery) WithUserx(opts ...func(*AsUserQuery)) *AsPersonQuery {
	query := &AsUserQuery{config: apq.config}
	for _, opt := range opts {
		opt(query)
	}
	apq.withUserx = query
	return apq
}

// WithAgencys tells the query-builder to eager-load the nodes that are connected to
// the "agencys" edge. The optional arguments are used to configure the query builder of the edge.
func (apq *AsPersonQuery) WithAgencys(opts ...func(*AsInnerAgencyQuery)) *AsPersonQuery {
	query := &AsInnerAgencyQuery{config: apq.config}
	for _, opt := range opts {
		opt(query)
	}
	apq.withAgencys = query
	return apq
}

// WithJobs tells the query-builder to eager-load the nodes that are connected to
// the "jobs" edge. The optional arguments are used to configure the query builder of the edge.
func (apq *AsPersonQuery) WithJobs(opts ...func(*AsJobQuery)) *AsPersonQuery {
	query := &AsJobQuery{config: apq.config}
	for _, opt := range opts {
		opt(query)
	}
	apq.withJobs = query
	return apq
}

// ThenUserx tells the query-builder to eager-load the nodes that are connected to
// the "userx" edge. The optional arguments are used to configure the query builder of the edge.
func (apq *AsPersonQuery) ThenUserx(opts ...func(*AsUserQuery)) *AsPersonQuery {
	query := &AsUserQuery{config: apq.config}
	for _, opt := range opts {
		opt(query.Where(asuser.IsDeleted(0)))
	}
	apq.withUserx = query
	return apq
}

// ThenAgencys tells the query-builder to eager-load the nodes that are connected to
// the "agencys" edge. The optional arguments are used to configure the query builder of the edge.
func (apq *AsPersonQuery) ThenAgencys(opts ...func(*AsInnerAgencyQuery)) *AsPersonQuery {
	query := &AsInnerAgencyQuery{config: apq.config}
	for _, opt := range opts {
		opt(query.Where(asinneragency.IsDeleted(0)))
	}
	apq.withAgencys = query
	return apq
}

// ThenJobs tells the query-builder to eager-load the nodes that are connected to
// the "jobs" edge. The optional arguments are used to configure the query builder of the edge.
func (apq *AsPersonQuery) ThenJobs(opts ...func(*AsJobQuery)) *AsPersonQuery {
	query := &AsJobQuery{config: apq.config}
	for _, opt := range opts {
		opt(query.Where(asjob.IsDeleted(0)))
	}
	apq.withJobs = query
	return apq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		TenantCode string `json:"tenantCode"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.AsPerson.Query().
//		GroupBy(asperson.FieldTenantCode).
//		Aggregate(schema.Count()).
//		Scan(ctx, &v)
//
func (apq *AsPersonQuery) GroupBy(field string, fields ...string) *AsPersonGroupBy {
	group := &AsPersonGroupBy{config: apq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := apq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return apq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		TenantCode string `json:"tenantCode"`
//	}
//
//	client.AsPerson.Query().
//		Select(asperson.FieldTenantCode).
//		Scan(ctx, &v)
//
func (apq *AsPersonQuery) Select(fields ...string) *AsPersonSelect {
	apq.fields = append(apq.fields, fields...)
	return &AsPersonSelect{AsPersonQuery: apq}
}

func (apq *AsPersonQuery) prepareQuery(ctx context.Context) error {
	for _, f := range apq.fields {
		if !asperson.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("schema: invalid field %q for query", f)}
		}
	}
	if apq.path != nil {
		prev, err := apq.path(ctx)
		if err != nil {
			return err
		}
		apq.sql = prev
	}
	return nil
}

func (apq *AsPersonQuery) sqlAll(ctx context.Context) ([]*AsPerson, error) {
	var (
		nodes       = []*AsPerson{}
		_spec       = apq.querySpec()
		loadedTypes = [3]bool{
			apq.withUserx != nil,
			apq.withAgencys != nil,
			apq.withJobs != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &AsPerson{config: apq.config}
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
	if err := sqlgraph.QueryNodes(ctx, apq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := apq.withUserx; query != nil {
		ids := make([]int64, 0, len(nodes))
		nodeids := make(map[int64][]*AsPerson)
		for i := range nodes {
			fk := nodes[i].UserID
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(asuser.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "user_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Userx = n
			}
		}
	}

	if query := apq.withAgencys; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		ids := make(map[int64]*AsPerson, len(nodes))
		for _, node := range nodes {
			ids[node.ID] = node
			fks = append(fks, node.ID)
			node.Edges.Agencys = []*AsInnerAgency{}
		}
		var (
			edgeids []int64
			edges   = make(map[int64][]*AsPerson)
		)
		_spec := &sqlgraph.EdgeQuerySpec{
			Edge: &sqlgraph.EdgeSpec{
				Inverse: true,
				Table:   asperson.AgencysTable,
				Columns: asperson.AgencysPrimaryKey,
			},
			Predicate: func(s *sql.Selector) {
				s.Where(sql.InValues(asperson.AgencysPrimaryKey[1], fks...))
			},
			ScanValues: func() [2]interface{} {
				return [2]interface{}{new(sql.NullInt64), new(sql.NullInt64)}
			},
			Assign: func(out, in interface{}) error {
				eout, ok := out.(*sql.NullInt64)
				if !ok || eout == nil {
					return fmt.Errorf("unexpected id value for edge-out")
				}
				ein, ok := in.(*sql.NullInt64)
				if !ok || ein == nil {
					return fmt.Errorf("unexpected id value for edge-in")
				}
				outValue := eout.Int64
				inValue := ein.Int64
				node, ok := ids[outValue]
				if !ok {
					return fmt.Errorf("unexpected node id in edges: %v", outValue)
				}
				if _, ok := edges[inValue]; !ok {
					edgeids = append(edgeids, inValue)
				}
				edges[inValue] = append(edges[inValue], node)
				return nil
			},
		}
		if err := sqlgraph.QueryEdges(ctx, apq.driver, _spec); err != nil {
			return nil, fmt.Errorf(`query edges "agencys": %w`, err)
		}
		query.Where(asinneragency.IDIn(edgeids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := edges[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "agencys" node returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Agencys = append(nodes[i].Edges.Agencys, n)
			}
		}
	}

	if query := apq.withJobs; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		ids := make(map[int64]*AsPerson, len(nodes))
		for _, node := range nodes {
			ids[node.ID] = node
			fks = append(fks, node.ID)
			node.Edges.Jobs = []*AsJob{}
		}
		var (
			edgeids []int64
			edges   = make(map[int64][]*AsPerson)
		)
		_spec := &sqlgraph.EdgeQuerySpec{
			Edge: &sqlgraph.EdgeSpec{
				Inverse: true,
				Table:   asperson.JobsTable,
				Columns: asperson.JobsPrimaryKey,
			},
			Predicate: func(s *sql.Selector) {
				s.Where(sql.InValues(asperson.JobsPrimaryKey[1], fks...))
			},
			ScanValues: func() [2]interface{} {
				return [2]interface{}{new(sql.NullInt64), new(sql.NullInt64)}
			},
			Assign: func(out, in interface{}) error {
				eout, ok := out.(*sql.NullInt64)
				if !ok || eout == nil {
					return fmt.Errorf("unexpected id value for edge-out")
				}
				ein, ok := in.(*sql.NullInt64)
				if !ok || ein == nil {
					return fmt.Errorf("unexpected id value for edge-in")
				}
				outValue := eout.Int64
				inValue := ein.Int64
				node, ok := ids[outValue]
				if !ok {
					return fmt.Errorf("unexpected node id in edges: %v", outValue)
				}
				if _, ok := edges[inValue]; !ok {
					edgeids = append(edgeids, inValue)
				}
				edges[inValue] = append(edges[inValue], node)
				return nil
			},
		}
		if err := sqlgraph.QueryEdges(ctx, apq.driver, _spec); err != nil {
			return nil, fmt.Errorf(`query edges "jobs": %w`, err)
		}
		query.Where(asjob.IDIn(edgeids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := edges[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "jobs" node returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Jobs = append(nodes[i].Edges.Jobs, n)
			}
		}
	}

	return nodes, nil
}

func (apq *AsPersonQuery) sqlCount(ctx context.Context) (int64, error) {
	_spec := apq.querySpec()
	_spec.Node.Columns = apq.fields
	if len(apq.fields) > 0 {
		_spec.Unique = apq.unique != nil && *apq.unique
	}
	c, err := sqlgraph.CountNodes(ctx, apq.driver, _spec)
	return int64(c), err
}

func (apq *AsPersonQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := apq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("schema: check existence: %w", err)
	}
	return n > 0, nil
}

func (apq *AsPersonQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   asperson.Table,
			Columns: asperson.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: asperson.FieldID,
			},
		},
		From:   apq.sql,
		Unique: true,
	}
	if unique := apq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := apq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, asperson.FieldID)
		for i := range fields {
			if fields[i] != asperson.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := apq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := apq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := apq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := apq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (apq *AsPersonQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(apq.driver.Dialect())
	t1 := builder.Table(asperson.Table)
	columns := apq.fields
	if len(columns) == 0 {
		columns = asperson.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if apq.sql != nil {
		selector = apq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if apq.unique != nil && *apq.unique {
		selector.Distinct()
	}
	for _, p := range apq.predicates {
		p(selector)
	}
	for _, p := range apq.order {
		p(selector)
	}
	if offset := apq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := apq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AsPersonGroupBy is the group-by builder for AsPerson entities.
type AsPersonGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (apgb *AsPersonGroupBy) Aggregate(fns ...AggregateFunc) *AsPersonGroupBy {
	apgb.fns = append(apgb.fns, fns...)
	return apgb
}

// Scan applies the group-by query and scans the result into the given value.
func (apgb *AsPersonGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := apgb.path(ctx)
	if err != nil {
		return err
	}
	apgb.sql = query
	return apgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (apgb *AsPersonGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := apgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (apgb *AsPersonGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(apgb.fields) > 1 {
		return nil, errors.New("schema: AsPersonGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := apgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (apgb *AsPersonGroupBy) StringsX(ctx context.Context) []string {
	v, err := apgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (apgb *AsPersonGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = apgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{asperson.Label}
	default:
		err = fmt.Errorf("schema: AsPersonGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (apgb *AsPersonGroupBy) StringX(ctx context.Context) string {
	v, err := apgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (apgb *AsPersonGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(apgb.fields) > 1 {
		return nil, errors.New("schema: AsPersonGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := apgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (apgb *AsPersonGroupBy) IntsX(ctx context.Context) []int {
	v, err := apgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (apgb *AsPersonGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = apgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{asperson.Label}
	default:
		err = fmt.Errorf("schema: AsPersonGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (apgb *AsPersonGroupBy) IntX(ctx context.Context) int {
	v, err := apgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (apgb *AsPersonGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(apgb.fields) > 1 {
		return nil, errors.New("schema: AsPersonGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := apgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (apgb *AsPersonGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := apgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (apgb *AsPersonGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = apgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{asperson.Label}
	default:
		err = fmt.Errorf("schema: AsPersonGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (apgb *AsPersonGroupBy) Float64X(ctx context.Context) float64 {
	v, err := apgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (apgb *AsPersonGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(apgb.fields) > 1 {
		return nil, errors.New("schema: AsPersonGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := apgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (apgb *AsPersonGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := apgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (apgb *AsPersonGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = apgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{asperson.Label}
	default:
		err = fmt.Errorf("schema: AsPersonGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (apgb *AsPersonGroupBy) BoolX(ctx context.Context) bool {
	v, err := apgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int64s returns list of int64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (apgb *AsPersonGroupBy) Int64s(ctx context.Context) ([]int64, error) {
	if len(apgb.fields) > 1 {
		return nil, errors.New("schema: AsPersonGroupBy.Int64s is not achievable when grouping more than 1 field")
	}
	var v []int64
	if err := apgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Int64sX is like Int64s, but panics if an error occurs.
func (apgb *AsPersonGroupBy) Int64sX(ctx context.Context) []int64 {
	v, err := apgb.Int64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int64 returns a single int64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (apgb *AsPersonGroupBy) Int64(ctx context.Context) (_ int64, err error) {
	var v []int64
	if v, err = apgb.Int64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{asperson.Label}
	default:
		err = fmt.Errorf("schema: AsPersonGroupBy.Int64s returned %d results when one was expected", len(v))
	}
	return
}

// Int64X is like Int64, but panics if an error occurs.
func (apgb *AsPersonGroupBy) Int64X(ctx context.Context) int64 {
	v, err := apgb.Int64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (apgb *AsPersonGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range apgb.fields {
		if !asperson.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := apgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := apgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (apgb *AsPersonGroupBy) sqlQuery() *sql.Selector {
	selector := apgb.sql.Select()
	aggregation := make([]string, 0, len(apgb.fns))
	for _, fn := range apgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(apgb.fields)+len(apgb.fns))
		for _, f := range apgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(apgb.fields...)...)
}

// AsPersonSelect is the builder for selecting fields of AsPerson entities.
type AsPersonSelect struct {
	*AsPersonQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (aps *AsPersonSelect) Scan(ctx context.Context, v interface{}) error {
	if err := aps.prepareQuery(ctx); err != nil {
		return err
	}
	aps.sql = aps.AsPersonQuery.sqlQuery(ctx)
	return aps.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (aps *AsPersonSelect) ScanX(ctx context.Context, v interface{}) {
	if err := aps.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (aps *AsPersonSelect) Strings(ctx context.Context) ([]string, error) {
	if len(aps.fields) > 1 {
		return nil, errors.New("schema: AsPersonSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := aps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (aps *AsPersonSelect) StringsX(ctx context.Context) []string {
	v, err := aps.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (aps *AsPersonSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = aps.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{asperson.Label}
	default:
		err = fmt.Errorf("schema: AsPersonSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (aps *AsPersonSelect) StringX(ctx context.Context) string {
	v, err := aps.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (aps *AsPersonSelect) Ints(ctx context.Context) ([]int, error) {
	if len(aps.fields) > 1 {
		return nil, errors.New("schema: AsPersonSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := aps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (aps *AsPersonSelect) IntsX(ctx context.Context) []int {
	v, err := aps.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (aps *AsPersonSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = aps.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{asperson.Label}
	default:
		err = fmt.Errorf("schema: AsPersonSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (aps *AsPersonSelect) IntX(ctx context.Context) int {
	v, err := aps.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (aps *AsPersonSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(aps.fields) > 1 {
		return nil, errors.New("schema: AsPersonSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := aps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (aps *AsPersonSelect) Float64sX(ctx context.Context) []float64 {
	v, err := aps.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (aps *AsPersonSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = aps.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{asperson.Label}
	default:
		err = fmt.Errorf("schema: AsPersonSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (aps *AsPersonSelect) Float64X(ctx context.Context) float64 {
	v, err := aps.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (aps *AsPersonSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(aps.fields) > 1 {
		return nil, errors.New("schema: AsPersonSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := aps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (aps *AsPersonSelect) BoolsX(ctx context.Context) []bool {
	v, err := aps.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (aps *AsPersonSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = aps.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{asperson.Label}
	default:
		err = fmt.Errorf("schema: AsPersonSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (aps *AsPersonSelect) BoolX(ctx context.Context) bool {
	v, err := aps.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int64s returns list of int64s from a selector. It is only allowed when selecting one field.
func (aps *AsPersonSelect) Int64s(ctx context.Context) ([]int64, error) {
	if len(aps.fields) > 1 {
		return nil, errors.New("schema: AsPersonSelect.Int64s is not achievable when selecting more than 1 field")
	}
	var v []int64
	if err := aps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Int64sX is like Int64s, but panics if an error occurs.
func (aps *AsPersonSelect) Int64sX(ctx context.Context) []int64 {
	v, err := aps.Int64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int64 returns a single int64 from a selector. It is only allowed when selecting one field.
func (aps *AsPersonSelect) Int64(ctx context.Context) (_ int64, err error) {
	var v []int64
	if v, err = aps.Int64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{asperson.Label}
	default:
		err = fmt.Errorf("schema: AsPersonSelect.Int64s returned %d results when one was expected", len(v))
	}
	return
}

// Int64X is like Int64, but panics if an error occurs.
func (aps *AsPersonSelect) Int64X(ctx context.Context) int64 {
	v, err := aps.Int64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (aps *AsPersonSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := aps.sql.Query()
	if err := aps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
