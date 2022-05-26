package bag

func (b *Bag) StoreProperty(key string, value interface{}) {
	b.propertiesMu.Lock()
	defer b.propertiesMu.Unlock()

	b.properties[key] = value
}

func (b *Bag) GetProperty(key string) (interface{}, bool) {
	b.propertiesMu.RLock()
	defer b.propertiesMu.RUnlock()

	if a, ok := b.properties[key]; ok {
		return a, true
	}
	return nil, false
}

func (b *Bag) GetProperties() map[string]interface{} {
	b.propertiesMu.RLock()
	defer b.propertiesMu.RUnlock()

	m := make(map[string]interface{})
	for key, value := range b.properties {
		m[key] = value
	}

	return m
}
