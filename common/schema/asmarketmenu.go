// Code generated by entc, DO NOT EDIT.

package schema

import (
	"fmt"
	"orginone/common/schema/asmarketapp"
	"orginone/common/schema/asmarketmenu"
	"orginone/common/tools/date"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// AsMarketMenu is the model entity for the AsMarketMenu schema.
type AsMarketMenu struct {
	config `json:"-"`
	// ID of the ent.
	// 主键
	ID int64 `json:"id,string"`
	// AppID holds the value of the "app_id" field.
	// 应用主键
	AppID int64 `json:"appId,string"`
	// ParentID holds the value of the "parent_id" field.
	// 上级菜单
	ParentID int64 `json:"parentId,string"`
	// MenuName holds the value of the "menu_name" field.
	// 菜单名称
	MenuName string `json:"menuName"`
	// MenuURL holds the value of the "menu_url" field.
	// 菜单url
	MenuURL string `json:"menuUrl"`
	// MenuColumn holds the value of the "menu_column" field.
	// 菜单字段
	MenuColumn string `json:"menuColumn"`
	// Icon holds the value of the "icon" field.
	// 菜单图标
	Icon string `json:"icon"`
	// Sort holds the value of the "sort" field.
	// 排序字段
	Sort int64 `json:"sort"`
	// HTTPSMenuURL holds the value of the "https_menu_url" field.
	// 菜单url
	HTTPSMenuURL string `json:"httpsMenuUrl"`
	// ReformStatus holds the value of the "reform_status" field.
	// 整改状态;0-已认证,1-整改中
	ReformStatus int64 `json:"reformStatus"`
	// OutIPMenuURL holds the value of the "out_ip_menu_url" field.
	// 外网IPurl
	OutIPMenuURL string `json:"outIpMenuUrl"`
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
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AsMarketMenuQuery when eager-loading is set.
	Edges AsMarketMenuEdges `json:"edges"`
}

// AsMarketMenuEdges holds the relations/edges for other nodes in the graph.
type AsMarketMenuEdges struct {
	// Parent holds the value of the parent edge.
	Parent *AsMarketMenu `json:"parent"`
	// Childrens holds the value of the childrens edge.
	Childrens []*AsMarketMenu `json:"childrens"`
	// Appx holds the value of the appx edge.
	Appx *AsMarketApp `json:"appx"`
	// RoleMenus holds the value of the roleMenus edge.
	RoleMenus []*AsMarketRoleMenu `json:"rolemenus"`
	// Roles holds the value of the roles edge.
	Roles []*AsMarketAppRole `json:"roles"`
	// UserSorts holds the value of the UserSorts edge.
	UserSorts []*AsMarketMenuUserSort `json:"usersorts"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [6]bool
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AsMarketMenuEdges) ParentOrErr() (*AsMarketMenu, error) {
	if e.loadedTypes[0] {
		if e.Parent == nil {
			// The edge parent was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: asmarketmenu.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// ChildrensOrErr returns the Childrens value or an error if the edge
// was not loaded in eager-loading.
func (e AsMarketMenuEdges) ChildrensOrErr() ([]*AsMarketMenu, error) {
	if e.loadedTypes[1] {
		return e.Childrens, nil
	}
	return nil, &NotLoadedError{edge: "childrens"}
}

// AppxOrErr returns the Appx value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AsMarketMenuEdges) AppxOrErr() (*AsMarketApp, error) {
	if e.loadedTypes[2] {
		if e.Appx == nil {
			// The edge appx was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: asmarketapp.Label}
		}
		return e.Appx, nil
	}
	return nil, &NotLoadedError{edge: "appx"}
}

// RoleMenusOrErr returns the RoleMenus value or an error if the edge
// was not loaded in eager-loading.
func (e AsMarketMenuEdges) RoleMenusOrErr() ([]*AsMarketRoleMenu, error) {
	if e.loadedTypes[3] {
		return e.RoleMenus, nil
	}
	return nil, &NotLoadedError{edge: "roleMenus"}
}

// RolesOrErr returns the Roles value or an error if the edge
// was not loaded in eager-loading.
func (e AsMarketMenuEdges) RolesOrErr() ([]*AsMarketAppRole, error) {
	if e.loadedTypes[4] {
		return e.Roles, nil
	}
	return nil, &NotLoadedError{edge: "roles"}
}

// UserSortsOrErr returns the UserSorts value or an error if the edge
// was not loaded in eager-loading.
func (e AsMarketMenuEdges) UserSortsOrErr() ([]*AsMarketMenuUserSort, error) {
	if e.loadedTypes[5] {
		return e.UserSorts, nil
	}
	return nil, &NotLoadedError{edge: "UserSorts"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AsMarketMenu) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case asmarketmenu.FieldID, asmarketmenu.FieldAppID, asmarketmenu.FieldParentID, asmarketmenu.FieldSort, asmarketmenu.FieldReformStatus, asmarketmenu.FieldIsDeleted, asmarketmenu.FieldStatus, asmarketmenu.FieldCreateUser, asmarketmenu.FieldUpdateUser:
			values[i] = new(sql.NullInt64)
		case asmarketmenu.FieldMenuName, asmarketmenu.FieldMenuURL, asmarketmenu.FieldMenuColumn, asmarketmenu.FieldIcon, asmarketmenu.FieldHTTPSMenuURL, asmarketmenu.FieldOutIPMenuURL:
			values[i] = new(sql.NullString)
		case asmarketmenu.FieldCreateTime, asmarketmenu.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type AsMarketMenu", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AsMarketMenu fields.
func (amm *AsMarketMenu) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case asmarketmenu.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			amm.ID = int64(value.Int64)
		case asmarketmenu.FieldAppID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value.Valid {
				amm.AppID = value.Int64
			}
		case asmarketmenu.FieldParentID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field parent_id", values[i])
			} else if value.Valid {
				amm.ParentID = value.Int64
			}
		case asmarketmenu.FieldMenuName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field menu_name", values[i])
			} else if value.Valid {
				amm.MenuName = value.String
			}
		case asmarketmenu.FieldMenuURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field menu_url", values[i])
			} else if value.Valid {
				amm.MenuURL = value.String
			}
		case asmarketmenu.FieldMenuColumn:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field menu_column", values[i])
			} else if value.Valid {
				amm.MenuColumn = value.String
			}
		case asmarketmenu.FieldIcon:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field icon", values[i])
			} else if value.Valid {
				amm.Icon = value.String
			}
		case asmarketmenu.FieldSort:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sort", values[i])
			} else if value.Valid {
				amm.Sort = value.Int64
			}
		case asmarketmenu.FieldHTTPSMenuURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field https_menu_url", values[i])
			} else if value.Valid {
				amm.HTTPSMenuURL = value.String
			}
		case asmarketmenu.FieldReformStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field reform_status", values[i])
			} else if value.Valid {
				amm.ReformStatus = value.Int64
			}
		case asmarketmenu.FieldOutIPMenuURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field out_ip_menu_url", values[i])
			} else if value.Valid {
				amm.OutIPMenuURL = value.String
			}
		case asmarketmenu.FieldIsDeleted:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field is_deleted", values[i])
			} else if value.Valid {
				amm.IsDeleted = value.Int64
			}
		case asmarketmenu.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				amm.Status = value.Int64
			}
		case asmarketmenu.FieldCreateUser:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_user", values[i])
			} else if value.Valid {
				amm.CreateUser = value.Int64
			}
		case asmarketmenu.FieldUpdateUser:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field update_user", values[i])
			} else if value.Valid {
				amm.UpdateUser = value.Int64
			}
		case asmarketmenu.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				amm.CreateTime = date.DateTime(value.Time)
			}
		case asmarketmenu.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				amm.UpdateTime = date.DateTime(value.Time)
			}
		}
	}
	return nil
}

// QueryParent queries the "parent" edge of the AsMarketMenu entity.
func (amm *AsMarketMenu) QueryParent() *AsMarketMenuQuery {
	return (&AsMarketMenuClient{config: amm.config}).QueryParent(amm)
}

// QueryChildrens queries the "childrens" edge of the AsMarketMenu entity.
func (amm *AsMarketMenu) QueryChildrens() *AsMarketMenuQuery {
	return (&AsMarketMenuClient{config: amm.config}).QueryChildrens(amm)
}

// QueryAppx queries the "appx" edge of the AsMarketMenu entity.
func (amm *AsMarketMenu) QueryAppx() *AsMarketAppQuery {
	return (&AsMarketMenuClient{config: amm.config}).QueryAppx(amm)
}

// QueryRoleMenus queries the "roleMenus" edge of the AsMarketMenu entity.
func (amm *AsMarketMenu) QueryRoleMenus() *AsMarketRoleMenuQuery {
	return (&AsMarketMenuClient{config: amm.config}).QueryRoleMenus(amm)
}

// QueryRoles queries the "roles" edge of the AsMarketMenu entity.
func (amm *AsMarketMenu) QueryRoles() *AsMarketAppRoleQuery {
	return (&AsMarketMenuClient{config: amm.config}).QueryRoles(amm)
}

// QueryUserSorts queries the "UserSorts" edge of the AsMarketMenu entity.
func (amm *AsMarketMenu) QueryUserSorts() *AsMarketMenuUserSortQuery {
	return (&AsMarketMenuClient{config: amm.config}).QueryUserSorts(amm)
}

// Update returns a builder for updating this AsMarketMenu.
// Note that you need to call AsMarketMenu.Unwrap() before calling this method if this AsMarketMenu
// was returned from a transaction, and the transaction was committed or rolled back.
func (amm *AsMarketMenu) Update() *AsMarketMenuUpdateOne {
	return (&AsMarketMenuClient{config: amm.config}).UpdateOne(amm)
}

// Unwrap unwraps the AsMarketMenu entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (amm *AsMarketMenu) Unwrap() *AsMarketMenu {
	tx, ok := amm.config.driver.(*txDriver)
	if !ok {
		panic("schema: AsMarketMenu is not a transactional entity")
	}
	amm.config.driver = tx.drv
	return amm
}

// String implements the fmt.Stringer.
func (amm *AsMarketMenu) String() string {
	var builder strings.Builder
	builder.WriteString("AsMarketMenu(")
	builder.WriteString(fmt.Sprintf("id=%v", amm.ID))
	builder.WriteString(", app_id=")
	builder.WriteString(fmt.Sprintf("%v", amm.AppID))
	builder.WriteString(", parent_id=")
	builder.WriteString(fmt.Sprintf("%v", amm.ParentID))
	builder.WriteString(", menu_name=")
	builder.WriteString(amm.MenuName)
	builder.WriteString(", menu_url=")
	builder.WriteString(amm.MenuURL)
	builder.WriteString(", menu_column=")
	builder.WriteString(amm.MenuColumn)
	builder.WriteString(", icon=")
	builder.WriteString(amm.Icon)
	builder.WriteString(", sort=")
	builder.WriteString(fmt.Sprintf("%v", amm.Sort))
	builder.WriteString(", https_menu_url=")
	builder.WriteString(amm.HTTPSMenuURL)
	builder.WriteString(", reform_status=")
	builder.WriteString(fmt.Sprintf("%v", amm.ReformStatus))
	builder.WriteString(", out_ip_menu_url=")
	builder.WriteString(amm.OutIPMenuURL)
	builder.WriteString(", is_deleted=")
	builder.WriteString(fmt.Sprintf("%v", amm.IsDeleted))
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", amm.Status))
	builder.WriteString(", create_user=")
	builder.WriteString(fmt.Sprintf("%v", amm.CreateUser))
	builder.WriteString(", update_user=")
	builder.WriteString(fmt.Sprintf("%v", amm.UpdateUser))
	builder.WriteString(", create_time=")
	builder.WriteString(fmt.Sprintf("%v", amm.CreateTime))
	builder.WriteString(", update_time=")
	builder.WriteString(fmt.Sprintf("%v", amm.UpdateTime))
	builder.WriteByte(')')
	return builder.String()
}

// AsMarketMenus is a parsable slice of AsMarketMenu.
type AsMarketMenus []*AsMarketMenu

func (amm AsMarketMenus) config(cfg config) {
	for _i := range amm {
		amm[_i].config = cfg
	}
}