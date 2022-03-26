// Code generated by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/sombriks/rosetta-beer-store/beer-store-service-go-gin-ent/service/models/beer"
)

// BeerCreate is the builder for creating a Beer entity.
type BeerCreate struct {
	config
	mutation *BeerMutation
	hooks    []Hook
}

// SetIdbeer sets the "idbeer" field.
func (bc *BeerCreate) SetIdbeer(i int) *BeerCreate {
	bc.mutation.SetIdbeer(i)
	return bc
}

// SetCreationdatebeer sets the "creationdatebeer" field.
func (bc *BeerCreate) SetCreationdatebeer(t time.Time) *BeerCreate {
	bc.mutation.SetCreationdatebeer(t)
	return bc
}

// SetTitlebeer sets the "titlebeer" field.
func (bc *BeerCreate) SetTitlebeer(s string) *BeerCreate {
	bc.mutation.SetTitlebeer(s)
	return bc
}

// SetDescriptionbeer sets the "descriptionbeer" field.
func (bc *BeerCreate) SetDescriptionbeer(s string) *BeerCreate {
	bc.mutation.SetDescriptionbeer(s)
	return bc
}

// SetIdmedia sets the "idmedia" field.
func (bc *BeerCreate) SetIdmedia(i int) *BeerCreate {
	bc.mutation.SetIdmedia(i)
	return bc
}

// Mutation returns the BeerMutation object of the builder.
func (bc *BeerCreate) Mutation() *BeerMutation {
	return bc.mutation
}

// Save creates the Beer in the database.
func (bc *BeerCreate) Save(ctx context.Context) (*Beer, error) {
	var (
		err  error
		node *Beer
	)
	if len(bc.hooks) == 0 {
		if err = bc.check(); err != nil {
			return nil, err
		}
		node, err = bc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BeerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = bc.check(); err != nil {
				return nil, err
			}
			bc.mutation = mutation
			if node, err = bc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(bc.hooks) - 1; i >= 0; i-- {
			if bc.hooks[i] == nil {
				return nil, fmt.Errorf("models: uninitialized hook (forgotten import models/runtime?)")
			}
			mut = bc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (bc *BeerCreate) SaveX(ctx context.Context) *Beer {
	v, err := bc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bc *BeerCreate) Exec(ctx context.Context) error {
	_, err := bc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bc *BeerCreate) ExecX(ctx context.Context) {
	if err := bc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bc *BeerCreate) check() error {
	if _, ok := bc.mutation.Idbeer(); !ok {
		return &ValidationError{Name: "idbeer", err: errors.New(`models: missing required field "Beer.idbeer"`)}
	}
	if _, ok := bc.mutation.Creationdatebeer(); !ok {
		return &ValidationError{Name: "creationdatebeer", err: errors.New(`models: missing required field "Beer.creationdatebeer"`)}
	}
	if _, ok := bc.mutation.Titlebeer(); !ok {
		return &ValidationError{Name: "titlebeer", err: errors.New(`models: missing required field "Beer.titlebeer"`)}
	}
	if _, ok := bc.mutation.Descriptionbeer(); !ok {
		return &ValidationError{Name: "descriptionbeer", err: errors.New(`models: missing required field "Beer.descriptionbeer"`)}
	}
	if _, ok := bc.mutation.Idmedia(); !ok {
		return &ValidationError{Name: "idmedia", err: errors.New(`models: missing required field "Beer.idmedia"`)}
	}
	return nil
}

func (bc *BeerCreate) sqlSave(ctx context.Context) (*Beer, error) {
	_node, _spec := bc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (bc *BeerCreate) createSpec() (*Beer, *sqlgraph.CreateSpec) {
	var (
		_node = &Beer{config: bc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: beer.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: beer.FieldID,
			},
		}
	)
	if value, ok := bc.mutation.Idbeer(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: beer.FieldIdbeer,
		})
		_node.Idbeer = value
	}
	if value, ok := bc.mutation.Creationdatebeer(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: beer.FieldCreationdatebeer,
		})
		_node.Creationdatebeer = value
	}
	if value, ok := bc.mutation.Titlebeer(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: beer.FieldTitlebeer,
		})
		_node.Titlebeer = value
	}
	if value, ok := bc.mutation.Descriptionbeer(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: beer.FieldDescriptionbeer,
		})
		_node.Descriptionbeer = value
	}
	if value, ok := bc.mutation.Idmedia(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: beer.FieldIdmedia,
		})
		_node.Idmedia = value
	}
	return _node, _spec
}

// BeerCreateBulk is the builder for creating many Beer entities in bulk.
type BeerCreateBulk struct {
	config
	builders []*BeerCreate
}

// Save creates the Beer entities in the database.
func (bcb *BeerCreateBulk) Save(ctx context.Context) ([]*Beer, error) {
	specs := make([]*sqlgraph.CreateSpec, len(bcb.builders))
	nodes := make([]*Beer, len(bcb.builders))
	mutators := make([]Mutator, len(bcb.builders))
	for i := range bcb.builders {
		func(i int, root context.Context) {
			builder := bcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BeerMutation)
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
					_, err = mutators[i+1].Mutate(root, bcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, bcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bcb *BeerCreateBulk) SaveX(ctx context.Context) []*Beer {
	v, err := bcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bcb *BeerCreateBulk) Exec(ctx context.Context) error {
	_, err := bcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcb *BeerCreateBulk) ExecX(ctx context.Context) {
	if err := bcb.Exec(ctx); err != nil {
		panic(err)
	}
}
