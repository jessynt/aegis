package bag

import (
	"context"
	"fmt"
	"sync"

	"aegis/internal/proto"
)

type Bag struct {
	propertiesMu sync.RWMutex
	properties   map[string]interface{}

	abstractionMu sync.RWMutex
	abstractions  map[string]interface{}

	activationsMu sync.RWMutex
	activations   map[string]*proto.RiskObject
}

func NewBag() *Bag {
	return &Bag{
		properties:   make(map[string]interface{}),
		abstractions: make(map[string]interface{}),
		activations:  make(map[string]*proto.RiskObject),
	}
}

type contextKey struct{}

func New(parent context.Context) (context.Context, *Bag) {
	bag := NewBag()
	return context.WithValue(parent, contextKey{}, bag), bag
}

func FromContext(ctx context.Context) *Bag {
	value := ctx.Value(contextKey{})
	if bag, ok := value.(*Bag); ok {
		return bag
	}

	panic(fmt.Errorf("no bag binded to context: %+v", ctx))
}
