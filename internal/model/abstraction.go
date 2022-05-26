package model

import (
	"time"
)

// Abstraction 特征工程
type Abstraction struct {
	Id                     int64  `db:"id,pk"`
	Name                   string `db:"name"`
	Label                  string `db:"label"`
	ModelId                int64  `db:"model_id"`
	AggregateType          int8   `db:"aggregate_type"`           // 聚合类型
	AggregateField         string `db:"aggregate_field"`          // 聚合字段
	AggregateIntervalType  int8   `db:"aggregate_interval_type"`  // 聚合时间片类型
	AggregateIntervalValue int64  `db:"aggregate_interval_value"` // 聚合时间片长度
	SearchField            string `db:"search_field"`             // 搜索字段
	FilterExpression       string `db:"filter_expression"`        // 过滤条件表达式

	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
