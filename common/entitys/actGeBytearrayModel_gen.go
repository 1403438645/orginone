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
	actGeBytearrayFieldNames          = builder.RawFieldNames(&ActGeBytearray{})
	actGeBytearrayRows                = strings.Join(actGeBytearrayFieldNames, ",")
	actGeBytearrayRowsExpectAutoSet   = strings.Join(stringx.Remove(actGeBytearrayFieldNames, "`create_time`", "`update_time`"), ",")
	actGeBytearrayRowsWithPlaceHolder = strings.Join(stringx.Remove(actGeBytearrayFieldNames, "`ID_`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheAssetActGeBytearrayIDPrefix = "cache:asset:actGeBytearray:iD:"
)

type (
	actGeBytearrayModel interface {
		Insert(ctx context.Context, data *ActGeBytearray) (sql.Result, error)
		FindOne(ctx context.Context, iD string) (*ActGeBytearray, error)
		Update(ctx context.Context, data *ActGeBytearray) error
		Delete(ctx context.Context, iD string) error
	}

	defaultActGeBytearrayModel struct {
		sqlc.CachedConn
		table string
	}

	ActGeBytearray struct {
		ID           string         `db:"ID_"`
		REV          sql.NullInt64  `db:"REV_"`
		NAME         sql.NullString `db:"NAME_"`
		DEPLOYMENTID sql.NullString `db:"DEPLOYMENT_ID_"`
		BYTES        sql.NullString `db:"BYTES_"`
		GENERATED    sql.NullInt64  `db:"GENERATED_"`
	}
)

func newActGeBytearrayModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultActGeBytearrayModel {
	return &defaultActGeBytearrayModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`act_ge_bytearray`",
	}
}

func (m *defaultActGeBytearrayModel) Insert(ctx context.Context, data *ActGeBytearray) (sql.Result, error) {
	assetActGeBytearrayIDKey := fmt.Sprintf("%s%v", cacheAssetActGeBytearrayIDPrefix, data.ID)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, actGeBytearrayRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.ID, data.REV, data.NAME, data.DEPLOYMENTID, data.BYTES, data.GENERATED)
	}, assetActGeBytearrayIDKey)
	return ret, err
}

func (m *defaultActGeBytearrayModel) FindOne(ctx context.Context, iD string) (*ActGeBytearray, error) {
	assetActGeBytearrayIDKey := fmt.Sprintf("%s%v", cacheAssetActGeBytearrayIDPrefix, iD)
	var resp ActGeBytearray
	err := m.QueryRowCtx(ctx, &resp, assetActGeBytearrayIDKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `ID_` = ? limit 1", actGeBytearrayRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, iD)
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

func (m *defaultActGeBytearrayModel) Update(ctx context.Context, data *ActGeBytearray) error {
	assetActGeBytearrayIDKey := fmt.Sprintf("%s%v", cacheAssetActGeBytearrayIDPrefix, data.ID)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `ID_` = ?", m.table, actGeBytearrayRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.REV, data.NAME, data.DEPLOYMENTID, data.BYTES, data.GENERATED, data.ID)
	}, assetActGeBytearrayIDKey)
	return err
}

func (m *defaultActGeBytearrayModel) Delete(ctx context.Context, iD string) error {
	assetActGeBytearrayIDKey := fmt.Sprintf("%s%v", cacheAssetActGeBytearrayIDPrefix, iD)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `ID_` = ?", m.table)
		return conn.ExecCtx(ctx, query, iD)
	}, assetActGeBytearrayIDKey)
	return err
}

func (m *defaultActGeBytearrayModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheAssetActGeBytearrayIDPrefix, primary)
}

func (m *defaultActGeBytearrayModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `ID_` = ? limit 1", actGeBytearrayRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultActGeBytearrayModel) tableName() string {
	return m.table
}