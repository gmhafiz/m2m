// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/bug/ent/certificatetype"
	"entgo.io/bug/ent/leaguecertificatetype"
	"entgo.io/bug/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CertificateTypeUpdate is the builder for updating CertificateType entities.
type CertificateTypeUpdate struct {
	config
	hooks    []Hook
	mutation *CertificateTypeMutation
}

// Where appends a list predicates to the CertificateTypeUpdate builder.
func (ctu *CertificateTypeUpdate) Where(ps ...predicate.CertificateType) *CertificateTypeUpdate {
	ctu.mutation.Where(ps...)
	return ctu
}

// SetName sets the "name" field.
func (ctu *CertificateTypeUpdate) SetName(s string) *CertificateTypeUpdate {
	ctu.mutation.SetName(s)
	return ctu
}

// AddLeagueCertificateTypeLeagueTypeIDIDs adds the "league_certificate_type_league_type_id" edge to the LeagueCertificateType entity by IDs.
func (ctu *CertificateTypeUpdate) AddLeagueCertificateTypeLeagueTypeIDIDs(ids ...int) *CertificateTypeUpdate {
	ctu.mutation.AddLeagueCertificateTypeLeagueTypeIDIDs(ids...)
	return ctu
}

// AddLeagueCertificateTypeLeagueTypeID adds the "league_certificate_type_league_type_id" edges to the LeagueCertificateType entity.
func (ctu *CertificateTypeUpdate) AddLeagueCertificateTypeLeagueTypeID(l ...*LeagueCertificateType) *CertificateTypeUpdate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return ctu.AddLeagueCertificateTypeLeagueTypeIDIDs(ids...)
}

// Mutation returns the CertificateTypeMutation object of the builder.
func (ctu *CertificateTypeUpdate) Mutation() *CertificateTypeMutation {
	return ctu.mutation
}

// ClearLeagueCertificateTypeLeagueTypeID clears all "league_certificate_type_league_type_id" edges to the LeagueCertificateType entity.
func (ctu *CertificateTypeUpdate) ClearLeagueCertificateTypeLeagueTypeID() *CertificateTypeUpdate {
	ctu.mutation.ClearLeagueCertificateTypeLeagueTypeID()
	return ctu
}

// RemoveLeagueCertificateTypeLeagueTypeIDIDs removes the "league_certificate_type_league_type_id" edge to LeagueCertificateType entities by IDs.
func (ctu *CertificateTypeUpdate) RemoveLeagueCertificateTypeLeagueTypeIDIDs(ids ...int) *CertificateTypeUpdate {
	ctu.mutation.RemoveLeagueCertificateTypeLeagueTypeIDIDs(ids...)
	return ctu
}

