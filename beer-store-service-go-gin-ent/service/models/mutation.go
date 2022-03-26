// Code generated by entc, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/sombriks/rosetta-beer-store/beer-store-service-go-gin-ent/service/models/beer"
	"github.com/sombriks/rosetta-beer-store/beer-store-service-go-gin-ent/service/models/media"
	"github.com/sombriks/rosetta-beer-store/beer-store-service-go-gin-ent/service/models/predicate"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeBeer  = "Beer"
	TypeMedia = "Media"
)

// BeerMutation represents an operation that mutates the Beer nodes in the graph.
type BeerMutation struct {
	config
	op               Op
	typ              string
	id               *int
	idbeer           *int
	addidbeer        *int
	creationdatebeer *time.Time
	titlebeer        *string
	descriptionbeer  *string
	idmedia          *int
	addidmedia       *int
	clearedFields    map[string]struct{}
	done             bool
	oldValue         func(context.Context) (*Beer, error)
	predicates       []predicate.Beer
}

var _ ent.Mutation = (*BeerMutation)(nil)

// beerOption allows management of the mutation configuration using functional options.
type beerOption func(*BeerMutation)

// newBeerMutation creates new mutation for the Beer entity.
func newBeerMutation(c config, op Op, opts ...beerOption) *BeerMutation {
	m := &BeerMutation{
		config:        c,
		op:            op,
		typ:           TypeBeer,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withBeerID sets the ID field of the mutation.
func withBeerID(id int) beerOption {
	return func(m *BeerMutation) {
		var (
			err   error
			once  sync.Once
			value *Beer
		)
		m.oldValue = func(ctx context.Context) (*Beer, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Beer.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withBeer sets the old Beer of the mutation.
func withBeer(node *Beer) beerOption {
	return func(m *BeerMutation) {
		m.oldValue = func(context.Context) (*Beer, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m BeerMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m BeerMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("models: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *BeerMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *BeerMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Beer.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetIdbeer sets the "idbeer" field.
func (m *BeerMutation) SetIdbeer(i int) {
	m.idbeer = &i
	m.addidbeer = nil
}

// Idbeer returns the value of the "idbeer" field in the mutation.
func (m *BeerMutation) Idbeer() (r int, exists bool) {
	v := m.idbeer
	if v == nil {
		return
	}
	return *v, true
}

// OldIdbeer returns the old "idbeer" field's value of the Beer entity.
// If the Beer object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *BeerMutation) OldIdbeer(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldIdbeer is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldIdbeer requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldIdbeer: %w", err)
	}
	return oldValue.Idbeer, nil
}

// AddIdbeer adds i to the "idbeer" field.
func (m *BeerMutation) AddIdbeer(i int) {
	if m.addidbeer != nil {
		*m.addidbeer += i
	} else {
		m.addidbeer = &i
	}
}

// AddedIdbeer returns the value that was added to the "idbeer" field in this mutation.
func (m *BeerMutation) AddedIdbeer() (r int, exists bool) {
	v := m.addidbeer
	if v == nil {
		return
	}
	return *v, true
}

// ResetIdbeer resets all changes to the "idbeer" field.
func (m *BeerMutation) ResetIdbeer() {
	m.idbeer = nil
	m.addidbeer = nil
}

// SetCreationdatebeer sets the "creationdatebeer" field.
func (m *BeerMutation) SetCreationdatebeer(t time.Time) {
	m.creationdatebeer = &t
}

// Creationdatebeer returns the value of the "creationdatebeer" field in the mutation.
func (m *BeerMutation) Creationdatebeer() (r time.Time, exists bool) {
	v := m.creationdatebeer
	if v == nil {
		return
	}
	return *v, true
}

// OldCreationdatebeer returns the old "creationdatebeer" field's value of the Beer entity.
// If the Beer object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *BeerMutation) OldCreationdatebeer(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreationdatebeer is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreationdatebeer requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreationdatebeer: %w", err)
	}
	return oldValue.Creationdatebeer, nil
}

// ResetCreationdatebeer resets all changes to the "creationdatebeer" field.
func (m *BeerMutation) ResetCreationdatebeer() {
	m.creationdatebeer = nil
}

// SetTitlebeer sets the "titlebeer" field.
func (m *BeerMutation) SetTitlebeer(s string) {
	m.titlebeer = &s
}

// Titlebeer returns the value of the "titlebeer" field in the mutation.
func (m *BeerMutation) Titlebeer() (r string, exists bool) {
	v := m.titlebeer
	if v == nil {
		return
	}
	return *v, true
}

// OldTitlebeer returns the old "titlebeer" field's value of the Beer entity.
// If the Beer object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *BeerMutation) OldTitlebeer(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldTitlebeer is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldTitlebeer requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTitlebeer: %w", err)
	}
	return oldValue.Titlebeer, nil
}

// ResetTitlebeer resets all changes to the "titlebeer" field.
func (m *BeerMutation) ResetTitlebeer() {
	m.titlebeer = nil
}

// SetDescriptionbeer sets the "descriptionbeer" field.
func (m *BeerMutation) SetDescriptionbeer(s string) {
	m.descriptionbeer = &s
}

// Descriptionbeer returns the value of the "descriptionbeer" field in the mutation.
func (m *BeerMutation) Descriptionbeer() (r string, exists bool) {
	v := m.descriptionbeer
	if v == nil {
		return
	}
	return *v, true
}

// OldDescriptionbeer returns the old "descriptionbeer" field's value of the Beer entity.
// If the Beer object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *BeerMutation) OldDescriptionbeer(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldDescriptionbeer is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldDescriptionbeer requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDescriptionbeer: %w", err)
	}
	return oldValue.Descriptionbeer, nil
}

// ResetDescriptionbeer resets all changes to the "descriptionbeer" field.
func (m *BeerMutation) ResetDescriptionbeer() {
	m.descriptionbeer = nil
}

// SetIdmedia sets the "idmedia" field.
func (m *BeerMutation) SetIdmedia(i int) {
	m.idmedia = &i
	m.addidmedia = nil
}

// Idmedia returns the value of the "idmedia" field in the mutation.
func (m *BeerMutation) Idmedia() (r int, exists bool) {
	v := m.idmedia
	if v == nil {
		return
	}
	return *v, true
}

// OldIdmedia returns the old "idmedia" field's value of the Beer entity.
// If the Beer object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *BeerMutation) OldIdmedia(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldIdmedia is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldIdmedia requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldIdmedia: %w", err)
	}
	return oldValue.Idmedia, nil
}

