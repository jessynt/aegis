package aggregation

import (
	"context"
	"fmt"
	"math"
	"reflect"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/spf13/cast"
	"gonum.org/v1/gonum/stat"
)

type Opt func(a *Aggregator)

type Aggregator struct {
	clickHouseConn sqlx.ExtContext
}

func NewAggregator(opts ...Opt) *Aggregator {
	a := &Aggregator{}
	for _, opt := range opts {
		opt(a)
	}
	return a
}

func WithClickHouseConn(db sqlx.ExtContext) Opt {
	return func(a *Aggregator) {
		a.clickHouseConn = db
	}
}

func (a *Aggregator) Count(
	ctx context.Context, modelId int64, searchField string, searchFieldValue interface{}, startsAt, endsAt time.Time,
) (rv float64, err error) {
	sql := fmt.Sprintf("SELECT count(*) FROM event_dist WHERE ModelId = ? AND %s AND Time >= ? AND Time <= ?", generateQuery(searchField, searchFieldValue))
	err = a.clickHouseConn.
		QueryRowxContext(ctx, sql, modelId, startsAt, endsAt).
		Scan(&rv)
	return
}

func (a *Aggregator) DistinctCount(
	ctx context.Context, modelId int64, searchField string, searchFieldValue interface{}, startsAt, endsAt time.Time, aggregateField string,
) (rv float64, err error) {
	sql := fmt.Sprintf("SELECT uniqExact(%s) FROM event_dist WHERE ModelId = ? AND %s AND Time >= ? AND Time <= ?", aggregateField, generateQuery(searchField, searchFieldValue))
	err = a.clickHouseConn.
		QueryRowxContext(ctx, sql, modelId, startsAt, endsAt).
		Scan(&rv)
	return
}

func (a *Aggregator) Average(
	ctx context.Context, modelId int64, searchField string, searchFieldValue interface{}, startsAt, endsAt time.Time, aggregateField string,
) (rv float64, err error) {
	sql := fmt.Sprintf("SELECT avg(%s) FROM event_dist WHERE ModelId = ? AND %s AND Time >= ? AND Time <= ?", aggregateField, generateQuery(searchField, searchFieldValue))
	err = a.clickHouseConn.
		QueryRowxContext(ctx, sql, modelId, startsAt, endsAt).
		Scan(&rv)
	return
}

func (a *Aggregator) Sum(ctx context.Context, modelId int64, searchField string, searchFieldValue interface{}, startsAt, endsAt time.Time, aggregateField string) (rv float64, err error) {
	sql := fmt.Sprintf("SELECT sum(%s) FROM event_dist WHERE ModelId = ? AND %s AND Time >= ? AND Time <= ?", aggregateField, generateQuery(searchField, searchFieldValue))
	err = a.clickHouseConn.
		QueryRowxContext(ctx, sql, modelId, startsAt, endsAt).
		Scan(&rv)
	return
}

func (a *Aggregator) Min(ctx context.Context, modelId int64, searchField string, searchFieldValue interface{}, startsAt, endsAt time.Time, aggregateField string) (rv float64, err error) {
	sql := fmt.Sprintf("SELECT min(%s) FROM event_dist WHERE ModelId = ? AND %s AND Time >= ? AND Time <= ?", aggregateField, generateQuery(searchField, searchFieldValue))
	err = a.clickHouseConn.
		QueryRowxContext(ctx, sql, modelId, startsAt, endsAt).
		Scan(&rv)
	return
}

func (a *Aggregator) Max(ctx context.Context, modelId int64, searchField string, searchFieldValue interface{}, startsAt, endsAt time.Time, aggregateField string) (rv float64, err error) {
	sql := fmt.Sprintf("SELECT max(%s) FROM event_dist WHERE ModelId = ? AND %s AND Time >= ? AND Time <= ?", aggregateField, generateQuery(searchField, searchFieldValue))
	err = a.clickHouseConn.
		QueryRowxContext(ctx, sql, modelId, startsAt, endsAt).
		Scan(&rv)
	return
}

// SD 求标准差
func (a *Aggregator) SD(ctx context.Context, modelId int64, searchField string, searchFieldValue interface{}, startsAt, endsAt time.Time, aggregateField string) (rv float64, err error) {
	variance, err := a.Variance(ctx, modelId, searchField, searchFieldValue, startsAt, endsAt, aggregateField)
	if err != nil {
		return 0, err
	}
	return math.Sqrt(variance), nil
}

// Variance 方差 \sum_i w_i (x_i - mean)^2 / (sum_i w_i - 1)
func (a *Aggregator) Variance(ctx context.Context, modelId int64, searchField string, searchFieldValue interface{}, startsAt, endsAt time.Time, aggregateField string) (rv float64, err error) {
	sql := fmt.Sprintf("SELECT %s FROM event_dist WHERE ModelId = ? AND %s AND Time >= ? AND Time <= ?", aggregateField, generateQuery(searchField, searchFieldValue))

	rows, err := a.clickHouseConn.QueryContext(ctx, sql, modelId, startsAt, endsAt)
	if err != nil {
		return 0, err
	}

	var records []float64
	for rows.Next() {
		var x float64
		if err := rows.Scan(&x); err != nil {
			return 0, err
		}
		records = append(records, x)
	}

	return stat.Variance(records, nil), nil
}

// Deviation 偏离率
func (a *Aggregator) Deviation(ctx context.Context, modelId int64, searchField string, searchFieldValue interface{}, startsAt, endsAt time.Time, aggregateField string, aggregateFieldValue float64) (rv float64, err error) {
	avg, err := a.Average(ctx, modelId, searchField, searchFieldValue, startsAt, endsAt, aggregateField)
	if err != nil {
		return 0, err
	}

	deviationVal := math.Abs(aggregateFieldValue - avg)
	return deviationVal, nil
}

func (a *Aggregator) Median(ctx context.Context, modelId int64, searchField string, searchFieldValue interface{}, startsAt, endsAt time.Time, aggregateField string) (rv float64, err error) {
	sql := fmt.Sprintf("SELECT quantileExact(0.5)(%s) FROM event_dist WHERE ModelId = ? AND %s AND Time >= ? AND Time <= ?", aggregateField, generateQuery(searchField, searchFieldValue))
	err = a.clickHouseConn.
		QueryRowxContext(ctx, sql, modelId, startsAt, endsAt).
		Scan(&rv)
	return
}

func quoted(s string) string { return "'" + s + "'" }

// FIXME: dirty hack
func generateQuery(field string, value interface{}) (stmt string) {
	value = indirect(value)
	switch value.(type) {
	case int, int64, int32, int16, int8, uint, uint64, uint32, uint16, uint8:
		return fmt.Sprintf("`%s` = %d", field, value)
	case float32, float64:
		return fmt.Sprintf("`%s` = %f", field, value)
	default:
		return fmt.Sprintf("`%s` = %s", field, quoted(cast.ToString(value)))
	}
}

// From html/template/content.go
// Copyright 2011 The Go Authors. All rights reserved.
// indirect returns the value, after dereferencing as many times
// as necessary to reach the base type (or nil).
func indirect(a interface{}) interface{} {
	if a == nil {
		return nil
	}
	if t := reflect.TypeOf(a); t.Kind() != reflect.Ptr {
		// Avoid creating a reflect.Value if it's not a pointer.
		return a
	}
	v := reflect.ValueOf(a)
	for v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
	}
	return v.Interface()
}
