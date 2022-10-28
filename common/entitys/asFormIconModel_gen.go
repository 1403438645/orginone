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
	asFormIconFieldNames          = builder.RawFieldNames(&AsFormIcon{})
	asFormIconRows                = strings.Join(asFormIconFieldNames, ",")
	asFormIconRowsExpectAutoSet   = strings.Join(stringx.Remove(asFormIconFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	asFormIconRowsWithPlaceHolder = strings.Join(stringx.Remove(asFormIconFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheAssetAsFormIconIdPrefix = "cache:asset:asFormIcon:id:"
)

type (
	asFormIconModel interface {
		Insert(ctx context.Context, data *AsFormIcon) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*AsFormIcon, error)
		Update(ctx context.Context, data *AsFormIcon) error
		Delete(ctx context.Context, id int64) error
	}

	defaultAsFormIconModel struct {
		sqlc.CachedConn
		table string
	}

	AsFormIcon struct {
		Id         int64          `db:"id"`
		TenantCode sql.NullString `db:"tenant_code"`
		UserId     sql.NullInt64  `db:"user_id"`
		IconUrl    sql.NullString `db:"icon_url"`
		IsDeleted  int64          `db:"is_deleted"`
	}
)

func newAsFormIconModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultAsFormIconModel {
	return &defaultAsFormIconModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`as_form_icon`",
	}
}

func (m *defaultAsFormIconModel) Insert(ctx context.Context, data *AsFormIcon) (sql.Result, error) {
	assetAsFormIconIdKey := fmt.Sprintf("%s%v", cacheAssetAsFormIconIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, asFormIconRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.TenantCode, data.UserId, data.IconUrl, data.IsDeleted)
	}, assetAsFormIconIdKey)
	return ret, err
}

func (m *defaultAsFormIconModel) FindOne(ctx context.Context, id int64) (*AsFormIcon, error) {
	assetAsFormIconIdKey := fmt.Sprintf("%s%v", cacheAssetAsFormIconIdPrefix, id)
	var resp AsFormIcon
	err := m.QueryRowCtx(ctx, &resp, assetAsFormIconIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", asFormIconRows, m.table)
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

func (m *defaultAsFormIconModel) Update(ctx context.Context, data *AsFormIcon) error {
	assetAsFormIconIdKey := fmt.Sprintf("%s%v", cacheAssetAsFormIconIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, asFormIconRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.TenantCode, data.UserId, data.IconUrl, data.IsDeleted, data.Id)
	}, assetAsFormIconIdKey)
	return err
}

func (m *defaultAsFormIconModel) Delete(ctx context.Context, id int64) error {
	assetAsFormIconIdKey := fmt.Sprintf("%s%v", cacheAssetAsFormIconIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, assetAsFormIconIdKey)
	return err
}

func (m *defaultAsFormIconModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheAssetAsFormIconIdPrefix, primary)
}

func (m *defaultAsFormIconModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", asFormIconRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultAsFormIconModel) tableName() string {
	return m.table
}