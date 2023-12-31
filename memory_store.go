package bitcask

import "fmt"

type MemoryStore struct {
	data map[string]string
}

func NewMemoryStore() Bitcask {
	return &MemoryStore{data: make(map[string]string)}
}

// Get returns the value for the given key.
func (m *MemoryStore) Get(key string) (string, error) {
	value, found := m.data[key]
	if !found {
		return "", fmt.Errorf("key %q not found", key)
	}
	return value, nil
}

// Put stores a new key and value and returns the old value if the key
// already exists, otherwise an empty string.
func (m *MemoryStore) Put(key string, value string) (string, error) {
	oldValue, _ := m.Get(key)
	m.data[key] = value
	return oldValue, nil
}

// Delete removes a key-value pair from the store.
func (m *MemoryStore) Delete(key string) (string, error) {
	value, found := m.data[key]
	if !found {
		return "", fmt.Errorf("key %q not found", key)
	}
	delete(m.data, key)
	return value, nil
}

// ListKeys returns a slice containing all active keys.
func (m *MemoryStore) ListKeys() []string {
	var out []string
	for key := range m.data {
		out = append(out, key)
	}

	return out
}

// Close closes the store.
func (m *MemoryStore) Close() error {
	return nil
}
