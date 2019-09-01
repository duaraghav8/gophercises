package apiserver

import (
	"github.com/duaraghav8/gophercises/ex17_secretstore/pkg/secretstore"
	"github.com/gorilla/mux"
	"net/http"
)

type APIServer struct {
	Datastore secretstore.KVStore
	*mux.Router
}

func (server *APIServer) HandlerListSecrets(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusAccepted)
}
func (server *APIServer) HandlerReadSecret(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusAccepted)
}
func (server *APIServer) HandlerWriteSecret(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusAccepted)
}

func NewAPIServer(kvstore secretstore.KVStore) (http.Handler, error) {
	server := APIServer{Datastore: kvstore, Router: mux.NewRouter()}

	registerAPIRoutes(server)
	return server, nil
}
