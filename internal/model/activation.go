package model

import (
	"time"
)

type Activation struct {
	Id           int64     `db:"id,pk"`
	Name         string    `db:"name"`
	Label        string    `db:"label"`
	ModelId      int64     `db:"model_id"`
	WarningScore int64     `db:"warning_score"` // 警戒值
	BlockScore   int64     `db:"block_score"`   // 拒绝值
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}
