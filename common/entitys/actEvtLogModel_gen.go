// Code generated by goctl. DO NOT EDIT!

package entitys

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	actEvtLogFieldNames          = builder.RawFieldNames(&ActEvtLog{})
	actEvtLogRows                = strings.Join(actEvtLogFieldNames, ",")
	actEvtLogRowsExpectAutoSet   = strings.Join(stringx.Remove(actEvtLogFieldNames, "`LOG_NR_`", "`create_time`", "`update_time`"), ",")
	actEvtLogRowsWithPlaceHolder = strings.Join(stringx.Remove(actEvtLogFieldNames, "`LOG_NR_`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheAssetActEvtLogLOGNRPrefix = "cache:asset:actEvtLog:lOGNR:"
)

type (
	actEvtLogModel interface {
		Insert(ctx context.Context, data *ActEvtLog) (sql.Result, error)
		FindOne(ctx context.Context, lOGNR int64) (*ActEvtLog, error)
		Update(ctx context.Context, data *ActEvtLog) error
		Delete(ctx context.Context, lOGNR int64) error
	}

	defaultActEvtLogModel struct {
		sqlc.CachedConn
		table string
	}

	ActEvtLog struct {
		LOGNR       int64          `db:"LOG_NR_"`
		TYPE        sql.NullString `db:"TYPE_"`
		PROCDEFID   sql.NullString `db:"PROC_DEF_ID_"`
		PROCINSTID  sql.NullString `db:"PROC_INST_ID_"`
		EXECUTIONID sql.NullString `db:"EXECUTION_ID_"`
		TASKID      sql.NullString `db:"TASK_ID_"`
		TIMESTAMP   time.Time      `db:"TIME_STAMP_"`
		USERID      sql.NullString `db:"USER_ID_"`
		DATA        sql.NullString `db:"DATA_"`
		LOCKOWNER   sql.NullString `db:"LOCK_OWNER_"`
		LOCKTIME    sql.NullTime   `db:"LOCK_TIME_"`
		ISPROCESSED int64          `db:"IS_PROCESSED_"`
	}
)

func newActEvtLogModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultActEvtLogModel {
	return &defaultActEvtLogModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`act_evt_log`",
	}
}

func (m *defaultActEvtLogModel) Insert(ctx context.Context, data *ActEvtLog) (sql.Result, error) {
	assetActEvtLogLOGNRKey := fmt.Sprintf("%s%v", cacheAssetActEvtLogLOGNRPrefix, data.LOGNR)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, actEvtLogRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.TYPE, data.PROCDEFID, data.PROCINSTID, data.EXECUTIONID, data.TASKID, data.TIMESTAMP, data.USERID, data.DATA, data.LOCKOWNER, data.LOCKTIME, data.ISPROCESSED)
	}, assetActEvtLogLOGNRKey)
	return ret, err
}

func (m *defaultActEvtLogModel) FindOne(ctx context.Context, lOGNR int64) (*ActEvtLog, error) {
	assetActEvtLogLOGNRKey := fmt.Sprintf("%s%v", cacheAssetActEvtLogLOGNRPrefix, lOGNR)
	var resp ActEvtLog
	err := m.QueryRowCtx(ctx, &resp, assetActEvtLogLOGNRKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `LOG_NR_` = ? limit 1", actEvtLogRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, lOGNR)
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

func (m *defaultActEvtLogModel) Update(ctx context.Context, data *ActEvtLog) error {
	assetActEvtLogLOGNRKey := fmt.Sprintf("%s%v", cacheAssetActEvtLogLOGNRPrefix, data.LOGNR)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `LOG_NR_` = ?", m.table, actEvtLogRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.TYPE, data.PROCDEFID, data.PROCINSTID, data.EXECUTIONID, data.TASKID, data.TIMESTAMP, data.USERID, data.DATA, data.LOCKOWNER, data.LOCKTIME, data.ISPROCESSED, data.LOGNR)
	}, assetActEvtLogLOGNRKey)
	return err
}

func (m *defaultActEvtLogModel) Delete(ctx context.Context, lOGNR int64) error {
	assetActEvtLogLOGNRKey := fmt.Sprintf("%s%v", cacheAssetActEvtLogLOGNRPrefix, lOGNR)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `LOG_NR_` = ?", m.table)
		return conn.ExecCtx(ctx, query, lOGNR)
	}, assetActEvtLogLOGNRKey)
	return err
}

func (m *defaultActEvtLogModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheAssetActEvtLogLOGNRPrefix, primary)
}

func (m *defaultActEvtLogModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `LOG_NR_` = ? limit 1", actEvtLogRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultActEvtLogModel) tableName() string {
	return m.table
}