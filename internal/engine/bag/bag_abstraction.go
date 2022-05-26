package bag

func (b *Bag) StoreAbstraction(key string, value interface{}) {
	b.abstractionMu.Lock()
	defer b.abstractionMu.Unlock()

	b.abstractions[key] = value
}

func (b *Bag) GetAbstractionByKey(key string) (interface{}, bool) {
	b.abstractionMu.RLock()
	defer b.abstractionMu.RUnlock()
	if a, ok := b.abstractions[key]; ok {
		return a, true
	}
	return nil, false
}

func (b *Bag) GetAbstractions() map[string]interface{} {
	b.abstractionMu.RLock()
	defer b.abstractionMu.RUnlock()

	m := make(map[string]interface{})
	for k, v := range b.abstractions {
		m[k] = v
	}

	return m
}
