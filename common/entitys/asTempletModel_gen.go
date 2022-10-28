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
	asTempletFieldNames          = builder.RawFieldNames(&AsTemplet{})
	asTempletRows                = strings.Join(asTempletFieldNames, ",")
	asTempletRowsExpectAutoSet   = strings.Join(stringx.Remove(asTempletFieldNames, "`create_time`", "`update_time`"), ",")
	asTempletRowsWithPlaceHolder = strings.Join(stringx.Remove(asTempletFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheAssetAsTempletIdPrefix = "cache:asset:asTemplet:id:"
)

type (
	asTempletModel interface {
		Insert(ctx context.Context, data *AsTemplet) (sql.Result, error)
		FindOne(ctx context.Context, id string) (*AsTemplet, error)
		Update(ctx context.Context, data *AsTemplet) error
		Delete(ctx context.Context, id string) error
	}

	defaultAsTempletModel struct {
		sqlc.CachedConn
		table string
	}

	AsTemplet struct {
		Id              string         `db:"id"`                 // 模板Id
		TempletName     string         `db:"templet_name"`       // 模板的名称
		IconCls         string         `db:"icon_cls"`           // 模板的图标
		LinkAppId       string         `db:"link_app_id"`        // 模板下对应的应用Id
		LinkFormModelId string         `db:"link_form_model_id"` // 模板下对应的表单模型Id
		LinkProcModelId string         `db:"link_proc_model_id"` // 模板下对应的流程模型Id
		Status          int64          `db:"status"`             // 模板状态
		Committer       sql.NullString `db:"committer"`          // 发布模板的人的Id
		CreateTime      sql.NullTime   `db:"create_time"`        // 模板发布时间
	}
)

func newAsTempletModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultAsTempletModel {
	return &defaultAsTempletModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`as_templet`",
	}
}

func (m *defaultAsTempletModel) Insert(ctx context.Context, data *AsTemplet) (sql.Result, error) {
	assetAsTempletIdKey := fmt.Sprintf("%s%v", cacheAssetAsTempletIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", m.table, asTempletRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Id, data.TempletName, data.IconCls, data.LinkAppId, data.LinkFormModelId, data.LinkProcModelId, data.Status, data.Committer)
	}, assetAsTempletIdKey)
	return ret, err
}

func (m *defaultAsTempletModel) FindOne(ctx context.Context, id string) (*AsTemplet, error) {
	assetAsTempletIdKey := fmt.Sprintf("%s%v", cacheAssetAsTempletIdPrefix, id)
	var resp AsTemplet
	err := m.QueryRowCtx(ctx, &resp, assetAsTempletIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", asTempletRows, m.table)
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

func (m *defaultAsTempletModel) Update(ctx context.Context, data *AsTemplet) error {
	assetAsTempletIdKey := fmt.Sprintf("%s%v", cacheAssetAsTempletIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, asTempletRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.TempletName, data.IconCls, data.LinkAppId, data.LinkFormModelId, data.LinkProcModelId, data.Status, data.Committer, data.Id)
	}, assetAsTempletIdKey)
	return err
}

func (m *defaultAsTempletModel) Delete(ctx context.Context, id string) error {
	assetAsTempletIdKey := fmt.Sprintf("%s%v", cacheAssetAsTempletIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, assetAsTempletIdKey)
	return err
}

func (m *defaultAsTempletModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheAssetAsTempletIdPrefix, primary)
}

func (m *defaultAsTempletModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", asTempletRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultAsTempletModel) tableName() string {
	return m.table
}