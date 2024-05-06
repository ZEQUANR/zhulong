// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/ZEQUANR/zhulong/ent/teachers"
	"github.com/ZEQUANR/zhulong/ent/user"
)

// Teachers is the model entity for the Teachers schema.
type Teachers struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// College holds the value of the "college" field.
	College string `json:"college,omitempty"`
	// Phone holds the value of the "phone" field.
	Phone string `json:"phone,omitempty"`
	// Identity holds the value of the "identity" field.
	Identity string `json:"identity,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TeachersQuery when eager-loading is set.
	Edges         TeachersEdges `json:"edges"`
	user_teachers *int
	selectValues  sql.SelectValues
}

// TeachersEdges holds the relations/edges for other nodes in the graph.
type TeachersEdges struct {
	// Users holds the value of the users edge.
	Users *User `json:"users,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UsersOrErr returns the Users value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TeachersEdges) UsersOrErr() (*User, error) {
	if e.Users != nil {
		return e.Users, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "users"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Teachers) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case teachers.FieldID:
			values[i] = new(sql.NullInt64)
		case teachers.FieldName, teachers.FieldCollege, teachers.FieldPhone, teachers.FieldIdentity:
			values[i] = new(sql.NullString)
		case teachers.ForeignKeys[0]: // user_teachers
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Teachers fields.
func (t *Teachers) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case teachers.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = int(value.Int64)
		case teachers.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				t.Name = value.String
			}
		case teachers.FieldCollege:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field college", values[i])
			} else if value.Valid {
				t.College = value.String
			}
		case teachers.FieldPhone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone", values[i])
			} else if value.Valid {
				t.Phone = value.String
			}
		case teachers.FieldIdentity:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field identity", values[i])
			} else if value.Valid {
				t.Identity = value.String
			}
		case teachers.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_teachers", value)
			} else if value.Valid {
				t.user_teachers = new(int)
				*t.user_teachers = int(value.Int64)
			}
		default:
			t.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Teachers.
// This includes values selected through modifiers, order, etc.
func (t *Teachers) Value(name string) (ent.Value, error) {
	return t.selectValues.Get(name)
}

// QueryUsers queries the "users" edge of the Teachers entity.
func (t *Teachers) QueryUsers() *UserQuery {
	return NewTeachersClient(t.config).QueryUsers(t)
}

// Update returns a builder for updating this Teachers.
// Note that you need to call Teachers.Unwrap() before calling this method if this Teachers
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Teachers) Update() *TeachersUpdateOne {
	return NewTeachersClient(t.config).UpdateOne(t)
}

// Unwrap unwraps the Teachers entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Teachers) Unwrap() *Teachers {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Teachers is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Teachers) String() string {
	var builder strings.Builder
	builder.WriteString("Teachers(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("name=")
	builder.WriteString(t.Name)
	builder.WriteString(", ")
	builder.WriteString("college=")
	builder.WriteString(t.College)
	builder.WriteString(", ")
	builder.WriteString("phone=")
	builder.WriteString(t.Phone)
	builder.WriteString(", ")
	builder.WriteString("identity=")
	builder.WriteString(t.Identity)
	builder.WriteByte(')')
	return builder.String()
}

// TeachersSlice is a parsable slice of Teachers.
type TeachersSlice []*Teachers