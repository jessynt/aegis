package rule

import (
	"aegis/internal/model"
	"aegis/internal/proto"
)

type Rule struct {
	m *model.Rule
}

func NewRuleFromModel(m *model.Rule) *Rule {
	return &Rule{m: m}
}

func (r Rule) GetRuleLabel() string {
	return r.m.Label
}

func (r Rule) GetActivationId() int64 {
	return r.m.ActivationId
}

func (r Rule) GetAbstractionId() int64 {
	return r.m.AbstractionId.Int64
}

func (r Rule) GetBaseScore() int64 {
	return r.m.BaseScore
}

func (r Rule) GetBaseNum() int64 {
	return r.m.BaseNum
}

func (r Rule) GetRate() float64 {
	return float64(r.m.Rate) * 0.01
}

func (r Rule) GetOperatorType() proto.OperatorType {
	return proto.OperatorType(r.m.Operator)
}

func (r Rule) GetExpression() string {
	return r.m.Expression
}
