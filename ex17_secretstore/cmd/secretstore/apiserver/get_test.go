package apiserver

import (
	"encoding/json"
	"fmt"
	"github.com/duaraghav8/gophercises/ex17_secretstore/pkg/secretstore"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestAPIServer_HandlerReadSecret(t *testing.T) {

	t.Run("return 404 when secret doesn't exist", func(t *testing.T) {
		t.Parallel()
		request, _ := http.NewRequest(http.MethodGet, "/secret/nonexistent", nil)
		response := httptest.NewRecorder()

		server, _ := NewAPIServer(secretstore.NewInMemoryKVStore())

		request.Header.Set("X-SECRETSTORE-ENCODING-KEY", "random-encoding-key")
		server.ServeHTTP(response, request)
		assertStatus(t, response.Code, http.StatusNotFound)
	})

	t.Run("return expected value when secret exists", func(t *testing.T) {
		t.Parallel()
		var received map[string]string

		expected := map[string]string{"key": "foo", "value": "bar"}
		encodingKey := "gopher"

		request, _ := http.NewRequest(
			http.MethodGet,
			fmt.Sprintf("/secret/%s", expected["key"]),
			nil,
		)
		response := httptest.NewRecorder()

		ciphertext, err := secretstore.Encrypt(expected["value"], encodingKey)
		if err != nil {
			t.Fatalf("Failed to encrypt test data: %v", err)
		}

		storage := secretstore.NewInMemoryKVStore()
		_ = storage.Put(expected["key"], ciphertext)

		server, _ := NewAPIServer(storage)

		request.Header.Set(EncodingKeyHeader, encodingKey)
		server.ServeHTTP(response, request)

		if err := json.NewDecoder(response.Body).Decode(&received); err != nil {
			t.Fatalf(
				"Unable to parse response %q from server into string slice: %v",
				response.Body,
				err,
			)
		}

		assertStatus(t, response.Code, http.StatusOK)

		if !reflect.DeepEqual(received, expected) {
			t.Errorf("Expected %v but received %v", expected, received)
		}
	})

	t.Run("return 401 Unauthorized if encoding key not provided", func(t *testing.T) {
		t.Parallel()
		request, _ := http.NewRequest(http.MethodGet, "/secret/randomkey", nil)
		response := httptest.NewRecorder()

		server, _ := NewAPIServer(secretstore.NewInMemoryKVStore())

		server.ServeHTTP(response, request)
		assertStatus(t, response.Code, http.StatusUnauthorized)
	})

	t.Run("return 401 Unauthorized if incorrect encoding key provided", func(t *testing.T) {
		t.Parallel()

		var (
			key         = "foo"
			value       = "bar"
			encodingKey = "gopher"
		)

		request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/secret/%s", key), nil)
		response := httptest.NewRecorder()

		ciphertext, err := secretstore.Encrypt(value, encodingKey)
		if err != nil {
			t.Fatalf("Failed to encrypt test data: %v", err)
		}

		storage := secretstore.NewInMemoryKVStore()
		_ = storage.Put(key, ciphertext)

		server, _ := NewAPIServer(storage)

		request.Header.Set(EncodingKeyHeader, "INCORRECT-ENCODING-KEY")
		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusUnauthorized)
	})

}
