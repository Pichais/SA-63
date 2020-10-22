// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/Pichais/app/ent/spaciality"
	"github.com/Pichais/app/ent/typedisease"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// TypeDiseaseCreate is the builder for creating a TypeDisease entity.
type TypeDiseaseCreate struct {
	config
	mutation *TypeDiseaseMutation
	hooks    []Hook
}

// SetDiseaseName sets the DiseaseName field.
func (tdc *TypeDiseaseCreate) SetDiseaseName(s string) *TypeDiseaseCreate {
	tdc.mutation.SetDiseaseName(s)
	return tdc
}

// AddFortypeIDs adds the fortype edge to Spaciality by ids.
func (tdc *TypeDiseaseCreate) AddFortypeIDs(ids ...int) *TypeDiseaseCreate {
	tdc.mutation.AddFortypeIDs(ids...)
	return tdc
}

// AddFortype adds the fortype edges to Spaciality.
func (tdc *TypeDiseaseCreate) AddFortype(s ...*Spaciality) *TypeDiseaseCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return tdc.AddFortypeIDs(ids...)
}

// Mutation returns the TypeDiseaseMutation object of the builder.
func (tdc *TypeDiseaseCreate) Mutation() *TypeDiseaseMutation {
	return tdc.mutation
}

// Save creates the TypeDisease in the database.
func (tdc *TypeDiseaseCreate) Save(ctx context.Context) (*TypeDisease, error) {
	if _, ok := tdc.mutation.DiseaseName(); !ok {
		return nil, &ValidationError{Name: "DiseaseName", err: errors.New("ent: missing required field \"DiseaseName\"")}
	}
	if v, ok := tdc.mutation.DiseaseName(); ok {
		if err := typedisease.DiseaseNameValidator(v); err != nil {
			return nil, &ValidationError{Name: "DiseaseName", err: fmt.Errorf("ent: validator failed for field \"DiseaseName\": %w", err)}
		}
	}
	var (
		err  error
		node *TypeDisease
	)
	if len(tdc.hooks) == 0 {
		node, err = tdc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TypeDiseaseMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tdc.mutation = mutation
			node, err = tdc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tdc.hooks) - 1; i >= 0; i-- {
			mut = tdc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tdc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tdc *TypeDiseaseCreate) SaveX(ctx context.Context) *TypeDisease {
	v, err := tdc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (tdc *TypeDiseaseCreate) sqlSave(ctx context.Context) (*TypeDisease, error) {
	td, _spec := tdc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tdc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	td.ID = int(id)
	return td, nil
}

func (tdc *TypeDiseaseCreate) createSpec() (*TypeDisease, *sqlgraph.CreateSpec) {
	var (
		td    = &TypeDisease{config: tdc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: typedisease.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: typedisease.FieldID,
			},
		}
	)
	if value, ok := tdc.mutation.DiseaseName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: typedisease.FieldDiseaseName,
		})
		td.DiseaseName = value
	}
	if nodes := tdc.mutation.FortypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   typedisease.FortypeTable,
			Columns: []string{typedisease.FortypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: spaciality.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return td, _spec
}