package secretstore

import "fmt"

// KVStore interface defines the functionality a Key-Value store should provide
// in a thread-safe manner.
type KVStore interface {
	Get(key string) (string, error)
	Put(key, value string) error
	List() ([]string, error)
}

type InMemoryKVStore struct {
	data map[string]string
}

func (store *InMemoryKVStore) Get(key string) (string, error) {
	return "Hello", nil
}

func (store *InMemoryKVStore) Put(key, value string) error {
	return nil
}

func (store *InMemoryKVStore) List() ([]string, error) {
	return []string{}, nil
}

func NewInMemoryKVStore() *InMemoryKVStore {
	return new(InMemoryKVStore)
}

func Encrypt(plaintext, encodingKey string) (string, error) {
	return "", nil
}
func Decrypt(ciphertext, encodingKey string) (string, error) {
	return "", nil
}

type ErrKVStoreKeyNotFound struct {
	Key string
}

func (err ErrKVStoreKeyNotFound) Error() string {
	return fmt.Sprintf("%s: key does not exist", err.Key)
}

type ErrCryptoIncorrectEncodingKey struct {}

func (err ErrCryptoIncorrectEncodingKey) Error() string {
	return "incorrect encoding key"
}