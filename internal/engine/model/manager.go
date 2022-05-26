package model

import (
	"sync"

	"aegis/internal/engine/contract"
	"aegis/internal/engine/store"
)

type Opt func(m *Manager)

type Manager struct {
	modelsByGUID *sync.Map
	store        *store.Store
}

func NewManager(opts ...Opt) *Manager {
	m := &Manager{
		modelsByGUID: &sync.Map{},
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

func (m *Manager) Init() error {
	modelsByGUID := &sync.Map{}
	dbModels := m.store.GetModels()
	for _, dbModel := range dbModels {
		modelsByGUID.Store(dbModel.GUID, dbModel)
	}
	m.modelsByGUID = modelsByGUID
	return nil
}

func WithStore(s *store.Store) Opt {
	return func(m *Manager) {
		m.store = s
	}
}

func (m *Manager) GetModelByGuid(guid string) (contract.Model, bool) {
	v, ok := m.modelsByGUID.Load(guid)
	if !ok {
		return nil, false
	}
	return v.(contract.Model), true
}
