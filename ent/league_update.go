// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/bug/ent/league"
	"entgo.io/bug/ent/leaguecertificatetype"
	"entgo.io/bug/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LeagueUpdate is the builder for updating League entities.
type LeagueUpdate struct {
	config
	hooks    []Hook
	mutation *LeagueMutation
}

// Where appends a list predicates to the LeagueUpdate builder.
func (lu *LeagueUpdate) Where(ps ...predicate.League) *LeagueUpdate {
	lu.mutation.Where(ps...)
	return lu
}

// SetName sets the "name" field.
func (lu *LeagueUpdate) SetName(s string) *LeagueUpdate {
	lu.mutation.SetName(s)
	return lu
}

// AddLeagueCertificateTypeIDs adds the "league_certificate_type" edge to the LeagueCertificateType entity by IDs.
func (lu *LeagueUpdate) AddLeagueCertificateTypeIDs(ids ...int) *LeagueUpdate {
	lu.mutation.AddLeagueCertificateTypeIDs(ids...)
	return lu
}

// AddLeagueCertificateType adds the "league_certificate_type" edges to the LeagueCertificateType entity.
func (lu *LeagueUpdate) AddLeagueCertificateType(l ...*LeagueCertificateType) *LeagueUpdate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return lu.AddLeagueCertificateTypeIDs(ids...)
}

// Mutation returns the LeagueMutation object of the builder.
func (lu *LeagueUpdate) Mutation() *LeagueMutation {
	return lu.mutation
}

// ClearLeagueCertificateType clears all "league_certificate_type" edges to the LeagueCertificateType entity.
func (lu *LeagueUpdate) ClearLeagueCertificateType() *LeagueUpdate {
	lu.mutation.ClearLeagueCertificateType()
	return lu
}

// RemoveLeagueCertificateTypeIDs removes the "league_certificate_type" edge to LeagueCertificateType entities by IDs.
func (lu *LeagueUpdate) RemoveLeagueCertificateTypeIDs(ids ...int) *LeagueUpdate {
	lu.mutation.RemoveLeagueCertificateTypeIDs(ids...)
	return lu
}

// RemoveLeagueCertificateType removes "league_certificate_type" edges to LeagueCertificateType entities.
func (lu *LeagueUpdate) RemoveLeagueCertificateType(l ...*LeagueCertificateType) *LeagueUpdate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return lu.RemoveLeagueCertificateTypeIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lu *LeagueUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(lu.hooks) == 0 {
		affected, err = lu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LeagueMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			lu.mutation = mutation
			affected, err = lu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(lu.hooks) - 1; i >= 0; i-- {
			if lu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = lu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, lu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (lu *LeagueUpdate) SaveX(ctx context.Context) int {
	affected, err := lu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lu *LeagueUpdate) Exec(ctx context.Context) error {
	_, err := lu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lu *LeagueUpdate) ExecX(ctx context.Context) {
	if err := lu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (lu *LeagueUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   league.Table,
			Columns: league.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: league.FieldID,
			},
		},
	}
	if ps := lu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: league.FieldName,
		})
	}
	if lu.mutation.LeagueCertificateTypeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   league.LeagueCertificateTypeTable,
			Columns: []string{league.LeagueCertificateTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: leaguecertificatetype.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.RemovedLeagueCertificateTypeIDs(); len(nodes) > 0 && !lu.mutation.LeagueCertificateTypeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   league.LeagueCertificateTypeTable,
			Columns: []string{league.LeagueCertificateTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: leaguecertificatetype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.LeagueCertificateTypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   league.LeagueCertificateTypeTable,
			Columns: []string{league.LeagueCertificateTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: leaguecertificatetype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, lu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{league.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// LeagueUpdateOne is the builder for updating a single League entity.
type LeagueUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *LeagueMutation
}

// SetName sets the "name" field.
func (luo *LeagueUpdateOne) SetName(s string) *LeagueUpdateOne {
	luo.mutation.SetName(s)
	return luo
}

// AddLeagueCertificateTypeIDs adds the "league_certificate_type" edge to the LeagueCertificateType entity by IDs.
func (luo *LeagueUpdateOne) AddLeagueCertificateTypeIDs(ids ...int) *LeagueUpdateOne {
	luo.mutation.AddLeagueCertificateTypeIDs(ids...)
	return luo
}

// AddLeagueCertificateType adds the "league_certificate_type" edges to the LeagueCertificateType entity.
func (luo *LeagueUpdateOne) AddLeagueCertificateType(l ...*LeagueCertificateType) *LeagueUpdateOne {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return luo.AddLeagueCertificateTypeIDs(ids...)
}

// Mutation returns the LeagueMutation object of the builder.
func (luo *LeagueUpdateOne) Mutation() *LeagueMutation {
	return luo.mutation
}

// ClearLeagueCertificateType clears all "league_certificate_type" edges to the LeagueCertificateType entity.
func (luo *LeagueUpdateOne) ClearLeagueCertificateType() *LeagueUpdateOne {
	luo.mutation.ClearLeagueCertificateType()
	return luo
}

// RemoveLeagueCertificateTypeIDs removes the "league_certificate_type" edge to LeagueCertificateType entities by IDs.
func (luo *LeagueUpdateOne) RemoveLeagueCertificateTypeIDs(ids ...int) *LeagueUpdateOne {
	luo.mutation.RemoveLeagueCertificateTypeIDs(ids...)
	return luo
}

// RemoveLeagueCertificateType removes "league_certificate_type" edges to LeagueCertificateType entities.
func (luo *LeagueUpdateOne) RemoveLeagueCertificateType(l ...*LeagueCertificateType) *LeagueUpdateOne {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return luo.RemoveLeagueCertificateTypeIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (luo *LeagueUpdateOne) Select(field string, fields ...string) *LeagueUpdateOne {
	luo.fields = append([]string{field}, fields...)
	return luo
}

// Save executes the query and returns the updated League entity.
func (luo *LeagueUpdateOne) Save(ctx context.Context) (*League, error) {
	var (
		err  error
		node *League
	)
	if len(luo.hooks) == 0 {
		node, err = luo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LeagueMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			luo.mutation = mutation
			node, err = luo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(luo.hooks) - 1; i >= 0; i-- {
			if luo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = luo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, luo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*League)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from LeagueMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (luo *LeagueUpdateOne) SaveX(ctx context.Context) *League {
	node, err := luo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (luo *LeagueUpdateOne) Exec(ctx context.Context) error {
	_, err := luo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (luo *LeagueUpdateOne) ExecX(ctx context.Context) {
	if err := luo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (luo *LeagueUpdateOne) sqlSave(ctx context.Context) (_node *League, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   league.Table,
			Columns: league.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: league.FieldID,
			},
		},
	}
	id, ok := luo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "League.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := luo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, league.FieldID)
		for _, f := range fields {
			if !league.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != league.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := luo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := luo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: league.FieldName,
		})
	}
	if luo.mutation.LeagueCertificateTypeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   league.LeagueCertificateTypeTable,
			Columns: []string{league.LeagueCertificateTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: leaguecertificatetype.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.RemovedLeagueCertificateTypeIDs(); len(nodes) > 0 && !luo.mutation.LeagueCertificateTypeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   league.LeagueCertificateTypeTable,
			Columns: []string{league.LeagueCertificateTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: leaguecertificatetype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.LeagueCertificateTypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   league.LeagueCertificateTypeTable,
			Columns: []string{league.LeagueCertificateTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: leaguecertificatetype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &League{config: luo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, luo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{league.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
