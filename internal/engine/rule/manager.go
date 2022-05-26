package rule

import (
	"sync"

	"aegis/internal/engine/contract"
	"aegis/internal/engine/store"
)

type Opt func(m *Manager)

type Manager struct {
	rules *sync.Map
	store *store.Store
}

func NewManager(opts ...Opt) *Manager {
	m := &Manager{
		rules: &sync.Map{},
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
	rules := &sync.Map{}
	dbRules := m.store.GetRules()
	for id, dbRule := range dbRules {
		rules.Store(id, NewRuleFromModel(dbRule))
	}
	m.rules = rules
	return nil
}

func (m Manager) GetRulesByActivationId(activationId int64) (rv []contract.Rule) {
	m.rules.Range(func(key, value interface{}) bool {
		r := value.(contract.Rule)
		if r.GetActivationId() == activationId {
			rv = append(rv, r)
		}
		return true
	})
	return
}
