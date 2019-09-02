package apiserver

import (
	"github.com/duaraghav8/gophercises/ex17_secretstore/pkg/secretstore"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

type APIServer struct {
	Datastore secretstore.KVStore
	*mux.Router
}

const EncodingKeyHeader = "X-SECRETSTORE-ENCODING-KEY"

var logger = log.New(os.Stdout, "http: ", log.LstdFlags)

func NewAPIServer(kvStore secretstore.KVStore) (http.Handler, error) {
	server := APIServer{Datastore: kvStore, Router: mux.NewRouter()}

	registerAPIRoutes(server)
	return server, nil
}
