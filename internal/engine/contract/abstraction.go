package contract

import (
	"aegis/internal/proto"
)

type Abstraction interface {
	Id() int64
	GetAbstractionName() string
	GetModelId() int64
	GetAggregateType() proto.AggregateType
	GetAggregateField() string
	GetSearchField() string
	GetAggregateIntervalType() proto.AggregateIntervalType
	GetAggregateIntervalValue() int64
	GetFilterExpression() string
}
