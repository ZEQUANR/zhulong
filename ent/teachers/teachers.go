// Code generated by ent, DO NOT EDIT.

package teachers

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the teachers type in the database.
	Label = "teachers"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldCollege holds the string denoting the college field in the database.
	FieldCollege = "college"
	// FieldPhone holds the string denoting the phone field in the database.
	FieldPhone = "phone"
	// FieldIdentity holds the string denoting the identity field in the database.
	FieldIdentity = "identity"
	// EdgeUsers holds the string denoting the users edge name in mutations.
	EdgeUsers = "users"
	// Table holds the table name of the teachers in the database.
	Table = "teachers"
	// UsersTable is the table that holds the users relation/edge.
	UsersTable = "teachers"
	// UsersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UsersInverseTable = "users"
	// UsersColumn is the table column denoting the users relation/edge.
	UsersColumn = "user_teachers"
)

// Columns holds all SQL columns for teachers fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldCollege,
	FieldPhone,
	FieldIdentity,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "teachers"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_teachers",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the Teachers queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByCollege orders the results by the college field.
func ByCollege(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCollege, opts...).ToFunc()
}

// ByPhone orders the results by the phone field.
func ByPhone(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPhone, opts...).ToFunc()
}

// ByIdentity orders the results by the identity field.
func ByIdentity(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIdentity, opts...).ToFunc()
}

// ByUsersField orders the results by users field.
func ByUsersField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUsersStep(), sql.OrderByField(field, opts...))
	}
}
func newUsersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UsersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, UsersTable, UsersColumn),
	)
}
