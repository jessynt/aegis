package property

import (
	"sync"

	"aegis/internal/engine/contract"
	"aegis/internal/engine/store"
)

type Opt func(m *Manager)

type Manager struct {
	properties *sync.Map
	store      *store.Store
}

func NewManager(opts ...Opt) *Manager {
	m := &Manager{
		properties: &sync.Map{},
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

// Init 初始化
func (m *Manager) Init() error {
	properties := &sync.Map{}
	dbProperties := m.store.GetProperties()
	for _, dbProperties := range dbProperties {
		p := NewPropertyFromModel(dbProperties)
		properties.Store(p.GetPropertyName(), p)
	}
	m.properties = properties
	return nil
}

// GetProperties 获取属性列表
func (m Manager) GetProperties() (properties []contract.Property) {
	m.properties.Range(func(_, v interface{}) bool {
		properties = append(properties, v.(contract.Property))
		return true
	})
	return properties
}

func (m Manager) GetById(id int64) (rv contract.Property) {
	m.properties.Range(func(_, value interface{}) bool {
		if value.(contract.Property).Id() == id {
			rv = value.(contract.Property)
		}
		return true
	})
	return
}

// GetByName 通过属性名称获取属性
func (m Manager) GetByName(name string) (property contract.Property, ok bool) {
	v, ok := m.properties.Load(name)
	if !ok {
		return nil, false
	}
	return v.(contract.Property), true
}
