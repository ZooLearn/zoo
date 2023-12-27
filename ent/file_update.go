// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ZooLearn/zoo/ent/file"
	"github.com/ZooLearn/zoo/ent/predicate"
)

// FileUpdate is the builder for updating File entities.
type FileUpdate struct {
	config
	hooks    []Hook
	mutation *FileMutation
}

// Where appends a list predicates to the FileUpdate builder.
func (fu *FileUpdate) Where(ps ...predicate.File) *FileUpdate {
	fu.mutation.Where(ps...)
	return fu
}

// SetUpdatedAt sets the "updated_at" field.
func (fu *FileUpdate) SetUpdatedAt(t time.Time) *FileUpdate {
	fu.mutation.SetUpdatedAt(t)
	return fu
}

// SetDeleteAt sets the "delete_at" field.
func (fu *FileUpdate) SetDeleteAt(t time.Time) *FileUpdate {
	fu.mutation.SetDeleteAt(t)
	return fu
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (fu *FileUpdate) SetNillableDeleteAt(t *time.Time) *FileUpdate {
	if t != nil {
		fu.SetDeleteAt(*t)
	}
	return fu
}

// ClearDeleteAt clears the value of the "delete_at" field.
func (fu *FileUpdate) ClearDeleteAt() *FileUpdate {
	fu.mutation.ClearDeleteAt()
	return fu
}

// SetCode sets the "code" field.
func (fu *FileUpdate) SetCode(s string) *FileUpdate {
	fu.mutation.SetCode(s)
	return fu
}

// SetNillableCode sets the "code" field if the given value is not nil.
func (fu *FileUpdate) SetNillableCode(s *string) *FileUpdate {
	if s != nil {
		fu.SetCode(*s)
	}
	return fu
}

// SetName sets the "name" field.
func (fu *FileUpdate) SetName(s string) *FileUpdate {
	fu.mutation.SetName(s)
	return fu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (fu *FileUpdate) SetNillableName(s *string) *FileUpdate {
	if s != nil {
		fu.SetName(*s)
	}
	return fu
}

// SetFileType sets the "file_type" field.
func (fu *FileUpdate) SetFileType(u uint8) *FileUpdate {
	fu.mutation.ResetFileType()
	fu.mutation.SetFileType(u)
	return fu
}

// SetNillableFileType sets the "file_type" field if the given value is not nil.
func (fu *FileUpdate) SetNillableFileType(u *uint8) *FileUpdate {
	if u != nil {
		fu.SetFileType(*u)
	}
	return fu
}

// AddFileType adds u to the "file_type" field.
func (fu *FileUpdate) AddFileType(u int8) *FileUpdate {
	fu.mutation.AddFileType(u)
	return fu
}

// SetSize sets the "size" field.
func (fu *FileUpdate) SetSize(u uint64) *FileUpdate {
	fu.mutation.ResetSize()
	fu.mutation.SetSize(u)
	return fu
}

// SetNillableSize sets the "size" field if the given value is not nil.
func (fu *FileUpdate) SetNillableSize(u *uint64) *FileUpdate {
	if u != nil {
		fu.SetSize(*u)
	}
	return fu
}

// AddSize adds u to the "size" field.
func (fu *FileUpdate) AddSize(u int64) *FileUpdate {
	fu.mutation.AddSize(u)
	return fu
}

// SetPath sets the "path" field.
func (fu *FileUpdate) SetPath(s string) *FileUpdate {
	fu.mutation.SetPath(s)
	return fu
}

// SetNillablePath sets the "path" field if the given value is not nil.
func (fu *FileUpdate) SetNillablePath(s *string) *FileUpdate {
	if s != nil {
		fu.SetPath(*s)
	}
	return fu
}

// ClearPath clears the value of the "path" field.
func (fu *FileUpdate) ClearPath() *FileUpdate {
	fu.mutation.ClearPath()
	return fu
}

// SetTempPath sets the "temp_path" field.
func (fu *FileUpdate) SetTempPath(s string) *FileUpdate {
	fu.mutation.SetTempPath(s)
	return fu
}

// SetNillableTempPath sets the "temp_path" field if the given value is not nil.
func (fu *FileUpdate) SetNillableTempPath(s *string) *FileUpdate {
	if s != nil {
		fu.SetTempPath(*s)
	}
	return fu
}

// SetUserID sets the "user_id" field.
func (fu *FileUpdate) SetUserID(s string) *FileUpdate {
	fu.mutation.SetUserID(s)
	return fu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (fu *FileUpdate) SetNillableUserID(s *string) *FileUpdate {
	if s != nil {
		fu.SetUserID(*s)
	}
	return fu
}

// Mutation returns the FileMutation object of the builder.
func (fu *FileUpdate) Mutation() *FileMutation {
	return fu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (fu *FileUpdate) Save(ctx context.Context) (int, error) {
	fu.defaults()
	return withHooks(ctx, fu.sqlSave, fu.mutation, fu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fu *FileUpdate) SaveX(ctx context.Context) int {
	affected, err := fu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fu *FileUpdate) Exec(ctx context.Context) error {
	_, err := fu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fu *FileUpdate) ExecX(ctx context.Context) {
	if err := fu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fu *FileUpdate) defaults() {
	if _, ok := fu.mutation.UpdatedAt(); !ok {
		v := file.UpdateDefaultUpdatedAt()
		fu.mutation.SetUpdatedAt(v)
	}
}

func (fu *FileUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(file.Table, file.Columns, sqlgraph.NewFieldSpec(file.FieldID, field.TypeUUID))
	if ps := fu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fu.mutation.UpdatedAt(); ok {
		_spec.SetField(file.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := fu.mutation.DeleteAt(); ok {
		_spec.SetField(file.FieldDeleteAt, field.TypeTime, value)
	}
	if fu.mutation.DeleteAtCleared() {
		_spec.ClearField(file.FieldDeleteAt, field.TypeTime)
	}
	if value, ok := fu.mutation.Code(); ok {
		_spec.SetField(file.FieldCode, field.TypeString, value)
	}
	if value, ok := fu.mutation.Name(); ok {
		_spec.SetField(file.FieldName, field.TypeString, value)
	}
	if value, ok := fu.mutation.FileType(); ok {
		_spec.SetField(file.FieldFileType, field.TypeUint8, value)
	}
	if value, ok := fu.mutation.AddedFileType(); ok {
		_spec.AddField(file.FieldFileType, field.TypeUint8, value)
	}
	if value, ok := fu.mutation.Size(); ok {
		_spec.SetField(file.FieldSize, field.TypeUint64, value)
	}
	if value, ok := fu.mutation.AddedSize(); ok {
		_spec.AddField(file.FieldSize, field.TypeUint64, value)
	}
	if value, ok := fu.mutation.Path(); ok {
		_spec.SetField(file.FieldPath, field.TypeString, value)
	}
	if fu.mutation.PathCleared() {
		_spec.ClearField(file.FieldPath, field.TypeString)
	}
	if value, ok := fu.mutation.TempPath(); ok {
		_spec.SetField(file.FieldTempPath, field.TypeString, value)
	}
	if value, ok := fu.mutation.UserID(); ok {
		_spec.SetField(file.FieldUserID, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, fu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{file.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	fu.mutation.done = true
	return n, nil
}

// FileUpdateOne is the builder for updating a single File entity.
type FileUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *FileMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (fuo *FileUpdateOne) SetUpdatedAt(t time.Time) *FileUpdateOne {
	fuo.mutation.SetUpdatedAt(t)
	return fuo
}

// SetDeleteAt sets the "delete_at" field.
func (fuo *FileUpdateOne) SetDeleteAt(t time.Time) *FileUpdateOne {
	fuo.mutation.SetDeleteAt(t)
	return fuo
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (fuo *FileUpdateOne) SetNillableDeleteAt(t *time.Time) *FileUpdateOne {
	if t != nil {
		fuo.SetDeleteAt(*t)
	}
	return fuo
}

// ClearDeleteAt clears the value of the "delete_at" field.
func (fuo *FileUpdateOne) ClearDeleteAt() *FileUpdateOne {
	fuo.mutation.ClearDeleteAt()
	return fuo
}

// SetCode sets the "code" field.
func (fuo *FileUpdateOne) SetCode(s string) *FileUpdateOne {
	fuo.mutation.SetCode(s)
	return fuo
}

// SetNillableCode sets the "code" field if the given value is not nil.
func (fuo *FileUpdateOne) SetNillableCode(s *string) *FileUpdateOne {
	if s != nil {
		fuo.SetCode(*s)
	}
	return fuo
}

// SetName sets the "name" field.
func (fuo *FileUpdateOne) SetName(s string) *FileUpdateOne {
	fuo.mutation.SetName(s)
	return fuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (fuo *FileUpdateOne) SetNillableName(s *string) *FileUpdateOne {
	if s != nil {
		fuo.SetName(*s)
	}
	return fuo
}

// SetFileType sets the "file_type" field.
func (fuo *FileUpdateOne) SetFileType(u uint8) *FileUpdateOne {
	fuo.mutation.ResetFileType()
	fuo.mutation.SetFileType(u)
	return fuo
}

// SetNillableFileType sets the "file_type" field if the given value is not nil.
func (fuo *FileUpdateOne) SetNillableFileType(u *uint8) *FileUpdateOne {
	if u != nil {
		fuo.SetFileType(*u)
	}
	return fuo
}

// AddFileType adds u to the "file_type" field.
func (fuo *FileUpdateOne) AddFileType(u int8) *FileUpdateOne {
	fuo.mutation.AddFileType(u)
	return fuo
}

// SetSize sets the "size" field.
func (fuo *FileUpdateOne) SetSize(u uint64) *FileUpdateOne {
	fuo.mutation.ResetSize()
	fuo.mutation.SetSize(u)
	return fuo
}

// SetNillableSize sets the "size" field if the given value is not nil.
func (fuo *FileUpdateOne) SetNillableSize(u *uint64) *FileUpdateOne {
	if u != nil {
		fuo.SetSize(*u)
	}
	return fuo
}

// AddSize adds u to the "size" field.
func (fuo *FileUpdateOne) AddSize(u int64) *FileUpdateOne {
	fuo.mutation.AddSize(u)
	return fuo
}

// SetPath sets the "path" field.
func (fuo *FileUpdateOne) SetPath(s string) *FileUpdateOne {
	fuo.mutation.SetPath(s)
	return fuo
}

// SetNillablePath sets the "path" field if the given value is not nil.
func (fuo *FileUpdateOne) SetNillablePath(s *string) *FileUpdateOne {
	if s != nil {
		fuo.SetPath(*s)
	}
	return fuo
}

// ClearPath clears the value of the "path" field.
func (fuo *FileUpdateOne) ClearPath() *FileUpdateOne {
	fuo.mutation.ClearPath()
	return fuo
}

// SetTempPath sets the "temp_path" field.
func (fuo *FileUpdateOne) SetTempPath(s string) *FileUpdateOne {
	fuo.mutation.SetTempPath(s)
	return fuo
}

// SetNillableTempPath sets the "temp_path" field if the given value is not nil.
func (fuo *FileUpdateOne) SetNillableTempPath(s *string) *FileUpdateOne {
	if s != nil {
		fuo.SetTempPath(*s)
	}
	return fuo
}

// SetUserID sets the "user_id" field.
func (fuo *FileUpdateOne) SetUserID(s string) *FileUpdateOne {
	fuo.mutation.SetUserID(s)
	return fuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (fuo *FileUpdateOne) SetNillableUserID(s *string) *FileUpdateOne {
	if s != nil {
		fuo.SetUserID(*s)
	}
	return fuo
}

// Mutation returns the FileMutation object of the builder.
func (fuo *FileUpdateOne) Mutation() *FileMutation {
	return fuo.mutation
}

// Where appends a list predicates to the FileUpdate builder.
func (fuo *FileUpdateOne) Where(ps ...predicate.File) *FileUpdateOne {
	fuo.mutation.Where(ps...)
	return fuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (fuo *FileUpdateOne) Select(field string, fields ...string) *FileUpdateOne {
	fuo.fields = append([]string{field}, fields...)
	return fuo
}

// Save executes the query and returns the updated File entity.
func (fuo *FileUpdateOne) Save(ctx context.Context) (*File, error) {
	fuo.defaults()
	return withHooks(ctx, fuo.sqlSave, fuo.mutation, fuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fuo *FileUpdateOne) SaveX(ctx context.Context) *File {
	node, err := fuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (fuo *FileUpdateOne) Exec(ctx context.Context) error {
	_, err := fuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fuo *FileUpdateOne) ExecX(ctx context.Context) {
	if err := fuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fuo *FileUpdateOne) defaults() {
	if _, ok := fuo.mutation.UpdatedAt(); !ok {
		v := file.UpdateDefaultUpdatedAt()
		fuo.mutation.SetUpdatedAt(v)
	}
}

func (fuo *FileUpdateOne) sqlSave(ctx context.Context) (_node *File, err error) {
	_spec := sqlgraph.NewUpdateSpec(file.Table, file.Columns, sqlgraph.NewFieldSpec(file.FieldID, field.TypeUUID))
	id, ok := fuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "File.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := fuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, file.FieldID)
		for _, f := range fields {
			if !file.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != file.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := fuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fuo.mutation.UpdatedAt(); ok {
		_spec.SetField(file.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := fuo.mutation.DeleteAt(); ok {
		_spec.SetField(file.FieldDeleteAt, field.TypeTime, value)
	}
	if fuo.mutation.DeleteAtCleared() {
		_spec.ClearField(file.FieldDeleteAt, field.TypeTime)
	}
	if value, ok := fuo.mutation.Code(); ok {
		_spec.SetField(file.FieldCode, field.TypeString, value)
	}
	if value, ok := fuo.mutation.Name(); ok {
		_spec.SetField(file.FieldName, field.TypeString, value)
	}
	if value, ok := fuo.mutation.FileType(); ok {
		_spec.SetField(file.FieldFileType, field.TypeUint8, value)
	}
	if value, ok := fuo.mutation.AddedFileType(); ok {
		_spec.AddField(file.FieldFileType, field.TypeUint8, value)
	}
	if value, ok := fuo.mutation.Size(); ok {
		_spec.SetField(file.FieldSize, field.TypeUint64, value)
	}
	if value, ok := fuo.mutation.AddedSize(); ok {
		_spec.AddField(file.FieldSize, field.TypeUint64, value)
	}
	if value, ok := fuo.mutation.Path(); ok {
		_spec.SetField(file.FieldPath, field.TypeString, value)
	}
	if fuo.mutation.PathCleared() {
		_spec.ClearField(file.FieldPath, field.TypeString)
	}
	if value, ok := fuo.mutation.TempPath(); ok {
		_spec.SetField(file.FieldTempPath, field.TypeString, value)
	}
	if value, ok := fuo.mutation.UserID(); ok {
		_spec.SetField(file.FieldUserID, field.TypeString, value)
	}
	_node = &File{config: fuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, fuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{file.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	fuo.mutation.done = true
	return _node, nil
}