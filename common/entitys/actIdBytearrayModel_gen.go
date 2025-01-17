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
	actIdBytearrayFieldNames          = builder.RawFieldNames(&ActIdBytearray{})
	actIdBytearrayRows                = strings.Join(actIdBytearrayFieldNames, ",")
	actIdBytearrayRowsExpectAutoSet   = strings.Join(stringx.Remove(actIdBytearrayFieldNames, "`create_time`", "`update_time`"), ",")
	actIdBytearrayRowsWithPlaceHolder = strings.Join(stringx.Remove(actIdBytearrayFieldNames, "`ID_`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheAssetActIdBytearrayIDPrefix = "cache:asset:actIdBytearray:iD:"
)

type (
	actIdBytearrayModel interface {
		Insert(ctx context.Context, data *ActIdBytearray) (sql.Result, error)
		FindOne(ctx context.Context, iD string) (*ActIdBytearray, error)
		Update(ctx context.Context, data *ActIdBytearray) error
		Delete(ctx context.Context, iD string) error
	}

	defaultActIdBytearrayModel struct {
		sqlc.CachedConn
		table string
	}

	ActIdBytearray struct {
		ID    string         `db:"ID_"`
		REV   sql.NullInt64  `db:"REV_"`
		NAME  sql.NullString `db:"NAME_"`
		BYTES sql.NullString `db:"BYTES_"`
	}
)

func newActIdBytearrayModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultActIdBytearrayModel {
	return &defaultActIdBytearrayModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`act_id_bytearray`",
	}
}

func (m *defaultActIdBytearrayModel) Insert(ctx context.Context, data *ActIdBytearray) (sql.Result, error) {
	assetActIdBytearrayIDKey := fmt.Sprintf("%s%v", cacheAssetActIdBytearrayIDPrefix, data.ID)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, actIdBytearrayRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.ID, data.REV, data.NAME, data.BYTES)
	}, assetActIdBytearrayIDKey)
	return ret, err
}

func (m *defaultActIdBytearrayModel) FindOne(ctx context.Context, iD string) (*ActIdBytearray, error) {
	assetActIdBytearrayIDKey := fmt.Sprintf("%s%v", cacheAssetActIdBytearrayIDPrefix, iD)
	var resp ActIdBytearray
	err := m.QueryRowCtx(ctx, &resp, assetActIdBytearrayIDKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `ID_` = ? limit 1", actIdBytearrayRows, m.table)
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

func (m *defaultActIdBytearrayModel) Update(ctx context.Context, data *ActIdBytearray) error {
	assetActIdBytearrayIDKey := fmt.Sprintf("%s%v", cacheAssetActIdBytearrayIDPrefix, data.ID)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `ID_` = ?", m.table, actIdBytearrayRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.REV, data.NAME, data.BYTES, data.ID)
	}, assetActIdBytearrayIDKey)
	return err
}

func (m *defaultActIdBytearrayModel) Delete(ctx context.Context, iD string) error {
	assetActIdBytearrayIDKey := fmt.Sprintf("%s%v", cacheAssetActIdBytearrayIDPrefix, iD)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `ID_` = ?", m.table)
		return conn.ExecCtx(ctx, query, iD)
	}, assetActIdBytearrayIDKey)
	return err
}

func (m *defaultActIdBytearrayModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheAssetActIdBytearrayIDPrefix, primary)
}

func (m *defaultActIdBytearrayModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `ID_` = ? limit 1", actIdBytearrayRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultActIdBytearrayModel) tableName() string {
	return m.table
}
