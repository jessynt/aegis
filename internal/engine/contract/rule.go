package contract

import (
	"aegis/internal/proto"
)

type Rule interface {
	GetRuleLabel() string
	GetActivationId() int64
	GetAbstractionId() int64
	GetBaseScore() int64
	GetBaseNum() int64
	GetRate() float64
	GetOperatorType() proto.OperatorType
	GetExpression() string
}
