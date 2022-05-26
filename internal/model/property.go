package model

import (
	"time"
)

// Property 事件属性
type Property struct {
	Id           int64  `db:"id,pk"`
	Name         string `db:"name"`
	Label        string `db:"label"`
	Type         int32  `db:"type"`
	ValidateType int32  `db:"validate_type"`
	ValidateArgs string `db:"validate_args"`

	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
