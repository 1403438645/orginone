// Code generated by goctl. DO NOT EDIT!

package entitys

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	asMarketMenuFieldNames          = builder.RawFieldNames(&AsMarketMenu{})
	asMarketMenuRows                = strings.Join(asMarketMenuFieldNames, ",")
	asMarketMenuRowsExpectAutoSet   = strings.Join(stringx.Remove(asMarketMenuFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	asMarketMenuRowsWithPlaceHolder = strings.Join(stringx.Remove(asMarketMenuFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheAssetAsMarketMenuIdPrefix = "cache:asset:asMarketMenu:id:"
)

type (
	asMarketMenuModel interface {
		Insert(ctx context.Context, data *AsMarketMenu) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*AsMarketMenu, error)
		Update(ctx context.Context, data *AsMarketMenu) error
		Delete(ctx context.Context, id int64) error
	}

	defaultAsMarketMenuModel struct {
		sqlc.CachedConn
		table string
	}

	AsMarketMenu struct {
		Id           int64          `db:"id"`          // 主键
		AppId        sql.NullInt64  `db:"app_id"`      // 应用主键
		CreateUser   sql.NullInt64  `db:"create_user"` // 创建人
		CreateTime   sql.NullTime   `db:"create_time"` // 创建时间
		UpdateUser   sql.NullInt64  `db:"update_user"` // 修改人
		UpdateTime   sql.NullTime   `db:"update_time"` // 修改时间
		Status       sql.NullInt64  `db:"status"`      // 状态
		IsDeleted    sql.NullInt64  `db:"is_deleted"`  // 是否已删除
		ParentId     sql.NullInt64  `db:"parent_id"`   // 上级菜单
		MenuName     sql.NullString `db:"menu_name"`   // 菜单名称
		MenuUrl      sql.NullString `db:"menu_url"`    // 菜单url
		MenuColumn   sql.NullString `db:"menu_column"` // 菜单字段
		Icon         sql.NullString `db:"icon"`        // 菜单图标
		Sort         sql.NullInt64  `db:"sort"`        // 排序字段
		HttpsMenuUrl string         `db:"https_menu_url"`
		ReformStatus int64          `db:"reform_status"`   // 整改状态：0-已认证，1-整改中
		OutIpMenuUrl sql.NullString `db:"out_ip_menu_url"` // 外网IPurl
	}
)

func newAsMarketMenuModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultAsMarketMenuModel {
	return &defaultAsMarketMenuModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`as_market_menu`",
	}
}

func (m *defaultAsMarketMenuModel) Insert(ctx context.Context, data *AsMarketMenu) (sql.Result, error) {
	assetAsMarketMenuIdKey := fmt.Sprintf("%s%v", cacheAssetAsMarketMenuIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, asMarketMenuRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.AppId, data.CreateUser, data.UpdateUser, data.Status, data.IsDeleted, data.ParentId, data.MenuName, data.MenuUrl, data.MenuColumn, data.Icon, data.Sort, data.HttpsMenuUrl, data.ReformStatus, data.OutIpMenuUrl)
	}, assetAsMarketMenuIdKey)
	return ret, err
}

func (m *defaultAsMarketMenuModel) FindOne(ctx context.Context, id int64) (*AsMarketMenu, error) {
	assetAsMarketMenuIdKey := fmt.Sprintf("%s%v", cacheAssetAsMarketMenuIdPrefix, id)
	var resp AsMarketMenu
	err := m.QueryRowCtx(ctx, &resp, assetAsMarketMenuIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", asMarketMenuRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultAsMarketMenuModel) Update(ctx context.Context, data *AsMarketMenu) error {
	assetAsMarketMenuIdKey := fmt.Sprintf("%s%v", cacheAssetAsMarketMenuIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, asMarketMenuRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.AppId, data.CreateUser, data.UpdateUser, data.Status, data.IsDeleted, data.ParentId, data.MenuName, data.MenuUrl, data.MenuColumn, data.Icon, data.Sort, data.HttpsMenuUrl, data.ReformStatus, data.OutIpMenuUrl, data.Id)
	}, assetAsMarketMenuIdKey)
	return err
}

func (m *defaultAsMarketMenuModel) Delete(ctx context.Context, id int64) error {
	assetAsMarketMenuIdKey := fmt.Sprintf("%s%v", cacheAssetAsMarketMenuIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, assetAsMarketMenuIdKey)
	return err
}

func (m *defaultAsMarketMenuModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheAssetAsMarketMenuIdPrefix, primary)
}

func (m *defaultAsMarketMenuModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", asMarketMenuRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultAsMarketMenuModel) tableName() string {
	return m.table
}