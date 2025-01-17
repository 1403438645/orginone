// Code generated by entc, DO NOT EDIT.

package astenanticon

import (
	"orginone/common/tools/date"
)

const (
	// Label holds the string label denoting the astenanticon type in the database.
	Label = "as_tenant_icon"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTenantCode holds the string denoting the tenant_code field in the database.
	FieldTenantCode = "tenant_code"
	// FieldIcon holds the string denoting the icon field in the database.
	FieldIcon = "icon"
	// FieldIsDeleted holds the string denoting the is_deleted field in the database.
	FieldIsDeleted = "is_deleted"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldCreateUser holds the string denoting the create_user field in the database.
	FieldCreateUser = "create_user"
	// FieldUpdateUser holds the string denoting the update_user field in the database.
	FieldUpdateUser = "update_user"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// Table holds the table name of the astenanticon in the database.
	Table = "as_tenant_icon"
)

// Columns holds all SQL columns for astenanticon fields.
var Columns = []string{
	FieldID,
	FieldTenantCode,
	FieldIcon,
	FieldIsDeleted,
	FieldStatus,
	FieldCreateUser,
	FieldUpdateUser,
	FieldCreateTime,
	FieldUpdateTime,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultIsDeleted holds the default value on creation for the "is_deleted" field.
	DefaultIsDeleted int64
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus int64
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() date.DateTime
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() date.DateTime
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() date.DateTime
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() int64
)
