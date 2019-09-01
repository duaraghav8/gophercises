package apiserver

import (
	"encoding/json"
	"net/http"
)

func (server *APIServer) HandlerListSecrets(w http.ResponseWriter, r *http.Request) {
	list, err := server.Datastore.List()
	if err != nil {
		logger.Printf("List: Datastore operation failed: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(list); err != nil {
		logger.Printf("List: Failed to encode %v: %v\n", list, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
