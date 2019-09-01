package secretstore

type KVStore interface {
	Get(key, encodingKey string) (string, error)
	Put(key, value, encodingKey string) error
	List() ([]string, error)
}

type InMemoryKVStore struct {
	data map[string]string
}

func (store *InMemoryKVStore) Get(key, encodingKey string) (string, error) {
	return "Hello", nil
}

func (store *InMemoryKVStore) Put(key, value, encodingKey string) error {
	return nil
}

func (store *InMemoryKVStore) List() ([]string, error) {
	return []string{}, nil
}

func NewInMemoryKVStore() *InMemoryKVStore {
	return new(InMemoryKVStore)
}
