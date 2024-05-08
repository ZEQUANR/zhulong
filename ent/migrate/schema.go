// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AdministratorsColumns holds the columns for the "administrators" table.
	AdministratorsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "college", Type: field.TypeString},
		{Name: "phone", Type: field.TypeString},
		{Name: "identity", Type: field.TypeString},
		{Name: "user_administrators", Type: field.TypeInt, Unique: true},
	}
	// AdministratorsTable holds the schema information for the "administrators" table.
	AdministratorsTable = &schema.Table{
		Name:       "administrators",
		Columns:    AdministratorsColumns,
		PrimaryKey: []*schema.Column{AdministratorsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "administrators_users_administrators",
				Columns:    []*schema.Column{AdministratorsColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// StudentsColumns holds the columns for the "students" table.
	StudentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "college", Type: field.TypeString},
		{Name: "phone", Type: field.TypeString},
		{Name: "subject", Type: field.TypeString},
		{Name: "class", Type: field.TypeString},
		{Name: "identity", Type: field.TypeString},
		{Name: "user_students", Type: field.TypeInt, Unique: true},
	}
	// StudentsTable holds the schema information for the "students" table.
	StudentsTable = &schema.Table{
		Name:       "students",
		Columns:    StudentsColumns,
		PrimaryKey: []*schema.Column{StudentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "students_users_students",
				Columns:    []*schema.Column{StudentsColumns[7]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// TeachersColumns holds the columns for the "teachers" table.
	TeachersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "college", Type: field.TypeString},
		{Name: "phone", Type: field.TypeString},
		{Name: "identity", Type: field.TypeString},
		{Name: "user_teachers", Type: field.TypeInt, Unique: true},
	}
	// TeachersTable holds the schema information for the "teachers" table.
	TeachersTable = &schema.Table{
		Name:       "teachers",
		Columns:    TeachersColumns,
		PrimaryKey: []*schema.Column{TeachersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "teachers_users_teachers",
				Columns:    []*schema.Column{TeachersColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// ThesesColumns holds the columns for the "theses" table.
	ThesesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "time", Type: field.TypeTime},
		{Name: "url", Type: field.TypeString, Unique: true},
		{Name: "type", Type: field.TypeInt},
		{Name: "status", Type: field.TypeInt},
		{Name: "user_files", Type: field.TypeInt, Nullable: true},
	}
	// ThesesTable holds the schema information for the "theses" table.
	ThesesTable = &schema.Table{
		Name:       "theses",
		Columns:    ThesesColumns,
		PrimaryKey: []*schema.Column{ThesesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "theses_users_files",
				Columns:    []*schema.Column{ThesesColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "account", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "role", Type: field.TypeInt},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AdministratorsTable,
		StudentsTable,
		TeachersTable,
		ThesesTable,
		UsersTable,
	}
)

func init() {
	AdministratorsTable.ForeignKeys[0].RefTable = UsersTable
	StudentsTable.ForeignKeys[0].RefTable = UsersTable
	TeachersTable.ForeignKeys[0].RefTable = UsersTable
	ThesesTable.ForeignKeys[0].RefTable = UsersTable
}
