package abstraction

import (
	"sync"

	"aegis/internal/engine/contract"
	"aegis/internal/engine/store"
)

type Opt func(*Manager)

type Manager struct {
	abstractions *sync.Map
	store        *store.Store
}

func NewManager(opts ...Opt) *Manager {
	m := &Manager{
		abstractions: &sync.Map{},
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
	abstractions := &sync.Map{}
	dbAbstractions := m.store.GetAbstractions()
	for id, dbAbstraction := range dbAbstractions {
		abstractions.Store(id, NewAbstractionFromModel(dbAbstraction))
	}
	m.abstractions = abstractions
	return nil
}

func (m *Manager) GetAbstractionsByModelId(modelId int64) (rv []contract.Abstraction) {
	m.abstractions.Range(func(_, value interface{}) bool {
		a := value.(*Abstraction)
		if a.GetModelId() == modelId {
			rv = append(rv, a)
		}
		return true
	})
	return
}

func (m *Manager) GetAbstractionById(id int64) contract.Abstraction {
	rv, loaded := m.abstractions.Load(id)
	if !loaded {
		return nil
	}
	return rv.(contract.Abstraction)
}

func (m *Manager) GetAbstractions() (rv []contract.Abstraction) {
	m.abstractions.Range(func(_, value interface{}) bool {
		a := value.(contract.Abstraction)
		rv = append(rv, a)
		return true
	})
	return
}
