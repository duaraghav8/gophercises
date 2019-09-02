package apiserver

import (
	"encoding/json"
	"github.com/duaraghav8/gophercises/ex17_secretstore/pkg/secretstore"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestAPIServer_HandlerListSecrets(t *testing.T) {

	t.Run("return empty list when no secrets stored", func(t *testing.T) {
		t.Parallel()
		var received []string

		request, _ := http.NewRequest(http.MethodGet, "/keys", nil)
		response := httptest.NewRecorder()

		storage := secretstore.NewInMemoryKVStore()
		server, _ := NewAPIServer(storage)

		server.ServeHTTP(response, request)

		if err := json.NewDecoder(response.Body).Decode(&received); err != nil {
			t.Fatalf(
				"Unable to parse response %q from server into string slice: %v",
				response.Body,
				err,
			)
		}

		assertStatus(t, response.Code, http.StatusOK)

		if len(received) != 0 {
			t.Errorf("Expected empty slice, received %v", received)
		}
	})

	t.Run("return all keys stored", func(t *testing.T) {
		t.Parallel()
		var received []string

		expected := []string{"foo", "harry", "a380", "a-89hy&hashi##"}

		request, _ := http.NewRequest(http.MethodGet, "/keys", nil)
		response := httptest.NewRecorder()

		storage := secretstore.NewInMemoryKVStore()
		for _, key := range expected {
			_ = storage.Put(key, "random-value")
		}

		server, _ := NewAPIServer(storage)

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

}
