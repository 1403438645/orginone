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
	asMarketAppApplyFieldNames          = builder.RawFieldNames(&AsMarketAppApply{})
	asMarketAppApplyRows                = strings.Join(asMarketAppApplyFieldNames, ",")
	asMarketAppApplyRowsExpectAutoSet   = strings.Join(stringx.Remove(asMarketAppApplyFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	asMarketAppApplyRowsWithPlaceHolder = strings.Join(stringx.Remove(asMarketAppApplyFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheAssetAsMarketAppApplyIdPrefix = "cache:asset:asMarketAppApply:id:"
)

type (
	asMarketAppApplyModel interface {
		Insert(ctx context.Context, data *AsMarketAppApply) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*AsMarketAppApply, error)
		Update(ctx context.Context, data *AsMarketAppApply) error
		Delete(ctx context.Context, id int64) error
	}

	defaultAsMarketAppApplyModel struct {
		sqlc.CachedConn
		table string
	}

	AsMarketAppApply struct {
		Id             int64          `db:"id"`              // 主键
		CreateUser     sql.NullInt64  `db:"create_user"`     // 创建人
		CreateTime     sql.NullTime   `db:"create_time"`     // 创建时间
		UpdateUser     sql.NullInt64  `db:"update_user"`     // 修改人
		UpdateTime     sql.NullTime   `db:"update_time"`     // 修改时间
		IsDeleted      sql.NullInt64  `db:"is_deleted"`      // 是否已删除
		AppName        sql.NullString `db:"app_name"`        // 应用名称，申请时填写
		AppDescription sql.NullString `db:"app_description"` // 应用说明，申请时填写
		DeveloperId    sql.NullInt64  `db:"developer_id"`    // 开发者id
		ApplyDate      sql.NullTime   `db:"apply_date"`      // 申请时间
		LicenseDate    sql.NullTime   `db:"license_date"`    // 许可开始时间
		Status         sql.NullInt64  `db:"status"`          // 状态（申请、许可生效、许可注销、申请失败）
		Feedback       sql.NullString `db:"feedback"`        // 审批反馈
		AppId          sql.NullInt64  `db:"app_id"`          // 应用id
	}
)

func newAsMarketAppApplyModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultAsMarketAppApplyModel {
	return &defaultAsMarketAppApplyModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`as_market_app_apply`",
	}
}

func (m *defaultAsMarketAppApplyModel) Insert(ctx context.Context, data *AsMarketAppApply) (sql.Result, error) {
	assetAsMarketAppApplyIdKey := fmt.Sprintf("%s%v", cacheAssetAsMarketAppApplyIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, asMarketAppApplyRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.CreateUser, data.UpdateUser, data.IsDeleted, data.AppName, data.AppDescription, data.DeveloperId, data.ApplyDate, data.LicenseDate, data.Status, data.Feedback, data.AppId)
	}, assetAsMarketAppApplyIdKey)
	return ret, err
}

func (m *defaultAsMarketAppApplyModel) FindOne(ctx context.Context, id int64) (*AsMarketAppApply, error) {
	assetAsMarketAppApplyIdKey := fmt.Sprintf("%s%v", cacheAssetAsMarketAppApplyIdPrefix, id)
	var resp AsMarketAppApply
	err := m.QueryRowCtx(ctx, &resp, assetAsMarketAppApplyIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", asMarketAppApplyRows, m.table)
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

func (m *defaultAsMarketAppApplyModel) Update(ctx context.Context, data *AsMarketAppApply) error {
	assetAsMarketAppApplyIdKey := fmt.Sprintf("%s%v", cacheAssetAsMarketAppApplyIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, asMarketAppApplyRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.CreateUser, data.UpdateUser, data.IsDeleted, data.AppName, data.AppDescription, data.DeveloperId, data.ApplyDate, data.LicenseDate, data.Status, data.Feedback, data.AppId, data.Id)
	}, assetAsMarketAppApplyIdKey)
	return err
}

func (m *defaultAsMarketAppApplyModel) Delete(ctx context.Context, id int64) error {
	assetAsMarketAppApplyIdKey := fmt.Sprintf("%s%v", cacheAssetAsMarketAppApplyIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, assetAsMarketAppApplyIdKey)
	return err
}

func (m *defaultAsMarketAppApplyModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheAssetAsMarketAppApplyIdPrefix, primary)
}

func (m *defaultAsMarketAppApplyModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", asMarketAppApplyRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultAsMarketAppApplyModel) tableName() string {
	return m.table
}