// AddIdmedia adds i to the "idmedia" field.
func (m *BeerMutation) AddIdmedia(i int) {
	if m.addidmedia != nil {
		*m.addidmedia += i
	} else {
		m.addidmedia = &i
	}
}

// AddedIdmedia returns the value that was added to the "idmedia" field in this mutation.
func (m *BeerMutation) AddedIdmedia() (r int, exists bool) {
	v := m.addidmedia
	if v == nil {
		return
	}
	return *v, true
}

// ResetIdmedia resets all changes to the "idmedia" field.
func (m *BeerMutation) ResetIdmedia() {
	m.idmedia = nil
	m.addidmedia = nil
}

// Where appends a list predicates to the BeerMutation builder.
func (m *BeerMutation) Where(ps ...predicate.Beer) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *BeerMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Beer).
func (m *BeerMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *BeerMutation) Fields() []string {
	fields := make([]string, 0, 5)
	if m.idbeer != nil {
		fields = append(fields, beer.FieldIdbeer)
	}
	if m.creationdatebeer != nil {
		fields = append(fields, beer.FieldCreationdatebeer)
	}
	if m.titlebeer != nil {
		fields = append(fields, beer.FieldTitlebeer)
	}
	if m.descriptionbeer != nil {
		fields = append(fields, beer.FieldDescriptionbeer)
	}
	if m.idmedia != nil {
		fields = append(fields, beer.FieldIdmedia)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *BeerMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case beer.FieldIdbeer:
		return m.Idbeer()
	case beer.FieldCreationdatebeer:
		return m.Creationdatebeer()
	case beer.FieldTitlebeer:
		return m.Titlebeer()
	case beer.FieldDescriptionbeer:
		return m.Descriptionbeer()
	case beer.FieldIdmedia:
		return m.Idmedia()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *BeerMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case beer.FieldIdbeer:
		return m.OldIdbeer(ctx)
	case beer.FieldCreationdatebeer:
		return m.OldCreationdatebeer(ctx)
	case beer.FieldTitlebeer:
		return m.OldTitlebeer(ctx)
	case beer.FieldDescriptionbeer:
		return m.OldDescriptionbeer(ctx)
	case beer.FieldIdmedia:
		return m.OldIdmedia(ctx)
	}
	return nil, fmt.Errorf("unknown Beer field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *BeerMutation) SetField(name string, value ent.Value) error {
	switch name {
	case beer.FieldIdbeer:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetIdbeer(v)
		return nil
	case beer.FieldCreationdatebeer:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreationdatebeer(v)
		return nil
	case beer.FieldTitlebeer:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTitlebeer(v)
		return nil
	case beer.FieldDescriptionbeer:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDescriptionbeer(v)
		return nil
	case beer.FieldIdmedia:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetIdmedia(v)
		return nil
	}
	return fmt.Errorf("unknown Beer field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *BeerMutation) AddedFields() []string {
	var fields []string
	if m.addidbeer != nil {
		fields = append(fields, beer.FieldIdbeer)
	}
	if m.addidmedia != nil {
		fields = append(fields, beer.FieldIdmedia)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *BeerMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case beer.FieldIdbeer:
		return m.AddedIdbeer()
	case beer.FieldIdmedia:
		return m.AddedIdmedia()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *BeerMutation) AddField(name string, value ent.Value) error {
	switch name {
	case beer.FieldIdbeer:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddIdbeer(v)
		return nil
	case beer.FieldIdmedia:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddIdmedia(v)
		return nil
	}
	return fmt.Errorf("unknown Beer numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *BeerMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *BeerMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *BeerMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Beer nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *BeerMutation) ResetField(name string) error {
	switch name {
	case beer.FieldIdbeer:
		m.ResetIdbeer()
		return nil
	case beer.FieldCreationdatebeer:
		m.ResetCreationdatebeer()
		return nil
	case beer.FieldTitlebeer:
		m.ResetTitlebeer()
		return nil
	case beer.FieldDescriptionbeer:
		m.ResetDescriptionbeer()
		return nil
	case beer.FieldIdmedia:
		m.ResetIdmedia()
		return nil
	}
	return fmt.Errorf("unknown Beer field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *BeerMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *BeerMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *BeerMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *BeerMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *BeerMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *BeerMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *BeerMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Beer unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *BeerMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Beer edge %s", name)
}

// MediaMutation represents an operation that mutates the Media nodes in the graph.
type MediaMutation struct {
	config
	op                Op
	typ               string
	id                *int
	idmedia           *int
	addidmedia        *int
	creationdatemedia *time.Time
	datamedia         *[]byte
	nomemedia         *string
	mimemedia         *string
	clearedFields     map[string]struct{}
	done              bool
	oldValue          func(context.Context) (*Media, error)
	predicates        []predicate.Media
}

var _ ent.Mutation = (*MediaMutation)(nil)

// mediaOption allows management of the mutation configuration using functional options.
type mediaOption func(*MediaMutation)

// newMediaMutation creates new mutation for the Media entity.
func newMediaMutation(c config, op Op, opts ...mediaOption) *MediaMutation {
	m := &MediaMutation{
		config:        c,
		op:            op,
		typ:           TypeMedia,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withMediaID sets the ID field of the mutation.
func withMediaID(id int) mediaOption {
	return func(m *MediaMutation) {
		var (
			err   error
			once  sync.Once
			value *Media
		)
		m.oldValue = func(ctx context.Context) (*Media, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Media.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withMedia sets the old Media of the mutation.
func withMedia(node *Media) mediaOption {
	return func(m *MediaMutation) {
		m.oldValue = func(context.Context) (*Media, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m MediaMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m MediaMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("models: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *MediaMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *MediaMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Media.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetIdmedia sets the "idmedia" field.
func (m *MediaMutation) SetIdmedia(i int) {
	m.idmedia = &i
	m.addidmedia = nil
}

// Idmedia returns the value of the "idmedia" field in the mutation.
func (m *MediaMutation) Idmedia() (r int, exists bool) {
	v := m.idmedia
	if v == nil {
		return
	}
	return *v, true
}

// OldIdmedia returns the old "idmedia" field's value of the Media entity.
// If the Media object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MediaMutation) OldIdmedia(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldIdmedia is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldIdmedia requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldIdmedia: %w", err)
	}
	return oldValue.Idmedia, nil
}

// AddIdmedia adds i to the "idmedia" field.
func (m *MediaMutation) AddIdmedia(i int) {
	if m.addidmedia != nil {
		*m.addidmedia += i
	} else {
		m.addidmedia = &i
	}
}

// AddedIdmedia returns the value that was added to the "idmedia" field in this mutation.
func (m *MediaMutation) AddedIdmedia() (r int, exists bool) {
	v := m.addidmedia
	if v == nil {
		return
	}
	return *v, true
}

// ResetIdmedia resets all changes to the "idmedia" field.
func (m *MediaMutation) ResetIdmedia() {
	m.idmedia = nil
	m.addidmedia = nil
}

// SetCreationdatemedia sets the "creationdatemedia" field.
func (m *MediaMutation) SetCreationdatemedia(t time.Time) {
	m.creationdatemedia = &t
}

// Creationdatemedia returns the value of the "creationdatemedia" field in the mutation.
func (m *MediaMutation) Creationdatemedia() (r time.Time, exists bool) {
	v := m.creationdatemedia
	if v == nil {
		return
	}
	return *v, true
}

// OldCreationdatemedia returns the old "creationdatemedia" field's value of the Media entity.
// If the Media object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MediaMutation) OldCreationdatemedia(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreationdatemedia is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreationdatemedia requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreationdatemedia: %w", err)
	}
	return oldValue.Creationdatemedia, nil
}

// ResetCreationdatemedia resets all changes to the "creationdatemedia" field.
func (m *MediaMutation) ResetCreationdatemedia() {
	m.creationdatemedia = nil
}

// SetDatamedia sets the "datamedia" field.
func (m *MediaMutation) SetDatamedia(b []byte) {
	m.datamedia = &b
}

// Datamedia returns the value of the "datamedia" field in the mutation.
func (m *MediaMutation) Datamedia() (r []byte, exists bool) {
	v := m.datamedia
	if v == nil {
		return
	}
	return *v, true
}

// OldDatamedia returns the old "datamedia" field's value of the Media entity.
// If the Media object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MediaMutation) OldDatamedia(ctx context.Context) (v []byte, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldDatamedia is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldDatamedia requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDatamedia: %w", err)
	}
	return oldValue.Datamedia, nil
}

// ResetDatamedia resets all changes to the "datamedia" field.
func (m *MediaMutation) ResetDatamedia() {
	m.datamedia = nil
}

// SetNomemedia sets the "nomemedia" field.
func (m *MediaMutation) SetNomemedia(s string) {
	m.nomemedia = &s
}

// Nomemedia returns the value of the "nomemedia" field in the mutation.
func (m *MediaMutation) Nomemedia() (r string, exists bool) {
	v := m.nomemedia
	if v == nil {
		return
	}
	return *v, true
}

// OldNomemedia returns the old "nomemedia" field's value of the Media entity.
// If the Media object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MediaMutation) OldNomemedia(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldNomemedia is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldNomemedia requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldNomemedia: %w", err)
	}
	return oldValue.Nomemedia, nil
}

// ResetNomemedia resets all changes to the "nomemedia" field.
func (m *MediaMutation) ResetNomemedia() {
	m.nomemedia = nil
}

// SetMimemedia sets the "mimemedia" field.
func (m *MediaMutation) SetMimemedia(s string) {
	m.mimemedia = &s
}

// Mimemedia returns the value of the "mimemedia" field in the mutation.
func (m *MediaMutation) Mimemedia() (r string, exists bool) {
	v := m.mimemedia
	if v == nil {
		return
	}
	return *v, true
}

// OldMimemedia returns the old "mimemedia" field's value of the Media entity.
// If the Media object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MediaMutation) OldMimemedia(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldMimemedia is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldMimemedia requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldMimemedia: %w", err)
	}
	return oldValue.Mimemedia, nil
}

// ResetMimemedia resets all changes to the "mimemedia" field.
func (m *MediaMutation) ResetMimemedia() {
	m.mimemedia = nil
}

// Where appends a list predicates to the MediaMutation builder.
func (m *MediaMutation) Where(ps ...predicate.Media) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *MediaMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Media).
func (m *MediaMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *MediaMutation) Fields() []string {
	fields := make([]string, 0, 5)
	if m.idmedia != nil {
		fields = append(fields, media.FieldIdmedia)
	}
	if m.creationdatemedia != nil {
		fields = append(fields, media.FieldCreationdatemedia)
	}
	if m.datamedia != nil {
		fields = append(fields, media.FieldDatamedia)
	}
	if m.nomemedia != nil {
		fields = append(fields, media.FieldNomemedia)
	}
	if m.mimemedia != nil {
		fields = append(fields, media.FieldMimemedia)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *MediaMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case media.FieldIdmedia:
		return m.Idmedia()
	case media.FieldCreationdatemedia:
		return m.Creationdatemedia()
	case media.FieldDatamedia:
		return m.Datamedia()
	case media.FieldNomemedia:
		return m.Nomemedia()
	case media.FieldMimemedia:
		return m.Mimemedia()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *MediaMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case media.FieldIdmedia:
		return m.OldIdmedia(ctx)
	case media.FieldCreationdatemedia:
		return m.OldCreationdatemedia(ctx)
	case media.FieldDatamedia:
		return m.OldDatamedia(ctx)
	case media.FieldNomemedia:
		return m.OldNomemedia(ctx)
	case media.FieldMimemedia:
		return m.OldMimemedia(ctx)
	}
	return nil, fmt.Errorf("unknown Media field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *MediaMutation) SetField(name string, value ent.Value) error {
	switch name {
	case media.FieldIdmedia:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetIdmedia(v)
		return nil
	case media.FieldCreationdatemedia:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreationdatemedia(v)
		return nil
	case media.FieldDatamedia:
		v, ok := value.([]byte)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDatamedia(v)
		return nil
	case media.FieldNomemedia:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetNomemedia(v)
		return nil
	case media.FieldMimemedia:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetMimemedia(v)
		return nil
	}
	return fmt.Errorf("unknown Media field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *MediaMutation) AddedFields() []string {
	var fields []string
	if m.addidmedia != nil {
		fields = append(fields, media.FieldIdmedia)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *MediaMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case media.FieldIdmedia:
		return m.AddedIdmedia()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *MediaMutation) AddField(name string, value ent.Value) error {
	switch name {
	case media.FieldIdmedia:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddIdmedia(v)
		return nil
	}
	return fmt.Errorf("unknown Media numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *MediaMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *MediaMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *MediaMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Media nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *MediaMutation) ResetField(name string) error {
	switch name {
	case media.FieldIdmedia:
		m.ResetIdmedia()
		return nil
	case media.FieldCreationdatemedia:
		m.ResetCreationdatemedia()
		return nil
	case media.FieldDatamedia:
		m.ResetDatamedia()
		return nil
	case media.FieldNomemedia:
		m.ResetNomemedia()
		return nil
	case media.FieldMimemedia:
		m.ResetMimemedia()
		return nil
	}
	return fmt.Errorf("unknown Media field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *MediaMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *MediaMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *MediaMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *MediaMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *MediaMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *MediaMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *MediaMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Media unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *MediaMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Media edge %s", name)
}