// RemoveLeagueCertificateTypeLeagueTypeID removes "league_certificate_type_league_type_id" edges to LeagueCertificateType entities.
func (ctu *CertificateTypeUpdate) RemoveLeagueCertificateTypeLeagueTypeID(l ...*LeagueCertificateType) *CertificateTypeUpdate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return ctu.RemoveLeagueCertificateTypeLeagueTypeIDIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ctu *CertificateTypeUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ctu.hooks) == 0 {
		affected, err = ctu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CertificateTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ctu.mutation = mutation
			affected, err = ctu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ctu.hooks) - 1; i >= 0; i-- {
			if ctu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ctu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ctu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ctu *CertificateTypeUpdate) SaveX(ctx context.Context) int {
	affected, err := ctu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ctu *CertificateTypeUpdate) Exec(ctx context.Context) error {
	_, err := ctu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ctu *CertificateTypeUpdate) ExecX(ctx context.Context) {
	if err := ctu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ctu *CertificateTypeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   certificatetype.Table,
			Columns: certificatetype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: certificatetype.FieldID,
			},
		},
	}
	if ps := ctu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ctu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: certificatetype.FieldName,
		})
	}
	if ctu.mutation.LeagueCertificateTypeLeagueTypeIDCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   certificatetype.LeagueCertificateTypeLeagueTypeIDTable,
			Columns: []string{certificatetype.LeagueCertificateTypeLeagueTypeIDColumn},
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
	if nodes := ctu.mutation.RemovedLeagueCertificateTypeLeagueTypeIDIDs(); len(nodes) > 0 && !ctu.mutation.LeagueCertificateTypeLeagueTypeIDCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   certificatetype.LeagueCertificateTypeLeagueTypeIDTable,
			Columns: []string{certificatetype.LeagueCertificateTypeLeagueTypeIDColumn},
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
	if nodes := ctu.mutation.LeagueCertificateTypeLeagueTypeIDIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   certificatetype.LeagueCertificateTypeLeagueTypeIDTable,
			Columns: []string{certificatetype.LeagueCertificateTypeLeagueTypeIDColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, ctu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{certificatetype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// CertificateTypeUpdateOne is the builder for updating a single CertificateType entity.
type CertificateTypeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CertificateTypeMutation
}

// SetName sets the "name" field.
func (ctuo *CertificateTypeUpdateOne) SetName(s string) *CertificateTypeUpdateOne {
	ctuo.mutation.SetName(s)
	return ctuo
}

// AddLeagueCertificateTypeLeagueTypeIDIDs adds the "league_certificate_type_league_type_id" edge to the LeagueCertificateType entity by IDs.
func (ctuo *CertificateTypeUpdateOne) AddLeagueCertificateTypeLeagueTypeIDIDs(ids ...int) *CertificateTypeUpdateOne {
	ctuo.mutation.AddLeagueCertificateTypeLeagueTypeIDIDs(ids...)
	return ctuo
}

// AddLeagueCertificateTypeLeagueTypeID adds the "league_certificate_type_league_type_id" edges to the LeagueCertificateType entity.
func (ctuo *CertificateTypeUpdateOne) AddLeagueCertificateTypeLeagueTypeID(l ...*LeagueCertificateType) *CertificateTypeUpdateOne {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return ctuo.AddLeagueCertificateTypeLeagueTypeIDIDs(ids...)
}

// Mutation returns the CertificateTypeMutation object of the builder.
func (ctuo *CertificateTypeUpdateOne) Mutation() *CertificateTypeMutation {
	return ctuo.mutation
}

// ClearLeagueCertificateTypeLeagueTypeID clears all "league_certificate_type_league_type_id" edges to the LeagueCertificateType entity.
func (ctuo *CertificateTypeUpdateOne) ClearLeagueCertificateTypeLeagueTypeID() *CertificateTypeUpdateOne {
	ctuo.mutation.ClearLeagueCertificateTypeLeagueTypeID()
	return ctuo
}

// RemoveLeagueCertificateTypeLeagueTypeIDIDs removes the "league_certificate_type_league_type_id" edge to LeagueCertificateType entities by IDs.
func (ctuo *CertificateTypeUpdateOne) RemoveLeagueCertificateTypeLeagueTypeIDIDs(ids ...int) *CertificateTypeUpdateOne {
	ctuo.mutation.RemoveLeagueCertificateTypeLeagueTypeIDIDs(ids...)
	return ctuo
}

// RemoveLeagueCertificateTypeLeagueTypeID removes "league_certificate_type_league_type_id" edges to LeagueCertificateType entities.
func (ctuo *CertificateTypeUpdateOne) RemoveLeagueCertificateTypeLeagueTypeID(l ...*LeagueCertificateType) *CertificateTypeUpdateOne {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return ctuo.RemoveLeagueCertificateTypeLeagueTypeIDIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ctuo *CertificateTypeUpdateOne) Select(field string, fields ...string) *CertificateTypeUpdateOne {
	ctuo.fields = append([]string{field}, fields...)
	return ctuo
}

// Save executes the query and returns the updated CertificateType entity.
func (ctuo *CertificateTypeUpdateOne) Save(ctx context.Context) (*CertificateType, error) {
	var (
		err  error
		node *CertificateType
	)
	if len(ctuo.hooks) == 0 {
		node, err = ctuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CertificateTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ctuo.mutation = mutation
			node, err = ctuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ctuo.hooks) - 1; i >= 0; i-- {
			if ctuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ctuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ctuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*CertificateType)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from CertificateTypeMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ctuo *CertificateTypeUpdateOne) SaveX(ctx context.Context) *CertificateType {
	node, err := ctuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ctuo *CertificateTypeUpdateOne) Exec(ctx context.Context) error {
	_, err := ctuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ctuo *CertificateTypeUpdateOne) ExecX(ctx context.Context) {
	if err := ctuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ctuo *CertificateTypeUpdateOne) sqlSave(ctx context.Context) (_node *CertificateType, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   certificatetype.Table,
			Columns: certificatetype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: certificatetype.FieldID,
			},
		},
	}
	id, ok := ctuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "CertificateType.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ctuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, certificatetype.FieldID)
		for _, f := range fields {
			if !certificatetype.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != certificatetype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ctuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ctuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: certificatetype.FieldName,
		})
	}
	if ctuo.mutation.LeagueCertificateTypeLeagueTypeIDCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   certificatetype.LeagueCertificateTypeLeagueTypeIDTable,
			Columns: []string{certificatetype.LeagueCertificateTypeLeagueTypeIDColumn},
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
	if nodes := ctuo.mutation.RemovedLeagueCertificateTypeLeagueTypeIDIDs(); len(nodes) > 0 && !ctuo.mutation.LeagueCertificateTypeLeagueTypeIDCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   certificatetype.LeagueCertificateTypeLeagueTypeIDTable,
			Columns: []string{certificatetype.LeagueCertificateTypeLeagueTypeIDColumn},
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
	if nodes := ctuo.mutation.LeagueCertificateTypeLeagueTypeIDIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   certificatetype.LeagueCertificateTypeLeagueTypeIDTable,
			Columns: []string{certificatetype.LeagueCertificateTypeLeagueTypeIDColumn},
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
	_node = &CertificateType{config: ctuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ctuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{certificatetype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
