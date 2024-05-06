// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/ZEQUANR/zhulong/ent/administrators"
	"github.com/ZEQUANR/zhulong/ent/students"
	"github.com/ZEQUANR/zhulong/ent/teachers"
	"github.com/ZEQUANR/zhulong/ent/user"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Account holds the value of the "account" field.
	Account string `json:"account,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"password,omitempty"`
	// Role holds the value of the "role" field.
	Role int `json:"role,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges        UserEdges `json:"edges"`
	selectValues sql.SelectValues
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Administrators holds the value of the administrators edge.
	Administrators *Administrators `json:"administrators,omitempty"`
	// Students holds the value of the students edge.
	Students *Students `json:"students,omitempty"`
	// Teachers holds the value of the teachers edge.
	Teachers *Teachers `json:"teachers,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// AdministratorsOrErr returns the Administrators value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) AdministratorsOrErr() (*Administrators, error) {
	if e.Administrators != nil {
		return e.Administrators, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: administrators.Label}
	}
	return nil, &NotLoadedError{edge: "administrators"}
}

// StudentsOrErr returns the Students value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) StudentsOrErr() (*Students, error) {
	if e.Students != nil {
		return e.Students, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: students.Label}
	}
	return nil, &NotLoadedError{edge: "students"}
}

// TeachersOrErr returns the Teachers value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) TeachersOrErr() (*Teachers, error) {
	if e.Teachers != nil {
		return e.Teachers, nil
	} else if e.loadedTypes[2] {
		return nil, &NotFoundError{label: teachers.Label}
	}
	return nil, &NotLoadedError{edge: "teachers"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldID, user.FieldRole:
			values[i] = new(sql.NullInt64)
		case user.FieldAccount, user.FieldPassword:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int(value.Int64)
		case user.FieldAccount:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field account", values[i])
			} else if value.Valid {
				u.Account = value.String
			}
		case user.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				u.Password = value.String
			}
		case user.FieldRole:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field role", values[i])
			} else if value.Valid {
				u.Role = int(value.Int64)
			}
		default:
			u.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the User.
// This includes values selected through modifiers, order, etc.
func (u *User) Value(name string) (ent.Value, error) {
	return u.selectValues.Get(name)
}

// QueryAdministrators queries the "administrators" edge of the User entity.
func (u *User) QueryAdministrators() *AdministratorsQuery {
	return NewUserClient(u.config).QueryAdministrators(u)
}

// QueryStudents queries the "students" edge of the User entity.
func (u *User) QueryStudents() *StudentsQuery {
	return NewUserClient(u.config).QueryStudents(u)
}

// QueryTeachers queries the "teachers" edge of the User entity.
func (u *User) QueryTeachers() *TeachersQuery {
	return NewUserClient(u.config).QueryTeachers(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return NewUserClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("account=")
	builder.WriteString(u.Account)
	builder.WriteString(", ")
	builder.WriteString("password=")
	builder.WriteString(u.Password)
	builder.WriteString(", ")
	builder.WriteString("role=")
	builder.WriteString(fmt.Sprintf("%v", u.Role))
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User