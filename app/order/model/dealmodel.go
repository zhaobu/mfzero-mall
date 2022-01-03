package model

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/tal-tech/go-zero/core/stores/builder"
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
)

var (
	dealFieldNames          = builder.RawFieldNames(&Deal{})
	dealRows                = strings.Join(dealFieldNames, ",")
	dealRowsExpectAutoSet   = strings.Join(stringx.Remove(dealFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	dealRowsWithPlaceHolder = strings.Join(stringx.Remove(dealFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheDealIdPrefix = "cache:deal:id:"
)

type (
	DealModel interface {
		Insert(data *Deal) (sql.Result, error)
		FindOne(id int64) (*Deal, error)
		Update(data *Deal) error
		Delete(id int64) error
		FindAllByUid(uid int64) ([]*Deal, error)
	}

	defaultDealModel struct {
		sqlc.CachedConn
		table string
	}

	Deal struct {
		Id         int64     `db:"id"`
		Uid        int64     `db:"uid"`    // 用户ID
		Pid        int64     `db:"pid"`    // 产品ID
		Amount     int64     `db:"amount"` // 订单金额
		Status     int64     `db:"status"` // 订单状态
		CreateTime time.Time `db:"create_time"`
		UpdateTime time.Time `db:"update_time"`
	}
)

func NewDealModel(conn sqlx.SqlConn, c cache.CacheConf) DealModel {
	return &defaultDealModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`deal`",
	}
}

func (m *defaultDealModel) Insert(data *Deal) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, dealRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.Uid, data.Pid, data.Amount, data.Status)

	return ret, err
}

func (m *defaultDealModel) FindOne(id int64) (*Deal, error) {
	dealIdKey := fmt.Sprintf("%s%v", cacheDealIdPrefix, id)
	var resp Deal
	err := m.QueryRow(&resp, dealIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", dealRows, m.table)
		return conn.QueryRow(v, query, id)
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

func (m *defaultDealModel) Update(data *Deal) error {
	dealIdKey := fmt.Sprintf("%s%v", cacheDealIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, dealRowsWithPlaceHolder)
		return conn.Exec(query, data.Uid, data.Pid, data.Amount, data.Status, data.Id)
	}, dealIdKey)
	return err
}

func (m *defaultDealModel) Delete(id int64) error {

	dealIdKey := fmt.Sprintf("%s%v", cacheDealIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, dealIdKey)
	return err
}

func (m *defaultDealModel) FindAllByUid(uid int64) ([]*Deal, error) {
	var resp []*Deal

	query := fmt.Sprintf("select %s from %s where `uid` = ?", dealRows, m.table)
	err := m.QueryRowsNoCache(&resp, query, uid)

	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultDealModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheDealIdPrefix, primary)
}

func (m *defaultDealModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", dealRows, m.table)
	return conn.QueryRow(v, query, primary)
}
