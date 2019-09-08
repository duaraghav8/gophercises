package secretstore

// InMemoryKVStore manages a Key-Value store entirely in memory with
// no persistence. It primarily serves as a concept KV store
// implementation that provides basic CRUD functionality.
type InMemoryKVStore struct {
	data map[string]string
}

func (store *InMemoryKVStore) Get(key string) (string, error) {
	var value string

	value, ok := store.data[key]
	if !ok {
		return "", &ErrKVStoreKeyNotFound{key}
	}

	return value, nil
}

func (store *InMemoryKVStore) Put(key, value string) error {
	store.data[key] = value
	return nil
}

func (store *InMemoryKVStore) List() ([]string, error) {
	var keys = make([]string, 0, len(store.data))
	for k := range store.data {
		keys = append(keys, k)
	}
	return keys, nil
}

func NewInMemoryKVStore() *InMemoryKVStore {
	return &InMemoryKVStore{
		data: make(map[string]string, 0),
	}
}
