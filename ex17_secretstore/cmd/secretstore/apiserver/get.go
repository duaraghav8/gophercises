package apiserver

import (
	"encoding/json"
	"github.com/duaraghav8/gophercises/ex17_secretstore/pkg/secretstore"
	"github.com/gorilla/mux"
	"net/http"
)

type ReadSecretResponse struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (server *APIServer) HandlerReadSecret(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["key"]
	encodingKey := r.Header.Get(EncodingKeyHeader)

	valueCiphertext, err := server.Datastore.Get(key)
	if err != nil {
		switch err.(type) {
		case *secretstore.ErrKVStoreKeyNotFound:
			http.Error(w, err.Error(), http.StatusNotFound)
			break
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}

		return
	}

	valuePlaintext, err := secretstore.Decrypt(valueCiphertext, encodingKey)
	if err != nil {
		switch err.(type) {
		case *secretstore.ErrCryptoIncorrectEncodingKey:
			http.Error(w, err.Error(), http.StatusUnauthorized)
			break
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}

		return
	}

	response := &ReadSecretResponse{key, valuePlaintext}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		logger.Printf("Get: Failed to encode %v: %v\n", response, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
