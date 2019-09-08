package secretstore

import "testing"

func TestInMemoryKVStore_Get(t *testing.T) {

	t.Run("returns value if key exists", func(t *testing.T) {
		t.Parallel()

		kv := NewInMemoryKVStore()
		kv.data["foo"] = "bar"

		value, _ := kv.Get("foo")
		if value != "bar" {
			t.Errorf("Expected value bar but received %s", value)
		}
	})

	t.Run("returns custom error when key doesn't exist", func(t *testing.T) {
		t.Parallel()

		kv := NewInMemoryKVStore()

		_, err := kv.Get("foo")
		if err == nil {
			t.Fatal("No error was returned in case of missing key")
		}

		switch err.(type) {
		case *ErrKVStoreKeyNotFound:
			break
		default:
			t.Errorf("Expected KeyNotFound error but received %v", err)
		}
	})

}

func TestInMemoryKVStore_Put(t *testing.T) {

	t.Run("puts value for new key", func(t *testing.T) {
		t.Parallel()

		kv := NewInMemoryKVStore()

		err := kv.Put("foo", "bar")
		if err != nil {
			t.Errorf("Received unexpected error: %v", err)
		}

		val, ok := kv.data["foo"]
		if !ok {
			t.Error("Key does not exist")
		}
		if val != "bar" {
			t.Errorf("Expected value bar but received %s", val)
		}
	})

}

func TestInMemoryKVStore_List(t *testing.T) {

	t.Run("returns a list of all keys", func(t *testing.T) {
		t.Parallel()

		kv := NewInMemoryKVStore()

		kv.data["hello"] = "world"
		kv.data["lorem"] = "ipsum"

		keys, err := kv.List()
		if err != nil {
			t.Fatalf("An unexpected error occured: %v", err)
		}

		if len(keys) != 2 {
			t.Errorf("Unexpected number of keys: %v", keys)
		}
	})

}
