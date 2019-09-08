package apiserver

import (
	"encoding/json"
	"github.com/duaraghav8/gophercises/ex17_secretstore/pkg/secretstore"
	"github.com/gorilla/mux"
	"net/http"
)

type WriteSecretRequest struct {
	Value string `json:"value"`
}

func (server *APIServer) HandlerWriteSecret(w http.ResponseWriter, r *http.Request) {
	var payload WriteSecretRequest

	key := mux.Vars(r)["key"]
	encodingKey := r.Header.Get(EncodingKeyHeader)

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if encodingKey == "" {
		http.Error(w, "Invalid Encryption key", http.StatusBadRequest)
		return
	}

	cypherText, err := secretstore.Encrypt(payload.Value, encodingKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := server.Datastore.Put(key, cypherText); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
