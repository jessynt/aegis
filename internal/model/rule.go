package model

import (
	"database/sql"
	"time"
)

// Rule 风控规则
type Rule struct {
	Id            int64         `db:"id,pk"`
	Label         string        `db:"label"`
	ActivationId  int64         `db:"activation_id"`
	AbstractionId sql.NullInt64 `db:"abstraction_id"` // 指标字段（为空时按总体计算 i.e. COUNT(*)）
	BaseScore     int64         `db:"base_score"`     // 命中初始得分
	BaseNum       int64         `db:"base_num"`       // 命中基数（配合运算符，与指标字段进行计算）
	Expression    string        `db:"expression"`     // 表达式
	Operator      int64         `db:"operator"`       // 运算符
	Rate          int64         `db:"rate"`           // 比率
	CreatedAt     time.Time     `db:"created_at"`
	UpdatedAt     time.Time     `db:"updated_at"`
}
