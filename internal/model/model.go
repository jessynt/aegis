package model

import (
	"time"
)

// Model 风控模型
type Model struct {
	Id     int64  `db:"id,pk"`
	Name   string `db:"name"`
	Label  string `db:"label"`
	GUID   string `db:"guid"`
	Status int8   `db:"status"`

	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// GetModelName 获取风控模型名称
func (m Model) GetModelName() string {
	return m.Name
}

// GetModelId 获取风控模型 ID
func (m Model) GetModelId() int64 {
	return m.Id
}

// GetModelGUID 获取风控模型 GUID
func (m Model) GetModelGUID() string {
	return m.GUID
}
