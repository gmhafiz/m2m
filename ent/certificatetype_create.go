// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/bug/ent/certificatetype"
	"entgo.io/bug/ent/leaguecertificatetype"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CertificateTypeCreate is the builder for creating a CertificateType entity.
type CertificateTypeCreate struct {
	config
	mutation *CertificateTypeMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (ctc *CertificateTypeCreate) SetName(s string) *CertificateTypeCreate {
	ctc.mutation.SetName(s)
	return ctc
}

// SetID sets the "id" field.
func (ctc *CertificateTypeCreate) SetID(u uint64) *CertificateTypeCreate {
	ctc.mutation.SetID(u)
	return ctc
}

// AddLeagueCertificateTypeLeagueTypeIDIDs adds the "league_certificate_type_league_type_id" edge to the LeagueCertificateType entity by IDs.
func (ctc *CertificateTypeCreate) AddLeagueCertificateTypeLeagueTypeIDIDs(ids ...int) *CertificateTypeCreate {
	ctc.mutation.AddLeagueCertificateTypeLeagueTypeIDIDs(ids...)
	return ctc
}

// AddLeagueCertificateTypeLeagueTypeID adds the "league_certificate_type_league_type_id" edges to the LeagueCertificateType entity.
func (ctc *CertificateTypeCreate) AddLeagueCertificateTypeLeagueTypeID(l ...*LeagueCertificateType) *CertificateTypeCreate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return ctc.AddLeagueCertificateTypeLeagueTypeIDIDs(ids...)
}

// Mutation returns the CertificateTypeMutation object of the builder.
func (ctc *CertificateTypeCreate) Mutation() *CertificateTypeMutation {
	return ctc.mutation
}

// Save creates the CertificateType in the database.
func (ctc *CertificateTypeCreate) Save(ctx context.Context) (*CertificateType, error) {
	var (
		err  error
		node *CertificateType
	)
	if len(ctc.hooks) == 0 {
		if err = ctc.check(); err != nil {
			return nil, err
		}
		node, err = ctc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CertificateTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ctc.check(); err != nil {
				return nil, err
			}
			ctc.mutation = mutation
			if node, err = ctc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ctc.hooks) - 1; i >= 0; i-- {
			if ctc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ctc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ctc.mutation)
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

// SaveX calls Save and panics if Save returns an error.
func (ctc *CertificateTypeCreate) SaveX(ctx context.Context) *CertificateType {
	v, err := ctc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ctc *CertificateTypeCreate) Exec(ctx context.Context) error {
	_, err := ctc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ctc *CertificateTypeCreate) ExecX(ctx context.Context) {
	if err := ctc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ctc *CertificateTypeCreate) check() error {
	if _, ok := ctc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "CertificateType.name"`)}
	}
	return nil
}

func (ctc *CertificateTypeCreate) sqlSave(ctx context.Context) (*CertificateType, error) {
	_node, _spec := ctc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ctc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	return _node, nil
}

func (ctc *CertificateTypeCreate) createSpec() (*CertificateType, *sqlgraph.CreateSpec) {
	var (
		_node = &CertificateType{config: ctc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: certificatetype.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: certificatetype.FieldID,
			},
		}
	)
	if id, ok := ctc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ctc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: certificatetype.FieldName,
		})
		_node.Name = value
	}
	if nodes := ctc.mutation.LeagueCertificateTypeLeagueTypeIDIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CertificateTypeCreateBulk is the builder for creating many CertificateType entities in bulk.
type CertificateTypeCreateBulk struct {
	config
	builders []*CertificateTypeCreate
}

// Save creates the CertificateType entities in the database.
func (ctcb *CertificateTypeCreateBulk) Save(ctx context.Context) ([]*CertificateType, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ctcb.builders))
	nodes := make([]*CertificateType, len(ctcb.builders))
	mutators := make([]Mutator, len(ctcb.builders))
	for i := range ctcb.builders {
		func(i int, root context.Context) {
			builder := ctcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CertificateTypeMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ctcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ctcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint64(id)
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
		if _, err := mutators[0].Mutate(ctx, ctcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ctcb *CertificateTypeCreateBulk) SaveX(ctx context.Context) []*CertificateType {
	v, err := ctcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ctcb *CertificateTypeCreateBulk) Exec(ctx context.Context) error {
	_, err := ctcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ctcb *CertificateTypeCreateBulk) ExecX(ctx context.Context) {
	if err := ctcb.Exec(ctx); err != nil {
		panic(err)
	}
}