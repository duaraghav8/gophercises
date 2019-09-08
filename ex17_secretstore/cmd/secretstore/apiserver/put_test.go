package apiserver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/duaraghav8/gophercises/ex17_secretstore/pkg/secretstore"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPIServer_HandlerWriteSecret(t *testing.T) {

	t.Run("writes key-value pair when a valid request is made", func(t *testing.T) {
		t.Parallel()

		const (
			key   = "ruby"
			value = "rails"
		)

		storage := secretstore.NewInMemoryKVStore()
		server, _ := NewAPIServer(storage)
		reqUrl := fmt.Sprintf("/secret/%s", key)
		payload, _ := json.Marshal(map[string]string{"value": value})

		request, _ := http.NewRequest(http.MethodPost, reqUrl, bytes.NewBuffer(payload))
		response := httptest.NewRecorder()

		request.Header.Set("X-SECRETSTORE-ENCODING-KEY", "random")
		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusCreated)

		// Check for the existence of the key
		_, err := storage.Get(key)
		if err != nil {
			switch err.(type) {
			case *secretstore.ErrKVStoreKeyNotFound:
				t.Error("Key was not written to the KV store")
				break
			default:
				t.Fatalf("Unexpected error occured while fetching written key: %v", err)
			}
		}

	})

	t.Run("overwrites an existing key silently", func(t *testing.T) {
		t.Parallel()

		const (
			key   = "ruby"
			value = "rails"
		)

		reqUrl := fmt.Sprintf("/secret/%s", key)
		payload, _ := json.Marshal(map[string]string{"value": value})

		// Putting the value in plaintext is fine because here
		// we aren't checking for value, it could be anything
		storage := secretstore.NewInMemoryKVStore()
		_ = storage.Put(key, value)

		server, _ := NewAPIServer(storage)
		request, _ := http.NewRequest(http.MethodPost, reqUrl, bytes.NewBuffer(payload))
		response := httptest.NewRecorder()

		request.Header.Set("X-SECRETSTORE-ENCODING-KEY", "random")
		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusCreated)

		// Ensure that the key's value now is not the same as
		// the value stored originally, since the store secret
		// middleware stored the encrypted version of that value.
		// This proves that the Key's value was overwritten.
		datastoreValue, _ := storage.Get(key)
		if datastoreValue == value {
			t.Errorf("Value was not overwritten for an existing key")
		}
	})

	t.Run("return 400 if encoding key is empty or not provided", func(t *testing.T) {
		t.Parallel()

		const (
			key   = "ruby"
			value = "rails"
		)

		server, _ := NewAPIServer(secretstore.NewInMemoryKVStore())
		reqUrl := fmt.Sprintf("/secret/%s", key)
		payload, _ := json.Marshal(map[string]string{"value": value})

		request, _ := http.NewRequest(http.MethodPost, reqUrl, bytes.NewBuffer(payload))
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assertStatus(t, response.Code, http.StatusBadRequest)
	})

}
