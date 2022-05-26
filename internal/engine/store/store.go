package store

import (
	"context"
	"sync"

	"github.com/jmoiron/sqlx"

	"aegis/internal/model"
)

type Store struct {
	mysqlConn       *sqlx.DB
	clickHouseConn  *sqlx.DB
	models          *sync.Map
	properties      *sync.Map
	abstractions    *sync.Map
	activations     *sync.Map
	rules           *sync.Map
	collections     *sync.Map
	collectionItems *sync.Map
}

func NewStore(mysqlConn *sqlx.DB, clickHouseConn *sqlx.DB) *Store {
	return &Store{
		mysqlConn:       mysqlConn,
		clickHouseConn:  clickHouseConn,
		models:          &sync.Map{},
		properties:      &sync.Map{},
		abstractions:    &sync.Map{},
		activations:     &sync.Map{},
		rules:           &sync.Map{},
		collections:     &sync.Map{},
		collectionItems: &sync.Map{},
	}
}

func (s *Store) Init(ctx context.Context) error {
	if err := loadModels(ctx, s); err != nil {
		return err
	}
	if err := loadProperties(ctx, s); err != nil {
		return err
	}
	if err := loadAbstractions(ctx, s); err != nil {
		return err
	}
	if err := loadActivations(ctx, s); err != nil {
		return err
	}
	if err := loadRules(ctx, s); err != nil {
		return err
	}
	if err := loadCollections(ctx, s); err != nil {
		return err
	}
	if err := loadCollectionItems(ctx, s); err != nil {
		return err
	}
	return nil
}

func (s *Store) GetModels() map[string]*model.Model {
	rv := make(map[string]*model.Model)
	s.models.Range(func(k, v interface{}) bool {
		rv[k.(string)] = v.(*model.Model)
		return true
	})
	return rv
}

func (s *Store) GetProperties() map[string]*model.Property {
	rv := make(map[string]*model.Property)
	s.properties.Range(func(k, v interface{}) bool {
		rv[k.(string)] = v.(*model.Property)
		return true
	})
	return rv
}

func (s *Store) GetAbstractions() map[int64]*model.Abstraction {
	rv := make(map[int64]*model.Abstraction)
	s.abstractions.Range(func(k, v interface{}) bool {
		rv[k.(int64)] = v.(*model.Abstraction)
		return true
	})
	return rv
}

func (s *Store) GetActivations() map[string]*model.Activation {
	rv := make(map[string]*model.Activation)
	s.activations.Range(func(k, v interface{}) bool {
		rv[k.(string)] = v.(*model.Activation)
		return true
	})
	return rv
}

func (s *Store) GetRules() map[int64]*model.Rule {
	rv := make(map[int64]*model.Rule)
	s.rules.Range(func(k, v interface{}) bool {
		rv[k.(int64)] = v.(*model.Rule)
		return true
	})
	return rv
}

func (s *Store) GetCollections() map[int64]*model.Collection {
	rv := make(map[int64]*model.Collection)
	s.collections.Range(func(k, v interface{}) bool {
		rv[k.(int64)] = v.(*model.Collection)
		return true
	})
	return rv
}

func (s *Store) GetCollectionItems() map[int64]*model.CollectionItem {
	rv := make(map[int64]*model.CollectionItem)
	s.collectionItems.Range(func(k, v interface{}) bool {
		rv[k.(int64)] = v.(*model.CollectionItem)
		return true
	})
	return rv
}

func (s *Store) ClickHouseConn() *sqlx.DB { return s.clickHouseConn }
func (s *Store) MysqlConn() *sqlx.DB      { return s.mysqlConn }

func loadModels(ctx context.Context, s *Store) error {
	var models []*model.Model

	if err := sqlx.SelectContext(ctx, s.mysqlConn, &models, "SELECT * FROM models"); err != nil {
		return err
	}

	for _, m := range models {
		s.models.Store(m.Name, m)
	}
	return nil
}

func loadProperties(ctx context.Context, s *Store) error {
	var properties []*model.Property

	if err := sqlx.SelectContext(ctx, s.mysqlConn, &properties, "SELECT * FROM properties"); err != nil {
		return err
	}

	for _, property := range properties {
		s.properties.Store(property.Name, property)
	}
	return nil
}

func loadActivations(ctx context.Context, s *Store) error {
	as := &sync.Map{}
	var activations []*model.Activation
	if err := sqlx.SelectContext(ctx, s.mysqlConn, &activations, "SELECT * FROM activations"); err != nil {
		return err
	}
	for _, activation := range activations {
		as.Store(activation.Name, activation)
	}
	s.activations = as
	return nil
}

func loadRules(ctx context.Context, s *Store) error {
	rs := &sync.Map{}
	var rules []*model.Rule

	if err := sqlx.SelectContext(ctx, s.mysqlConn, &rules, "SELECT * FROM rules"); err != nil {
		return err
	}
	for _, rule := range rules {
		rs.Store(rule.Id, rule)
	}
	s.rules = rs
	return nil
}

func loadAbstractions(ctx context.Context, s *Store) error {
	as := &sync.Map{}
	var abstractions []*model.Abstraction

	if err := sqlx.SelectContext(ctx, s.mysqlConn, &abstractions, "SELECT * FROM abstractions"); err != nil {
		return err
	}

	for _, abstraction := range abstractions {
		as.Store(abstraction.Id, abstraction)
	}
	s.abstractions = as
	return nil
}

func loadCollections(ctx context.Context, s *Store) error {
	cs := &sync.Map{}
	var collections []*model.Collection

	if err := sqlx.SelectContext(ctx, s.mysqlConn, &collections, "SELECT * FROM collections"); err != nil {
		return err
	}

	for _, collection := range collections {
		cs.Store(collection.Id, collection)
	}

	s.collections = cs
	return nil
}

func loadCollectionItems(ctx context.Context, s *Store) error {
	cts := &sync.Map{}
	var collectionItems []*model.CollectionItem

	if err := sqlx.SelectContext(ctx, s.mysqlConn, &collectionItems, "SELECT * FROM collection_items"); err != nil {
		return err
	}

	for _, collectionItem := range collectionItems {
		cts.Store(collectionItem.Id, collectionItem)
	}
	s.collectionItems = cts
	return nil
}
