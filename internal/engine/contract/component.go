package contract

import (
	"context"
	"time"
)

type ModelManager interface {
	Init() error
	GetModelByGuid(guid string) (model Model, ok bool)
}

type PropertyManager interface {
	Init() error
	GetProperties() []Property
	GetById(id int64) Property
	GetByName(name string) (property Property, ok bool)
}

type AbstractionManager interface {
	Init() error
	GetAbstractionById(id int64) Abstraction
	GetAbstractions() (rv []Abstraction)
	GetAbstractionsByModelId(modelId int64) (rv []Abstraction)
}

type ActivationManager interface {
	Init() error
	GetActivationsByModelId(modelId int64) (rv []Activation)
}

type RuleManager interface {
	Init() error
	GetRulesByActivationId(activationId int64) (rv []Rule)
}

type Aggregator interface {
	Count(ctx context.Context, modelId int64, searchField string, searchFieldValue interface{}, startsAt, endsAt time.Time) (rv float64, err error)
	DistinctCount(ctx context.Context, modelId int64, searchField string, searchFieldValue interface{}, startsAt, endsAt time.Time, aggregateField string) (rv float64, err error)
	Average(ctx context.Context, modelId int64, searchField string, searchFieldValue interface{}, startsAt, endsAt time.Time, aggregateField string) (rv float64, err error)
	Sum(ctx context.Context, modelId int64, searchField string, searchFieldValue interface{}, startsAt, endsAt time.Time, aggregateField string) (rv float64, err error)
	Min(ctx context.Context, modelId int64, searchField string, searchFieldValue interface{}, startsAt, endsAt time.Time, aggregateField string) (rv float64, err error)
	Max(ctx context.Context, modelId int64, searchField string, searchFieldValue interface{}, startsAt, endsAt time.Time, aggregateField string) (rv float64, err error)
	SD(ctx context.Context, modelId int64, searchField string, searchFieldValue interface{}, startsAt, endsAt time.Time, aggregateField string) (rv float64, err error)
	Variance(ctx context.Context, modelId int64, searchField string, searchFieldValue interface{}, startsAt, endsAt time.Time, aggregateField string) (rv float64, err error)
	Deviation(ctx context.Context, modelId int64, searchField string, searchFieldValue interface{}, startsAt, endsAt time.Time, aggregateField string, aggregateFieldValue float64) (rv float64, err error)
	Median(ctx context.Context, modelId int64, searchField string, searchFieldValue interface{}, startsAt, endsAt time.Time, aggregateField string) (rv float64, err error)
}

type CollectionManager interface {
	Init() error
	GetCollections() []Collection
}
