// Code generated by entc, DO NOT EDIT.

package schema

import (
	"fmt"
	"orginone/common/schema/astenanticon"
	"orginone/common/tools/date"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// AsTenantIcon is the model entity for the AsTenantIcon schema.
type AsTenantIcon struct {
	config `json:"-"`
	// ID of the ent.
	// 主键
	ID int64 `json:"id,string"`
	// TenantCode holds the value of the "tenant_code" field.
	// 租户编码
	TenantCode string `json:"tenantCode"`
	// Icon holds the value of the "icon" field.
	// 图标
	Icon string `json:"icon"`
	// IsDeleted holds the value of the "is_deleted" field.
	// 是否被删除
	IsDeleted int64 `json:"isDeleted"`
	// Status holds the value of the "status" field.
	// 状态
	Status int64 `json:"status"`
	// CreateUser holds the value of the "create_user" field.
	// 创建用户
	CreateUser int64 `json:"createUser"`
	// UpdateUser holds the value of the "update_user" field.
	// 更新用户
	UpdateUser int64 `json:"updateUser"`
	// CreateTime holds the value of the "create_time" field.
	// 创建时间
	CreateTime date.DateTime `json:"createTime"`
	// UpdateTime holds the value of the "update_time" field.
	// 更新时间
	UpdateTime date.DateTime `json:"updateTime"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AsTenantIcon) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case astenanticon.FieldID, astenanticon.FieldIsDeleted, astenanticon.FieldStatus, astenanticon.FieldCreateUser, astenanticon.FieldUpdateUser:
			values[i] = new(sql.NullInt64)
		case astenanticon.FieldTenantCode, astenanticon.FieldIcon:
			values[i] = new(sql.NullString)
		case astenanticon.FieldCreateTime, astenanticon.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type AsTenantIcon", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AsTenantIcon fields.
func (ati *AsTenantIcon) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case astenanticon.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ati.ID = int64(value.Int64)
		case astenanticon.FieldTenantCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tenant_code", values[i])
			} else if value.Valid {
				ati.TenantCode = value.String
			}
		case astenanticon.FieldIcon:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field icon", values[i])
			} else if value.Valid {
				ati.Icon = value.String
			}
		case astenanticon.FieldIsDeleted:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field is_deleted", values[i])
			} else if value.Valid {
				ati.IsDeleted = value.Int64
			}
		case astenanticon.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				ati.Status = value.Int64
			}
		case astenanticon.FieldCreateUser:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_user", values[i])
			} else if value.Valid {
				ati.CreateUser = value.Int64
			}
		case astenanticon.FieldUpdateUser:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field update_user", values[i])
			} else if value.Valid {
				ati.UpdateUser = value.Int64
			}
		case astenanticon.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				ati.CreateTime = date.DateTime(value.Time)
			}
		case astenanticon.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				ati.UpdateTime = date.DateTime(value.Time)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this AsTenantIcon.
// Note that you need to call AsTenantIcon.Unwrap() before calling this method if this AsTenantIcon
// was returned from a transaction, and the transaction was committed or rolled back.
func (ati *AsTenantIcon) Update() *AsTenantIconUpdateOne {
	return (&AsTenantIconClient{config: ati.config}).UpdateOne(ati)
}

// Unwrap unwraps the AsTenantIcon entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ati *AsTenantIcon) Unwrap() *AsTenantIcon {
	tx, ok := ati.config.driver.(*txDriver)
	if !ok {
		panic("schema: AsTenantIcon is not a transactional entity")
	}
	ati.config.driver = tx.drv
	return ati
}

// String implements the fmt.Stringer.
func (ati *AsTenantIcon) String() string {
	var builder strings.Builder
	builder.WriteString("AsTenantIcon(")
	builder.WriteString(fmt.Sprintf("id=%v", ati.ID))
	builder.WriteString(", tenant_code=")
	builder.WriteString(ati.TenantCode)
	builder.WriteString(", icon=")
	builder.WriteString(ati.Icon)
	builder.WriteString(", is_deleted=")
	builder.WriteString(fmt.Sprintf("%v", ati.IsDeleted))
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", ati.Status))
	builder.WriteString(", create_user=")
	builder.WriteString(fmt.Sprintf("%v", ati.CreateUser))
	builder.WriteString(", update_user=")
	builder.WriteString(fmt.Sprintf("%v", ati.UpdateUser))
	builder.WriteString(", create_time=")
	builder.WriteString(fmt.Sprintf("%v", ati.CreateTime))
	builder.WriteString(", update_time=")
	builder.WriteString(fmt.Sprintf("%v", ati.UpdateTime))
	builder.WriteByte(')')
	return builder.String()
}

// AsTenantIcons is a parsable slice of AsTenantIcon.
type AsTenantIcons []*AsTenantIcon

func (ati AsTenantIcons) config(cfg config) {
	for _i := range ati {
		ati[_i].config = cfg
	}
}