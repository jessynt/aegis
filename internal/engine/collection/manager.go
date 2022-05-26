package collection

import (
	"sync"

	"aegis/internal/engine/contract"
	"aegis/internal/engine/store"
	"aegis/internal/model"
)

type Opt func(m *Manager)

type Manager struct {
	collections *sync.Map
	store       *store.Store
}

func NewManager(opts ...Opt) *Manager {
	m := &Manager{
		collections: &sync.Map{},
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
	collections := &sync.Map{}
	collectionItemsByCollectionId := make(map[int64][]*model.CollectionItem)
	for _, item := range m.store.GetCollectionItems() {
		collectionItemsByCollectionId[item.CollectionId] = append(collectionItemsByCollectionId[item.CollectionId], item)
	}
	for _, c := range m.store.GetCollections() {
		collections.Store(c.Id, NewFromModel(c, collectionItemsByCollectionId[c.Id]))
	}
	m.collections = collections
	return nil
}

func (m *Manager) GetCollections() (rv []contract.Collection) {
	m.collections.Range(func(_, value interface{}) bool {
		rv = append(rv, value.(contract.Collection))
		return true
	})
	return
}
