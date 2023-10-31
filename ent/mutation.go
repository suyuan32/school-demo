// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	uuid "github.com/gofrs/uuid/v5"
	"github.com/suyuan32/simple-admin-school/ent/predicate"
	"github.com/suyuan32/simple-admin-school/ent/student"
	"github.com/suyuan32/simple-admin-school/ent/teacher"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeStudent = "Student"
	TypeTeacher = "Teacher"
)

// StudentMutation represents an operation that mutates the Student nodes in the graph.
type StudentMutation struct {
	config
	op            Op
	typ           string
	id            *uuid.UUID
	created_at    *time.Time
	updated_at    *time.Time
	name          *string
	age           *int16
	addage        *int16
	address       *string
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Student, error)
	predicates    []predicate.Student
}

var _ ent.Mutation = (*StudentMutation)(nil)

// studentOption allows management of the mutation configuration using functional options.
type studentOption func(*StudentMutation)

// newStudentMutation creates new mutation for the Student entity.
func newStudentMutation(c config, op Op, opts ...studentOption) *StudentMutation {
	m := &StudentMutation{
		config:        c,
		op:            op,
		typ:           TypeStudent,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withStudentID sets the ID field of the mutation.
func withStudentID(id uuid.UUID) studentOption {
	return func(m *StudentMutation) {
		var (
			err   error
			once  sync.Once
			value *Student
		)
		m.oldValue = func(ctx context.Context) (*Student, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Student.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withStudent sets the old Student of the mutation.
func withStudent(node *Student) studentOption {
	return func(m *StudentMutation) {
		m.oldValue = func(context.Context) (*Student, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m StudentMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m StudentMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Student entities.
func (m *StudentMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *StudentMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *StudentMutation) IDs(ctx context.Context) ([]uuid.UUID, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []uuid.UUID{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Student.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetCreatedAt sets the "created_at" field.
func (m *StudentMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *StudentMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Student entity.
// If the Student object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *StudentMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *StudentMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *StudentMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *StudentMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the Student entity.
// If the Student object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *StudentMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *StudentMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// SetName sets the "name" field.
func (m *StudentMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *StudentMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Student entity.
// If the Student object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *StudentMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *StudentMutation) ResetName() {
	m.name = nil
}

// SetAge sets the "age" field.
func (m *StudentMutation) SetAge(i int16) {
	m.age = &i
	m.addage = nil
}

// Age returns the value of the "age" field in the mutation.
func (m *StudentMutation) Age() (r int16, exists bool) {
	v := m.age
	if v == nil {
		return
	}
	return *v, true
}

// OldAge returns the old "age" field's value of the Student entity.
// If the Student object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *StudentMutation) OldAge(ctx context.Context) (v int16, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldAge is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldAge requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAge: %w", err)
	}
	return oldValue.Age, nil
}

// AddAge adds i to the "age" field.
func (m *StudentMutation) AddAge(i int16) {
	if m.addage != nil {
		*m.addage += i
	} else {
		m.addage = &i
	}
}

// AddedAge returns the value that was added to the "age" field in this mutation.
func (m *StudentMutation) AddedAge() (r int16, exists bool) {
	v := m.addage
	if v == nil {
		return
	}
	return *v, true
}

// ResetAge resets all changes to the "age" field.
func (m *StudentMutation) ResetAge() {
	m.age = nil
	m.addage = nil
}

// SetAddress sets the "address" field.
func (m *StudentMutation) SetAddress(s string) {
	m.address = &s
}

// Address returns the value of the "address" field in the mutation.
func (m *StudentMutation) Address() (r string, exists bool) {
	v := m.address
	if v == nil {
		return
	}
	return *v, true
}

// OldAddress returns the old "address" field's value of the Student entity.
// If the Student object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *StudentMutation) OldAddress(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldAddress is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldAddress requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAddress: %w", err)
	}
	return oldValue.Address, nil
}

// ClearAddress clears the value of the "address" field.
func (m *StudentMutation) ClearAddress() {
	m.address = nil
	m.clearedFields[student.FieldAddress] = struct{}{}
}

// AddressCleared returns if the "address" field was cleared in this mutation.
func (m *StudentMutation) AddressCleared() bool {
	_, ok := m.clearedFields[student.FieldAddress]
	return ok
}

// ResetAddress resets all changes to the "address" field.
func (m *StudentMutation) ResetAddress() {
	m.address = nil
	delete(m.clearedFields, student.FieldAddress)
}

// Where appends a list predicates to the StudentMutation builder.
func (m *StudentMutation) Where(ps ...predicate.Student) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the StudentMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *StudentMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Student, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *StudentMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *StudentMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Student).
func (m *StudentMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *StudentMutation) Fields() []string {
	fields := make([]string, 0, 5)
	if m.created_at != nil {
		fields = append(fields, student.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, student.FieldUpdatedAt)
	}
	if m.name != nil {
		fields = append(fields, student.FieldName)
	}
	if m.age != nil {
		fields = append(fields, student.FieldAge)
	}
	if m.address != nil {
		fields = append(fields, student.FieldAddress)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *StudentMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case student.FieldCreatedAt:
		return m.CreatedAt()
	case student.FieldUpdatedAt:
		return m.UpdatedAt()
	case student.FieldName:
		return m.Name()
	case student.FieldAge:
		return m.Age()
	case student.FieldAddress:
		return m.Address()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *StudentMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case student.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case student.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	case student.FieldName:
		return m.OldName(ctx)
	case student.FieldAge:
		return m.OldAge(ctx)
	case student.FieldAddress:
		return m.OldAddress(ctx)
	}
	return nil, fmt.Errorf("unknown Student field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *StudentMutation) SetField(name string, value ent.Value) error {
	switch name {
	case student.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case student.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	case student.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case student.FieldAge:
		v, ok := value.(int16)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAge(v)
		return nil
	case student.FieldAddress:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAddress(v)
		return nil
	}
	return fmt.Errorf("unknown Student field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *StudentMutation) AddedFields() []string {
	var fields []string
	if m.addage != nil {
		fields = append(fields, student.FieldAge)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *StudentMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case student.FieldAge:
		return m.AddedAge()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *StudentMutation) AddField(name string, value ent.Value) error {
	switch name {
	case student.FieldAge:
		v, ok := value.(int16)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddAge(v)
		return nil
	}
	return fmt.Errorf("unknown Student numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *StudentMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(student.FieldAddress) {
		fields = append(fields, student.FieldAddress)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *StudentMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *StudentMutation) ClearField(name string) error {
	switch name {
	case student.FieldAddress:
		m.ClearAddress()
		return nil
	}
	return fmt.Errorf("unknown Student nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *StudentMutation) ResetField(name string) error {
	switch name {
	case student.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case student.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	case student.FieldName:
		m.ResetName()
		return nil
	case student.FieldAge:
		m.ResetAge()
		return nil
	case student.FieldAddress:
		m.ResetAddress()
		return nil
	}
	return fmt.Errorf("unknown Student field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *StudentMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *StudentMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *StudentMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *StudentMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *StudentMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *StudentMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *StudentMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Student unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *StudentMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Student edge %s", name)
}

// TeacherMutation represents an operation that mutates the Teacher nodes in the graph.
type TeacherMutation struct {
	config
	op            Op
	typ           string
	id            *uint64
	created_at    *time.Time
	updated_at    *time.Time
	name          *string
	age           *int16
	addage        *int16
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Teacher, error)
	predicates    []predicate.Teacher
}

var _ ent.Mutation = (*TeacherMutation)(nil)

// teacherOption allows management of the mutation configuration using functional options.
type teacherOption func(*TeacherMutation)

// newTeacherMutation creates new mutation for the Teacher entity.
func newTeacherMutation(c config, op Op, opts ...teacherOption) *TeacherMutation {
	m := &TeacherMutation{
		config:        c,
		op:            op,
		typ:           TypeTeacher,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withTeacherID sets the ID field of the mutation.
func withTeacherID(id uint64) teacherOption {
	return func(m *TeacherMutation) {
		var (
			err   error
			once  sync.Once
			value *Teacher
		)
		m.oldValue = func(ctx context.Context) (*Teacher, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Teacher.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withTeacher sets the old Teacher of the mutation.
func withTeacher(node *Teacher) teacherOption {
	return func(m *TeacherMutation) {
		m.oldValue = func(context.Context) (*Teacher, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m TeacherMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m TeacherMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Teacher entities.
func (m *TeacherMutation) SetID(id uint64) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *TeacherMutation) ID() (id uint64, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *TeacherMutation) IDs(ctx context.Context) ([]uint64, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []uint64{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Teacher.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetCreatedAt sets the "created_at" field.
func (m *TeacherMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *TeacherMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Teacher entity.
// If the Teacher object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TeacherMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *TeacherMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *TeacherMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *TeacherMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the Teacher entity.
// If the Teacher object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TeacherMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *TeacherMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// SetName sets the "name" field.
func (m *TeacherMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *TeacherMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Teacher entity.
// If the Teacher object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TeacherMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *TeacherMutation) ResetName() {
	m.name = nil
}

// SetAge sets the "age" field.
func (m *TeacherMutation) SetAge(i int16) {
	m.age = &i
	m.addage = nil
}

// Age returns the value of the "age" field in the mutation.
func (m *TeacherMutation) Age() (r int16, exists bool) {
	v := m.age
	if v == nil {
		return
	}
	return *v, true
}

// OldAge returns the old "age" field's value of the Teacher entity.
// If the Teacher object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TeacherMutation) OldAge(ctx context.Context) (v int16, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldAge is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldAge requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAge: %w", err)
	}
	return oldValue.Age, nil
}

// AddAge adds i to the "age" field.
func (m *TeacherMutation) AddAge(i int16) {
	if m.addage != nil {
		*m.addage += i
	} else {
		m.addage = &i
	}
}

// AddedAge returns the value that was added to the "age" field in this mutation.
func (m *TeacherMutation) AddedAge() (r int16, exists bool) {
	v := m.addage
	if v == nil {
		return
	}
	return *v, true
}

// ResetAge resets all changes to the "age" field.
func (m *TeacherMutation) ResetAge() {
	m.age = nil
	m.addage = nil
}

// Where appends a list predicates to the TeacherMutation builder.
func (m *TeacherMutation) Where(ps ...predicate.Teacher) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the TeacherMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *TeacherMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Teacher, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *TeacherMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *TeacherMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Teacher).
func (m *TeacherMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *TeacherMutation) Fields() []string {
	fields := make([]string, 0, 4)
	if m.created_at != nil {
		fields = append(fields, teacher.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, teacher.FieldUpdatedAt)
	}
	if m.name != nil {
		fields = append(fields, teacher.FieldName)
	}
	if m.age != nil {
		fields = append(fields, teacher.FieldAge)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *TeacherMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case teacher.FieldCreatedAt:
		return m.CreatedAt()
	case teacher.FieldUpdatedAt:
		return m.UpdatedAt()
	case teacher.FieldName:
		return m.Name()
	case teacher.FieldAge:
		return m.Age()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *TeacherMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case teacher.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case teacher.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	case teacher.FieldName:
		return m.OldName(ctx)
	case teacher.FieldAge:
		return m.OldAge(ctx)
	}
	return nil, fmt.Errorf("unknown Teacher field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *TeacherMutation) SetField(name string, value ent.Value) error {
	switch name {
	case teacher.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case teacher.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	case teacher.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case teacher.FieldAge:
		v, ok := value.(int16)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAge(v)
		return nil
	}
	return fmt.Errorf("unknown Teacher field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *TeacherMutation) AddedFields() []string {
	var fields []string
	if m.addage != nil {
		fields = append(fields, teacher.FieldAge)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *TeacherMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case teacher.FieldAge:
		return m.AddedAge()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *TeacherMutation) AddField(name string, value ent.Value) error {
	switch name {
	case teacher.FieldAge:
		v, ok := value.(int16)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddAge(v)
		return nil
	}
	return fmt.Errorf("unknown Teacher numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *TeacherMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *TeacherMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *TeacherMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Teacher nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *TeacherMutation) ResetField(name string) error {
	switch name {
	case teacher.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case teacher.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	case teacher.FieldName:
		m.ResetName()
		return nil
	case teacher.FieldAge:
		m.ResetAge()
		return nil
	}
	return fmt.Errorf("unknown Teacher field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *TeacherMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *TeacherMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *TeacherMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *TeacherMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *TeacherMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *TeacherMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *TeacherMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Teacher unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *TeacherMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Teacher edge %s", name)
}
