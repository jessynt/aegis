package activation

import (
	"sync"

	"aegis/internal/engine/contract"
	"aegis/internal/engine/store"
)

type Opt func(m *Manager)

type Manager struct {
	activations *sync.Map
	store       *store.Store
}

func NewManager(opts ...Opt) *Manager {
	m := &Manager{
		activations: &sync.Map{},
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

func WithStore(s *store.Store) Opt {
	return func(m *Manager) {
		m.store = s
	}
}

func (m *Manager) Init() error {
	activations := &sync.Map{}
	dbActivations := m.store.GetActivations()
	for _, dbActivation := range dbActivations {
		activations.Store(dbActivation.Name, NewActivationFromModel(dbActivation))
	}
	m.activations = activations
	return nil
}

func (m Manager) GetActivationsByModelId(modelId int64) (rv []contract.Activation) {
	m.activations.Range(func(_, value interface{}) bool {
		a := value.(contract.Activation)
		if a.GetModelId() == modelId {
			rv = append(rv, a)
		}
		return true
	})
	return
}
