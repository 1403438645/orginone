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
	actIdPrivFieldNames          = builder.RawFieldNames(&ActIdPriv{})
	actIdPrivRows                = strings.Join(actIdPrivFieldNames, ",")
	actIdPrivRowsExpectAutoSet   = strings.Join(stringx.Remove(actIdPrivFieldNames, "`create_time`", "`update_time`"), ",")
	actIdPrivRowsWithPlaceHolder = strings.Join(stringx.Remove(actIdPrivFieldNames, "`ID_`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheAssetActIdPrivIDPrefix   = "cache:asset:actIdPriv:iD:"
	cacheAssetActIdPrivNAMEPrefix = "cache:asset:actIdPriv:nAME:"
)

type (
	actIdPrivModel interface {
		Insert(ctx context.Context, data *ActIdPriv) (sql.Result, error)
		FindOne(ctx context.Context, iD string) (*ActIdPriv, error)
		FindOneByNAME(ctx context.Context, nAME string) (*ActIdPriv, error)
		Update(ctx context.Context, data *ActIdPriv) error
		Delete(ctx context.Context, iD string) error
	}

	defaultActIdPrivModel struct {
		sqlc.CachedConn
		table string
	}

	ActIdPriv struct {
		ID   string `db:"ID_"`
		NAME string `db:"NAME_"`
	}
)

func newActIdPrivModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultActIdPrivModel {
	return &defaultActIdPrivModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`act_id_priv`",
	}
}

func (m *defaultActIdPrivModel) Insert(ctx context.Context, data *ActIdPriv) (sql.Result, error) {
	assetActIdPrivIDKey := fmt.Sprintf("%s%v", cacheAssetActIdPrivIDPrefix, data.ID)
	assetActIdPrivNAMEKey := fmt.Sprintf("%s%v", cacheAssetActIdPrivNAMEPrefix, data.NAME)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, actIdPrivRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.ID, data.NAME)
	}, assetActIdPrivIDKey, assetActIdPrivNAMEKey)
	return ret, err
}

func (m *defaultActIdPrivModel) FindOne(ctx context.Context, iD string) (*ActIdPriv, error) {
	assetActIdPrivIDKey := fmt.Sprintf("%s%v", cacheAssetActIdPrivIDPrefix, iD)
	var resp ActIdPriv
	err := m.QueryRowCtx(ctx, &resp, assetActIdPrivIDKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `ID_` = ? limit 1", actIdPrivRows, m.table)
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

func (m *defaultActIdPrivModel) FindOneByNAME(ctx context.Context, nAME string) (*ActIdPriv, error) {
	assetActIdPrivNAMEKey := fmt.Sprintf("%s%v", cacheAssetActIdPrivNAMEPrefix, nAME)
	var resp ActIdPriv
	err := m.QueryRowIndexCtx(ctx, &resp, assetActIdPrivNAMEKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `NAME_` = ? limit 1", actIdPrivRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, nAME); err != nil {
			return nil, err
		}
		return resp.ID, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultActIdPrivModel) Update(ctx context.Context, data *ActIdPriv) error {
	assetActIdPrivIDKey := fmt.Sprintf("%s%v", cacheAssetActIdPrivIDPrefix, data.ID)
	assetActIdPrivNAMEKey := fmt.Sprintf("%s%v", cacheAssetActIdPrivNAMEPrefix, data.NAME)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `ID_` = ?", m.table, actIdPrivRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.NAME, data.ID)
	}, assetActIdPrivIDKey, assetActIdPrivNAMEKey)
	return err
}

func (m *defaultActIdPrivModel) Delete(ctx context.Context, iD string) error {
	data, err := m.FindOne(ctx, iD)
	if err != nil {
		return err
	}

	assetActIdPrivIDKey := fmt.Sprintf("%s%v", cacheAssetActIdPrivIDPrefix, iD)
	assetActIdPrivNAMEKey := fmt.Sprintf("%s%v", cacheAssetActIdPrivNAMEPrefix, data.NAME)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `ID_` = ?", m.table)
		return conn.ExecCtx(ctx, query, iD)
	}, assetActIdPrivIDKey, assetActIdPrivNAMEKey)
	return err
}

func (m *defaultActIdPrivModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheAssetActIdPrivIDPrefix, primary)
}

func (m *defaultActIdPrivModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `ID_` = ? limit 1", actIdPrivRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultActIdPrivModel) tableName() string {
	return m.table
}
