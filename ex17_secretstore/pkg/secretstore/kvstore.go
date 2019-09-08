package secretstore

import "fmt"

// KVStore interface defines the functionality a Key-Value store should provide
// in a thread-safe manner.
type KVStore interface {
	Get(key string) (string, error)
	Put(key, value string) error
	List() ([]string, error)
}

type KVStoreError struct {
	Key string
}

type ErrKVStoreKeyNotFound KVStoreError
type ErrCryptoIncorrectEncodingKey KVStoreError

func (err ErrKVStoreKeyNotFound) Error() string {
	return fmt.Sprintf("%s: key does not exist", err.Key)
}

func (err ErrCryptoIncorrectEncodingKey) Error() string {
	return "incorrect encoding key"
}
