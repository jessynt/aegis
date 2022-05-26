package bag

import (
	"aegis/internal/proto"
)

func (b *Bag) StoreActivation(key string, value *proto.RiskObject) {
	b.activationsMu.Lock()
	defer b.activationsMu.Unlock()

	b.activations[key] = value
}

func (b *Bag) GetsActivations() map[string]*proto.RiskObject {
	b.activationsMu.RLock()
	defer b.activationsMu.RUnlock()

	m := make(map[string]*proto.RiskObject)
	for key, value := range b.activations {
		m[key] = value
	}

	return m
}
