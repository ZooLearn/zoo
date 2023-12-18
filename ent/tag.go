// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/ZooLearn/zoo/ent/tag"
)

// Tag is the model entity for the Tag schema.
type Tag struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// Create Time
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Update Time
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeleteAt holds the value of the "delete_at" field.
	DeleteAt *time.Time `json:"delete_at,omitempty"`
	// Key holds the value of the "key" field.
	Key string `json:"key,omitempty"`
	// Alias holds the value of the "alias" field.
	Alias string `json:"alias,omitempty"`
	// Color holds the value of the "color" field.
	Color string `json:"color,omitempty"`
	// ImageUrl holds the value of the "imageUrl" field.
	ImageUrl     string `json:"imageUrl,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Tag) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case tag.FieldID:
			values[i] = new(sql.NullInt64)
		case tag.FieldKey, tag.FieldAlias, tag.FieldColor, tag.FieldImageUrl:
			values[i] = new(sql.NullString)
		case tag.FieldCreatedAt, tag.FieldUpdatedAt, tag.FieldDeleteAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Tag fields.
func (t *Tag) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case tag.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = uint64(value.Int64)
		case tag.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				t.CreatedAt = value.Time
			}
		case tag.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				t.UpdatedAt = value.Time
			}
		case tag.FieldDeleteAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field delete_at", values[i])
			} else if value.Valid {
				t.DeleteAt = new(time.Time)
				*t.DeleteAt = value.Time
			}
		case tag.FieldKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field key", values[i])
			} else if value.Valid {
				t.Key = value.String
			}
		case tag.FieldAlias:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field alias", values[i])
			} else if value.Valid {
				t.Alias = value.String
			}
		case tag.FieldColor:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field color", values[i])
			} else if value.Valid {
				t.Color = value.String
			}
		case tag.FieldImageUrl:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field imageUrl", values[i])
			} else if value.Valid {
				t.ImageUrl = value.String
			}
		default:
			t.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Tag.
// This includes values selected through modifiers, order, etc.
func (t *Tag) Value(name string) (ent.Value, error) {
	return t.selectValues.Get(name)
}

// Update returns a builder for updating this Tag.
// Note that you need to call Tag.Unwrap() before calling this method if this Tag
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Tag) Update() *TagUpdateOne {
	return NewTagClient(t.config).UpdateOne(t)
}

// Unwrap unwraps the Tag entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Tag) Unwrap() *Tag {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Tag is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Tag) String() string {
	var builder strings.Builder
	builder.WriteString("Tag(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("created_at=")
	builder.WriteString(t.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(t.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := t.DeleteAt; v != nil {
		builder.WriteString("delete_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("key=")
	builder.WriteString(t.Key)
	builder.WriteString(", ")
	builder.WriteString("alias=")
	builder.WriteString(t.Alias)
	builder.WriteString(", ")
	builder.WriteString("color=")
	builder.WriteString(t.Color)
	builder.WriteString(", ")
	builder.WriteString("imageUrl=")
	builder.WriteString(t.ImageUrl)
	builder.WriteByte(')')
	return builder.String()
}

// Tags is a parsable slice of Tag.
type Tags []*Tag
