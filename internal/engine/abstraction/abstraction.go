package abstraction

import (
	"aegis/internal/model"
	"aegis/internal/proto"
)

type Abstraction struct {
	m *model.Abstraction
}

func NewAbstractionFromModel(m *model.Abstraction) *Abstraction {
	return &Abstraction{
		m: m,
	}
}

func (a Abstraction) Id() int64 {
	return a.m.Id
}

func (a Abstraction) GetAbstractionName() string {
	return a.m.Name
}

// GetModelId 获取模型 ID
func (a Abstraction) GetModelId() int64 {
	return a.m.ModelId
}

// GetAggregateIntervalValue 获取聚合类型
func (a Abstraction) GetAggregateType() proto.AggregateType {
	return proto.AggregateType(a.m.AggregateType)
}

// GetAggregateIntervalValue 获取聚合字段
func (a Abstraction) GetAggregateField() string {
	return a.m.AggregateField
}

// GetSearchField 获取搜索字段
func (a Abstraction) GetSearchField() string {
	return a.m.SearchField
}

// GetAggregateIntervalValue 获取聚合时间片类型
func (a Abstraction) GetAggregateIntervalType() proto.AggregateIntervalType {
	return proto.AggregateIntervalType(a.m.AggregateIntervalType)
}

// GetAggregateIntervalValue 获取聚合时间片长度
func (a Abstraction) GetAggregateIntervalValue() int64 {
	return a.m.AggregateIntervalValue
}

// GetFilterExpression 获取过滤条件表达式
func (a Abstraction) GetFilterExpression() string {
	return a.m.FilterExpression
}
