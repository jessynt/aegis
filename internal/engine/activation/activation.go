package activation

import (
	"aegis/internal/model"
)

type Activation struct {
	m *model.Activation
}

func NewActivationFromModel(m *model.Activation) *Activation {
	return &Activation{m: m}
}

func (a Activation) Id() int64 {
	return a.m.Id
}

func (a Activation) GetActivationName() string {
	return a.m.Name
}

func (a Activation) GetModelId() int64 {
	return a.m.ModelId
}

func (a Activation) GetWarningScore() int64 {
	return a.m.WarningScore
}

func (a Activation) GetBlockScore() int64 {
	return a.m.BlockScore
}
